package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "produtores/producer"

	"google.golang.org/grpc"
)

func main() {
	// Configuração da conexão com o servidor de pão
	conn, err := grpc.Dial("produtor-de-pao:50054", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Erro ao conectar: %v", err)
	}
	defer conn.Close()

	client := pb.NewProducerClient(conn)

	// Seleção do tipo de carga
	fmt.Println("Escolha o tipo de carga:")
	fmt.Println("1 - Carga Simples")
	fmt.Println("2 - Carga Sequencial e Constante")
	fmt.Println("3 - Carga Concorrente (Picos de Carga)")
	fmt.Println("4 - Carga Contínua (Loop Infinito)")
	var escolha int
	fmt.Scanln(&escolha)

	switch escolha {
	case 1:
		cargaSimples(client)
	case 2:
		cargaSequencial(client, 10) // Ajuste o número de requisições conforme necessário
	case 3:
		cargaConcorrente(client, 5, 10) // 5 picos, com até 10 requisições cada
	case 4:
		cargaContinua(client, 500*time.Millisecond) // Intervalo entre requisições
	default:
		fmt.Println("Opção inválida")
	}
}

// Carga Simples - Uma única requisição
func cargaSimples(client pb.ProducerClient) {
	response, err := client.GetResource(context.Background(), &pb.ResourceRequest{Quantity: 10})
	if err != nil {
		log.Printf("Erro ao solicitar recurso: %v", err)
		return
	}
	fmt.Printf("Recurso recebido: %s, Quantidade: %d\n", response.ResourceName, response.Quantity)
}

// Carga Sequencial e Constante - Envia várias requisições em sequência
func cargaSequencial(client pb.ProducerClient, numRequests int) {
	for i := 0; i < numRequests; i++ {
		response, err := client.GetResource(context.Background(), &pb.ResourceRequest{Quantity: 10})
		if err != nil {
			log.Printf("Erro ao solicitar recurso na requisição %d: %v", i, err)
			continue
		}
		fmt.Printf("Requisição %d - Recurso recebido: %s, Quantidade: %d\n", i+1, response.ResourceName, response.Quantity)
	}
}

// Carga Concorrente (Picos de Carga) - Envia picos de requisições simultâneas
func cargaConcorrente(client pb.ProducerClient, numPicos int, maxRequests int) {
	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup

	for i := 0; i < numPicos; i++ {
		numRequests := rand.Intn(maxRequests) + 1 // Número de requisições neste pico
		wg.Add(numRequests)

		for j := 0; j < numRequests; j++ {
			go func(j int) {
				defer wg.Done()
				response, err := client.GetResource(context.Background(), &pb.ResourceRequest{Quantity: 10})
				if err != nil {
					log.Printf("Erro ao solicitar recurso na requisição %d do pico %d: %v", j, i, err)
					return
				}
				fmt.Printf("Pico %d - Requisição %d - Recurso recebido: %s, Quantidade: %d\n", i+1, j+1, response.ResourceName, response.Quantity)
			}(j)
		}

		// Espera um tempo aleatório entre picos de carga
		time.Sleep(time.Duration(rand.Intn(2000)+500) * time.Millisecond)
	}

	wg.Wait() // Aguarda todas as requisições terminarem
}

// Carga Contínua (Loop Infinito) - Envia requisições constantemente em intervalos definidos
func cargaContinua(client pb.ProducerClient, intervalo time.Duration) {
	for {
		response, err := client.GetResource(context.Background(), &pb.ResourceRequest{Quantity: 10})
		if err != nil {
			log.Printf("Erro ao solicitar recurso: %v", err)
			continue
		}
		fmt.Printf("Recurso recebido: %s, Quantidade: %d\n", response.ResourceName, response.Quantity)
		time.Sleep(intervalo) // Intervalo entre requisições
	}
}
