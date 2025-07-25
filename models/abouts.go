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
	Page       string `json:"page"`
}

func FindAbouts() []About {
	var a []About
	DB.Select("id,title_cn,page").Find(&a)
	return a
}

func FindAboutByPage(page interface{}) About {
	var a About
	DB.Where("page = ?", page).First(&a)
	return a
}
func FindAboutByPageLanguage(page interface{}, lang string) About {
	var a About
	if lang == "" {
		lang = "cn"
	}
	if lang == "en" {
		DB.Select("css_js,title_en,keywords_en,desc_en,html_en").Where("page = ?", page).First(&a)
	} else {
		DB.Select("css_js,title_cn,keywords_cn,desc_cn,html_cn").Where("page = ?", page).First(&a)
	}
	return a
}
func UpdateAbout(page string, title_cn string, title_en string, keywords_cn string, keywords_en string, desc_cn string, desc_en string, css_js string, html_cn string, html_en string) {
	c := &About{
		TitleCn:    title_cn,
		TitleEn:    title_en,
		KeywordsCn: keywords_cn,
		KeywordsEn: keywords_en,
		DescCn:     desc_cn,
		DescEn:     desc_en,
		CssJs:      css_js,
		HtmlCn:     html_cn,
		HtmlEn:     html_en,
	}
	DB.Model(c).Where("page = ?", page).Update(c)
}
