package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"

	"code.google.com/p/go.net/html"
)

var books = map[string]string{
	"julius_ceasar": "http://shakespeare.mit.edu/julius_caesar/full.html",
	"romeo_juliet": "http://shakespeare.mit.edu/romeo_juliet/full.html",
	"macbeth": "http://shakespeare.mit.edu/macbeth/full.html",
}

var speechRe = regexp.MustCompile(`^\d+\.\d+\.\d+$`)

func kaboom(pat string, args ...interface{}) {
	panic(fmt.Sprintf(pat, args...))
}

func parse(file string) *html.Node {
	f, err := os.Open(file)
	if err != nil {
		kaboom("Could not open file '%s': %s", file, err)
	}
	defer f.Close()

	doc, err := html.Parse(f)
	if err != nil {
		kaboom("Failed to parse '%s': %s", file, err)
	}
	return doc
}

func downloadSource(file string, url string) {
	var err error
	if _, err = os.Stat(file); err == nil {
		fmt.Fprintf(os.Stderr, "file %s already exists\n", file)
		return
	}

	res, err := http.Get(url)
	if err != nil {
		kaboom("failed to download '%s': %s", url, err)
	}

	var f *os.File
	f, err = os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		kaboom("failed to create file: %s", err)
	}
	defer f.Close()

	io.Copy(f, res.Body)
}

func main() {
	defer recover() // so defer's run :)

	book := os.Args[1]
	file := fmt.Sprintf("%s.html", book)
	url  := books[book]

	downloadSource(file, url)
	doc := parse(file)

	traverse(doc, &bookCtx{ sourceUrl: url })
}

type bookCtx struct {
	sourceUrl string
}

func traverse(n *html.Node, ctx *bookCtx) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, v := range n.Attr {
			if v.Key != "name" {
				continue
			}

			if speechRe.MatchString(v.Val) {
				key := fmt.Sprintf("%s#%s", ctx.sourceUrl, v.Val)
				fmt.Printf("%s\000%s\n", n.FirstChild.Data, key)
				return
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, ctx)
	}
}