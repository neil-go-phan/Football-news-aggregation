package serverhelper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var PATH = "./testJSON/"
var PATH_FAIL = "./testJSON/fail/"
func TestReadAdminJSONSuccess(t *testing.T) {
	adminConfig, err := ReadAdminJSON(PATH)
	assert.Nil(t, err)
	assert.Equal(t, "admin2023", adminConfig.Username)
	assert.Equal(t, "password_encrypted", adminConfig.Password)
}

func TestReadAdminJSONFail(t *testing.T) {
	_, got := ReadAdminJSON(PATH_FAIL)
	want := "file json not found"
	assert.EqualError(t, got, want, "ReadAdminJSONFail is suppose to %s, but got %s", want, got)
}

func TestReadleaguesJSON(t *testing.T) {
	leagues, err := ReadleaguesJSON(PATH)
	assert.Nil(t, err)
	assert.Equal(t, 2, len(leagues.Leagues))
	assert.Equal(t, "Tin tức bóng đá", leagues.Leagues[0].LeagueName)
	assert.Equal(t, true, leagues.Leagues[0].Active)
	assert.Equal(t, false, leagues.Leagues[1].Active)
	assert.Equal(t, "La Liga", leagues.Leagues[1].LeagueName)
}

func TestReadleaguesJSONFail(t *testing.T) {
	_, got := ReadleaguesJSON(PATH_FAIL)
	want := "file json not found"
	assert.EqualError(t, got, want, "ReadleaguesJSONFail is suppose to %s, but got %s", want, got)
}


func TestReadTagsJSON(t *testing.T) {
	tags, err := ReadTagsJSON(PATH)
	assert.Nil(t, err)
	assert.Equal(t, 4, len(tags.Tags))
	assert.Equal(t, "tin tuc bong da", tags.Tags[0])
	assert.Equal(t, "premier league", tags.Tags[1])
	assert.Equal(t, "v-league", tags.Tags[2])
	assert.Equal(t, "ban ket", tags.Tags[3])
}

func TestReadTagsJSONFail(t *testing.T) {
	_, got := ReadTagsJSON(PATH_FAIL)
	want := "file json not found"
	assert.EqualError(t, got, want, "ReadTagsJSONFail is suppose to %s, but got %s", want, got)
}

func TestReadHtmlClassJSON(t *testing.T) {
	htmlClasses, err := ReadHtmlClassJSON(PATH)
	assert.Nil(t, err)
	assert.Equal(t, "SoaBEf", htmlClasses.ArticleClass)
	assert.Equal(t, "GI74Re nDgy9d", htmlClasses.ArticleDescriptionClass)
	assert.Equal(t, "WlydOe", htmlClasses.ArticleLinkClass)
	assert.Equal(t, "mCBkyc ynAwRc MBeuO nDgy9d", htmlClasses.ArticleTitleClass)
}

func TestReadHtmlClassJSONFail(t *testing.T) {
	_, got := ReadHtmlClassJSON(PATH_FAIL)
	want := "file json not found"
	assert.EqualError(t, got, want, "ReadHtmlClassJSONFail is suppose to %s, but got %s", want, got)
}