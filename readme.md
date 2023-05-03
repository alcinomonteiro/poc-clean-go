# **POC-CLEAN-GO**
### **Pré-requisito**: Docker
Executar o comando abaixo para subir a Infra necessária para a aplicação funcionar
```
docker-compose -f "docker-compose-infra.yaml" up -d
```
## **Execução local**
## Passo 1
executar o comando abaixo para baixar os pacotes requeridos pelo projeto
```
go mod tidy
```
## Passo 2
executar o comando abaixo para subir a aplicação
```
go run .\app\infrastructure\main.go
```
## **Execução Docker**
Executar o comando abaixo para fazer o build da aplicação em container
```
docker-compose up -d --build
```
## **Mensagem de exemplo do kafka**
```
{"name":"nome","code":1111,"type":"tipo"}
```
## **Container health check**
- http://localhost:8086/live
- http://localhost:8086/ready

**Ps:** para ter uma visualização mais detalhara basta colocar. ```?full=1```