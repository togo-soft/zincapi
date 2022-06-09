package zincapi

import (
	"fmt"
	"testing"
)

func TestIndex_List(t *testing.T) {
	d := New("admin", "admin", "localhost:4080")
	indexList, err := d.Index().List()
	if err != nil {
		panic(err)
	}
	for _, index := range indexList {
		fmt.Printf("%+v\n", index)
	}
}
