/**
AUTHOR:  zeng_xiao_yu
GITHUB:  https://github.com/zengxiaolou
EMAIL:   zengevent@gmail.com
TIME:    2020/11/4-22:07
**/

package config

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Instance *Config

type Config struct {
	ENV 		string `yaml:"ENV"` 		//  环境： prod、 dev
	BaseUrl 	string `yaml:"BaseUrl"` 	//	base url
	Port		string `yaml:"Port"`		//  端口
	LogFile		string `yaml:"LogFile"`		//	日志文件
	ShowSql		bool   `yaml:"ShowSql"`		//  是否显示日志
	StaticPath	string `yaml:"StaticPath"`	//	静态文件目录

	MySqlUrl	string `yaml:"MySqlUrl"`	//	数据库连接地址

	Sentry 		string `yaml:"Sentry"`		//  错误日志DNS

	JWTKey		string  `yaml:"JWTKey"`		// 	JWT 密钥

	// github 第三方登录
	Github	struct {
		ClientID 		string `yaml:"ClientID"`		//  客户端ID
		ClientSecret	string `yaml:"ClientSecret"`	//	客户端密钥
	} `yaml:"Github"`

	// QQ登录
	QQLogin struct {
		AppID			string `yaml:"AppID"`			//	客户端ID
		AppSecret		string `yaml:"AppSecret"`		// 	客户端密钥
	} `yaml:"QQLogin"`

	// 七牛云OSS配置
	Uploader struct {
		Enable		string `yaml:"Enable"`
		QiNiuYunOss struct {
			QiNiuAccessKey 		string `yaml:"QiNiuAccessKey"` 		// 七牛云ID
			QiNiuSecretKey		string `yaml:"QiNiuSecretKey"`		// 七牛云密钥
			QiNiuBucketName		string `yaml:"QiNiuBucketName"`		// 	七牛云文件桶
		} `yaml:"QiNiuYunOss"`
	} `yaml:"Uploader"`

	// Sms
	Sms struct {
		SecretID		string `yaml:"SecretID"`		// 腾讯短信ID
		SecretKey		string `yaml:"SecretKey"`		// 腾讯云密钥
		SmsSDKAppID		string `yaml:"SmsSDKAppID"`		// 短信SDK ID
		Sign 			string `yaml:"Sign"`			// 签名
		TemplateID		string `yaml:"TemplateID"`		// 短信模版ID
	} `yaml:"Sms"`

	// smtp
	Smtp struct {
		Host 			string `yaml:"Host"`			// 邮箱服务器
		Port			string `yaml:"Port"`			// 邮箱服务器端口
		Username		string `yaml:"UserName"`		// 邮箱用户名
		Password 		string `yaml:"Password"`		// 邮箱密码
		SSL 			bool   `yaml:"SSL"`				// 是否开启证书
	} `yaml:"Smtp"`
}

func Init(filename string) *Config  {
	Instance = &Config{}
	if yamlFile, err := ioutil.ReadFile(filename); err != nil {
		logrus.Error(err)
	}else if err = yaml.Unmarshal(yamlFile, Instance); err != nil {
		logrus.Error(err)
	}
	return Instance
}