package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gookit/color"
)

const (
	antiDupPath = "antiDuplicated.txt"
)

var (
	db         string
	resultFile string
)

func init() {
	resultFile = fmt.Sprintf("%s Phone Numbers.txt", time.Now().Format("2006-Jan-02"))
	db = readScrapped(antiDupPath)
}

func phoneGrabber(id chan int) {
	for {
		Id := <-id
		resp, _ := getNumText(Id)
		if resp.Data.PostContact.ContactMobile == "" {
			fmt.Println(fmt.Sprintf("[%d]", Id), color.FgRed.Render("[NO PHONE NUM]"))
		} else {
			if !strings.Contains(db, fmt.Sprintf("\"%s\"", resp.Data.PostContact.ContactMobile)) {
				db += "\"" + resp.Data.PostContact.ContactMobile + "\""
				AppendText(resultFile, resp.Data.PostContact.ContactMobile+"\r\n")
				AppendText(antiDupPath, resp.Data.PostContact.ContactMobile+"\r\n")
				fmt.Println(fmt.Sprintf("[%d]", Id), color.FgGreen.Render(resp.Data.PostContact.ContactMobile))
			} else {
				fmt.Println(fmt.Sprintf("[%d]", Id), color.FgYellow.Render("[DUPLICATED NUM]"))
			}
		}
	}
}

func main() {
	id := make(chan int)
	for i := 0; i != 20; i++ {
		go phoneGrabber(id)
	}
	for _, key := range ReadAllLines("keywords.txt") {
		go scrap(key, id)
	}
	for {
		time.Sleep(time.Second)
	}
}

func scrap(key string, id chan int) {
	hasMore := true
	page := 1
	for hasMore {
		resp, err := searchHaraj(key, page)
		if err != nil {
			panic(err)
		}
		for _, item := range resp.Data.Posts.Items {
			id <- item.ID
		}
		hasMore = resp.Data.Posts.PageInfo.HasNextPage
		page++
	}
}

func getNumText(id int) (numText numText, err error) {
	Map := map[string]interface{}{
		"query": "query postContact($postId: Int!) {\n\t\t\n\t\tpostContact(postId: $postId)\n\t\t{ \n            contactText\n            contactMobile\n             }\n\t\n\t}",
		"variables": map[string]interface{}{
			"postId": id,
		},
	}
	data, err := json.Marshal(Map)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "https://graphql.haraj.com.sa/", strings.NewReader(string(data)))
	if err != nil {
		return numText, err
	}
	req.Header.Set("Host", "graphql.haraj.com.sa")
	req.Header.Set("Content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return numText, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return numText, err
	}
	err = json.Unmarshal(body, &numText)
	return numText, err
}

func searchHaraj(key string, page int) (search search, err error) {
	Map := map[string]interface{}{
		"query": "query($search:String!,$page:Int,$orderByPostId:Boolean) {  posts: search( search:$search, page:$page, orderByPostId:$orderByPostId) {\n\t\titems {\n\t\t\tid status authorUsername title city postDate updateDate hasImage thumbURL authorId bodyTEXT city tags imagesList commentStatus commentCount upRank downRank geoHash\n\t\t}\n\t\tpageInfo {\n\t\t\thasNextPage\n\t\t}\n\t\t} }",
		"variables": map[string]interface{}{
			"search":        key,
			"page":          page,
			"orderByPostId": false,
		},
	}
	data, err := json.Marshal(Map)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest("POST", "https://graphql.haraj.com.sa/", strings.NewReader(string(data)))
	if err != nil {
		return search, err
	}
	req.Header.Set("Host", "graphql.haraj.com.sa")
	req.Header.Set("Content-type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return search, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return search, err
	}
	err = json.Unmarshal(body, &search)
	return search, err
}
