package main

import (
	"context"
	"fmt"
	"os"

	protos "../protos/currency"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()

	cc, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Error("Unable to connect server", "error", err)
		os.Exit(1)
	}
	defer cc.Close()
	c := protos.NewCurrencyClient(cc)
	rr := &protos.RateRequest{
		Base:        protos.Currencies(protos.Currencies_value["BRL"]),
		Destination: protos.Currencies(protos.Currencies_value["PKR"]),
	}
	res, err := c.GetRate(context.Background(), rr)
	if err != nil {
		log.Error("got an error", "error", err)
		os.Exit(1)
	}
	fmt.Printf("conversion rate %v", conversion(res.Rate, 1))
}
func conversion(base float32, des int32) float32 {
	return base * float32(des)
}
