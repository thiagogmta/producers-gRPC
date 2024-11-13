

## Configuração do ambiente e Criação do Arquivo Protobuf:

O arquivo .proto define um serviço chamado Producer que possui uma função remota GetResource. Essa função permite que um cliente envie uma solicitação (informando uma quantidade desejada) e receba uma resposta contendo o nome do recurso e a quantidade disponível.

Funcionamento:
- O cliente faz uma requisição ResourceRequest com a quantidade que deseja.
- O serviço Producer processa a requisição e retorna uma resposta ResourceResponse com o nome do recurso e a quantidade.

**Requisitos:**
- Linguagem: Go (go version go1.23.2)
- Framework gRPC (protoc, protoc-gen-go)

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH=$PATH:/usr/local/go/bin
```

```bash
protoc --go_out=. --go-grpc_out=. producer.proto
```

Inici

### Implementação dos Serviços Produtores em Go:

Compilar os serviços:

```bash
go build -o produtor-de-graos grains.go
go build -o produtor-de-agua water.go
go build -o produtor-de-farinha flour.go
go build -o produtor-de-pao bread.go
```

Abrir um terminal para cada produtor e mais um terminal para executar o cliente. Executar em cada terminal os comandos:

```bash
./produtor-de-graos
./produtor-de-agua
./produtor-de-farinha
./produtor-de-pao
go run client.go
```

Construir as imagens Docker

```bash
# Para o produtor de grãos
docker build -f Dockerfile.grains -t produtor-de-graos .

# Para o produtor de água
docker build -f Dockerfile.water -t produtor-de-agua .

# Para o produtor de farinha
docker build -f Dockerfile.flour -t produtor-de-farinha .

# Para o produtor de pão
docker build -f Dockerfile.bread -t produtor-de-pao .

# Para o cliente
docker build -f Dockerfile.client -t client .
```

Executar os contêineres

```bash
docker-compose up

para com
```

Deverá receber uma saída como:

```bash
Creating produtores_produtor-de-agua_1    ... done
Creating produtores_produtor-de-pao_1     ... done
Creating produtores_produtor-de-farinha_1 ... done
Creating produtores_produtor-de-graos_1   ... done
Creating produtores_client_1              ... done
Attaching to produtores_produtor-de-pao_1, produtores_produtor-de-agua_1, produtores_produtor-de-graos_1, produtores_produtor-de-farinha_1, produtores_client_1
```

Para executar o cliente em outro terminal

```bash
docker-compose run client
```

