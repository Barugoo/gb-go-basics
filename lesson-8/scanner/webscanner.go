package scanner

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

type RequestResult struct {
	ElapsedTime  time.Duration
	RequestedURL *url.URL
	Status       string
	Err          error
}

type WebScanner struct {
	*sync.WaitGroup
	http.Client
	resultCh chan *RequestResult
}

func NewWebScanner() *WebScanner {
	return &WebScanner{
		WaitGroup: &sync.WaitGroup{},
		resultCh:  make(chan *RequestResult),
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (ws *WebScanner) MakeRequestsAndCountTime(urls []*url.URL) []*RequestResult {
	go func(resultCh chan *RequestResult) {
		ws.Wait()
		close(resultCh)
	}(ws.resultCh)

	for _, u := range urls {
		ws.Add(1)
		go ws.makeRequestAndCountTime(u)
	}

	var resultSlice []*RequestResult
	for result := range ws.resultCh {
		resultSlice = append(resultSlice, result)
	}
	return resultSlice
}

func (ws *WebScanner) makeRequestAndCountTime(parsedURL *url.URL) {
	var msg string
	var err error

	start := time.Now()
	defer func() {
		ws.resultCh <- &RequestResult{
			ElapsedTime:  time.Now().Sub(start),
			RequestedURL: parsedURL,
			Status:       msg,
			Err:          err,
		}
		ws.Done()
	}()

	msg, err = ws.makeRequest(parsedURL)
}

func (ws *WebScanner) makeRequest(parsedURL *url.URL) (string, error) {
	resp, err := ws.Get(parsedURL.String())
	if err != nil {
		return "", err
	}
	return resp.Status, nil
}
