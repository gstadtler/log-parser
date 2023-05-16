# Dependências
- Golang

# Disclaimer
A branch ```main``` contém apenas o programa que retorna o output esperado na descrição do desafio;
Já a branch ```develop``` possui o programa que fornece os dados do output esperado numa API http, que é consumida pelo
cliente (frontend) contido no repositório: https://github.com/gstadtler/tibia-server-log

# Passos para rodar o projeto
- Criar um Fork e/ou Clonar este repositório
- Acessar o diretório "main" 

### Usando a API
- Fazer checkout na branch ```develop```
- Acessar o terminal e rodar o comando ```go run main.go``` para iniciar o servidor da API

Obs: comandos $curl podem ser utilizados para acessar a api

### Usando a branch main 
- Para obter apenas o output esperado basta:
- Acessar o terminal e rodar o comando ```go run main.go```
