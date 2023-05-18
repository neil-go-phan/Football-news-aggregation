package serverhelper

import "github.com/spf13/viper"

type EnvConfig struct {
	ElasticsearchAddress string `mapstructure:"ELASTICSEARCH_ADDRESS"`
	Port                 string `mapstructure:"PORT"`
	CrawlerAddress       string `mapstructure:"CRAWLER_ADDRESS"`
	JsonPath             string `mapstructure:"JSON_PATH"`
	DBSource             string `mapstructure:"DB_SOURCE"`
	MigrationURL         string `mapstructure:"MIGRATION_URL"`
}

func LoadEnv(path string) (env EnvConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&env)
	return
}
