package zincapi

import (
	"fmt"
	"testing"
)

func TestBulkModels_Marshal(t *testing.T) {
	var models []BulkModel
	var insertObj = map[string]string{
		"hello": "world",
		"你好":    "世界",
	}
	models = append(models, NewBulkModelCreate("article", insertObj))
	models = append(models, NewBulkModelCreate("article", insertObj))
	models = append(models, NewBulkModelUpdate("article", "fjdskfjjf_1", insertObj))
	models = append(models, NewBulkModelCreate("article", insertObj))
	models = append(models, NewBulkModelDelete("article", "fjdskfjjf_1"))
	models = append(models, NewBulkModelCreate("article", insertObj))

	var m = BulkModels(models)
	fmt.Println(string(m.Marshal()))
}
