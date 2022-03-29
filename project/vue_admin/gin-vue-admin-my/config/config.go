/*
 * @Author: eeeecj
 * @Date: 2022-03-25 14:19:44
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-28 12:40:18
 * @Description:
 */

package config

type Server struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Zap   Zap   `mapstructure:"zap" json:"zap" yaml:"zap"`
}
