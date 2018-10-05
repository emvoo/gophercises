package main

import (
	"flag"
	"fmt"
		"golang.org/x/net/html"
	"bytes"
		"log"
	"github.com/davecgh/go-spew/spew"
)

type Link struct {
	Href string `json:"href,omitempty"`
	Text string `json:"value,omitempty"`
}

func main() {
	var domain string
	flag.StringVar(&domain, "url", "http://google.co.uk", "Makes it possible to pass the url the sitemap is to be build for.")
	flag.Parse()
	if len(domain) == 0 {
		fmt.Println("Url cannot be empty. Use '-url' flag to set the url at runtime.")
		return
	}

	reader := bytes.NewReader([]byte(domain))
	root, err := html.Parse(reader)
	if err != nil {
		panic(err)
	}

	var links []Link
	var f func(*html.Node)
	f = func(n *html.Node) {
		log.Printf("n data: %v", n.Data)
		if n.Type == html.ElementNode {
			if n.Data == "a" {
				link := new(Link)
				link.Text = n.FirstChild.Data
				for _, a := range n.Attr {
					link.Href = a.Val
				}
				links = append(links, *link)
				// TODO
			} else {
				f(n)
			}
			spew.Dump(links)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(root)
}
