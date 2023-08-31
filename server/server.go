package main

import (
	"context"
	"fmt"
	"net"
	controllers "netxdcustomer/controller"

	"github.com/sivasenthil09/netxd_dal/netxd_customer_dal/services"
	"github.com/sivasenthil09/netxd_project_config/config"
	"github.com/sivasenthil09/netxd_project_config/constants"
	pro "github.com/sivasenthil09/netxd_project_proto"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func initDatabase(client *mongo.Client) {
	customerCollection := config.GetCollection(client, "Bank", "profile")
	controllers.CustomerService = services.InitCustomerService(customerCollection, context.Background())
}

func main() {
	mongoclient, err := config.ConnectDatabase()
	defer mongoclient.Disconnect(context.TODO())
	if err != nil {
		panic(err)
	}
	initDatabase(mongoclient)
	lis, err := net.Listen("tcp", constants.Port)
	if err != nil {
		fmt.Printf("Failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	pro.RegisterCustomerServiceServer(s, &controllers.RPCserver{})

	fmt.Println("Server listening on", constants.Port)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Failed to serve: %v", err)
	}
}
