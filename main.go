package main

import (
	"adjust/service"
	"flag"
	"fmt"
	"os"
	"sync"
)

func main() {
	parallel := flag.Int("parallel", 10, "Parallel set max number of parallel execution. optional - example [-parallel=5] max - 100, default - 10")
	flag.Parse()

	values := flag.Args()
	if len(values) == 0 {
		fmt.Println("Usage: ./adjust [-parallel] urls ...")
		flag.PrintDefaults()
		os.Exit(1)
	}
	if *parallel > 100 {
		*parallel = 100
	}

	var urlsToProc []string
	for _, url := range values {
		urlsToProc = append(urlsToProc, url)
	}

	requestsService := service.New()

	start(*parallel, urlsToProc, requestsService)
}

type urlError struct {
	Err error
	Url string
}

func start(parallel int, urls []string, requestsService *service.Service) {
	workerLimit := make(chan struct{}, parallel)
	errorsCh := make(chan urlError, len(urls))
	wg := &sync.WaitGroup{}
	wg.Add(len(urls))
	for _, url := range urls {
		go processUrl(wg, workerLimit, errorsCh, url, requestsService)
	}
	wg.Wait()

	select {
	case err := <-errorsCh:
		fmt.Printf("the url [%s] can't be proccessed due error [%s] \n", err.Url, err.Err.Error())
	default:
	}
}

func processUrl(
	wg *sync.WaitGroup,
	workerLimit chan struct{},
	errorsCh chan urlError,
	url string,
	service *service.Service,
) {
	workerLimit <- struct{}{}
	defer wg.Done()

	hashStr, err := service.GetHash(url)
	if err != nil {
		errorsCh <- urlError{
			Err: err,
			Url: url,
		}
	} else {
		fmt.Printf("%s %s \n", url, hashStr)
	}

	<-workerLimit
	return
}
