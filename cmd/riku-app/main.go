package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"hello_gin/controller"

	"github.com/olivere/elastic/v7"
)

type User struct {
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Age       int       `json:"age"`
	Image     string    `json:"image,omitempty"`
	CreatedAt time.Time `json:"created,omitempty"`
	Comment   string    `json:"comment,omitempty"`
}

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

	// IndexExists　該当のindexが存在するか確認
	exists, err := client.IndexExists("user").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %t\n", exists)

	mapping := ""

	// もし存在しなかったら新しいindexを作る為の値を用意
	if !exists {
		mapping = `{
			"settings":{
				"number_of_shards":1,
				"number_of_replicas":0
			},
			"mappings":{
				"doc":{
					"properties":{
						"name":{
							"type":"keyword"
						},
						"address":{
							"type":"text",
							"store": true,
							"fielddata": true
						},
					"age":{
						"type":"long"
					}
					}
				}
			}
		}`
	}

	// CreateIndex　indexを作成
	createIndex(client, mapping)

	// サーバーを起動
	router := controller.GetRouter()
	router.Run("127.0.0.1:8888")
}

// createIndex indexを作成します
func createIndex(client *elastic.Client, mapping string) {
	createIndex, err := client.CreateIndex("user").Body(mapping).IncludeTypeName(true).Do(context.Background())
	if err != nil {
		fmt.Println("createIndexに失敗しました")
		panic(err)
	}
	if !createIndex.Acknowledged {
		fmt.Println("Not acknowledged")
	}
}
