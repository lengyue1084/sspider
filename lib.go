package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func MkDirForImages() error {
	rootpath, _ := os.Getwd()
	SaveFolder = rootpath + SaveFolder + time.Now().Format("2006-01-02") + "/"
	f, err := os.Open(SaveFolder)
	defer f.Close()
	if err != nil && os.IsExist(err) {
		return nil
	}
	if err := os.MkdirAll(SaveFolder, 0777); err != nil {
		fmt.Println("创建文件夹失败：", err)
		return err
	}
	os.Chmod(SaveFolder, 0777)
	fmt.Println("文件夹创建成功：", SaveFolder)
	return nil
}

func getAndSaveImages(i int) {
	//out := <- pageUrlChan
	//log.Println(out)
	for url := range pageUrlChan {
		fmt.Println("携程开启：", i)
		fmt.Println("携程开启爬取的url：", url)
		getPageUrlV2(url)
		//out := <- pageUrlChan
		//log.Println(out)
	}
	wg.Done()
}
func getPageUrlV2(url string) {
	// Request the HTML page.
	res, err := http.Get(url)
	if err != nil {
		log.Println("抓取分页请求失败：",err)
		return
		//log.Fatal("请求失败：",err)
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Println("读取分页html失败",err)
		}
		// Find the review items
		doc.Find("img").Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			imageSrc, _ := s.Attr("src")
			if imageSrc != "" {
				str := ""
				imageSrc = strings.Replace(imageSrc, " ", "", -1)
				if imageSrc[0:4] != "http" && imageSrc[0:5] != "https"  {
					if imageSrc[0:2] != "//" && imageSrc[0:3]!= "://"{
						str = Scheme + "://"
					}else if imageSrc[0:2] == "//" {
						str = Scheme + ":"
					}else if imageSrc[0:2] == "://" {
						str = Scheme
					}
				}
				imageSrc = str + imageSrc
				fmt.Printf("获取图片地址： %d: %s\n", i, imageSrc)
				saveImage(imageSrc)
			}
		})
	}
	return
	log.Println("status code error: %d %s", res.StatusCode, res.Status)
	//log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	// Load the HTML document

}

func saveImage(imageUrl string) {
	res := getReponseWithGlobalHeaders(imageUrl)
	if err := recover(); res == nil {
		log.Println("Skip panic2",  err)
		return
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 获取图片扩展名
	fileNameExt := path.Ext(imageUrl)

	if fileNameExt != ".png" || fileNameExt != ".jpg" || fileNameExt != ".bmp" || fileNameExt != ".gif" {
		fileNameExt = ".jpg"
	}
	// 图片保存的全路径
	savePath := path.Join(SaveFolder, strconv.Itoa(rand.Int())+fileNameExt)
	log.Println("savePath", savePath)
	imageWriter, _ := os.OpenFile(savePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	length, _ := io.Copy(imageWriter, res.Body)
	fmt.Println(savePath + " image saved! " + strconv.Itoa(int(length)) + " bytes." + imageUrl)
}

func getReponseWithGlobalHeaders(url string) *http.Response {
	req, _ := http.NewRequest("GET", url, nil)
	if headers != nil && len(headers) != 0 {
		for k, v := range headers {
			for _, val := range v {
				req.Header.Add(k, val)
			}
		}
	}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		if err := recover(); err != nil {
			log.Println("Skip panic",  err)
		}
		log.Println(err)
	}
	return res
}
func setNextPageUrl() {
	defer close(pageUrlChan)
	for {
		fmt.Println(CurentPageNum)
		pageUrlChan <- (StartUrlPre + strconv.Itoa(CurentPageNum) + StartUrlend)
		CurentPageNum++
		if CurentPageNum == MaxPageNum {
			break
		}
	}

}

func CountTime(num int) {
	if num > 0 {
		fmt.Printf("倒计时：%d秒\n",num)
		time.Sleep(time.Duration(1) * time.Second)
		CountTime(num - 1)
	}
}

func Dump(s string)  {
	fmt.Println("您输入的值为：",s)
}
