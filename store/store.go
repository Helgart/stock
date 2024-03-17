package store

type Store struct {
	Name string
}

func NewStore(name string) Store {
	return Store{
		Name: name,
	}
}
