package model

type Category int

type Pokemon struct {
	Name     string
	Height   string
	Weight   string
	Category Category
	Type     []string
	Waekness []string
}

const (
	Undefined Category = iota + 1
	Seed
	Lizard
	Flame
)

func (c Category) toString() string {
	return []string{
		"seed",
		"lizard",
		"flame",
	}[c]
}

func CategoryFromValue(str string) Category {
	switch str {
	case Seed.toString():
		return Seed
	case Lizard.toString():
		return Lizard
	}
	return Undefined
}
