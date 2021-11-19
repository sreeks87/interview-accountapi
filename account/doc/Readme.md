### About

The interview-account api is a client library exposed to interact with the Fake Account API

The library exposes three functions via.

Create an account

    The Create function expects the Data in the models.Data format to be passed
    
    Returns a created Data model (account).

Fetch details of the accoount

    The Fetch function expects a valid UUID to be passed as the parameter.
    
    Returns error if the UUID is invalid.

Delete an account
    
    Function expects a valid uuid of the account and the version of the account to be deleted.
    
    Returns error if any one of the parameters is missed/invalid  

Example

`go get github.com/sreeks87/interview-accountapi`

Usage

```go

package main

import (
	"fmt"

	api "github.com/sreeks87/interview-accountapi/account"
)

func main() {
	c := api.NewAcctController("http://localhost:8080")
	acct := &api.AccountData{
		ID:             "4bdf8de6-d36d-4912-95be-623007daabb4",
		OrganisationID: "4bdf8de6-d36d-4912-95be-623007daabb4",
		Attributes: api.AccountAttributes{
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
	payload := api.Data{
		Account: *acct,
	}
	id, _ := c.CreateAccount(payload)
	fmt.Println(id)
	fmt.Println("--------------created--------------")
	a, e := c.FetchAccount("4bdf8de6-d36d-4912-95be-623007daabb4")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(a)
	fmt.Println("-----------------------fetched--------------")

	e = c.DeleteAccount("4bdf8de6-d36d-4912-95be-623007daabb4", "0")
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println("Deleted")

}
```
    
Test

    Unit tests have been written in the service_test.go file to cover most of the cases in the service

    Functional tests have been covered in the /tests folder.
