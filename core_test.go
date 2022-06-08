package zincapi

import (
	"fmt"
	"testing"
)

func TestDriver_Version(t *testing.T) {
	d := New("admin", "admin", "localhost:4080")
	got, err := d.Version()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", got)
}

func TestDriver_Metrics(t *testing.T) {
	d := New("admin", "admin", "localhost:4080")
	got, err := d.Metrics()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", got)
}
