package zincapi

type Search struct {
	scope *Index // 文档归属索引
}

func (s *Search) Search() error {
	return nil
}
