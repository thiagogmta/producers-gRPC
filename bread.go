package main

import (
	"context"
	"log"
	"net"

	pb "produtores/producer"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProducerServer
}

// Solicita recursos de outros produtores
func requestResource(resource string, quantity int32, address string) int32 {
	// Conecta-se diretamente ao endereço (nome do serviço e porta)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Printf("Erro ao conectar ao servidor de %s: %v", resource, err)
		return 0
	}
	defer conn.Close()

	client := pb.NewProducerClient(conn)
	response, err := client.GetResource(context.Background(), &pb.ResourceRequest{Quantity: quantity})
	if err != nil {
		log.Printf("Erro ao solicitar %s: %v", resource, err)
		return 0
	}
	return response.Quantity
}

// Implementa o método GetResource para produzir pão
func (s *server) GetResource(ctx context.Context, in *pb.ResourceRequest) (*pb.ResourceResponse, error) {
	// Chamadas diretas usando endereço completo para cada recurso
	flour := requestResource("Farinha", in.Quantity, "produtor-de-farinha:50053")
	water := requestResource("Água", in.Quantity, "produtor-de-agua:50052")
	log.Printf("Quantidade de farinha: %d, Quantidade de água: %d", flour, water)
	breadQuantity := min(flour, water) // A quantidade de pão depende de quem tem menos
	return &pb.ResourceResponse{ResourceName: "Pão", Quantity: breadQuantity}, nil
}

func min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProducerServer(s, &server{})
	log.Println("Servidor Produtor de Pão rodando na porta 50054...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}
