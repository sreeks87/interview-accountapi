package account

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"
)

const (
	timeout     = 10 * time.Second
	retires     = 3
	api         = "/v1/organisation/accounts/"
	Accept      = "application/vnd.api+json"
	ContentType = "application/vnd.api+json"
)

// The service interface
// A service is supposed to implent the Create Fetch and delete functions
type Service interface {
	Fetch(string) (Data, error)
	Delete(string, string) error
	Create(*Data) (Data, error)
}

type accountService struct {
	client *http.Client
	url    *url.URL
}

// NewAccountService returns the object of a service which can be used by controlled to call the implemented methods.
func NewAccountService(c *http.Client, u string) Service {
	c.Timeout = timeout
	full, _ := getFullURL(u, api)
	return &accountService{
		client: c,
		url:    full,
	}
}

// Create function does the accounbt creation, it returns a Data object from the models.
// input - Data object
// Output - Data, error
func (s *accountService) Create(data *Data) (Data, error) {
	if data == nil {
		return Data{}, errors.New("input cant be nil/empty")
	}
	if data.Account.ID == "" {
		return Data{}, errors.New("uuid cant be nil/empty")
	}
	req, e := json.Marshal(data)
	if e != nil {
		return Data{}, e
	}
	resp, e := s.makeRequest("POST", s.url.String(), bytes.NewBuffer(req))
	if e != nil {
		return Data{}, e
	}
	body, e := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if e != nil {
		return Data{}, e
	}
	var newAcct Data
	json.Unmarshal(body, &newAcct)
	return newAcct, nil
}

// Fetch function fetches the account from API, it returns a Data object from the models.
// input - uuid of the account to be fetched
// Output - Data, error
func (s *accountService) Fetch(id string) (Data, error) {
	if id == "" {
		return Data{}, errors.New("id cant be nil/empty")
	}
	log.Println("fetching ", id)
	fullUrl, _ := getFullURL(s.url.String(), id)
	resp, e := s.makeRequest("GET", fullUrl.String(), nil)
	if e != nil {
		return Data{}, e
	}
	if resp.StatusCode != 200 {
		log.Println("error ", resp)
	}
	body, e := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if e != nil {
		return Data{}, e
	}
	var acctDetails Data
	json.Unmarshal(body, &acctDetails)
	return acctDetails, nil
}

// Delete deletes the created account
// input - uuid and version
// output - error, only in case it failed
func (s *accountService) Delete(id string, version string) error {
	if id == "" || version == "" {
		return errors.New("both id and version anre mandatory")
	}
	log.Println("deleting ", id)
	fullUrl, _ := getFullURL(s.url.String(), id)
	q := fullUrl.Query()      // Get a copy of the query values.
	q.Add("version", version) // Add a new value to the set.
	fullUrl.RawQuery = q.Encode()
	log.Println(fullUrl.String())
	resp, e := s.makeRequest("DELETE", fullUrl.String(), nil)
	if e != nil {
		return e
	}
	body, e := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if e != nil {
		return e
	}
	var acctDetails AccountData
	json.Unmarshal(body, &acctDetails)
	return nil
}

// getFullURL is the helper function for creating the full url
func getFullURL(baseURL string, pathURL string) (*url.URL, error) {
	log.SetPrefix("getFullURL")
	base, e := url.Parse(baseURL)
	if e != nil {
		log.Println("error occured ", e.Error())
		return nil, e
	}
	log.Println("base url :", base)
	base.Path = path.Join(base.Path, pathURL)
	return base, nil
}

// makeRequest is used to make a generic request to the API
func (s *accountService) makeRequest(method string, url string, payload io.Reader) (*http.Response, error) {
	req, e := http.NewRequest(method, url, payload)
	if e != nil {
		return nil, e
	}
	var resp *http.Response
	req.Header.Set("Accept", Accept)
	req.Header.Set("Content-Type", ContentType)
	for i := 0; i < retires; i++ {
		resp, e = s.client.Do(req)
		if e == nil {
			break
		}
	}
	if e != nil {
		return nil, e
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return resp, errors.New(resp.Status)
	}
	return resp, nil
}
