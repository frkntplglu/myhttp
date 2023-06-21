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
)

func main() {

	// Add flag to determine number of parallel process
	parallel := flag.Int("parallel", 10, "You can limit number of the goroutines working parallel at the same time")
	flag.Parse()

	sources := flag.Args()

	if len(sources) == 0 {
		log.Println("You should provide sources to make an http request")
		return
	}

	// Create semaphore to limit the maximum number of parallel process
	sem := semaphore.New(*parallel)

	wg := sync.WaitGroup{}

	for _, source := range sources {
		wg.Add(1)
		sem.Acquire()
		go func(source string) {
			defer sem.Release()
			defer wg.Done()

			client := httpclient.New(5 * time.Second)
			body, err := client.Get(source)
			if err != nil {
				log.Printf("%v error for %s", err, source)
			}

			result := hash.GetMd5Hash(string(body))

			fmt.Printf("%s %s \n", source, result)

		}(source)
	}

	wg.Wait()

}
