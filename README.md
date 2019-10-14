# godom

html parser for golang

```go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/GuNanHai/godom"
)

func main() {
	resp, err := http.Get("http://www.runoob.com/w3cnote/bubble-sort.html")
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err)
	}

	page := godom.Parser(string(body))
	elements := page.Find(".mobile-nav li a")

	for _, e := range elements {
		fmt.Printf("link: %s , Text: %s \n", e.Attr("href"), e.Text)
	}
}

```
