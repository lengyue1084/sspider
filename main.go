package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
func main() {
	// 从标准输入流中接收输入数据
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("###########################################:\n")
	fmt.Printf("胖达图片采集器v0.1\n")
	fmt.Printf("该版本支持采集列表规则，页码{{n}}为变量\n如：《https://www.58pic.com/collect/fav-{{n}}.html》\n其中n为页码，起始页码一般默认为1\n")
	fmt.Printf("###########################################:\n")
	fmt.Printf("命令列表：\n1、设置采集前缀 如：https://www.58pic.com/collect/fav-\n")
	fmt.Printf("2、设置采集后缀 如：.html,根据实际情况设置,默认为空\n")
	fmt.Printf("3、设置起始页码 如：1,默认值为1\n")
	fmt.Printf("4、设置最大页码 如：999,根据需要采集的页面列表采集设置默认为999\n")
	fmt.Printf("5、设置最大线程数默认8 \n")
	fmt.Printf("6、开始采集 \n")
	fmt.Printf("8、退出程序 \n")


	fmt.Printf("请输入命令：\n")
	// 逐行扫描
	for input.Scan() {
		cmd := input.Text()
		// 输入bye时 结束

		switch cmd {
		case "1":
			fmt.Printf("请输入采集前缀：")
			for input.Scan(){
				s := input.Text()
				StartUrlPre = s
				fmt.Println("您输入的值为采集前缀为：",s)
				fmt.Printf("请输入命令：\n")
				break;
			}
		case "2":
			fmt.Printf("请输入采集后缀：")
			for input.Scan(){
				s := input.Text()
				StartUrlend = s
				fmt.Println("您输入的值为采集后缀为：",s)
				fmt.Printf("请输入命令：\n")
				break;
			}
		case "3":
			fmt.Printf("设置起始页码：")
			for input.Scan(){
				s := input.Text()
				CurentPageNum, _ = strconv.Atoi(s)
				fmt.Println("您输入起始页码默认为：",s)
				fmt.Printf("请输入命令：\n")
				break;
			}
		case "4":
			fmt.Printf("设置最大页码：")
			for input.Scan(){
				s := input.Text()
				MaxPageNum, _ = strconv.Atoi(s)
				fmt.Println("您输入的最大页码为：",s)
				fmt.Printf("请输入命令：\n")
				break;
			}
		case "5":
			fmt.Printf("设置最大线程数：")
			for input.Scan(){
				s := input.Text()
				MaxPageNum, _ = strconv.Atoi(s)
				fmt.Println("您输入的最大线程数为：",s)
				fmt.Printf("请输入命令：\n")
				break;
			}
		case "6":
			goto start
		case "8":
			fmt.Printf("程序即将退出~\n")
			CountTime(5)
			os.Exit(0)
		default:
			fmt.Printf("请输入命令前的序号比如 1~\n")
		}
	}

	//fmt.Printf("Please type in something:\n")

	start:
	//判断存储文件夹是否存在
	if err := MkDirForImages(); err != nil {
		fmt.Println(err.Error())
	}
	wg.Add(MaxGoroutineNum)
	for i := 0; i < MaxGoroutineNum; i++ {
		go getAndSaveImages(i)
	}
	scheme := StartUrlPre[0:5]
	if scheme != Scheme{
		Scheme = "http"
	}

	fmt.Println(Scheme)
	setNextPageUrl()
	wg.Wait()
}
