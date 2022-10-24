package main

import (
	"fmt"
	"path/filepath"

	"github.com/gocolly/colly"
	"github.com/kennygrant/sanitize"
	"github.com/otiai10/gosseract"
)

func main() {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("src"))

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.OnResponse(func(r *colly.Response) {

		fileName := "image"
		dir := "/home/sametavcii/My-Repos/Python-Selenium"
		ext := filepath.Ext(r.Request.URL.String())
		cleanExt := sanitize.BaseName(ext)
		final_filename := dir + fmt.Sprintf("%s.%s", fileName, cleanExt[1:])
		r.Save(final_filename)
	})

	c.Visit("https://captcha.com/demos/features/captcha-demo.aspx")

	client := gosseract.NewClient()
	defer client.Close()
	client.SetImage("/home/sametavcii/My-Repos/Python-Seleniumimage.phpget-image-c-demoCaptcha-t-de11762af1bb85ff97ac1832f3830c75")
	text, _ := client.Text()
	fmt.Println(text)
}
