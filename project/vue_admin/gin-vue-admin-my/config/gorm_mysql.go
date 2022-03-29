/*
 * @Author: eeeecj
 * @Date: 2022-03-25 14:21:51
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-25 15:13:59
 * @Description:
 */

package config

import "fmt"

type Mysql struct {
	Host        string `mapstructure:"host" json:"host" yaml:"host"`
	Port        string `mapstructure:"port" json:"port" yaml:"port"`
	Config      string `mapstructure:"config" json:"config" yaml:"config"`
	DbName      string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	UserName    string `mapstructure:"username" json:"username" yaml:"username"`
	Password    string `mapstructure:"password" json:"password" yaml:"password"`
	MaxIdleConn int    `mapstructure:"maxIdleConn" json:"maxIdleConn" yaml:"maxIdleConn"`
	MaxOpenConn int    `mapstructure:"maxOpenConn" json:"maxOpenConn" yaml:"maxOpenConn"`
	LogMode     string `mapstructure:"logMode" json:"logMode" yaml:"logMode"`
	LogZap      bool   `mapstructure:"logZap" json:"logZap" yaml:"logZap"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", m.UserName, m.Password, m.Host, m.Port, m.DbName, m.Config)
}
