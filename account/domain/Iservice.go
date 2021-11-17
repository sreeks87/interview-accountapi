package domain

type Service interface {
	Fetch(string) (Data, error)
	Delete(string) error
	Create(*Data) (string, error)
}
