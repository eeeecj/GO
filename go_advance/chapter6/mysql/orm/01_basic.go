/*
 * @Author: eeeecj
 * @Date: 2022-03-19 20:24:51
 * @LastEditors: eeeecj
 * @LastEditTime: 2022-03-20 11:57:22
 * @Description:
 */

package main

import (
	"errors"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	// gorm.Model
	ID   uint
	Name string
}

func (p *User) Tabler() string {
	/**
	 * @name:
	 * @test:
	 * @msg:
	 * @param {*}
	 * @return {*}
	 */
	return "users"
}

func main() {
	dsn := "root:pNFHVsRw@tcp(39.101.135.40:3306)/gobasic?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err)
	}
	// db.Create(&User{Name: "xiaohua"})

	// db.Model(&User{}).Create(map[string]interface{}{
	// 	"Name": "阿克",
	// })

	// 查询第一个，id排序
	user1 := &User{}
	db.First(user1)
	log.Println(user1)

	// 查询一个，不排序
	var user2 User
	db.Take(&user2)
	log.Println(&user2)

	// 查询最后一个
	var user3 User
	result := db.Last(&user3)

	// 判断错误
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Println(result.Error)
	}
	log.Println(&user3)

	// 将结果存储在字典
	r := []map[string]interface{}{}
	db.Model(&User{}).Find(&r, []int{1, 2, 3})
	log.Println(r)

	// 使用Find或First时，要将结果存储至字典需要指定模型
	user4 := map[string]interface{}{}
	user_con := &User{Name: "xiaohua"}
	db.Model(&User{}).Where(user_con).Find(&user4)
	log.Println(user4)

	// 条件查询
	user5 := map[string]interface{}{}
	db.Model(&User{}).Find(&user5, "name=?", "xiaohua")
	log.Println(user5)

	// 多条件字典查询
	user6 := []map[string]interface{}{}
	db.Model(&User{}).Find(&user6, map[string]interface{}{"name": []string{"xiaohua", "阿克"}})
	log.Println(user6)
}
