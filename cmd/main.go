package main

import (
	"log"

	uiza "github.com/uizaio/api-wrapper-go"
	"github.com/uizaio/api-wrapper-go/entity"
	_ "github.com/uizaio/api-wrapper-go/testing"
)

func main() {
	params := &uiza.EntityRetrieveParams{ID: uiza.String("")}
	response, _ := entity.Retrieve(params)
	log.Printf("%s\n", response)
}
