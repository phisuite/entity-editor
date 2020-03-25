package main

import (
	"context"
	"github.com/phiskills/grpc-api.go"
	"github.com/phisuite/data.go"
	"log"
)

type entityServer struct {
	data.UnimplementedEntityWriteAPIServer
}

func (e entityServer) Create(_ context.Context, entity *data.Entity) (*data.Entity, error) {
	log.Printf("Create: %v", entity)
	return entity, nil
}

func (e entityServer) Update(_ context.Context, entity *data.Entity) (*data.Entity, error) {
	log.Printf("Update: %v", entity)
	return entity, nil
}

func main() {
	api := grpc.New()
	data.RegisterEntityWriteAPIServer(api.Server, &entityServer{})
	api.Start()
}
