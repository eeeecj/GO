/*
 * @Author: eeeecj
 * @Date: 2022-03-25 15:02:13
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-25 18:07:44
 * @Description:
 */

package initialize

import "gorm.io/gorm"

func Gorm() (*gorm.DB, error) {
	return GormMysql()
}

type _gorm struct{}

func (g *_gorm) Config() *gorm.Config {
	return &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}
}
