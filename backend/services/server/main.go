package main

import (
	"backend/services/server/entities"
	"backend/services/server/services"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EnvConfig struct {
	ElasticsearchAddress string `mapstructure:"ELASTICSEARCH_ADDRESS"`
	Port                 string `mapstructure:"PORT"`
	CrawlerAddress       string `mapstructure:"CRAWLER_ADDRESS"`
}

func main() {
	env, err := loadEnv(".")
	if err != nil {
		log.Fatal("cannot load env")
	}
	// load default config
	classConfig, keywordsconfig, err := readConfigFromJSON()
	if err != nil {
		log.Fatalln("Fail to read config from JSON: ", err)
	}
	conn := connectToCrawler(env)

	// declare services
	htmlClassesService := services.NewHtmlClassesService(classConfig)
	keywordsService := services.NewKeywordsService(keywordsconfig)
	articleService := services.NewArticleService(keywordsService, htmlClassesService, conn)

	// elastic search
	es, err := connectToElasticsearch(env)
	if err != nil {
		log.Println("error occurred while connecting to elasticsearch node: ", err)
	}
	createElaticsearchIndex(es, keywordsconfig)

	// cronjob
	cronjob := cron.New()
	
	articleService.GetArticlesEveryMinutes(cronjob)
	cronjob.Run()

}

func readConfigFromJSON() (entities.HtmlClasses, entities.Keywords, error) {
	var classConfig entities.HtmlClasses
	var keywordsConfig entities.Keywords

	classConfig, err := services.ReadHtmlClassJSON()
	if err != nil {
		log.Println("Fail to read htmlClassesConfig.json: ", err)
		return classConfig, keywordsConfig, err
	}

	keywordsconfig, err := services.ReadKeywordsJSON()
	if err != nil {
		log.Println("Fail to read keywordsConfig.json: ", err)
		return classConfig, keywordsconfig, err
	}

	return classConfig, keywordsconfig, nil
}

func loadEnv(path string) (env EnvConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}

func connectToCrawler(env EnvConfig) *grpc.ClientConn {
	// dial server
	conn, err := grpc.Dial(env.CrawlerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}
	return conn
}

func connectToElasticsearch(env EnvConfig) (*elasticsearch.Client, error) {
	var esConfig = elasticsearch.Config{
		Addresses: []string{
			env.ElasticsearchAddress,
		},
	}

	es, err := elasticsearch.NewClient(esConfig)
	if err != nil {
		return nil, err
	}
	return es, nil
}

func createElaticsearchIndex(es *elasticsearch.Client, keywords entities.Keywords) {
	var wg sync.WaitGroup
	for _, elasticIndex := range keywords.Keywords {
		wg.Add(1)
		go func(elasticIndex string) {
			indexName := formatIndexName(elasticIndex)
			// check if index exist. if not exist, create new one, if exist, skip it
			exists, err := checkIndexExists(es, indexName)
			if err != nil {
				log.Printf("Error checking if the index %s exists: %s\n", indexName, err)
			}

			if !exists {
				log.Printf("Index: %s is not exist, create a new one...\n", indexName)
				err = createIndex(es, indexName)
				if err != nil {
					log.Fatalf("Error createing index: %s", err)
				}
				wg.Done()
				return 
			}
			log.Printf("Index: %s is already exist, skip it...\n", indexName)
			wg.Done()
		}(elasticIndex)
	}

	wg.Wait()
}

func checkIndexExists(es *elasticsearch.Client, indexName string) (bool, error) {
	res, err := es.Indices.Exists([]string{indexName})
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return true, nil
	}

	return false, nil
}

func createIndex(es *elasticsearch.Client, indexName string) error {
	res, err := es.Indices.Create(indexName)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.IsError() {
		return fmt.Errorf("error creating index %s: %s", indexName, res.Status())
	}

	return nil
}

func formatIndexName(indexName string) string {
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
	indexName = reg_a.ReplaceAllLiteralString(indexName, "a")
	indexName = reg_e.ReplaceAllLiteralString(indexName, "e")
	indexName = reg_i.ReplaceAllLiteralString(indexName, "i")
	indexName = reg_o.ReplaceAllLiteralString(indexName, "o")
	indexName = reg_u.ReplaceAllLiteralString(indexName, "u")
	indexName = reg_y.ReplaceAllLiteralString(indexName, "y")
	indexName = reg_d.ReplaceAllLiteralString(indexName, "d")

	// regexp remove charaters in ()
	var RegexpPara = `\(.*\)`
	reg_para := regexp.MustCompile(RegexpPara)
	indexName = reg_para.ReplaceAllLiteralString(indexName, "")

	indexName = strings.ToLower(indexName)
	return strings.Replace(indexName, " ", "", -1)
}
