package pokemon

type Pokemon struct {
	Name     string
	Height   string
	Weight   string
	Category string
	Type     []string
	Waekness []string
}

func List() []Pokemon {
	var pokemons []Pokemon

	pokemons = append(pokemons, Pokemon{
		Name:     "Bulbasaur",
		Height:   "2' 04\"",
		Weight:   "15.2 lbs",
		Category: "seed",
	})

	return pokemons
}
