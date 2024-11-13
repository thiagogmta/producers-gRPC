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

// Solicita grãos para o servidor de grãos
func requestGrains(quantity int32) int32 {
	conn, err := grpc.Dial("produtor-de-graos:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("Erro ao conectar ao servidor de grãos: %v", err)
		return 0
	}
	defer conn.Close()

	client := pb.NewProducerClient(conn)
	response, err := client.GetResource(context.Background(), &pb.ResourceRequest{Quantity: quantity})
	if err != nil {
		log.Fatalf("Erro ao solicitar grãos: %v", err)
		return 0
	}
	return response.Quantity
}

// Implementa o método GetResource para produzir farinha
func (s *server) GetResource(ctx context.Context, in *pb.ResourceRequest) (*pb.ResourceResponse, error) {
	grains := requestGrains(in.Quantity)
	log.Printf("Quantidade de grãos recebida: %d", grains)
	return &pb.ResourceResponse{ResourceName: "Farinha", Quantity: grains}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProducerServer(s, &server{})
	log.Println("Servidor Produtor de Farinha rodando na porta 50053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao iniciar o servidor: %v", err)
	}
}
