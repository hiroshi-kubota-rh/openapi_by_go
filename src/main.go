package main

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"newview-backend-go/app/app"
	"newview-backend-go/app/config"
	"newview-backend-go/app/handler"
)

func main() {
	log.Fatalf("%+v", serve(context.Background()))
}

func serve(ctx context.Context) error {
	app, err := app.NewApp()
	if err != nil {
		return err
	}
	//config.Portでポート番号8080を取得して，:8080を作る
	addr := ":" + strconv.Itoa(config.Port())
	log.Printf("Serve on http://%s", addr)

	//:8080をlisten handerとしてNewRouterを呼ぶ．NewRouterの引数であるapp変数にはdbとの接続情報などが入っている．
	//hander.NewRouterはapp/handler/routor.goの中へ
	return http.ListenAndServe(addr, handler.NewRouter(app))
}
