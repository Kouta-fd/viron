package config

import (
	"fmt"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/cam-inc/viron/example/golang/pkg/constant"
	pkgConstant "github.com/cam-inc/viron/packages/golang/constant"
	"github.com/go-sql-driver/mysql"
)

type (
	Mode   string
	Config struct {
		StoreMode  Mode
		StoreMySQL *MySQL
		StoreMongo *Mongo
		Cors       *Cors
		Auth       *Auth
		Oas        *Oas
	}

	Store struct {
		Mode string
		*MySQL
		*Mongo
		options.ClientOptions
	}

	Mongo struct {
		URI                  string
		User                 string
		Password             string
		VironDB              string
		CasbinCollectionName string
	}

	MySQL struct {
		Dialect              string
		User                 string `yaml:"user"`
		Password             string `yaml:"password"`
		Net                  string `yaml:"net"`
		Host                 string `yaml:"host"`
		Port                 int    `yaml:"port"`
		DBName               string `yaml:"dbname"`
		TLSConfig            string `yaml:"tls"`
		AllowNativePasswords bool   `yaml:"native_password"`
		ParseTime            bool   `yaml:"parse_time"`
	}

	Cors struct {
		/*
			allowOrigins: ['https://localhost:8000'],
		*/
		AllowOrigins []string `yaml:"allowOrigins"`
	}
	JWT struct {
		Secret        string `yaml:"secret"`
		Provider      string `yaml:"provider"`
		ExpirationSec int    `yaml:"expirationSec"`
	}
	GoogleOAuth2 struct {
		ClientID          string
		ClientSecret      string
		AdditionalScope   []string `yaml:"additionalScopes"`
		UserHostedDomains []string `yaml:"userHostedDomains"`
	}
	Auth struct {
		JWT          *JWT
		GoogleOAuth2 *GoogleOAuth2
		/**
		  jwt: {
		     secret: 'XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX',
		     provider: 'viron-example-nodejs',
		     expirationSec: 24 * 60 * 60,
		   },
		   googleOAuth2: {
		     clientId: process.env.GOOGLE_OAUTH2_CLIENT_ID ?? '',

		     clientSecret: process.env.GOOGLE_OAUTH2_CLIENT_SECRET ?? '',
		     additionalScopes: [],
		     userHostedDomains: ['cam-inc.co.jp', 'cyberagent.co.jp'],
		   },
		*/

	}

	Oas struct {
		InfoExtensions map[string]interface{} `json:"infoExtensions"`
		/*
			infoExtentions: {
			        [OAS_X_THEME]: THEME.STANDARD,
			        [OAS_X_THUMBNAIL]: 'https://example.com/logo.png',
			        [OAS_X_TAGS]: ['example', 'nodejs'],
			      },
		*/

	}
)

const (
	StoreModeMongo Mode = "mongo"
	StoreModeMySQL Mode = "mysql"
)

func (m *MySQL) ToDriverConfig() *mysql.Config {
	return &mysql.Config{
		Addr:                 fmt.Sprintf("%s:%d", m.Host, m.Port),
		DBName:               m.DBName,
		User:                 m.User,
		Passwd:               m.Password,
		Net:                  m.Net,
		ParseTime:            m.ParseTime,
		TLSConfig:            m.TLSConfig,
		AllowNativePasswords: m.AllowNativePasswords,
	}
}

func New() *Config {

	mysqlPort, _ := strconv.Atoi(os.Getenv(constant.MYSQL_PORT))
	mode := StoreModeMongo
	if os.Getenv(pkgConstant.ENV_STORE_MODE) == string(StoreModeMySQL) {
		mode = StoreModeMySQL
	}
	// TODO: yaml -> statik で環境別設定
	return &Config{
		Auth: &Auth{
			JWT: &JWT{
				Secret:        "xxxxxxxxxxxxxxxxxxxx",
				Provider:      "viron_example",
				ExpirationSec: 24 * 60 * 60,
			},
		},
		Cors: &Cors{
			AllowOrigins: []string{"https://localhost:8000"},
		},
		StoreMode: mode,
		StoreMySQL: &MySQL{
			Dialect:              "mysql",
			Host:                 os.Getenv(constant.MYSQL_HOST),
			Port:                 mysqlPort,
			Net:                  "tcp",
			User:                 os.Getenv(constant.MYSQL_USER),
			Password:             os.Getenv(constant.MYSQL_PASSWORD),
			DBName:               os.Getenv(constant.MYSQL_DATABASE),
			AllowNativePasswords: true,
			ParseTime:            true,
		},
		StoreMongo: &Mongo{
			URI:                  "mongodb://mongo:27017",
			User:                 "root",
			Password:             "password",
			VironDB:              "viron_example",
			CasbinCollectionName: "casbin_rule_go",
		},
		Oas: &Oas{
			InfoExtensions: map[string]interface{}{
				pkgConstant.OAS_X_THEME:     pkgConstant.THEME_STANDARD,
				pkgConstant.OAS_X_THUMBNAIL: "https://example.com/logo.png",
				pkgConstant.OAS_X_TAGS:      []string{"example", "golang"},
			},
		},
	}
}
