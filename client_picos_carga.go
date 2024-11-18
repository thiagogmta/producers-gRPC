package main

import (
	"context"
	"fmt"
	"log"
	"sync"

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

	// Número de requisições concorrentes
	numRequests := 10
	var wg sync.WaitGroup

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			response, err := client.GetResource(context.Background(), &pb.ResourceRequest{Quantity: 10})
			if err != nil {
				log.Printf("Erro ao solicitar recurso na requisição %d: %v", i, err)
				return
			}
			fmt.Printf("Requisição %d - Recurso recebido: %s, Quantidade: %d\n", i+1, response.ResourceName, response.Quantity)
		}(i)
	}

	wg.Wait() // Aguarda todas as requisições concorrentes serem concluídas
}
