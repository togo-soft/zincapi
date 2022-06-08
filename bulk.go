package zincapi

import (
	"bytes"
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

// BulkModel 批量操作接口
type BulkModel interface {
	BulkCreate(indexName string, obj interface{})
	BulkUpdate(indexName string, id string, obj interface{})
	BulkDelete(indexName string, id string)
	Marshal() []byte
}

type BulkModels []BulkModel

func (b *BulkModels) Marshal() []byte {
	var buff bytes.Buffer
	for _, v := range *b {
		fmt.Println(string(v.Marshal()))
		buff.Write(v.Marshal())
	}
	return buff.Bytes()
}

type IndexActionProperties struct {
	ID        string `json:"_id,omitempty"`
	IndexName string `json:"_index"`
}

type IndexAction struct {
	Create *IndexActionProperties `json:"create,omitempty"`
	Update *IndexActionProperties `json:"update,omitempty"`
	Delete *IndexActionProperties `json:"delete,omitempty"`
}

type BulkObject struct {
	IndexAction
	DataLine interface{}
}

func (b *BulkObject) BulkCreate(indexName string, obj interface{}) {
	b.IndexAction.Create = &IndexActionProperties{
		IndexName: indexName,
	}
	b.DataLine = obj
}

func (b *BulkObject) BulkUpdate(indexName string, id string, obj interface{}) {
	b.IndexAction.Update = &IndexActionProperties{
		IndexName: indexName,
		ID:        id,
	}
	b.DataLine = obj
}

func (b *BulkObject) BulkDelete(indexName string, id string) {
	b.IndexAction.Delete = &IndexActionProperties{
		IndexName: indexName,
		ID:        id,
	}
}

func (b *BulkObject) Marshal() []byte {
	var buff bytes.Buffer
	if b.Create != nil {
		bs, _ := jsoniter.Marshal(b.IndexAction)
		buff.Write(bs)
		buff.WriteRune('\n')
		bs, _ = jsoniter.Marshal(b.DataLine)
		buff.Write(bs)
	}
	if b.Update != nil {
		bs, _ := jsoniter.Marshal(b.IndexAction)
		buff.Write(bs)
		buff.WriteRune('\n')
		bs, _ = jsoniter.Marshal(b.DataLine)
		buff.Write(bs)
	}
	if b.Delete != nil {
		b, _ := jsoniter.Marshal(b.IndexAction)
		buff.Write(b)
	}
	return buff.Bytes()
}

// NewBulkModelCreate 创建批量操作的新建模型
func NewBulkModelCreate(indexName string, createObject interface{}) BulkModel {
	var bm BulkModel
	var bb = &BulkObject{}
	bb.BulkCreate(indexName, createObject)
	bm = bb
	return bm
}

// NewBulkModelUpdate 创建批量操作的更新模型
func NewBulkModelUpdate(indexName, id string, createObject interface{}) BulkModel {
	var bm BulkModel
	var bb = &BulkObject{}
	bb.BulkUpdate(indexName, id, createObject)
	bm = bb
	return bm
}

// NewBulkModelDelete 创建批量操作的删除模型
func NewBulkModelDelete(indexName, id string) BulkModel {
	var bm BulkModel
	var bb = &BulkObject{}
	bb.BulkDelete(indexName, id)
	bm = bb
	return bm
}
