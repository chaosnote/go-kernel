package http

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// timeout ...
const timeout = 5

/*
Post example

	u := url.URL{
		Scheme: "http",
		Host: "localhost:port",
		Path: "" // option
	}

	http.Post(u).Request(map[string]string{}, url.Values{})

*/
type Post url.URL

// Request ...
func (v Post) Request(header map[string]string, body url.Values) ([]byte, error) {

	var c = &http.Client{
		Timeout: time.Second * timeout,
	}

	u := url.URL(v)
	req, err := http.NewRequest("POST", u.String(), strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for _key, _value := range header {
		req.Header.Set(_key, _value)
	}

	res, e := c.Do(req)
	if e != nil {
		return nil, e
	}

	defer res.Body.Close()
	b, e := ioutil.ReadAll(res.Body)

	if e != nil {
		return nil, e
	}

	return b, nil
}

/*
Get example

	q := url.Values{}
	q.Add("key", "value")

	u := url.URL{
		Scheme:   "http",
		Host:     "localhost:port",
		Path:     "", // option
		RawQuery: q.Encode(),
	}

	u.Request()

*/
type Get url.URL

// Request ...
func (v Get) Request() ([]byte, error) {
	var h = &http.Client{
		Timeout: time.Second * timeout,
	}

	u := url.URL(v)
	res, e := h.Get(u.String())

	if e != nil {
		return nil, e
	}

	defer res.Body.Close()
	b, e := ioutil.ReadAll(res.Body)

	if e != nil {
		return nil, e
	}

	return b, nil
}
