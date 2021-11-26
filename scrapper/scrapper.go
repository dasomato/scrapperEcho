package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type searchResults struct {
	id          string
	title       string
	company     string
	salary      string
	description string
}

//Scrape indeed by term
func Scrape(term string) {
	getPages(term)

}

func getPages(term string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + term

	pageC := make(chan []searchResults)
	extractedJobResults := []searchResults{}

	// writeC := make(chan bool)
	returnPages := 0
	fmt.Println("baseURL : ", baseURL)
	res, err := http.Get(baseURL)
	checkError(err)
	checkStatus(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		returnPages = s.Find("a").Length()
	})

	for i := 0; i < returnPages; i++ {
		go getPage(i, baseURL, pageC)
	}

	for i := 0; i < returnPages; i++ {
		jobs := <-pageC
		extractedJobResults = append(extractedJobResults, jobs...)
	}
	writeJobs(extractedJobResults)
}

func getPage(pageNumber int, baseURL string, pageC chan<- []searchResults) {
	c := make(chan searchResults)
	pageURL := baseURL + "&start=" + strconv.Itoa(pageNumber*50)
	fmt.Println("pageURL : ", pageURL)
	res, err := http.Get(pageURL)
	checkError(err)
	checkStatus(res)

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkError(err)

	extractedResults := []searchResults{}
	doc.Find(".mosaic-provider-jobcards>a").Each(func(i int, card *goquery.Selection) {
		go extractedJob(card, c)
	})

	for i := 0; i < doc.Find(".mosaic-provider-jobcards>a").Length(); i++ {
		job := <-c
		extractedResults = append(extractedResults, job)
	}
	pageC <- extractedResults
}

func extractedJob(card *goquery.Selection, c chan<- searchResults) {
	id, _ := card.Attr("data-jk")
	title := card.Find("div.heading4>h2.jobTitle>span").Text()
	company := card.Find("div.heading6").Find("span.companyName").Text()
	salary, _ := card.Find("div.heading6").Attr("aria-label")
	c <- searchResults{id: id, title: title, company: company, salary: salary}
}

func trimString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func writeJobs(jobs []searchResults) {
	// file, err := os.Open("jobs.csv")
	// if err != nil {
	// 	fmt.Println("err:", err)
	// 	file, _ = os.Create("jobs.csv")
	// } else {
	// 	utf8bom := []byte{0xEF, 0xBB, 0xBF}
	// 	file.Write(utf8bom)
	// }
	file, _ := os.Create("jobs.csv")
	w := csv.NewWriter(file)
	defer w.Flush()
	headers := []string{"ID", "TITLE", "COMPANY", "SALARY", "DESCRIPTION"}

	wErr := w.Write(headers)
	// for _, row := range rows {
	// 	fmt.Println("row : ", row)
	// 	wwJobErr := w.Write(row)
	// 	checkError(wwJobErr)
	// }

	checkError(wErr)
	for _, job := range jobs {
		jobSlice := []string{"https://kr.indeed.com/viewjob?jk=" + job.id, job.title, job.company, job.salary, job.description}
		wJobErr := w.Write(jobSlice)
		checkError(wJobErr)
	}
	// writeC <- true
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkStatus(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request faild with status:", res.StatusCode)
	}
}
