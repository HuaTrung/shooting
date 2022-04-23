// Conn represent a SQL connection.
package database

import (
	"fmt"
	"gorm.io/gorm"
	"sync"

	"gorm.io/driver/postgres"
)

var lock = &sync.Mutex{}
var pgDB *gorm.DB

func GetPgClient() *gorm.DB {
	if pgDB == nil {
		lock.Lock()
		defer lock.Unlock()
		if pgDB == nil {
			dsn := "host=ec2-54-80-122-11.compute-1.amazonaws.com user=ixubrxnkshmjeq password=b0f84a6f88d22519cc169ac58be06e239749b6eae8048bb359ffd2b098881e2d dbname=d17kn0qcpr45d7 port=5432 sslmode=require TimeZone=Asia/Ho_Chi_Minh"
			db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
			if err != nil {
				panic(err)
			}
			pgDB=db
		} else {
			fmt.Println("Single instance already created.")
		}
	} else {
		fmt.Println("Single instance already created.")
	}
	return pgDB
}
