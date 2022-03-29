/*
 * @Author: eeeecj
 * @Date: 2022-03-28 12:07:41
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-28 12:39:48
 * @Description:
 */
package config

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Directory     string `mapstructure:"directory" json:"directory" yaml:"directory"`
	ShowLine      bool   `mapstructure:"showLine" yaml:"showLine" json:"showLine"`
	EncodeLevel   string `mapstructure:"encodedLevel" json:"encodedLevel" yaml:"encodedLevel"`
	StacktraceKey string `mapstructure:"stacktraceKey" json:"stacktrace" yaml:"stacktrace"`
	LoginConsole  bool   `mapstructure:"loginConsole" json:"loginConsole" yaml:"loginConsole"`
}
