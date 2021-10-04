package main

import "github.com/amacneil/dbmate/pkg/dbmate"

func main() {
	db := dbmate.New("root:root@/todolist?charset=utf8&parseTime=True&loc=Local")
	db.CreateAndMigrate()
}
