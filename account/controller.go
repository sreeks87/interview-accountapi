package account

import (
	"net/http"

	"github.com/sreeks87/interview-accountapi/account/domain"
)

type accountController struct {
	svc domain.Service
}

func NewAcctController(base string) *accountController {
	return &accountController{
		svc: NewAccountService(http.DefaultClient, base),
	}
}

func (a *accountController) CreateAccount(data domain.Data) (string, error) {
	return a.svc.Create(&data)
}

func (a *accountController) DeleteAccount(acid string) error {
	return a.svc.Delete(acid)
}

func (a *accountController) FetchAccount(acid string) (domain.Data, error) {
	return a.svc.Fetch(acid)
}
