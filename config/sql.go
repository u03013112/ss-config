package config

import (
	"github.com/u03013112/ss-config/sql"
)

// Config : 数据库存储格式
type Config struct {
	sql.BaseModel
	Name        string `json:"name,omitempty"`
	Address     string `json:"address,omitempty"`
	Description string `json:"description,omitempty"`
	Domain      string `json:"domain,omitempty"`
	IP          string `json:"ip,omitempty"`
	Port        string `json:"port,omitempty"`
	Method      string `json:"method,omitempty"`
	Passwd      string `json:"passwd,omitempty"`
	Status      int64  `json:"status,omitempty"` //0代表不通，1代表通，负数代表各种错误，比如dns错误等。
}

// InitDB : 初始化表格，建议在整个数据库初始化之后调用
func InitDB() {
	sql.GetInstance().AutoMigrate(&Config{})
}

func getConfigList() []*Config {
	var configList []*Config
	db := sql.GetInstance().Model(&Config{})
	db.Find(&configList, "status = ? ", 1)
	return configList
}

func getConfigListAll() []*Config {
	var configList []*Config
	db := sql.GetInstance().Model(&Config{})
	db.Find(&configList)
	return configList
}

func updateConfig(c *Config) {
	sql.GetInstance().Model(c).Where("id = ?", c.ID).Update(Config{IP: c.IP, Status: c.Status})
}
