package serverhelper

import (
	serverhelper "server/helper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadAdminJSON(t *testing.T) {
	adminConfig, err := serverhelper.ReadAdminJSON("./")
	assert.Nil(t, err)
	assert.Equal(t, "admin2023", adminConfig.Username)
	assert.Equal(t, "password_encrypted", adminConfig.Password)
}

func TestReadleaguesJSON(t *testing.T) {
	leagues, err := serverhelper.ReadleaguesJSON("./")
	assert.Nil(t, err)
	assert.Equal(t, 2, len(leagues.Leagues))
	assert.Equal(t, "Tin tức bóng đá", leagues.Leagues[0].LeagueName)
	assert.Equal(t, true, leagues.Leagues[0].Active)
	assert.Equal(t, false, leagues.Leagues[1].Active)
	assert.Equal(t, "La Liga", leagues.Leagues[1].LeagueName)
}

func TestReadTagsJSON(t *testing.T) {
	tags, err := serverhelper.ReadTagsJSON("./")
	assert.Nil(t, err)
	assert.Equal(t, 4, len(tags.Tags))
	assert.Equal(t, "tin tuc bong da", tags.Tags[0])
	assert.Equal(t, "premier league", tags.Tags[1])
	assert.Equal(t, "v-league", tags.Tags[2])
	assert.Equal(t, "ban ket", tags.Tags[3])
}

func TestReadHtmlClassJSON(t *testing.T) {
	htmlClasses, err := serverhelper.ReadHtmlClassJSON("./")
	assert.Nil(t, err)
	assert.Equal(t, "SoaBEf", htmlClasses.ArticleClass)
	assert.Equal(t, "GI74Re nDgy9d", htmlClasses.ArticleDescriptionClass)
	assert.Equal(t, "WlydOe", htmlClasses.ArticleLinkClass)
	assert.Equal(t, "mCBkyc ynAwRc MBeuO nDgy9d", htmlClasses.ArticleTitleClass)
}
