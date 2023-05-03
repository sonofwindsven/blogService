package model

import "blogService/pkg/app"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}

// 仅仅用于文档返回
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
