package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	MonitorURLEnv = "MonitorURL"
)

func main() {
	if len(os.Args) == 1 {
		execPlugin()
	} else if len(os.Args) == 2 && os.Args[1] == "config" {
		execConfig()
	} else if len(os.Args) == 2 && os.Args[1] == "autoconf" {
		execAutoConfig()
	}
}

func execPlugin() {
	env := os.Getenv(MonitorURLEnv)
	req, err := http.NewRequest("GET", env, nil)
	if err != nil {
		fmt.Println("code.value 600")
		fmt.Println("response.value 0")
		os.Exit(0)
	}
	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: &http.Transport{
			MaxIdleConns:       10,
			IdleConnTimeout:    30 * time.Second,
			DisableCompression: true,
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
		},
	}
	start := time.Now()
	res, err := client.Do(req)
	dur := time.Since(start) / time.Millisecond
	if err != nil {
		fmt.Println("code.value 599")
		fmt.Printf("response.value %d\n", dur)
		os.Exit(0)
	}
	defer res.Body.Close()
	fmt.Printf("code.value %d\n", res.StatusCode)
	fmt.Printf("response.value %d\n", dur)
	os.Exit(0)
}

func execConfig() {
	fmt.Println("graph_title HTTP Performance")
	fmt.Println("graph_vlabel Response")
	fmt.Println("code.label code")
	fmt.Println("code.warning 300:399")
	fmt.Println("code.critical 400:1000")
	fmt.Println("response.label response time")
	fmt.Println("response.warning 7000:14999")
	fmt.Println("response.critical 15000:10000000")
	fmt.Println("graph_category webserver")
	os.Exit(0)
}

func execAutoConfig() {
	env := os.Getenv(MonitorURLEnv)
	if len(env) == 0 {
		fmt.Printf("no (No env var %s)\n", MonitorURLEnv)
	} else {
		_, err := http.NewRequest("GET", env, nil)
		if err != nil {
			fmt.Printf("no (Url error %s)\n", err.Error())
		} else {
			fmt.Println("yes")
		}
	}
	os.Exit(0)
}
