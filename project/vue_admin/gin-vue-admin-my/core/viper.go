/*
 * @Author: eeeecj
 * @Date: 2022-03-25 13:45:48
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-26 21:32:48
 * @Description:
 */
package core

import (
	"errors"
	"flag"
	"fmt"
	"project/vue_admin/gin-vue-admin-my/global"
	"project/vue_admin/gin-vue-admin-my/utils"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path ...string) (*viper.Viper, error) {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose a config file")
		flag.Parse()
		if config == "" {
			config = utils.ConfigFile
			fmt.Printf("正在使用配置文件%v\n", config)
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("正在使用传入的参数配置文件%v\n", path[0])
	}
	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yml")
	err := v.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("配置文件不存在")
		}
		return nil, err
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.GAB_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.GAB_CONFIG); err != nil {
		return nil, err
	}
	return v, nil
}
