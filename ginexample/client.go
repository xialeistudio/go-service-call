package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100        // 空闲连接数
	t.MaxConnsPerHost = 100     // 每个host的最大连接数
	t.MaxIdleConnsPerHost = 100 // 每个host的空闲连接数
	client := &http.Client{
		Timeout:   time.Second * 10,
		Transport: t,
	}
	// 表单请求
	params := url.Values{
		"username": {"admin"},
		"password": {"123456"},
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/user/login", strings.NewReader(params.Encode()))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// JSON响应
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		panic(err)
	}
	log.Printf("Response: %v", body)
	// JSON请求
	data, err := json.Marshal(map[string]string{
		"username": "admin",
		"password": "123456",
	})
	if err != nil {
		panic(err)
	}
	req, err = http.NewRequest("POST", "http://localhost:8080/user/register", strings.NewReader(string(data)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// JSON响应
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		panic(err)
	}
	log.Printf("Response: %v", body)
}
