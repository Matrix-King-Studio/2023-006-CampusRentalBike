package config

import (
	model "bike_manage/models"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/yaml.v2"
	"os"
)

var DB *gorm.DB

var Path string
var Port string
var SecretKey string
var AppID string
var AppSecret string
var Expires int

type DBConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}
type JWTConfig struct {
	SecretKey string `yaml:"secretKey"`
	Expires   int    `yaml:"expires"`
}
type GinSettings struct {
	Port string `yaml:"port"`
}
type WeChatSetting struct {
	AppID     string `yaml:"appid"`
	AppSecret string `yaml:"appsecret"`
}

// Config 用于保存所有配置信息
type Config struct {
	Database      DBConfig      `yaml:"database"`
	JWT           JWTConfig     `yaml:"jwt"`
	GinSettings   GinSettings   `yaml:"ginSettings"`
	WeChatSetting WeChatSetting `yaml:"WeChatSetting"`
}

func init() {
	flag.StringVar(&Path, "config", "config/config.yml", "Path to the configuration file")
	flag.Parse()
	var config_, _ = ReadConfig(Path)
	SecretKey = config_.JWT.SecretKey
	Port = config_.GinSettings.Port
	AppID = config_.WeChatSetting.AppID
	AppSecret = config_.WeChatSetting.AppSecret
	Expires = config_.JWT.Expires
}
func InitDB() {
	var err error

	// 从 YAML 文件中读取配置信息
	config, err := ReadConfig(Path)
	if err != nil {
		panic("failed to read config file: " + err.Error())
	}

	// 构建数据库连接字符串
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.DBName)
	// 连接到数据库
	DB, err = gorm.Open("mysql", connStr)
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}
	// 根据gormDB获取sqlDb
	sqlDb := DB.DB()
	// 设置最大连接数
	sqlDb.SetMaxOpenConns(100)

	// 设置最大闲置连接数
	sqlDb.SetMaxIdleConns(20)

	// 设置日志模式，以便在开发过程中查看SQL语句
	//DB.LogMode(true)

	// 自动迁移数据库，根据模型创建或更新表结构
	// 注意：将需要自动迁移的模型添加到此处
	// 例如：DB.AutoMigrate(&models.User{}, &models.Vehicle{}, &models.Staff{}, &models.Customer{})
	DB.AutoMigrate(&model.User{}, &model.Role{}, &model.Permission{}, &model.Location{}, &model.Bike{}, &model.Fence{}, &model.Invoices{}, &model.RegisterUser{}, &model.UserBill{}, &model.BikePrice{}, &model.Feedback{})
	//DB.Exec("ALTER TABLE bikes AUTO_INCREMENT = 100000000")
}

// ReadConfig 从 YAML 文件中读取配置信息
func ReadConfig(filename string) (*Config, error) {
	config := &Config{}

	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
