package http

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/chaosnote/go-kernel/net/http/method"
)

// timeout ...
const timeout = 15

type Post struct {
	url.URL
	url.Values
	Header map[string]string
}

// Request ...
func (v Post) Request() ([]byte, error) {
	c := &http.Client{
		Timeout: time.Second * timeout,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 3 * time.Second,
			}).Dial,
		},
	}

	u := url.URL(v.URL)
	req, err := http.NewRequest(method.POST.String(), u.String(), strings.NewReader(v.Values.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	for key, value := range v.Header {
		req.Header.Set(key, value)
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

	http.Get(u).Request()

*/
type Get url.URL

// Request ...
func (v Get) Request() ([]byte, error) {
	var c = &http.Client{
		Timeout: time.Second * timeout,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout: 3 * time.Second,
			}).Dial,
		},
	}

	u := url.URL(v)
	res, e := c.Get(u.String())

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
