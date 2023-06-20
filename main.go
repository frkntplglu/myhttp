package main

import (
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

func main() {

	/* sources := []string{"https://adjust.com", "https://google.com", "https://cimri.com"}
	 */
	parallel := flag.Int("parallel", 10, "You can limit number of the goroutines working parallel at the same time")
	flag.Parse()

	sources := flag.Args()

	if len(sources) == 0 {
		log.Println("You should provide sources to make an http request")
		return
	}

	wg := sync.WaitGroup{}

	wg.Add(len(sources))
	waitChan := make(chan struct{}, *parallel)

	for _, source := range sources {
		waitChan <- struct{}{}
		go func(source string) {

			defer func() {
				wg.Done()
				<-waitChan
			}()

			response := makeRequest(source)
			result := getMD5Hash(string(response))
			fmt.Printf("%s %s \n", source, result)

		}(source)
	}

	wg.Wait()

}

func makeRequest(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body
}

func getMD5Hash(value string) string {
	hash := md5.Sum([]byte(value))
	return hex.EncodeToString(hash[:])
}
