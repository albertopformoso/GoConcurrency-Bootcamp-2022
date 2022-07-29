package use_cases

import (
	"context"
	"log"
	"sort"
	"strings"
	"sync"

	"GoConcurrency-Bootcamp-2022/models"
)

type api interface {
	FetchPokemon(id int) (models.Pokemon, error)
}

type writer interface {
	Write(pokemons []models.Pokemon) error
}

type Fetcher struct {
	api     api
	storage writer
}

func NewFetcher(api api, storage writer) Fetcher {
	return Fetcher{api, storage}
}

type PokemonsByID []models.Pokemon

func (x PokemonsByID) Len() int           { return len(x) }
func (x PokemonsByID) Less(i, j int) bool { return x[i].ID < x[j].ID }
func (x PokemonsByID) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Result struct {
	Error error
	Pokemon models.Pokemon
}

func (f Fetcher) Fetch(ctx context.Context, from, to int) error {
	var pokemons []models.Pokemon
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	channel := GetResponses(ctx, from, to, f)

	for id := from; id <= to; id++ {
		result := <-channel
		if result.Error != nil {
			log.Printf("error: %v", result.Error)
			cancel()
			continue
		}
		pokemons = append(pokemons, result.Pokemon)
	}

	sort.Sort(PokemonsByID(pokemons))

	return f.storage.Write(pokemons)
}

func GetResponses(ctx context.Context, from, to int, f Fetcher) <-chan Result {
	results := make(chan Result)
	wg := sync.WaitGroup{}

	go func() {
		wg.Wait()
		close(results)
	}()

	for id := from; id <= to; id++ {
		wg.Add(1)
		go PingAPI(ctx, wg, id, f, results)
	}

	return results
}

func PingAPI(ctx context.Context, wg sync.WaitGroup, id int, f Fetcher, results chan Result) {
	defer wg.Done()
	var result Result
	pokemon, err := f.api.FetchPokemon(id)

	var flatAbilities []string
	for _, t := range pokemon.Abilities {
		flatAbilities = append(flatAbilities, t.Ability.URL)
	}
	pokemon.FlatAbilityURLs = strings.Join(flatAbilities, "|")

	result = Result{Error: err, Pokemon: pokemon}

	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		return
	case results <- result:
	}
}
