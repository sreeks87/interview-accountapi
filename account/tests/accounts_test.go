package tests

// This test runs against the actual API runnin on the docker container
// Expectation -> an env variable ACCT_SERVERADDRESS is setup to point to the API address
// if not default API address will be used. "http://localhost:8080/v1"
import (
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/sreeks87/interview-accountapi/account"
)

const SERVER = "http://localhost:8080/v1"

// setup and env variable with the address of the docker
// if not found will default to the localhost:8080 address
func setup() (add string) {
	add, b := os.LookupEnv("ACCT_SERVERADDRESS")
	if !b {
		log.Println("defaulting the server address to ", SERVER)
		add = SERVER
	}
	return add
}

// Test the functionality against the docker compose
func TestFunctional(t *testing.T) {
	a, e := AccountCreate()
	if e != nil {
		t.Fatal(e)
	}
	v := a.Account.Version
	if e = AccountFetch(a.Account.ID); e != nil {
		t.Fatal(e)
	}
	if e = AccountDelete(a.Account.ID, strconv.Itoa(int(v))); e != nil {
		t.Fatal(e)
	}
	log.Println("Tests complete")

}

func AccountCreate() (id account.Data, e error) {
	add := setup()
	c := account.NewAcctController(add)
	acct := &account.AccountData{
		ID:             "4bdf8de6-d36d-4912-95be-623007daabb4",
		OrganisationID: "4bdf8de6-d36d-4912-95be-623007daabb4",
		Attributes: account.AccountAttributes{
			AccountClassification:   "Personal",
			AccountMatchingOptOut:   false,
			AccountNumber:           "41426819",
			AlternativeNames:        []string{"Samantha Holder"},
			BankID:                  "400300",
			BankIDCode:              "GBDSC",
			BaseCurrency:            "GBP",
			Bic:                     "NWBKGB22",
			Country:                 "GB",
			Iban:                    "GB11NWBK40030041426819",
			JointAccount:            false,
			Name:                    []string{"Sam Holder"},
			SecondaryIdentification: "SECID",
			Status:                  "confirmed",
			Switched:                false,
		},
		Type:    "accounts",
		Version: 0,
	}
	payload := account.Data{
		Account: *acct,
	}
	createdAcc, e := c.CreateAccount(payload)
	if e != nil {
		return account.Data{}, e
	}
	fmt.Println(id)
	fmt.Println("Account created with UUID-->", createdAcc.Account.ID)
	return createdAcc, nil
}

func AccountFetch(id string) error {
	add := setup()
	c := account.NewAcctController(add)
	fetchedAcc, e := c.FetchAccount(id)
	if e != nil {
		return e
	}
	fmt.Println("Account fetched with UUID-->", id, fetchedAcc.Account)
	return nil
}

func AccountDelete(id string, version string) error {
	add := setup()
	c := account.NewAcctController(add)
	if e := c.DeleteAccount(id, version); e != nil {
		return e
	}
	fmt.Println("Account deleted with UUID and version-->", id, version)
	return nil
}
