package zincapi

import (
	"fmt"
	"testing"
)

func TestSearch_Search(t *testing.T) {
	d := New("admin", "admin", "localhost:4080")
	filter := QueryParam{
		SearchType: SearchTypeMatchPhrase,
		Query: map[string]interface{}{
			"term": "USA",
		},
		SortFields: nil,
		From:       0,
		MaxResults: 10,
	}
	searchList, err := d.Index().SetIndexName("olympics").Search().Search(filter)
	if err != nil {
		panic(err)
	}
	fmt.Println(searchList)
}
