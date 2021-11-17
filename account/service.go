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

	"github.com/sreeks87/interview-accountapi/account/domain"
)

const (
	timeout     = 10 * time.Second
	retires     = 3
	api         = "/v1/organisation/accounts/"
	Accept      = "application/vnd.api+json"
	ContentType = "application/vnd.api+json"
)

type accountService struct {
	client *http.Client
	url    *url.URL
}

func NewAccountService(c *http.Client, u string) domain.Service {
	c.Timeout = timeout
	full, _ := getFullURL(u, api)
	return &accountService{
		client: c,
		url:    full,
	}
}

func (s *accountService) Create(data *domain.Data) (domain.Data, error) {
	req, e := json.Marshal(data)
	if e != nil {
		return domain.Data{}, e
	}
	resp, e := s.makeRequest("POST", s.url.String(), bytes.NewBuffer(req))
	if e != nil {
		return domain.Data{}, e
	}
	body, e := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if e != nil {
		return domain.Data{}, e
	}
	var newAcct domain.Data
	json.Unmarshal(body, &newAcct)
	return newAcct, nil
}

func (s *accountService) Fetch(id string) (domain.Data, error) {
	log.Println("fetching ", id)
	fullUrl, _ := getFullURL(s.url.String(), id)
	resp, e := s.makeRequest("GET", fullUrl.String(), nil)
	if e != nil {
		return domain.Data{}, e
	}
	if resp.StatusCode != 200 {
		log.Println("error ", resp)
	}
	body, e := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if e != nil {
		return domain.Data{}, e
	}
	var acctDetails domain.Data
	json.Unmarshal(body, &acctDetails)
	return acctDetails, nil
}

func (s *accountService) Delete(id string, version string) error {
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
	var acctDetails domain.AccountData
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

func (s *accountService) makeRequest(method string, url string, payload io.Reader) (*http.Response, error) {
	req, e := http.NewRequest(method, url, payload)
	if e != nil {
		return nil, e
	}
	req.Header.Set("Accept", Accept)
	req.Header.Set("Content-Type", ContentType)

	resp, e := s.client.Do(req)
	if e != nil {
		return nil, e
	}
	if resp.StatusCode >= http.StatusBadRequest {
		return resp, errors.New(resp.Status)
	}
	return resp, nil
}
