package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Realiza la conexion
var dsn = "root:mypassword@tcp(localhost:3306)/gomysql?charset=utf8mb4&parseTime=True&loc=Local"
var Database = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error en la conexion", err)
		panic(err)
	} else {
		fmt.Println("Conexion exitosa")
		return db
	}
}()
