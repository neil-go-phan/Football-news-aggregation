package crawlerhelpers

import (
	"crawler/entities"
	"fmt"
	"io"
	log "github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func FormatClassName(class string) string {
	var classes string
	hashParts := strings.Split(class, " ")
	for _, s := range hashParts {
		classes = classes + "." + s
	}
	return classes
}

func FormatToSearch(keyword string) string {
	keyword = strings.ToLower(keyword)
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
	keyword = reg_a.ReplaceAllLiteralString(keyword, "a")
	keyword = reg_e.ReplaceAllLiteralString(keyword, "e")
	keyword = reg_i.ReplaceAllLiteralString(keyword, "i")
	keyword = reg_o.ReplaceAllLiteralString(keyword, "o")
	keyword = reg_u.ReplaceAllLiteralString(keyword, "u")
	keyword = reg_y.ReplaceAllLiteralString(keyword, "y")
	keyword = reg_d.ReplaceAllLiteralString(keyword, "d")

	// regexp remove charaters in ()
	var RegexpPara = `\(.*\)`
	reg_para := regexp.MustCompile(RegexpPara)
	keyword = reg_para.ReplaceAllLiteralString(keyword, "")
	
	return strings.Replace(keyword, " ", "+", -1)
}

func FormatDate(date string) string {
	dataPart := strings.Split(date, ",")
	if len(dataPart) > 1 {
		return strings.TrimSpace(dataPart[1]) 
	}
	return strings.TrimSpace(dataPart[0])
}

func ReadHtmlClassScheduleJSON() (entities.HtmlSchedulesClass, error){
	var classes entities.HtmlSchedulesClass
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	env, err := LoadEnv("./")
	if err != nil {
		log.Fatalln("cannot load env: ", err)
	}
	classesJson, err := os.Open(fmt.Sprintf("%shtmlSchedulesClass.json", env.JsonPath))
	if err != nil {
		log.Println(err)
		return classes, err
	}
	defer classesJson.Close()

	classesByte, err := io.ReadAll(classesJson)
	if err != nil {
		log.Println(err)
		return classes, err
	}

	err = json.Unmarshal(classesByte, &classes)
	if err != nil {
		log.Println(err)
		return classes, err
	}
	return classes, nil
}

func ReadXPathClassMatchDetailJSON() (entities.XPathMatchDetail, error){
	var classes entities.XPathMatchDetail
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	env, err := LoadEnv("./")
	if err != nil {
		log.Fatalln("cannot load env: ", err)
	}
	classesJson, err := os.Open(fmt.Sprintf("%sxPathMatchDetail.json", env.JsonPath))
	if err != nil {
		log.Println(err)
		return classes, err
	}
	defer classesJson.Close()

	classesByte, err := io.ReadAll(classesJson)
	if err != nil {
		log.Println(err)
		return classes, err
	}

	err = json.Unmarshal(classesByte, &classes)
	if err != nil {
		log.Println(err)
		return classes, err
	}
	return classes, nil
}

func ReadHtmlArticlesClassJSON() (entities.HtmlArticleClass, error) {
	var classes entities.HtmlArticleClass
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	env, err := LoadEnv("./")
	if err != nil {
		log.Fatalln("cannot load env: ", err)
	}
	classesJson, err := os.Open(fmt.Sprintf("%shtmlArticlesClassesConfig.json", env.JsonPath))
	if err != nil {
		log.Println(err)
		return classes, err
	}
	defer classesJson.Close()

	classesByte, err := io.ReadAll(classesJson)
	if err != nil {
		log.Println(err)
		return classes, err
	}

	err = json.Unmarshal(classesByte, &classes)
	if err != nil {
		log.Println(err)
		return classes, err
	}
	return classes, nil
}