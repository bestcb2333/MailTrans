package config

import (
	"encoding/json"
	"errors"
	"os"
)

// 对程序配置文件进行初始化
func ConfigInit() error {

	// 检测配置文件是否存在，不存在使用默认
	if _, err := os.Stat("config.json"); os.IsNotExist(err) {
		data, err := json.Marshal(&Config)
		if err != nil {
			return err
		}
		return os.WriteFile("config.json", data, 0644)
	} else if err == nil {
		data, err := os.ReadFile("config.json")
		if err != nil {
			return err
		}
		return json.Unmarshal(data, &Config)
	} else {
		return errors.New("未知的错误类型")
	}
}
