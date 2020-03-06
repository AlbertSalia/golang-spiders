package main

import (
	"golang-spiders/spiders/config"
	"golang-spiders/spiders/dispatch"
	"golang-spiders/spiders/engine"
	"golang-spiders/spiders/lagou/resolver"
	"golang-spiders/spiders/storage"
	"log"
)

func main() {
	itemChan, err := storage.ItemSaver(config.LaGuoIndex)
	if err != nil {
		log.Fatalf("es连接失败：%s", err)
		return
	}
	concurrendEngine := &engine.ConcurrendEngine{
		Dispatch:  &dispatch.QueuedDispatch{},
		WorkCount: 100,
		ItemChan:  itemChan,
	}
	concurrendEngine.Run(engine.Request{
		Url:       config.SeedUrl,
		ParseFunc: resolver.ParseCtiyList,
	})
}
