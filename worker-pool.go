package main

import (
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func crawli(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker Id: %d\n", wId)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{Status: resp.StatusCode, URL: site.URL}
	}
}

func main() {
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++ {
		go crawli(w, jobs, results)
	}

	urls := []string{
		"https://github.com",
		"https://google.com",
		"https://yandex.ru",
		"https://github.com",
		"https://google.com",
		"https://yandex.ru",
		"https://bing.com",
	}
	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for a := 1; a <= 7; a++ {
		result := <-results
		log.Println(result)
	}

}
