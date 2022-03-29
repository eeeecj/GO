/*
 * @Author: eeeecj
 * @Date: 2022-03-28 11:59:21
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-28 15:05:04
 * @Description:
 */
package utils

import (
	"errors"
	"os"
)

/**
 * @Description: 判断路径是否存在
 * @param {string} path
 * @return {*}
 */
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if fi.IsDir() {
		return true, nil
	}
	return false, errors.New("存在同名文件")
}
