package main

import (
	"bytes"
	"io"
	"log"

	"github.com/go-swagger/go-swagger/fixtures/bugs/1899/codegen/client"
	"github.com/go-swagger/go-swagger/fixtures/bugs/1899/codegen/client/connect"
)

func main() {
	tr := &client.TransportConfig{
		Host:     "localhost:9999",
		Schemes:  []string{"http"},
		BasePath: "/",
	}
	api := client.NewHTTPClientWithConfig(nil, tr)

	cfg := io.NopCloser(bytes.NewBufferString("abcd"))
	params := connect.NewConfigPutParams().WithInstance("my-instance").WithConfig(cfg)
	resp, err := api.Connect.ConfigPut(params)
	if err != nil {
		log.Printf("err: %v", err)

		return
	}

	log.Printf("resp: %#v", resp)
}
