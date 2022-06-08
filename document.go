package zincapi

import (
	"fmt"
	"os"

	"github.com/guonaihong/gout"
	jsoniter "github.com/json-iterator/go"
)

type Document struct {
	scope *Index // 文档归属索引
}

type DocumentResponse struct {
	Id          string `json:"id"`
	Index       string `json:"index"`
	Message     string `json:"message"`
	RecordCount int64  `json:"record_count"`
}

// InsertOrUpdate 插入或更新文档
// data 类型是需要插入的结构体
// https://docs.zincsearch.com/api/document/create/
func (d *Document) InsertOrUpdate(data interface{}) (*DocumentResponse, error) {
	b, err := jsoniter.Marshal(data)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc insert or update document: %v\n", err)
		return nil, err
	}
	var resp = new(DocumentResponse)
	if err := gout.PUT(getURI(fmt.Sprintf("/api/%s/_doc", d.scope.Name))).SetHeader(getHeader()).
		SetBody(b).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc insert or update document: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// InsertOrUpdateWithID 基于ID插入或更新文档
// 需要注意的是，当ID为空，将会抛错
// data 类型是需要插入的结构体
// https://docs.zincsearch.com/api/document/update-with-id/
func (d *Document) InsertOrUpdateWithID(id string, data interface{}) (*DocumentResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id field required")
	}
	b, err := jsoniter.Marshal(data)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc insert or update document: %v\n", err)
		return nil, err
	}
	var resp = new(DocumentResponse)
	if err := gout.PUT(getURI(fmt.Sprintf("/api/%s/_doc/%s", d.scope.Name, id))).SetHeader(getHeader()).
		SetBody(b).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc insert or update document: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// Delete 删除文档
// https://docs.zincsearch.com/api/document/delete/
func (d *Document) Delete(id string) (*DocumentResponse, error) {
	if id == "" {
		return nil, fmt.Errorf("id field required")
	}
	var resp = new(DocumentResponse)
	if err := gout.DELETE(getURI(fmt.Sprintf("/api/%s/_doc/%s", d.scope.Name, id))).SetHeader(getHeader()).
		BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc delete document: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// Bulk 批量操作
// https://docs.zincsearch.com/api/document/bulk/
func (d *Document) Bulk(models []BulkModel) (*DocumentResponse, error) {
	var resp = new(DocumentResponse)
	bulkModels := BulkModels(models)
	var request = bulkModels.Marshal()
	if err := gout.POST(getURI("/api/_bulk")).SetHeader(getHeader()).SetBody(request).
		BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc bulk write: %v\n", err)
		return nil, err
	}
	return resp, nil
}
