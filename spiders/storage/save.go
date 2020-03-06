package storage

import (
	"context"
	"errors"
	"golang-spiders/spiders/config"
	"golang-spiders/spiders/lagou/model"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

func ItemSaver(index string) (chan model.Position, error) {
	client, err := elastic.NewClient(elastic.SetURL(config.EsUrl), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan model.Position)
	go func() {
		itemCount := 0
		for {
			position := <-out
			log.Printf("Got item: #%d: %+v\n", itemCount, position)
			itemCount++

			err = es_save(client, index, position)
			if err != nil {
				log.Fatalf("Item Saver: error, saving item %v: %v", position, err)
			}
		}
	}()
	return out, nil
}

func es_save(client *elastic.Client, index string, item model.Position) error {

	if item.Type == "" {
		return errors.New("必须要类型")
	}

	indexService := client.Index().
		Index(index).   //数据库名
		Type(item.Type) //表名
	//Id(""). //编号  也可以由数据生成
	//BodyJson(item). //数据转成json
	//Do(context.Background())
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.BodyJson(item).Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}
