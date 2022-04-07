package method

import "strings"

// https://zh.wikipedia.org/wiki/%E5%A2%9E%E5%88%AA%E6%9F%A5%E6%94%B9
// https://www.restapitutorial.com/lessons/httpmethods.html
// RESTful
const (
	POST   classify = "Post"   // POST Create
	GET    classify = "Get"    // GET Read
	PUT    classify = "Put"    // PUT Update/Replace
	PATCH  classify = "Patch"  // PATCH Update/Modify
	DELETE classify = "Delete" // DELETE Delete
)

type classify string

func (v classify) Same(t string) bool {
	return strings.EqualFold(t, string(v))
}

func (v classify) String() string {
	return string(v)
}
