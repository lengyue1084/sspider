package main

/*const (
	MaxGoroutineNum = 500

)*/
var MaxGoroutineNum = 8
var Scheme = "http"
var SaveFolder      = "/images/"
var StartUrl = "https://www.58pic.com/collect/fav-1.html"

//var StartUrlPre = "https://www.58pic.com/collect/fav-"
//var StartUrlend =".html"

var StartUrlPre = "http://www.kongjie.com/home.php?mod=space&do=album&view=all&order=hot&page="
var StartUrlend =""
//`http://www.kongjie.com/home.php?mod=space&do=album&view=all&order=hot&page=1`
var CurentPageNum = 1
var MaxPageNum = 999
var pageUrlChan = make(chan string, 50)
var headers = map[string][]string{
	"Accept":                    []string{"text/html,application/xhtml+xml,application/xml", "q=0.9,image/webp,*/*;q=0.8"},
	"Accept-Encoding":           []string{"gzip, deflate, sdch"},
	"Accept-Language":           []string{"zh-CN,zh;q=0.8,en;q=0.6,zh-TW;q=0.4"},
	"Accept-Charset":            []string{"utf-8"},
	"Connection":                []string{"keep-alive"},
	"DNT":                       []string{"1"},
	"Host":                      []string{"www.58pic.com"},
	"Referer":                   []string{"https://www.58pic.com/collect/fav-1.html"},
	"Upgrade-Insecure-Requests": []string{"1"},
	"User-Agent":                []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36"},
}
