The interview-account api is a client library exposed to interact with the Fake Account API

The library exposes three functions via.

Create an account

    The Create function expects the Data in the models.Data fomrat to be passed
    Returns acreated Data (account) object

Fetch details of the accoount

    The Fetch function expects a valid UUID to be passed as the parameter.
    Returns error if the UUID is invalid.

Delete an account
    
    Function expects a valid uuid of the account and the version of the account to be deleted.
    Returns error if any one of the parameters is missed/invalid  

Example

`go get `

Usage

```go
import  "github.com/sreeks87/interview-accountapi/account"
func AccountCreate() {
c := account.NewAcctController(add)
    acct := &account.AccountData{
        ID:             ID,
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

    fetchedAcc, e := c.FetchAccount(id)
    
    e := c.DeleteAccount(id, version)

}
```
    
Test

    Unit tests have been written in the service_test.go file to cover most of the cases in the service

    Functional tests have been covered in the /tests folder.