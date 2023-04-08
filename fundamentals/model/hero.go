package model

type Hero struct {
	Name        string
	Age         int
	Nationality string
}

func NewHero(name string, age int, nationality string) *Hero {
	return &Hero{
		Name:        name,
		Age:         age,
		Nationality: nationality,
	}
}

func MockSliceOfHeros() []Hero {
	peterPark := Hero{
		Name: "Peter Park",
		Age:  10,
	}
	return []Hero{peterPark}
}
