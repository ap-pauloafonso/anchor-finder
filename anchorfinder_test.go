package anchorfinder

import (
	"reflect"
	"testing"
)

func TestFinder(t *testing.T) {

	data := []struct {
		name string
		in   string
		want []Link
	}{
		{
			name: "two anchors, with nested text",
			in: `<ul>
			<li><a href="https://store.toolband.com" target="_blank">Official <b>Store</b></a></li>
			<li><a href="/exclusive-products"><b>Exclusive</b> Products</a></li>
			</ul>`,
			want: []Link{{Url: "https://store.toolband.com", Text: "Official Store"}, {Url: "/exclusive-products", Text: "Exclusive Products"}},
		},
		{
			name: "one flat anchor",
			in:   `<li><a href="https://allwithinmyhands.org/" target="_blank">All Within My Hands Foundation</a></li>`,
			want: []Link{{Url: "https://allwithinmyhands.org/", Text: "All Within My Hands Foundation"}},
		},
	}

	for _, v := range data {
		t.Run(v.name, func(t *testing.T) {
			got, _ := Find(v.in)
			if !reflect.DeepEqual(got, v.want) {
				t.Errorf("got: %+v want %+v", got, v.want)
			}
		})
	}

}
