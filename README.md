
# Distributed Resource Producers with gRPC

Sistema distribuído de produtores de recursos usando gRPC em Go, que simula o fornecimento e consumo de itens como grãos, água, farinha e pão. Esse projeto utiliza o protocolo gRPC para comunicação entre serviços, permitindo ao cliente solicitar quantidades específicas de recursos e receber as informações de disponibilidade.

## Funcionalidade

Este projeto define um serviço chamado `Producer` por meio de um arquivo `.proto`, onde cada produtor (grãos, água, farinha e pão) processa solicitações enviadas pelo cliente e responde com o nome do recurso e a quantidade disponível.

### Funcionamento:

1. O cliente faz uma requisição `ResourceRequest` com a quantidade desejada de um recurso.
2. O serviço `Producer` processa a requisição e retorna uma resposta `ResourceResponse` com o nome do recurso e a quantidade disponível.

---

## Requisitos

- **Linguagem:** Go (versão 1.23.2)
- **Framework:** gRPC com `protoc` e `protoc-gen-go`

### Configuração do Ambiente

1. Instale o plugin `protoc-gen-go` para o Protobuf e o gRPC:

   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

2. Adicione o Go ao seu PATH:

   ```bash
   export PATH=$PATH:/usr/local/go/bin
   ```

3. Compile o arquivo `.proto` para gerar os arquivos Go necessários:

   ```bash
   protoc --go_out=. --go-grpc_out=. producer.proto
   ```

---

## Implementação dos Serviços Produtores em Go

Cada serviço produtor é implementado em Go para simular um recurso específico. Os arquivos principais para os serviços são:

- **`grains.go`** - Produtor de grãos
- **`water.go`** - Produtor de água
- **`flour.go`** - Produtor de farinha
- **`bread.go`** - Produtor de pão

Para compilar os serviços, execute:

```bash
go build -o produtor-de-graos grains.go
go build -o produtor-de-agua water.go
go build -o produtor-de-farinha flour.go
go build -o produtor-de-pao bread.go
```

Abra um terminal para cada produtor e execute os comandos para iniciar cada serviço:

```bash
./produtor-de-graos
./produtor-de-agua
./produtor-de-farinha
./produtor-de-pao
```

Em outro terminal, execute o cliente:

```bash
go run client.go
```

---

## Configuração com Docker

Para facilitar a execução e gerenciamento dos serviços, o projeto pode ser executado com Docker e Docker Compose.

### Construir as Imagens Docker

Para cada serviço produtor e para o cliente, crie uma imagem Docker:

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

### Executar os Contêineres com Docker Compose

Inicie todos os contêineres de uma vez usando o Docker Compose:

```bash
docker-compose up
```

Você deverá ver uma saída similar a:

```bash
Creating produtores_produtor-de-agua_1    ... done
Creating produtores_produtor-de-pao_1     ... done
Creating produtores_produtor-de-farinha_1 ... done
Creating produtores_produtor-de-graos_1   ... done
Creating produtores_client_1              ... done
Attaching to produtores_produtor-de-pao_1, produtores_produtor-de-agua_1, produtores_produtor-de-graos_1, produtores_produtor-de-farinha_1, produtores_client_1
```

Para executar o cliente em outro terminal, use:

```bash
docker-compose run client
```

---

## Estrutura do Projeto

- **`producer.proto`** - Arquivo Protobuf que define o serviço `Producer` e as mensagens `ResourceRequest` e `ResourceResponse`.
- **`grains.go`, `water.go`, `flour.go`, `bread.go`** - Implementações dos serviços de cada produtor.
- **`client.go`** - Cliente que envia solicitações para os produtores.
- **Dockerfiles** - Arquivos de configuração Docker para cada serviço.
- **`docker-compose.yml`** - Configuração Docker Compose para gerenciar todos os contêineres.

---

## Comandos Úteis

```bash
#Parar e remover contêineres:**
docker-compose down

# Reiniciar contêineres:
docker-compose restart

# Parar todos os contêineres em execução
docker stop $(docker ps -q)

# Remover Todos os Contêineres Parados
docker rm $(docker ps -aq)

# Limpar Volumes e Redes Antigos
docker volume prune -f
docker network prune -f

```

---

## Licença

Este projeto é disponibilizado sob a [MIT License](LICENSE).

---

## Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir uma issue ou enviar um pull request.

---
