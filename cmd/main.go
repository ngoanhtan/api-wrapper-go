package main

import (
	uiza "github.com/ngoanhtan/api-wrapper-go"
	"github.com/ngoanhtan/api-wrapper-go/entity"
	_ "github.com/ngoanhtan/api-wrapper-go/testing"
	"log"
)

func main() {
	params := &uiza.EntityRetrieveParams{ID: uiza.String("")}
	response, _ := entity.Retrieve(params)
	log.Printf("%s\n", response)
}
