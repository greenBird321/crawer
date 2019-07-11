package main

import (
	"crawler/zhenai/parser"
	"crawler/engine"
)

func main() {
	engine.Run(engine.Request{
		URL: "http://www.zhenai.com/zhenghun",
		// 使用何种方式 解析当前url页面
		ParserFunc: parser.CityListParse,
	})
}
