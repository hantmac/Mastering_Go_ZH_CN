package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
    "sync"
    "time"
)

var (
    myUrl   string
    delay   int = 5
    w       sync.WaitGroup
)

type myData struct {
    r   *http.Response
    err error
}

func connect(c context.Context) error {
    defer w.Done()
    data := make(chan myData, 1)
    tr := &http.Transport{}
    httpClient := &http.Client{Transport: tr}

	req, _ := http.NewRequest("GET", myUrl, nil)
	
    go func() {
        response, err := httpClient.Do(req)
        if err != nil {
            fmt.Println(err)
            data <- myData{nil, err}
            return
        } else {
            pack := myData{response, err}
            data <- pack
        }
    }()	
    select {
	case <-c.Done():
		tr.CancelRequest(req)
		<-data
		fmt.Println("The request was cancelled!")
		return c.Err()
	case ok := <-data:
		err := ok.err
		resp := ok.r
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		defer resp.Body.Close()

		realHTTPData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error select:", err)
			return err
		}
		fmt.Println("Server Response: %s\n", realHTTPData)
	}
	return nil
}

func main() {
    if len(os.Args) == 1 {
        fmt.Println("Need a URL and a delay!")
        return
    }

    myUrl = os.Args[1]
    if len(os.Args) == 3 {
        t. err := strconv.Atoi(os.Args[2])
        if err != nil {
            fmt.Println(err)
            return
        }
        delay = t
    }

    fmt.Println("Delay:", delay)
    c := context.Background()
    c, cancel := contedxt.WithTimeout(c, time.Duration(delay)*time.Second)
    defer cancel()

    fmt.Printf("Connecting to %s \n", myUrl)
    w.Add(1)
    go connect(c)
    w.Wait()
    fmt.Println("Exiting...")
}