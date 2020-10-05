package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"google.golang.org/grpc"
	IPW "ipw/packages/IPWhitelist"
)

func main() {
	fmt.Print("Starting server...")
	grpcServer := grpc.NewServer()
	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:8080",
	)
	router := NewRouter()
	if err = IPW.RegisterWhitelistServiceHandler(context.Background(), router, conn); err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}
	//log.Fatal(http.ListenAndServe(":8080", router))

	log.Fatal(http.ListenAndServe(":8080",  httpGrpcRouter(grpcServer, router)))
}

