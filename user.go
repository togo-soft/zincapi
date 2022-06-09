package zincapi

import (
	"fmt"
	"os"

	"github.com/guonaihong/gout"
	jsoniter "github.com/json-iterator/go"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	UserID   string `json:"_id"`
	Username string `json:"username"`
	Role     Role   `json:"role"`
	Password string `json:"password"`
}

type Hits struct {
	Total    Total   `json:"total"`
	MaxScore float64 `json:"max_score"`
	Hits     []Hit   `json:"hits"`
}

type Hit struct {
	Index     string  `json:"_index"`
	Type      string  `json:"_type"`
	ID        string  `json:"_id"`
	Score     float64 `json:"_score"`
	Timestamp string  `json:"@timestamp"`
	Source    Source  `json:"_source"`
}

type Source struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Total struct {
	Value int64 `json:"value"`
}

type Shards struct {
	Total      int64 `json:"total"`
	Successful int64 `json:"successful"`
	Skipped    int64 `json:"skipped"`
	Failed     int64 `json:"failed"`
}

type UserResponse struct {
	Message  Message `json:"message"`
	Took     int64   `json:"took"`
	TimedOut bool    `json:"timed_out"`
	Shards   Shards  `json:"_shards"`
	Hits     Hits    `json:"hits"`
	Error    string  `json:"error"`
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
// https://docs.zincsearch.com/api/user/create/
func (u *User) CreateOrUpdate(p *User) (*UserResponse, error) {
	b, err := jsoniter.Marshal(p)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc create or update user: %v\n", err)
		return nil, err
	}
	var resp = new(UserResponse)
	if err := gout.PUT(getURI("/api/user")).SetHeader(getHeader()).SetBody(b).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc create or update user: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// Delete 删除用户
// https://docs.zincsearch.com/api/user/delete/
func (u *User) Delete(uid string) (*UserResponse, error) {
	var resp = new(UserResponse)
	if err := gout.DELETE(getURI("/api/user/" + uid)).SetHeader(getHeader()).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc create or update user: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// List 获取用户列表
// https://docs.zincsearch.com/api/user/list/
func (u *User) List() (*UserResponse, error) {
	var resp = new(UserResponse)
	if err := gout.GET(getURI("/api/user")).SetHeader(getHeader()).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc get user list: %v\n", err)
		return nil, err
	}
	return resp, nil
}
