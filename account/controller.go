package account

import (
	"net/http"
)

type accountController struct {
	svc Service
}

func NewAcctController(base string) *accountController {
	return &accountController{
		svc: NewAccountService(http.DefaultClient, base),
	}
}

func (a *accountController) CreateAccount(data Data) (Data, error) {
	return a.svc.Create(&data)
}

func (a *accountController) DeleteAccount(acid string, version string) error {
	return a.svc.Delete(acid, version)
}

func (a *accountController) FetchAccount(acid string) (Data, error) {
	return a.svc.Fetch(acid)
}
