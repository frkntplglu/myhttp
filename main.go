package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/frkntplglu/myhttp/pkg/hash"
	"github.com/frkntplglu/myhttp/pkg/httpclient"
	"github.com/frkntplglu/myhttp/pkg/semaphore"
	"github.com/frkntplglu/myhttp/pkg/urls"
)

func main() {

	parallel := flag.Int("parallel", 10, "You can limit number of the goroutines working parallel at the same time")
	flag.Parse()
	sources := flag.Args()
	if len(sources) == 0 {
		log.Println("You should provide sources to make an http request")
		return
	}

	// First check protocol is exist or not
	sources = urls.ProtocolCheck(sources)
	// Then validate all urls
	var validSources []string

	for _, source := range sources {
		if !urls.IsValidUrl(source) {
			break
		}

		validSources = append(validSources, source)
	}

	sem := semaphore.New(*parallel)

	wg := sync.WaitGroup{}

	for _, source := range validSources {
		wg.Add(1)
		sem.Acquire()
		go func(source string) {
			defer sem.Release()
			defer wg.Done()

			client := httpclient.New(5 * time.Second)
			body, err := client.Get(source)
			if err != nil {
				// This only prints the error to not terminate the whole program so that even if an error occurred tool can continue with other parameters
				log.Printf("%v error for %s", err, source)
			} else {
				// This is for preventing to hash body as an empty string while an error occurred
				result := hash.GetMd5Hash(string(body))
				fmt.Printf("%s %s \n", source, result)
			}

		}(source)
	}

	wg.Wait()

}
