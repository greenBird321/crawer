package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	// 创建队列
	var requestQuenen []Request
	for _, r := range seeds {
		requestQuenen = append(requestQuenen, r)
	}

	for len(requestQuenen) > 0 {
		r := requestQuenen[0]
		requestQuenen = requestQuenen[1:]
		log.Printf("fetching url %s", r.URL)
		body, err := fetcher.Fetch(r.URL)
		if err != nil {
			log.Printf("Fetcher: err fetching url %s: %v", r.URL, err)
			continue
		}
		// 使用外面传递过来的处理方法来 处理解析后的页面
		parser := r.ParserFunc(body)
		// 把从网页上抓取出的数据加到quenen中
		requestQuenen = append(requestQuenen, parser.Requests...)
		for _, m := range parser.Items {
			log.Printf("Got item %s\n", m)
		}
	}
}
