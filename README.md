# go-rest-api

### BACKEND
- Estudos sobre como criar uma REST-API utilizando Golang
-- A aplicacao foi modelada utilizando, camada de models e services.
-- Foi utilizado o *GORM*, ORM do Go  que uma biblioteca que que permite 
   fazer uma relação dos objetos com os dados que os mesmos representam.
-- Foi utilizado *GORILLA/MUX* para criar as rotas da aplicacao
-- Foi utilizado tambem *GODOTENV* para referenciar as variaveis de ambiente da aplicacao

### INFRA
- Foi utilizado docker para criar um container com a imagem do postgres
- Foi criado um arquivo de docker compose para facilitar a criacao do conteiner

- Linux
```sh
  docker compose up -d
```
- Windows

```sh
  winpty docker compose up -d
```
