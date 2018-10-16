package goquery_util

import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"strings"
)

func Grap()  {
	doc, err := goquery.NewDocument("https://gocn.vip/explore/category-14")

	if err != nil {
		fmt.Printf("New document fail, err :%v\n", err)
		return
	}

	list := doc.Find("a:contains(GoCN)")
	if list == nil {
		fmt.Printf("list is empty")
		return
	}

	var newUrl string
	list.Each(func(i int, selection *goquery.Selection) {
		if len(newUrl) != 0 {
			return
		}
		val, _ := selection.Attr("href")
		if strings.Contains(val, "gocn.vip/question") == true {
			newUrl = val
		}
	})

	fmt.Printf("newUrl :%s\n", newUrl)

	newDoc, err := goquery.NewDocument(newUrl)
	if err != nil {
		fmt.Printf("open new url fail, err :%v\n", err)
		return
	}

	modClass := newDoc.Find("div.content").Find("li")
	if modClass == nil {
		fmt.Printf("mod class is nil\n")
	}

	modClass.Each(func(i int, selection * goquery.Selection) {
		fmt.Printf("%d. %s\n", i+1, selection.Text())
	})


}
