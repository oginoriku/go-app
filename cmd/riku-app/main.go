package main

import (
	"fmt"
	"log"

	"hello_gin/controller"
	"hello_gin/model"

	"github.com/olivere/elastic/v7"
)

func main() {
	// ES client生成
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Fatal(err)
	}

	// ESのversion 取得
	version, err := client.ElasticsearchVersion("http://127.0.0.1:9200")
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", version)

	// CreateIndex　indexを作成
	model.CreateIndex(client)

	// Put　indexに追加します
	model.Put(client)

	// サーバーを起動
	router := controller.GetRouter()
	router.Run("127.0.0.1:8888")
}
