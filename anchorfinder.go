package anchorfinder

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
	Url  string
	Text string
}

func dfsLinkNodes(node *html.Node) []Link {
	total := 0
	result := []Link{}
	var f func(node *html.Node)
	f = func(node *html.Node) {
		if node.Data == "a" {
			for _, a := range node.Attr {
				if a.Key == "href" {
					result = append(result, Link{Url: a.Val, Text: dfsTextContent(node.FirstChild)})
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

// Find  accepts a html string and finds all the links <a> alongside with its text content
func Find(htmlString string) ([]Link, error) {

	parsed, err := html.Parse(strings.NewReader(htmlString))
	if err != nil {
		return nil, err
	}

	result := dfsLinkNodes(parsed)
	return result, nil

}
