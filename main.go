package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
)

const apiURL = "https://pokeapi.co/api/v2/pokemon?limit=100"

// Estruturas para armazenar os dados dos Pokémon
type Pokemon struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonResponse struct {
	Results []Pokemon `json:"results"`
}

// Estrutura para armazenar detalhes do Pokémon
type PokemonDetails struct {
	Name  string
	Types []string
}

// Channel para enviar detalhes do Pokémon
type PokemonChannel struct {
	Type    string
	Pokemon string
}

// Mapa para armazenar os Pokémon separados por tipo
var pokemonTypes = make(map[string][]string)
var mutex = &sync.Mutex{}
var wg sync.WaitGroup

// Variável atômica para contar Pokémon processados
var pokemonCount int32

func main() {
	pokemonResponse, err := fetchPokemonList()
	if err != nil {
		fmt.Println("Erro ao obter lista de Pokémon:", err)
		return
	}

	// Channel para comunicação
	pokemonChan := make(chan PokemonChannel)

	// Inicia uma goroutine para processar os detalhes do Pokémon
	go processPokemonDetails(pokemonChan)

	// Para cada Pokémon, inicia uma goroutine para obter detalhes
	for _, pokemon := range pokemonResponse.Results {
		wg.Add(1)
		go fetchPokemonDetails(pokemon, pokemonChan)
	}

	// Aguarda todas as goroutines terminarem
	wg.Wait()
	close(pokemonChan)

	// Gera e exibe o relatório
	generateReport()
}

func fetchPokemonList() (PokemonResponse, error) {
	response, err := http.Get(apiURL)
	if err != nil {
		return PokemonResponse{}, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return PokemonResponse{}, err
	}

	var pokemonResponse PokemonResponse
	if err := json.Unmarshal(body, &pokemonResponse); err != nil {
		return PokemonResponse{}, err
	}

	return pokemonResponse, nil
}

func fetchPokemonDetails(pokemon Pokemon, pokemonChan chan PokemonChannel) {
	defer wg.Done()

	// Faz uma chamada à API para obter detalhes do Pokémon
	response, err := http.Get(pokemon.URL)
	if err != nil {
		fmt.Println("Erro ao fazer requisição:", err)
		return
	}
	defer response.Body.Close()

	var details struct {
		Types []struct {
			Type struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Erro ao ler resposta:", err)
		return
	}
	json.Unmarshal(body, &details)

	// Envia os detalhes do Pokémon pelo channel
	for _, t := range details.Types {
		pokemonChan <- PokemonChannel{Type: t.Type.Name, Pokemon: pokemon.Name}
	}

	// Atualiza o contador atômico
	atomic.AddInt32(&pokemonCount, 1)
}

func processPokemonDetails(pokemonChan chan PokemonChannel) {
	for pokemonDetail := range pokemonChan {
		mutex.Lock()
		pokemonTypes[pokemonDetail.Type] = append(pokemonTypes[pokemonDetail.Type], pokemonDetail.Pokemon)
		mutex.Unlock()
	}
}

func generateReport() {
	fmt.Println("Relatório de Pokémon por Tipo:")
	for tipo, pokemons := range pokemonTypes {
		fmt.Printf("Tipo: %s\nPokémons: %v\n", tipo, pokemons)
	}
	fmt.Printf("Total de Pokémon processados: %d\n", atomic.LoadInt32(&pokemonCount))
}
