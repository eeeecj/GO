/*
 * @Author: eeeecj
 * @Date: 2022-03-24 17:05:09
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-28 21:25:34
 * @Description:
 */
package main

import (
	"bytes"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"project/vue_admin/gin-vue-admin-my/config"
)

func main() {
	// var err error
	// global.GAB_VP, err = core.Viper()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// global.GAB_LOG = core.Zap()
	// zap.ReplaceGlobals(global.GAB_LOG)
	// global.GAB_DB, err = initialize.GormMysql()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	mySigningKey := []byte("AllYourBase")

	//type MapStructure struct {
	//	Foo string `json:"foo"`
	//	jwt.RegisteredClaims
	//}
	//// Create the Claims
	//claims := &MapStructure{
	//	"xiaoming",
	//	jwt.RegisteredClaims{
	//		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	//	},
	//}
	claims := &jwt.MapClaims{
		"Foo": "sss",
		"Age": 12,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)
}

func InitConfigs() {
	g := config.Server{
		Mysql: config.Mysql{
			Host:        "127.0.0.1",
			Port:        "3306",
			Config:      "charset=utf8mb4&parseTime=True&loc=Local",
			DbName:      "gobasic",
			UserName:    "root",
			Password:    "eeeecj",
			MaxIdleConn: 10,
			MaxOpenConn: 100,
			LogMode:     "",
			LogZap:      false,
		},
		Zap: config.Zap{
			Level:         "info",
			Format:        "console",
			Prefix:        "[gin-vue-admin/server]",
			Directory:     "log",
			ShowLine:      true,
			EncodeLevel:   "LowercaseColorLevelEncoder",
			StacktraceKey: "stacktrace",
			LoginConsole:  true,
		},
	}
	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	out, _ := yaml.Marshal(&g)
	v.ReadConfig(bytes.NewBuffer(out))
	v.WriteConfigAs("./config.yaml")
}
