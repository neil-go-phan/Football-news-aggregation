package helper

import (
	"regexp"
	"strings"
)

func FormatVietnamese(stringInput string) string {
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



	stringInput = strings.ToLower(stringInput)
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
	tagsSlice := strings.Split(tags, ",")
	return tagsSlice
}