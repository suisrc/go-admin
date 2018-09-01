package main

import (
	"net/http"
	"github.com/chenhg5/go-admin/framework/nethttp"
	"github.com/chenhg5/go-admin"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
)

func main() {
	mux := http.NewServeMux()

	ad := goAdmin.Default()

	cfg := config.Config{
		DATABASE: config.Database{
			IP:           "127.0.0.1",
			PORT:         "3306",
			USER:         "root",
			PWD:          "root",
			NAME:         "godmin",
			MAX_IDLE_CON: 50,
			MAX_OPEN_CON: 150,
			DRIVER:       "mysql",
		},
		AUTH_DOMAIN:  "localhost",
		LANGUAGE:     "cn",         // 语言
		ADMIN_PREFIX: "admin_goal", // 前缀
	}

	ad.AddConfig(cfg).AddPlugins(admin.NewAdmin(datamodel.TableFuncConfig)).Use(new(nethttp.Http), mux)

	http.ListenAndServe(":9002", mux)
}
