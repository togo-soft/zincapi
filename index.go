package zincapi

import (
	"fmt"
	"os"
	"time"

	"github.com/guonaihong/gout"
	jsoniter "github.com/json-iterator/go"
)

// StorageType 存储类型
type StorageType string

// FieldType 字段类型
type FieldType string

const (
	S3    StorageType = "s3"
	Disk  StorageType = "disk"
	Minio StorageType = "minio"
)

const (
	FieldTypeText    FieldType = "text"
	FieldTypeKeyword FieldType = "keyword"
	FieldTypeTime    FieldType = "time"
)

// IndexField 索引的字段
type IndexField struct {
	Type          FieldType `json:"type"`
	Index         bool      `json:"index"`
	Store         bool      `json:"store"`
	Sortable      bool      `json:"sortable"`
	Aggregatable  bool      `json:"aggregatable"`
	Highlightable bool      `json:"highlightable"`
	TimeFormat    string    `json:"format"`
}

type Property map[string]IndexField

type Mapping struct {
	Properties Property `json:"properties"`
}

type IndexSetting struct{}

// Index 索引
type Index struct {
	Name        string       `json:"name"`         // 索引名
	Storage     StorageType  `json:"storage_type"` // 存储类型
	Mappings    Mapping      `json:"mappings"`     // 字段映射关系
	Settings    IndexSetting `json:"settings"`     // 目前未知字段
	CreateAt    time.Time    `json:"create_at"`    // 创建时间
	UpdateAt    time.Time    `json:"update_at"`    // 更新时间
	DocsCount   int64        `json:"docs_count"`   // 文档数
	StorageSize int64        `json:"storage_size"` // 存储大小
}

type IndexResponse struct {
	IndexName string      `json:"index"`
	Message   string      `json:"message"`
	Storage   StorageType `json:"storage_type"`
}

// SetIndexName 设置索引名
func (i *Index) SetIndexName(name string) *Index {
	i.Name = name
	return i
}

// Create 创建索引
// https://docs.zincsearch.com/api/index/create/
func (i *Index) Create(record *Index) (*IndexResponse, error) {
	doc, err := jsoniter.Marshal(record)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc create index: %v\n", err)
		return nil, err
	}
	var resp = new(IndexResponse)
	if err := gout.PUT(getURI("/api/index")).SetHeader(getHeader()).SetBody(doc).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc create index: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// Delete 删除索引
// https://docs.zincsearch.com/API%20Reference/delete-index/
func (i *Index) Delete(indexName string) (*IndexResponse, error) {
	var resp = new(IndexResponse)
	if err := gout.DELETE(getURI("/api/index/" + indexName)).SetHeader(getHeader()).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc delete index: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// List 罗列索引
// https://docs.zincsearch.com/api/index/list/
func (i *Index) List() ([]*Index, error) {
	var resp []*Index
	if err := gout.GET(getURI("/api/index")).SetHeader(getHeader()).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc get index list: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// GetMappings 获取所有索引的映射
// https://docs.zincsearch.com/api/index/get-mapping/
func (i *Index) GetMappings() (*Index, error) {
	var resp *Index
	if err := gout.GET(getURI(fmt.Sprintf("/api/%s/_mappings", i.Name))).SetHeader(getHeader()).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc get index mappings: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// UpdateMappings 更新索引的映射
// https://docs.zincsearch.com/api/index/update-mapping/
func (i *Index) UpdateMappings(indexName string, index *Index) (*Index, error) {
	var resp *Index
	b, err := jsoniter.Marshal(index)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc create or update index mappings: %v\n", err)
		return nil, err
	}
	if err := gout.PUT(getURI(fmt.Sprintf("/api/%s/_mappings", indexName))).SetHeader(getHeader()).SetBody(b).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc create or update index mapping: %v\n", err)
		return nil, err
	}
	return resp, nil
}

// GetSettings 获取配置
// https://docs.zincsearch.com/api/index/get-settings/
func (i *Index) GetSettings() error {
	var resp = new(Index)
	if err := gout.GET(getURI(fmt.Sprintf("/api/%s/_settings", i.Name))).SetHeader(getHeader()).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc get index settings: %v\n", err)
		return err
	}
	return nil
}

// UpdateSettings 更新索引配置
// https://docs.zincsearch.com/api/index/update-settings/
func (i *Index) UpdateSettings() error {
	var resp = new(Index)
	if err := gout.PUT(getURI(fmt.Sprintf("/api/%s/_settings", i.Name))).SetHeader(getHeader()).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc get index settings: %v\n", err)
		return err
	}
	return nil
}

// Document 返回索引对应的文档操作流
func (i *Index) Document() *Document {
	return &Document{
		scope: i,
	}
}

// Search 返回搜索操作流
func (i *Index) Search() *Search {
	return &Search{
		scope: i,
	}
}
