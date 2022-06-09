package zincapi

import (
	"fmt"
	"testing"
)

func TestUser_List(t *testing.T) {
	d := New("admin", "admin", "localhost:4080")
	userList, err := d.User().List()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", userList)
}
