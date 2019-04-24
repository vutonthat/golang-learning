package ipinfo

import (
	"encoding/json"
	"log"
	"net/http"
)

type Adapter interface {
	GetIPInfo() *IPInfo
}

type adapter struct {
	baseURL string
}

func NewAdapter(baseURL string) Adapter {
	return &adapter{baseURL: baseURL}
}

func (a *adapter) GetIPInfo() *IPInfo {
	resp, err := http.Get(a.baseURL + "/json")
	if err != nil {
		log.Printf("Requesting /json: %v", err)
	}
	defer resp.Body.Close()
	var ipInfo IPInfo
	if err := json.NewDecoder(resp.Body).Decode(&ipInfo); err != nil {
		log.Printf("Decoding request body: %v", err)
	}
	return &ipInfo
}
