package main


import (
	"fmt"
	"time"
	"net/http"
	"io/ioutil"
	"context"
)

type Result struct {
	r *http.ReadResponse
	err error
}

func process() {
	ctx, cancel := context.WithTimeout(context.Background(), 2 * time.Second)
	defer cancel()
	tr := &http.Transport{}
	client := &http.Client{Timeout: tr}
	c := make(chan Result, 1)
	req, err := http.NewRequest("GET", "http://baidu.com", nil)
	if err != nil {
		fmt.Println("http request failed, err: ", err)
		return
	}

	go func() {
		resp, err := client.Do(req)
		pack := Result{r: resp, err: err}
		c <- pack
	}()
	
	select {
	case <- ctx.Done():
		tr.CancelRequest(req)
		<-c
		fmt.Println("Timeout!")
	case res := <-c:
		defer res.r.Body.Close()
		out, _ := ioutil.ReadAll(res.r.Body)
		fmt.Println("Server Response: %s", out)
		
	}

}