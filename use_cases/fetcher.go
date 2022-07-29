package use_cases

import (
	"log"
	"sort"
	"strings"

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

func (f Fetcher) Fetch(from, to int) error {
	var pokemons []models.Pokemon
	done := make(chan interface{})
	defer close(done)

	channel := GetResponses(from, to, f, done)

	for id := from; id <= to; id++ {
		result := <-channel
		if result.Error != nil {
			log.Printf("error: %v", result.Error)
		}
		pokemons = append(pokemons, result.Pokemon)
	}

	sort.Sort(PokemonsByID(pokemons))

	return f.storage.Write(pokemons)
}

func GetResponses(from, to int, f Fetcher, done chan interface{}) <-chan Result {
	results := make(chan Result)
	for id := from; id <= to; id++ {
		go PingAPI(id, f, done, results)
	}

	return results
}

func PingAPI(id int, f Fetcher, done chan interface{}, results chan Result) {
	var result Result
	pokemon, err := f.api.FetchPokemon(id)

	var flatAbilities []string
	for _, t := range pokemon.Abilities {
		flatAbilities = append(flatAbilities, t.Ability.URL)
	}
	pokemon.FlatAbilityURLs = strings.Join(flatAbilities, "|")

	result = Result{Error: err, Pokemon: pokemon}

	select {
	case <-done:
		return
	case results <- result:
	}
}
