package main

import (
	"context"
	"fmt"
	"log"
	"time"

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

	// Intervalo entre cada requisição no loop contínuo
	interval := 2 * time.Second

	for i := 1; ; i++ { // Loop infinito
		response, err := client.GetResource(context.Background(), &pb.ResourceRequest{Quantity: 10})
		if err != nil {
			log.Printf("Erro ao solicitar recurso na requisição %d: %v", i, err)
			continue
		}
		fmt.Printf("Requisição %d - Recurso recebido: %s, Quantidade: %d\n", i, response.ResourceName, response.Quantity)
		time.Sleep(interval) // Intervalo entre requisições
	}
}
