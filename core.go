package zincapi

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/guonaihong/gout"
)

var std *Driver

type Driver struct {
	Username string
	Password string
	Token    string
	Address  string
}

// New 初始化驱动
func New(username, password, address string) *Driver {
	std = &Driver{
		Username: username,
		Password: password,
		Address:  addressShaping(address),
		Token: base64.StdEncoding.
			EncodeToString([]byte(username + ":" + password)),
	}
	return std
}

// Get 获取驱动句柄 句柄未初始化时将 panic
func Get() *Driver {
	if std == nil {
		panic("zinc driver nil")
	}
	return std
}

// Index 索引操作
func (d *Driver) Index() *Index {
	return &Index{}
}

// Version zinc的版本
type Version struct {
	Branch     string `json:"Branch"`
	Build      string `json:"Build"`
	BuildDate  string `json:"BuildDate"`
	CommitHash string `json:"CommitHash"`
	Version    string `json:"Version"`
}

// Version 获取zinc的版本
func (d *Driver) Version() (*Version, error) {
	var resp = new(Version)
	if err := gout.GET(getURI("/version")).SetHeader(getHeader()).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc get version: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// Metrics 获取监控指标
func (d *Driver) Metrics() (string, error) {
	var resp string
	if err := gout.GET(getURI("/metrics")).SetHeader(getHeader()).BindBody(&resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc get version: %v\n", err)
		return "", err
	}
	return resp, nil
}

// User 用户操作
func (d *Driver) User() *User {
	return &User{}
}
