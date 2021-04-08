package http

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// timeout ...
const timeout = 5

// Post example
// @see url.URL
// u := url.URL{
// 	Scheme: "http",
// 	Host: "localhost:port",
// }
// u.Path = "aaa/bbb" // option
//
// http.Post(...).Request(...)
type Post url.URL

// Request ...
func (v Post) Request(p url.Values) ([]byte, error) {

	var h = &http.Client{
		Timeout: time.Second * timeout,
	}
	u := url.URL(v)
	res, e := h.PostForm(u.String(), p)

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

// Get example
// http.Get("https://www.google.com/search?q=golang").Request()
type Get string

// Request ...
func (v Get) Request() ([]byte, error) {
	var h = &http.Client{
		Timeout: time.Second * timeout,
	}
	res, e := h.Get(string(v))

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
