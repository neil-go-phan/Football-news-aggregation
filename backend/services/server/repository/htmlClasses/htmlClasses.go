package htmlclassesrepo

import (
	"server/entities"

)

type htmlClassesRepo struct {
	HtmlClasses entities.HtmlClasses
}

func NewHtmlClassesRepo(htmlClassesInput entities.HtmlClasses) *htmlClassesRepo {
	htmlClasses := &htmlClassesRepo{
		HtmlClasses: htmlClassesInput,
	}
	return htmlClasses
}

func (repo *htmlClassesRepo) GetHtmlClasses() entities.HtmlClasses {
	return repo.HtmlClasses
}