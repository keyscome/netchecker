// config/config.go
package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// NetworkConfig 定义了各个服务的地址列表，键为服务名称，值为字符串数组
type NetworkConfig map[string][]string

// LoadConfig 从指定文件中读取 YAML 配置，并解析为 NetworkConfig 类型
func LoadConfig(path string) (NetworkConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg NetworkConfig
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
