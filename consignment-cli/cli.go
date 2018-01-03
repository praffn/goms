package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/micro/go-micro/cmd"

	microclient "github.com/micro/go-micro/client"
	pb "github.com/praffn/goms/consignment-service/proto/consignment"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

func add(file string, client pb.ShippingServiceClient) {
	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}

	cmd.Init()
	client := pb.NewShippingServiceClient("go.micro.srv.consignment", microclient.DefaultClient)

	switch args[0] {
	case "add":
		if len(args) == 1 {
			log.Fatalf("You must pass a json file along to add consignment")
		}
		add(args[1], client)
		return
	case "list":
		getAll, err := client.GetConsignments(context.Background(), &pb.GetRequest{})
		if err != nil {
			log.Fatalf("Could not list consignments: %v", err)
		}
		for _, v := range getAll.Consignments {
			log.Println(v)
		}
	}
}
