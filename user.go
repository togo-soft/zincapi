package zincapi

import (
	"fmt"
	"os"

	"github.com/guonaihong/gout"
	jsoniter "github.com/json-iterator/go"
)

type User struct{}

type UserResponse struct {
	Message Message `json:"message"`
}

type Message struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	Salt      string `json:"salt"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// CreateOrUpdate 创建或更新用户信息
// https://docs.zincsearch.com/API%20Reference/update-user/
func (u *User) CreateOrUpdate(p *User) (*UserResponse, error) {
	b, err := jsoniter.Marshal(p)
	if err != nil {
		fmt.Fprintf(os.Stderr, "zinc create or update user: %v\n", err)
		return nil, err
	}
	var resp = new(UserResponse)
	if err := gout.PUT(getURI("/api/user")).SetHeader(getHeader()).SetBody(b).BindJSON(resp).Do(); err != nil {
		fmt.Fprintf(os.Stderr, "zinc create or update user: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// Delete 删除用户
// https://docs.zincsearch.com/API%20Reference/delete-user/
func (u *User) Delete(uid string) (*UserResponse, error) {
	var resp = new(UserResponse)
	if err := gout.DELETE(getURI("/api/user/" + uid)).SetHeader(getHeader()).BindJSON(resp).Do(); err != nil {
		fmt.Fprintf(os.Stderr, "zinc create or update user: %v\n", err)
		return nil, err
	}
	return resp, nil
}
