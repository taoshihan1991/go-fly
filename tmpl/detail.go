package tmpl

import (
	"github.com/gin-gonic/gin"
	"github.com/taoshihan1991/imaptool/models"
	"html"
	"html/template"
	"net/http"
)

func PageDetail(c *gin.Context) {
	if c.Request.RequestURI == "/favicon.ico" {
		return
	}
	page := c.Param("page")
	lang, _ := c.Get("lang")
	about := models.FindAboutByPageLanguage(page, lang.(string))
	cssJs := html.UnescapeString(about.CssJs)
	title := about.TitleCn
	keywords := about.KeywordsCn
	desc := html.UnescapeString(about.DescCn)
	content := html.UnescapeString(about.HtmlCn)
	if lang == "en" {
		title = about.TitleEn
		keywords = about.KeywordsEn
		desc = html.UnescapeString(about.DescEn)
		content = html.UnescapeString(about.HtmlEn)
	}
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"Lang":     lang,
		"Title":    title,
		"Keywords": keywords,
		"Desc":     desc,
		"Content":  template.HTML(content),
		"CssJs":    template.HTML(cssJs),
	})
}
