package main

import (
	"fmt"
	"todo_app/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)
	*/

	fmt.Println(models.Db)
}
