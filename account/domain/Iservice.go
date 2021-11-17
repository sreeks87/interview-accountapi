package domain

type Service interface {
	Fetch(string) (AccountData, error)
	Delete(string) error
	Create(*AccountData) (string, error)
}
