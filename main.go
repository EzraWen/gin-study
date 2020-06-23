package main

import db "com.ezra/gin-test/database"

func main() {
	defer db.Database.Close()
	router := initRouter()
	router.Run(":8080")
}
