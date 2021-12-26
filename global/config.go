package global

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Conf *Config

type Config struct {
	AppID  uint64 `yaml:"app_id"`
	AppKey string `yaml:"app_key"`
	Token  string `yaml:"token"`
	Debug  bool   `yaml:"debug"`
}

// ConfigGener ...
func ConfigGener() error {
	conf := &Config{
		AppID:  1234567890,
		AppKey: "此处填上机器人的签名的密钥",
		Token:  "此处填上机器人的Token",
		Debug:  false,
	}

	data, err := yaml.Marshal(conf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("config.yml", data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ConfigDealer ...
func ConfigDealer() (*Config, error) {
	data := &Config{}

	content, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Error("无法读取 config.yml，请检查填写内容")
		return data, err
	}

	err = yaml.Unmarshal(content, data)
	if err != nil {
		log.Error("解析 config.yml 失败，请检查填写内容")
		return data, err
	}

	Conf = data

	return data, nil
}
