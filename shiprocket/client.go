package shiprocket

import (
	"context"
	"fmt"
	"log"

	"github.com/shubham-scalent/testing-ship/shiprocket/apimodel"
)

type DependencyOptions struct {
	ShiprocketService ShiprockertService
}

type ShiprocketClient struct {
	Options DependencyOptions
	Token   string
}

func NewShiprocketClient(Options DependencyOptions) *ShiprocketClient {
	return &ShiprocketClient{Options: Options}
}

func ShiprocketCli(ctx context.Context, config apimodel.Config) (*ShiprocketClient, error) {

	ShiprocketClient, err := initServer()
	if err != nil {
		log.Fatal(err)
		fmt.Println("err", err)
	}

	ShiprocketClient.Token, err = ShiprocketClient.Options.ShiprocketService.GetToken(config)
	if err != nil {
		log.Fatal(err)
		fmt.Println("err", err)
	}

	return ShiprocketClient, nil
}

func initServer() (*ShiprocketClient, error) {
	shiprocketServiceImpl, err := NewShiprocketServiceImpl()
	if err != nil {
		return nil, err
	}

	dependencyOptions := DependencyOptions{
		ShiprocketService: shiprocketServiceImpl,
	}

	shiprocketClient := NewShiprocketClient(dependencyOptions)
	return shiprocketClient, nil
}
