package model

import (
	"context"
	"fmt"
	"time"

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

// createIndex indexを作成します
func CreateIndex(client *elastic.Client) {
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

	createIndex, err := client.CreateIndex("user").Body(mapping).IncludeTypeName(true).Do(context.Background())
	if err != nil {
		fmt.Println("createIndexに失敗しました")
		panic(err)
	}
	if !createIndex.Acknowledged {
		fmt.Println("Not acknowledged")
	}
}

// Put indexにデータを追加
func Put(client *elastic.Client) {
	user1 := User{Name: "里久", Address: "埼玉県", Age: 23}

	put1, err := client.Index().
		Index("user").
		Type("doc").
		Id("1").
		BodyJson(user1).
		Do(context.Background())
	if err != nil {
		fmt.Println("putに失敗しました")
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)

	user2 := `{"name" : "太郎", "image" : "写真.png"}`

	put2, err := client.Index().
		Index("user").
		Type("doc").
		Id("2").
		BodyString(user2).
		Do(context.Background())
	if err != nil {
		fmt.Println("putに失敗しました")
		panic(err)
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)
}
