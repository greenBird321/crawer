package engine

type Request struct {
	// 需要发起请求的URL
	URL string
	// 使用那种parse
	ParserFunc func([]byte) ParseResult
}

type ParseResult struct {
	// 需要处理的所有的url
	Requests []Request
	// 对业务有用的任何数据
	Items []interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}