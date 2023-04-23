package serverhelper

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func FormatVietnamese(stringInput string) string {
	stringInput = strings.ToLower(stringInput)
	var Regexp_A = `à|á|ạ|ã|ả|ă|ắ|ằ|ẳ|ẵ|ặ|â|ấ|ầ|ẩ|ẫ|ậ`
	var Regexp_E = `è|ẻ|ẽ|é|ẹ|ê|ề|ể|ễ|ế|ệ`
	var Regexp_I = `ì|ỉ|ĩ|í|ị`
	var Regexp_U = `ù|ủ|ũ|ú|ụ|ư|ừ|ử|ữ|ứ|ự`
	var Regexp_Y = `ỳ|ỷ|ỹ|ý|ỵ`
	var Regexp_O = `ò|ỏ|õ|ó|ọ|ô|ồ|ổ|ỗ|ố|ộ|ơ|ờ|ở|ỡ|ớ|ợ`
	var Regexp_D = `Đ|đ`
	reg_a := regexp.MustCompile(Regexp_A)
	reg_e := regexp.MustCompile(Regexp_E)
	reg_i := regexp.MustCompile(Regexp_I)
	reg_o := regexp.MustCompile(Regexp_O)
	reg_u := regexp.MustCompile(Regexp_U)
	reg_y := regexp.MustCompile(Regexp_Y)
	reg_d := regexp.MustCompile(Regexp_D)
	stringInput = reg_a.ReplaceAllLiteralString(stringInput, "a")
	stringInput = reg_e.ReplaceAllLiteralString(stringInput, "e")
	stringInput = reg_i.ReplaceAllLiteralString(stringInput, "i")
	stringInput = reg_o.ReplaceAllLiteralString(stringInput, "o")
	stringInput = reg_u.ReplaceAllLiteralString(stringInput, "u")
	stringInput = reg_y.ReplaceAllLiteralString(stringInput, "y")
	stringInput = reg_d.ReplaceAllLiteralString(stringInput, "d")

	
	return stringInput
}

func FormatElasticSearchIndexName(indexName string) string {
	formatedPhase1 := FormatVietnamese(indexName)
	// regexp remove charaters in ()
	var RegexpPara = `\(.*\)`
	reg_para := regexp.MustCompile(RegexpPara)
	formatedPhase1 = reg_para.ReplaceAllLiteralString(formatedPhase1, "")
	return strings.Replace(formatedPhase1, " ", "", -1)
}

func FortmatTagsFromRequest(tags string) []string {
	trim := strings.TrimSpace(tags)
	tagsSlice := strings.Split(trim, ",")
	tagsFormated := make([]string, 0)
	for _, tag := range tagsSlice {
		if strings.TrimSpace(tag) != "" {
			tagsFormated = append(tagsFormated, tag)
		}
	}
	return tagsFormated
}

func FormatDateToVietnamesDateSting(date time.Time) string {
	year, month, day := date.Date()
	return fmt.Sprintf("%v/%v/%v", day, int(month), year)
}

func FormatCacheKey(tags string) string {
	pharse1 := FormatVietnamese(tags)
	pharse2 := strings.Replace(pharse1, "+", "_", -1)
	pharse3 := strings.Replace(pharse2, "-", "_", -1)
	pharse4 := strings.TrimSpace(pharse3)
	pharse5 := strings.Replace(pharse4, " ", "_", -1)
	return pharse5
}
