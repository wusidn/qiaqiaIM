package main

import (
	"sync"
	"fmt"
	"github.com/wusidn/qiaqia/web"
	"github.com/wusidn/qiaqia/chat"
	"github.com/wusidn/qiaqia/daomysql"
)

func main() {

	dao, err := daomysql.Default()
	if err != nil{
		panic(fmt.Sprintf("init dao fail [%v]", err.Error()))
	}
	defer dao.Close()

	wg := &sync.WaitGroup{}

	//开启web服务器
	wg.Add(1)
	webService := web.New(dao)
	go webService.Run(wg)

	//开启socket服务器
	wg.Add(1)
	chatService := chat.New(dao)
	go chatService.Run(wg)

	// chatService.Done()

	wg.Wait()
}
