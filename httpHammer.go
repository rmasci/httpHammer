package main

import (
        "flag"
        "fmt"
        "io"
        "io/ioutil"
        "net/http"
        "os"
        "path/filepath"
        "regexp"
        "runtime"
        "strings"
        "time"
)

var runAs = filepath.Base(os.Args[0])

func main() {
        runtime.GOMAXPROCS(runtime.NumCPU())
        c := make(chan string)
        ttime := make(chan time.Duration)
        var totalT time.Duration
        var loop, z int
        var cont bool
        m := 1
        var x, i int64
        var url, rx string
        // define flags passed at runtime, and assign them to the variables defined above
        flag.StringVar(&rx, "r", "", "Regular Expression")
        flag.StringVar(&url, "u", "", "URL")
        flag.IntVar(&loop, "l", 1, "Number of loops to run...")
        flag.BoolVar(&cont, "c", false, "Run Continuously")
        flag.Int64Var(&x, "h", 1, "Number of hits")
        flag.Parse()
        if cont {
                loop=2
        }
        if url == "" {
                Usage()
                os.Exit(1)
        } else {
                test1 := strings.Split(url, ":")
                if test1[0] != "http" {
                        url = "http://" + url
                }
        }
        for z = 1; z <= loop; z++ {
                fmt.Printf("Loop: %v\n", m)
                m++
                for i = 1; i <= x; i++ {
                        go testUrl(rx, url, c, ttime)
                }
                for i = 1; i <= x; i++ {
                        fmt.Print(<-c)
                        totalT = (<-ttime + totalT)
                }
                tAveStr := fmt.Sprintf("%vns", totalT.Nanoseconds()/x)
                tAverage, _ := time.ParseDuration(tAveStr)
                fmt.Printf("Total Time: %s\nAverage Time: %v\n\n", totalT, tAverage)
                if cont == true {
                        z = 1
                }
        }
}

func Usage() {
        fmt.Printf("Usage: %v -u <url> -h <hits> (optional) -l <times to run> -r <regexp>\n", runAs)
        flag.PrintDefaults()
}

func testUrl(rx string, url string, c chan<- string, ttime chan<- time.Duration) {
        var n int64
        var search, etag string
        var dur time.Duration
        bod := ""
        tn := time.Now()

        res, err := http.Get(url)
        defer res.Body.Close()
        if err != nil {
                c <- fmt.Sprintf("URL: %s, Error: %s\n", url, err)
                os.Exit(1)
        }
        if rx == "" {
                n, _ = io.Copy(ioutil.Discard, res.Body)
                dur = time.Since(tn)
        } else {
                search = ""
                bt, _ := ioutil.ReadAll(res.Body)
                dur = time.Since(tn)
                n = int64(len(bt))
                bod = string(bt)
                regx := regexp.MustCompile(rx)
                search = regx.FindString(bod)
                if search == "" {
                        search = "\tCould not match string: " + rx
                } else {
                        search = "\tFound: " + search
                }
        }
        //etag = etagrx.FindString(res.Header)
        et, ok := res.Header["Etag"]
        if ok {
                etag = et[0]
                if etag != "" {
                        etag = "\tETag: " + etag
                }
        }

        c <- fmt.Sprintf("%s, downloaded %dk in %s %s %s\n", res.Status, n/1020, dur, etag, search)
        ttime <- dur
}
