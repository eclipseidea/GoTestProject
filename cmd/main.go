package main

import (
	"go_web_server/cmd/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	app.Run()
}
