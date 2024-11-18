package main

import (
	"context"
	"fmt"
	"log"

	pb "produtores/producer"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("produtor-de-pao:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewProducerClient(conn)

	// Número de requisições para Carga Simples
	numRequests := 10

	for i := 0; i < numRequests; i++ {
		response, err := client.GetResource(context.Background(), &pb.ResourceRequest{Quantity: 10})
		if err != nil {
			log.Printf("Erro ao solicitar recurso na requisição %d: %v", i, err)
			continue
		}
		fmt.Printf("Requisição %d - Recurso recebido: %s, Quantidade: %d\n", i+1, response.ResourceName, response.Quantity)
	}
}
