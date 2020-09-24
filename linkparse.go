package linkparse

import (
	"strings"

	"golang.org/x/net/html"
)

func dfsTextContent(node *html.Node) string {
	result := ""
	if node != nil {
		var f func(node *html.Node)
		f = func(node *html.Node) {
			if node.Type == html.TextNode {
				result = result + strings.ReplaceAll(node.Data, "\n", "")
			}

			if node.FirstChild != nil {
				f(node.FirstChild)
			}
			if node.NextSibling != nil {
				f(node.NextSibling)
			}
		}

		f(node)
	}
	return result
}

//Link struct
type Link struct {
	url  string
	text string
}

func dfsLinkNodes(node *html.Node) []Link {
	total := 0
	result := []Link{}
	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					result = append(result, Link{url: a.Val, text: dfsTextContent(node.FirstChild)})
					total++
					break
				}
			}
		}

		if node.FirstChild != nil {
			f(node.FirstChild)
		}
		if node.NextSibling != nil {
			f(node.NextSibling)
		}
	}
	f(node)
	return result
}

// Parse  accepts a html string and finds all the links <a> alongside with its text content
func Parse(htmlString string) []Link {

	parsed, _ := html.Parse(strings.NewReader(htmlString))

	result := dfsLinkNodes(parsed)
	return result

}
