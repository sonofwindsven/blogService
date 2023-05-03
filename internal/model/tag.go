package model

import "blogService/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t *Tag) TableName() string {
	return "blog_tag"
}

// 仅仅用于文档返回
type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
