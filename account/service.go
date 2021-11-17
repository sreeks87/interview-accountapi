package account

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/sreeks87/interview-accountapi/account/domain"
)

const (
	contentType = "application/json; charset=utf-8"
	timeout     = 10 * time.Second
	retires     = 3
	api         = "/v1/organisation/accounts/"
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

func (s *accountService) Create(data *domain.Data) (id string, e error) {
	req, e := json.Marshal(data)
	if e != nil {
		return "", e
	}
	payload := bytes.NewBuffer(req)
	resp, e := s.client.Post(s.url.String(), contentType, payload)
	if e != nil {
		return "", e
	}
	body, e := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if e != nil {
		return "", e
	}
	var newAcct domain.AccountData
	json.Unmarshal(body, &data)
	return newAcct.ID, nil
}

func (s *accountService) Fetch(id string) (domain.Data, error) {
	log.Println("fetching ", id)
	fullUrl, _ := getFullURL(s.url.String(), id)
	resp, e := s.client.Get(fullUrl.String())
	if e != nil {
		return domain.Data{}, e
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

func (s *accountService) Delete(id string) error {
	log.Println("deleting ", id)
	fullUrl, _ := getFullURL(s.url.String(), id)
	req, e := http.NewRequest("DELETE", fullUrl.String(), nil)
	if e != nil {
		return e
	}
	resp, e := s.client.Do(req)
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
