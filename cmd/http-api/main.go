package main

import (
	"log"

	"github.com/reyhanfahlevi/simple-app/app/api/http"
	"github.com/reyhanfahlevi/simple-app/config"
	postservice "github.com/reyhanfahlevi/simple-app/internal/service/post"
	userservice "github.com/reyhanfahlevi/simple-app/internal/service/user"
	bloguc "github.com/reyhanfahlevi/simple-app/internal/usecase/blog"
	"github.com/reyhanfahlevi/simple-app/pkg/safesql"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

func main() {
	err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	cfg := config.Get()

	/* initialize resource like db, elastic, redis, httpclient etc here */

	// db, err := safesql.OpenMasterDB("mysql", "root:root@tcp(localhost:3306)/test_db")
	db, err := safesql.OpenMasterDB("mysql", cfg.Databases.MySQL.Master)

	if err != nil {
		log.Fatal("failed connect to db. ", err)
	}

	/* initialize services */

	postSvc := postservice.New(db)
	userSvc := userservice.New(db)

	/* initialize usecase */

	blogUc := bloguc.New(postSvc, userSvc)

	/* initialize http handler */

	http.Init(blogUc)

	// run server
	http.Run(cfg.App.Port)
}
