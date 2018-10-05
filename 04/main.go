package main

import (
	"io/ioutil"
	"golang.org/x/net/html"
	"bytes"
	"github.com/davecgh/go-spew/spew"
)

type Link struct {
	Href string `json:"href,omitempty"`
	Text string `json:"value,omitempty"`
}

func main() {
	byt, err := ioutil.ReadFile("ex4.html")
	if err != nil {
		panic(err)
	}

	var links []Link
	var f func(n *html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			link := new(Link)
			link.Text = n.FirstChild.Data
			for _, a := range n.Attr {
				link.Href = a.Val
			}
			links = append(links, *link)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	n, err := html.Parse(bytes.NewReader(byt))
	if err != nil {
		panic(err)
	}

	f(n)

	spew.Dump(links)
}
