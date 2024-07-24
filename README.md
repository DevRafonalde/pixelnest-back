# Documentação da API CRUD em Golang com Gin

## Sumário

1. [Introdução](#introdução)
2. [Estrutura do Projeto](#estrutura-do-projeto)
3. [Configuração](#configuração)
4. [Endpoints da API](#endpoints-da-api)
    - [GET /api/usuario/:id](#get-apiusuarioid)
    - [GET /api/usuario/email](#get-apiusuarioemail)
    - [GET /api/usuario](#get-apitodosusuarios)
    - [GET /api/usuario/csv](#get-apitodosusuarioscsv)
    - [POST /api/usuario](#post-apiusuario)
    - [POST /api/usuario/csv](#post-apiusuariocsv)
    - [PUT /api/usuario/:id](#put-apiusuarioid)
    - [DELETE /api/usuario/:id](#delete-apiusuarioid)
    - [DELETE /api/usuario/all](#delete-apiallusuarios)
5. [Modelos de Dados](#modelos-de-dados)
6. [Serviços](#serviços)
7. [Controladores](#controladores)
8. [Instruções para Execução](#instruções-para-execução)

## Introdução

Esta documentação descreve uma API CRUD desenvolvida em Golang usando o framework Gin. A API permite criar, ler, atualizar e excluir usuários em um banco de dados PostgreSQL.

## Estrutura do Projeto

```plaintext
crud-rafael/
├── controller/
│   └── usuarioController.go
├── db/
│   └── connection.go
├── model/
│   └── usuario.go
├── service/
│   └── usuarioService.go
├── main.go
└── go.mod
```

## Configuração

### Banco de Dados

A configuração do banco de dados está definida no arquivo `db/connection.go`. Um banco de dados PostgreSQL é utilizado, com as seguintes credenciais:

- **Usuário**: postgres
- **Senha**: example
- **Host**: localhost
- **Porta**: 5432
- **Nome do Banco de Dados**: postgres (será criado um banco de dados chamado `usuario`)

### Docker Compose

Para rodar o banco de dados PostgreSQL e o Adminer, você pode usar o seguinte arquivo `docker-compose.yml`:

```yaml
version: '3.1'

services:
  db:
    image: postgres
    restart: always
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: example
      POSTGRES_DB: postgres
    ports:
      - "5432:5432"

  adminer:
    image: adminer
    restart: always
    ports:
      - "8080:8080"
```

## Endpoints da API

### GET /api/usuario/:id

#### Descrição

Recupera um usuário pelo ID.

#### Exemplo de Requisição

```http
GET /api/usuario/1
```

#### Exemplo de Resposta

```json
{
    "ID": 1,
    "Nome": "Rafael",
    "Email": "rafael@example.com"
}
```

### GET /api/usuario/email

#### Descrição

Recupera um usuário pelo Email.

#### Exemplo de Requisição

```json
{
    "Email": "rafael@example.com"
}
```

#### Exemplo de Resposta

```json
{
    "ID": 1,
    "Nome": "Rafael",
    "Email": "rafael@example.com"
}
```

### GET /api/usuario

#### Descrição

Recupera todos os usuários do banco.

#### Exemplo de Requisição

```http
GET /api/usuario/
```

#### Exemplo de Resposta

```json
[
    {
        "ID": 1,
        "Nome": "Rafael",
        "Email": "rafael@teste.com"
    },
    {
        "ID": 3,
        "Nome": "Camila",
        "Email": "camila@teste.com"
    }
]
```

### GET /api/usuario/csv

#### Descrição

Recupera todos os usuários do banco e exporta em um arquivo `.csv`.

#### Exemplo de Requisição

```http
GET /api/usuario/csv
```

#### Exemplo de Resposta

```csv
    ID,Nome,Email
    105,Ryan Hartman,xhughes@hernandez.net
    106,Joshua Nichols,cathy54@gmail.com
    107,Monica Ramirez,derek77@lewis.com
    108,Meredith Harris,velasquezstacy@becker.com
    109,Jennifer Elliott,hcherry@stevens-dunlap.com
    110,David Williams,michaelsalas@hotmail.com
```

### POST /api/usuario

#### Descrição

Cria um novo usuário.

#### Exemplo de Requisição

```json

{
    "Nome": "Rafael",
    "Email": "rafael@teste.com"
}
```

#### Exemplo de Resposta

```json
{
    "ID": 1,
    "Nome": "Rafael",
    "Email": "rafael@teste.com"
}
```

### POST /api/usuario

#### Descrição

Cria novos usuários em lote, baseados em um arquivo `.csv`.

#### Exemplo de Requisição

```csv
    Nome,Email
    Ryan Hartman,xhughes@hernandez.net
    Joshua Nichols,cathy54@gmail.com
    Monica Ramirez,derek77@lewis.com
    Meredith Harris,velasquezstacy@becker.com
    Jennifer Elliott,hcherry@stevens-dunlap.com
    David Williams,michaelsalas@hotmail.com
```

#### Exemplo de Resposta

```json
{
    "mensagem": "Arquivo carregado e processado com sucesso!"
}
```

### PUT /api/usuario/:id

#### Descrição

Atualiza um usuário existente pelo ID.

#### Exemplo de Requisição

```http
PUT /api/usuario/1
Content-Type: application/json

{
    "Nome": "Rafael Atualizado",
    "Email": "rafael.updated@example.com"
}
```

#### Exemplo de Resposta

```json
{
    "ID": 1,
    "Nome": "Rafael Atualizado",
    "Email": "rafael.updated@example.com"
}
```

### DELETE /api/usuario/:id

#### Descrição

Exclui um usuário pelo ID.

#### Exemplo de Requisição

```http
DELETE /api/usuario/1
```

#### Exemplo de Resposta

```json
{
    "deletado": true
}
```

### DELETE /api/usuario/all

#### Descrição

Exclui **TODOS** os usuários.

#### Exemplo de Requisição

```http
DELETE /api/usuario/all
```

#### Exemplo de Resposta

```json
{
    "mensagem": "Todos os usuários foram deletados com sucesso!"
}
```

## Modelos de Dados

### Usuario

O modelo `Usuario` está definido no arquivo `model/usuario.go`.

```go
package model

type Usuario struct {
    ID    uint64 `gorm:"primary_key,autoIncrement"`
    Nome, Email string
}
```

## Serviços

### UsuarioService

Os serviços relacionados ao usuário estão definidos no arquivo `service/usuarioService.go`.

- **FindUsuarioById(id uint64)**: Busca um usuário pelo ID.
- **CreateUsuario(usuario model.Usuario)**: Cria um novo usuário.
- **UpdateUsuario(usuarioRecebido model.Usuario, id uint64)**: Atualiza um usuário existente.
- **DeleteUsuarioById(id uint64)**: Exclui um usuário pelo ID.

## Controladores

### UsuarioController

Os controladores estão definidos no arquivo `controller/usuarioController.go`.

- **InitRoutes()**: Inicializa as rotas da API.
- **findUsuarioById(context *gin.Context)**: Recupera um usuário pelo ID.
- **findUsuarioByEmail(context *gin.Context)**: Recupera um usuário pelo e-mail.
- **findAllUsuarios(context *gin.Context)**: Recupera todos os usuários.
- **findAllUsuariosExportCSV(context *gin.Context)**: Recupera todos os usuários e os exporta em um arquivo csv.
- **createUsuario(context *gin.Context)**: Cria um registro de um único usuário.
- **createUsuarioByCSV(context *gin.Context)**: Cria um registro para cada usuário contido no arquivo `.csv` enviado.
- **updateUsuario(context *gin.Context)**: Atualiza um usuário existente pelo ID.
- **deleteUsuarioById(context *gin.Context)**: Exclui um usuário pelo ID.
- **deleteAllUsuarios(context *gin.Context)**: Exclui **TODOS** os usuários.

## Instruções para Execução

### Pré-requisitos

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Passos para Execução

1. **Inicie o banco de dados PostgreSQL e Adminer**:

   ```sh
   docker-compose up
   ```

2. **Execute o aplicativo Go**:

   ```sh
   go run main.go
   ```

3. **Acesse a API**:

   - A API estará disponível em `http://localhost:8601`.
   - O Adminer estará disponível em `http://localhost:8080` para gerenciar o banco de dados.

### Testando a API

Você pode usar ferramentas como `curl`, `Postman`, ou qualquer outra ferramenta de sua preferência para testar os endpoints da API.

Exemplo usando `curl`:

```sh
# Criar um novo usuário
curl -X POST http://localhost:8601/api/usuario -H "Content-Type: application/json" -d '{"Nome":"Rafael", "Email":"rafael@example.com"}'

# Buscar um usuário pelo ID
curl http://localhost:8601/api/usuario/1

# Atualizar um usuário pelo ID
curl -X PUT http://localhost:8601/api/usuario/1 -H "Content-Type: application/json" -d '{"Nome":"Rafael Atualizado", "Email":"rafael.updated@example.com"}'

# Excluir um usuário pelo ID
curl -X DELETE http://localhost:8601/api/usuario/1
```

Esta documentação cobre os principais aspectos da API CRUD desenvolvida em Go com o framework Gin. Ela inclui detalhes sobre a estrutura do projeto, configuração, endpoints da API, modelos de dados, serviços, controladores e instruções para execução.
