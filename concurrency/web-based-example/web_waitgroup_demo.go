package web_based_example

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{
	"https://educative.io",
	"https://educative.io/teach",
	"https://educative.io/assessments",
	"https://educative.io/projects",
	"https://educative.io/paths",
	"https://educative.io/learning-plans",
	"https://educative.io/learn",
	"https://educative.io/edpresso",
	"https://educative.io/explore",
	"https://educative.io/efer-a-friend",
	"https://google.com",
	"https://twitter.com",
}

func fetchUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, error := http.Get(url)
	if error != nil {
		fmt.Println("some error happened")
	}
}

func homeHandler(writer http.ResponseWriter, r *http.Request) {
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go fetchUrl(url, &wg)
	}
	wg.Wait()
	writer.Write([]byte("All responses received successfully."))
}

func main() {

	http.HandleFunc("/", homeHandler)

	err := http.ListenAndServe(":7079", nil)
	if err != nil {
		panic(err)
	}
}
