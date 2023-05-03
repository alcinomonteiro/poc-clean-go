# POC filial-go (caminho feliz)

**pré-requisito**: Docker

## Passo 1
Executar o comando abaixo para levantar a Infra necessária para a aplicação funcionar
```
docker-compose up -d
```
## Passo 2
Conectar ao Kafka e criar o tópico conforme indicado no .env

## Passo 3
Conectar ao MongoDB e criar a base de dados e a coleção conforme indicado no .env

## Passo 4
executar o comando abaixo para baixar os pacotes requeridos pelo projeto
```
go mod tidy
```
## Passo 5
executar o comando abaixo para subir a aplicação
```
go run main.go
```
## Mensagem de exemplo do kafka
```
{"name":"nome","code": 33333,"type":"tipo"}
```
