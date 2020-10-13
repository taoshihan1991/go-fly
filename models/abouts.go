package models

type About struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	TitleCn    string `json:"title_cn"`
	TitleEn    string `json:"title_en"`
	KeywordsCn string `json:"keywords_cn"`
	KeywordsEn string `json:"keywords_en"`
	DescCn     string `json:"desc_cn"`
	DescEn     string `json:"desc_en"`
	CssJs      string `json:"css_js"`
	HtmlCn     string `json:"html_cn"`
	HtmlEn     string `json:"html_en"`
}

func FindAboutByPage(page interface{}) About {
	var a About
	DB.Where("page = ?", page).First(&a)
	return a
}
