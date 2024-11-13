package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "produtores/producer"
)

type server struct {
    pb.UnimplementedProducerServer
}

// Implementa o método GetResource para produzir água
func (s *server) GetResource(ctx context.Context, in *pb.ResourceRequest) (*pb.ResourceResponse, error) {
    log.Printf("Recebendo solicitação de %d unidades de água", in.Quantity)
	return &pb.ResourceResponse{ResourceName: "Água", Quantity: in.Quantity}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("Falha ao escutar: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterProducerServer(s, &server{})
    log.Println("Servidor Produtor de Água rodando na porta 50052...")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("Falha ao iniciar o servidor: %v", err)
    }
}
