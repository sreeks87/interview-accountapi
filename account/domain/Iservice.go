package domain

type Service interface {
	Fetch(string) (Data, error)
	Delete(string, string) error
	Create(*Data) (Data, error)
}
