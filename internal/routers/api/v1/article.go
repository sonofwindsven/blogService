package v1

import (
	"blogService/pkg/app"
	"blogService/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// @Summary 获取文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [get]
func (a *Article) Get(c *gin.Context) {}

// @Summary 获取多个文章
// @Produce json
// @Param name query string false "文章名称" maxlength(100)
// @Param state query int false "状态" Enums(0,1) defalut(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article [get]
func (a *Article) List(c *gin.Context) {
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
}

// @Summary 新增文章
// @Produce json
// @Param name body string false "文章名称" minlength(3) maxlength(100)
// @Param desc body string false "文章描述" minlength(3) maxlength(100)
// @Param cover_image_url body string false "文章封面" minlength(3) maxlength(100)
// @Param content body string false "文章内容"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article [post]
func (a *Article) Create(c *gin.Context) {}

// @Summary 更新文章
// @Produce json
// @Param name body string false "文章名称" minlength(3) maxlength(100)
// @Param desc body string false "文章描述" minlength(3) maxlength(100)
// @Param cover_image_url body string false "文章封面" minlength(3) maxlength(100)
// @Param content body string false "文章内容"
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.Article "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [put]
func (a *Article) Update(c *gin.Context) {}

// @Summary 删除文章
// @Produce json
// @Param id path int true "文章ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/article/{id} [delete]
func (a *Article) Delete(c *gin.Context) {}
