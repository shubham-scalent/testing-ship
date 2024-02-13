package shiprocket

import (
	"context"
	"fmt"
	"log"

	"github.com/shubham-scalent/testing-ship/apimodel"
	"github.com/shubham-scalent/testing-ship/pkg"
)

type DependencyOptions struct {
	ShiprocketService pkg.ShiprockertService
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

	ShiprocketClient.Token, err = ShiprocketClient.Options.ShiprocketService.GetToken(config.BaseURL, config.Email, config.Password)
	if err != nil {
		log.Fatal(err)
		fmt.Println("err", err)
	}

	return ShiprocketClient, nil
}

func initServer() (*ShiprocketClient, error) {
	shiprocketServiceImpl, err := pkg.NewShiprocketServiceImpl()
	if err != nil {
		return nil, err
	}

	dependencyOptions := DependencyOptions{
		ShiprocketService: shiprocketServiceImpl,
	}

	shiprocketClient := NewShiprocketClient(dependencyOptions)
	return shiprocketClient, nil
}
