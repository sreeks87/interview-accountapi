package account

// This package contains the unit tests for the service
import (
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestCreate400NilBody(t *testing.T) {
	c := &http.Client{}
	s := NewAccountService(c, "http://localhost:1237")
	_, e := s.Create(nil)
	assert.NotEqual(t, e, nil)
}

func TestCreate400EmptyBody(t *testing.T) {
	c := &http.Client{}
	s := NewAccountService(c, "http://localhost:1237")
	_, e := s.Create(&Data{})
	assert.NotEqual(t, e, nil)
}

func TestCreate400IDnil(t *testing.T) {
	c := &http.Client{}
	s := NewAccountService(c, "http://localhost:1237")
	di := &Data{
		Account: AccountData{
			ID:             "",
			Attributes:     AccountAttributes{},
			OrganisationID: "",
			Type:           "Personal",
			Version:        1,
		},
	}
	_, e := s.Create(di)
	assert.NotEqual(t, e, nil)
}

func TestCreate200(t *testing.T) {
	d := &Data{
		Account: AccountData{
			ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			Attributes:     AccountAttributes{},
			OrganisationID: "",
			Type:           "Personal",
			Version:        1,
		},
	}
	r := []byte(`{
		"data": {
		  "type": "accounts",
		  "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		  "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		  "version": 0,
		  "attributes": {
			"country": "GB",
			"base_currency": "GBP",
			"account_number": "41426819",
			"bank_id": "400300",
			"bank_id_code": "GBDSC",
			"bic": "NWBKGB22",
			"iban": "GB11NWBK40030041426819",
			"title": "Ms",
			"first_name": "Samantha",
			"bank_account_name": "Samantha Holder",
			"alternative_bank_account_names": [
			  "Sam Holder"
			],
			"account_classification": "Personal",
			"joint_account": false,
			"account_matching_opt_out": false,
			"secondary_identification": "A1B2C3D4"
		  }
		}
	  }`)
	httpmock.Activate()
	httpmock.RegisterResponder("POST", "http://localhost:1237/v1/organisation/accounts",
		httpmock.NewBytesResponder(200, r))
	defer httpmock.DeactivateAndReset()
	c := &http.Client{}
	s := NewAccountService(c, "http://localhost:1237")

	res, _ := s.Create(d)
	assert.Equal(t, res.Account.ID, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
}

func TestFetch200(t *testing.T) {
	r := []byte(`{
		"data": {
		  "type": "accounts",
		  "id": "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		  "organisation_id": "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
		  "version": 0,
		  "attributes": {
			"country": "GB",
			"base_currency": "GBP",
			"account_number": "41426819",
			"bank_id": "400300",
			"bank_id_code": "GBDSC",
			"bic": "NWBKGB22",
			"iban": "GB11NWBK40030041426819",
			"title": "Ms",
			"first_name": "Samantha",
			"bank_account_name": "Samantha Holder",
			"alternative_bank_account_names": [
			  "Sam Holder"
			],
			"account_classification": "Personal",
			"joint_account": false,
			"account_matching_opt_out": false,
			"secondary_identification": "A1B2C3D4"
		  }
		}
	  }`)
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "http://localhost:1237/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		httpmock.NewBytesResponder(200, r))
	defer httpmock.DeactivateAndReset()
	c := &http.Client{}
	s := NewAccountService(c, "http://localhost:1237")

	res, _ := s.Fetch("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
	assert.Equal(t, res.Account.ID, "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc")
}

func TestFetch400(t *testing.T) {
	c := &http.Client{}
	s := NewAccountService(c, "http://localhost:1237")
	_, e := s.Fetch("")
	assert.NotEqual(t, e, nil)
}

func TestDelete400(t *testing.T) {
	c := &http.Client{}
	s := NewAccountService(c, "http://localhost:1237")
	e := s.Delete("", "")
	assert.NotEqual(t, e, nil)
}

func TestDelete400Emptyversion(t *testing.T) {
	c := &http.Client{}
	s := NewAccountService(c, "http://localhost:1237")
	e := s.Delete("15856685522", "")
	assert.NotEqual(t, e, nil)
}

func TestDelete204(t *testing.T) {
	httpmock.Activate()
	httpmock.RegisterResponder("GET", "http://localhost:1237/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		httpmock.NewBytesResponder(200, nil))
	defer httpmock.DeactivateAndReset()
	c := &http.Client{}
	s := NewAccountService(c, "http://localhost:1237")
	e := s.Delete("ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", "")
	assert.NotEqual(t, e, nil)
}
