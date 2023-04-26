package htmlclassesrepo

import (
	"fmt"
	"server/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHtmlClasses(t *testing.T) {
	assert := assert.New(t)
	want := entities.HtmlClasses{
		ArticleClass: "sample article",
		ArticleTitleClass: "sample title",
		ArticleDescriptionClass: "sample description",
		ArticleLinkClass: "sample link",
	}
	htmlClass := NewHtmlClassesRepo(want)
	got := htmlClass.GetHtmlClasses()

	assert.Equal(want, got, fmt.Sprintf("Method GetHtmlClasses is supose to %v, but got %s", want, got))
}
