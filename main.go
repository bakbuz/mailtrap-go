package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sort"

	"github.com/bakbuz/mailtrap-go/emails/model"
	"github.com/bakbuz/mailtrap-go/emails/request"
)

func main() {
	//emailval()
	//attachmentval()
	//requestlog()

	//mime()
	//mime2()

	DoWork()
}

func DoWork() {
	req := request.Create().
		WithFrom("john.doe@demomailtrap.com", "John Doe").
		WithTo("hero.bill@galaxy.net", "").
		WithSubject("Invitation to Earth").
		WithTextBody("Dear Bill,\n\nIt will be a great pleasure to see you on our blue planet next weekend.\n\nBest regards, John.")

	fmt.Println("request")
	reqjson, _ := json.MarshalIndent(req, "", "  ")
	fmt.Println(string(reqjson))
}

func requestlog() {
	req := request.Create()
	//req.From = model.NewEmailAddress("", "")
	req.From = model.NewEmailAddress("a@a.com", "deli")

	//req.ReplyTo = model.NewEmailAddress("", "")
	req.ReplyTo = model.NewEmailAddress("noreply@a.com", "noreply")

	//req.To = []model.EmailAddress{}
	req.To = []*model.EmailAddress{model.NewEmailAddress("target@a.com", "")}

	//req.Cc = []model.EmailAddress{}
	req.Cc = []*model.EmailAddress{model.NewEmailAddress("cc@a.com", "")}

	//req.Bcc = []model.EmailAddress{}
	req.Bcc = []*model.EmailAddress{model.NewEmailAddress("Bcc@a.com", "")}

	req.Attachments = []*model.Attachment{model.NewAttachment("içerik", "abc.txt", "", "", "")}

	req.Headers = map[string]string{"key": "value", "one": "1"}

	req.CustomVariables = map[string]string{"key": "value", "one": "1"}

	// req.Subject = "konu"
	// req.TextBody = "içerik"
	// req.Category = "kategori"

	req.TemplateId = "welcome"
	req.TemplateVariables = map[string]any{"key1": "val1", "key2": "val2"}

	fmt.Println("request")
	reqjson, _ := json.MarshalIndent(req, "", "  ")
	fmt.Println(string(reqjson))

	fmt.Println("response")

	err := req.Validate()
	fmt.Println(err)
}

func emailval() {

	emp := &model.EmailAddress{Email: "abcdefg"}
	result1 := emp.Validate()
	fmt.Println("empty input: " + result1.Error())

	inv := &model.EmailAddress{Email: "abcdefg"}
	result2 := inv.Validate()
	fmt.Println("invld input: " + result2.Error())

	recipient := &model.EmailAddress{Email: "a@a.com"}
	result3 := recipient.Validate()
	if result3 == nil {
		fmt.Println("valid input OK")
	} else {
		fmt.Println("valid input ERR: " + result3.Error())
	}

}

func attachmentval() {
	//attachment := model.NewAttachment("test content", "test.txt", common.DispositionType_Attachment, common.MimeType_ImagePNG, "s")

	attachment := model.NewAttachment("test content", "test.txt", "", "abc", "")

	// json3, _ := json.MarshalIndent(attachment, "", "  ")
	// fmt.Println(string(json3))

	err := attachment.Validate()
	if err != nil {
		log.Fatal(err)
	}
}

const dbURL = "https://cdn.jsdelivr.net/gh/jshttp/mime-db@master/db.json"

func mime() {
	fmt.Println("Downloading mime-db...")
	resp, err := http.Get(dbURL)
	if err != nil {
		fmt.Println("Download error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Bad status:", resp.Status)
		os.Exit(1)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		os.Exit(1)
	}

	var db map[string]interface{}
	if err := json.Unmarshal(body, &db); err != nil {
		fmt.Println("JSON parse error:", err)
		os.Exit(1)
	}

	minLen := -1
	shortest := []string{}
	for mime := range db {
		l := len(mime)
		if minLen == -1 || l < minLen {
			minLen = l
			shortest = []string{mime}
		} else if l == minLen {
			shortest = append(shortest, mime)
		}
	}

	fmt.Printf("Total MIME types scanned: %d\n", len(db))
	fmt.Printf("Shortest length: %d\n", minLen)
	sort.Strings(shortest)
	fmt.Println("MIME types with that length:")
	for _, m := range shortest {
		fmt.Println(" ", m)
	}
}

func mime2() {
	fmt.Println("Downloading mime-db...")
	resp, err := http.Get(dbURL)
	if err != nil {
		fmt.Println("Download error:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("Bad status:", resp.Status)
		os.Exit(1)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		os.Exit(1)
	}

	var db map[string]interface{}
	if err := json.Unmarshal(body, &db); err != nil {
		fmt.Println("JSON parse error:", err)
		os.Exit(1)
	}

	shortest := []string{}
	for mime := range db {
		shortest = append(shortest, mime)
	}

	// shortest listesini length'e göre artan sıralama
	sort.Slice(shortest, func(i, j int) bool {
		return len(shortest[i]) < len(shortest[j])
	})
	for i, m := range shortest {
		fmt.Println(m)
		if i == 10 {
			break
		}
	}
}
