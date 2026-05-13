package pokeapi

type RespPokemonEncounters struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}
