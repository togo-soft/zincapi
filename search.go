package zincapi

import (
	"fmt"
	"github.com/guonaihong/gout"
	jsoniter "github.com/json-iterator/go"
	"os"
)

type Search struct {
	scope *Index // 文档归属索引
}

type SearchResponse struct {
	Took     int64       `json:"took"`
	TimedOut bool        `json:"timed_out"`
	MaxScore float64     `json:"max_score"`
	Hits     Hits        `json:"hits"`
	Buckets  interface{} `json:"buckets"`
	Error    string      `json:"error"`
}

type SearchType string

const (
	SearchTypeMatch       SearchType = "match"
	SearchTypeMatchAll    SearchType = "matchall"
	SearchTypeMatchPhrase SearchType = "matchphrase"
	SearchTypeTerm        SearchType = "term"
	SearchTypeQueryString SearchType = "querystring"
	SearchTypePrefix      SearchType = "prefix"
	SearchTypeWildcard    SearchType = "wildcard"
	SearchTypeFuzzy       SearchType = "fuzzy"
	SearchTypeDateRange   SearchType = "daterange"
)

type QueryParam struct {
	SearchType SearchType             `json:"search_type"`
	Query      map[string]interface{} `json:"query"`
	SortFields []string               `json:"sort_fields,omitempty"`
	From       int64                  `json:"from"`
	MaxResults int64                  `json:"max_results"`
	Source     []string               `json:"_source,omitempty"`
	Aggregate  map[string]interface{} `json:"aggs,omitempty"`
	Highlight  map[string]interface{} `json:"highlight,omitempty"`
}

func (s *Search) Search(filter QueryParam) (*SearchResponse, error) {
	b, err := jsoniter.Marshal(filter)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc search: %v\n", err)
		return nil, err
	}
	var resp = new(SearchResponse)
	if err := gout.POST(getURI(fmt.Sprintf("/api/%s/_search", s.scope.Name))).Debug(true).SetHeader(getHeader()).SetBody(b).BindJSON(resp).Do(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "zinc search: %v\n", err)
		return nil, err
	}
	return resp, nil
}
