package zincapi

import (
	"fmt"
	"os"

	"github.com/guonaihong/gout"
	jsoniter "github.com/json-iterator/go"
)

// TODO UpdateDocument UpdateDocumentsBulk

type Document struct {
	scope *Index // 文档归属索引
}

type DocumentResponse struct {
	Id      string `json:"id"`
	Index   string `json:"index"`
	Message string `json:"message"`
}

// InsertOrUpdate 插入或更新文档
// 需要注意的是，当ID为空，将会抛错
// data 类型是需要插入的结构体
// https://docs.zincsearch.com/API%20Reference/update-document-with-id/
func (d *Document) InsertOrUpdate(id string, data interface{}) (*DocumentResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id field required")
	}
	b, err := jsoniter.Marshal(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "zinc insert or update document: %v\n", err)
		return nil, err
	}
	var resp = new(DocumentResponse)
	if err := gout.PUT(getURI(fmt.Sprintf("/api/%s/_doc/%s", d.scope.Name, id))).SetHeader(getHeader()).
		SetBody(b).BindJSON(resp).Do(); err != nil {
		fmt.Fprintf(os.Stderr, "zinc insert or update document: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// Delete 删除文档
// https://docs.zincsearch.com/API%20Reference/delete-document/
func (d *Document) Delete(id string) (*DocumentResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id field required")
	}
	var resp = new(DocumentResponse)
	if err := gout.DELETE(getURI(fmt.Sprintf("/api/%s/_doc/%s", d.scope.Name, id))).SetHeader(getHeader()).
		BindJSON(resp).Do(); err != nil {
		fmt.Fprintf(os.Stderr, "zinc delete document: %v\n", err)
		return nil, err
	}
	return resp, nil
}
