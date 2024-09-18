# Documentação da API Simfonia em Golang com Gin

## Sumário

1. [Introdução](#introducao)
2. [Estrutura do Projeto](#estrutura-do-projeto)
3. [Configuração](#configuracao)
    - [Banco de Dados](#banco-de-dados)
    - [Docker Compose](#docker-compose)
4. [Endpoints da API](#endpoints-da-api)
    1. [Cidades](#cidades)
        - [GET /simfonia/api/cidades/:id](#get-apicidadeid)
        - [GET /simfonia/api/cidades/nome](#get-apicidadenome)
        - [GET /simfonia/api/cidades/](#get-apiallcidades)
        - [GET /simfonia/api/cidades/csv](#get-apiallcidadescsv)
        - [POST /simfonia/api/cidades/](#post-apicreatecidade)
        - [POST /simfonia/api/cidades/csv](#post-apicreatecidadecsv)
        - [PUT /simfonia/api/cidades/:id](#put-apiupdatecidade)
        - [DELETE /simfonia/api/cidades/:id](#delete-apideletecidadeid)
    2. [Números Telefônicos](#numeros-telefonicos)
        - [GET /simfonia/api/numerostelefonicos/:id](#get-apinumerosid)
        - [GET /simfonia/api/numerostelefonicos/numero](#get-apinumerosnumero)
        - [GET /simfonia/api/numerostelefonicos/simcard/:id](#get-apinumerossimcardid)
        - [GET /simfonia/api/numerostelefonicos/simcard](#get-apinumerossimcard)
        - [GET /simfonia/api/numerostelefonicos/](#get-apiallnumeros)
        - [GET /simfonia/api/numerostelefonicos/csv](#get-apiallnumeroscsv)
        - [POST /simfonia/api/numerostelefonicos/](#post-apicreatenumero)
        - [POST /simfonia/api/numerostelefonicos/csv](#post-apicreatenumerocsv)
        - [PUT /simfonia/api/numerostelefonicos/:id](#put-apiupdatenumero)
        - [DELETE /simfonia/api/numerostelefonicos/:id](#delete-apideletenumeroid)
    3. [Operadoras](#operadoras)
        - [GET /simfonia/api/operadoras/:id](#get-apioperadorasid)
        - [GET /simfonia/api/operadoras/nome](#get-apioperadorasnome)
        - [GET /simfonia/api/operadoras/abreviacao](#get-apioperadorasabreviacao)
        - [GET /simfonia/api/operadoras/](#get-apialloperadoras)
        - [GET /simfonia/api/operadoras/csv](#get-apialloperadorascsv)
        - [POST /simfonia/api/operadoras/](#post-apicreateoperadoras)
        - [POST /simfonia/api/operadoras/csv](#post-apicreateoperadorascsv)
        - [PUT /simfonia/api/operadoras/:id](#put-apiupdateoperadoras)
        - [DELETE /simfonia/api/operadoras/:id](#get-apideleteoperadorasid)
    4. [SimCard](#simCard)
        - [GET /simfonia/api/simcard/:id](#get-apisimcardid)
        - [GET /simfonia/api/simcard/telefonianumero/:id](#get-apisimcardnumeroid)
        - [GET /simfonia/api/simcard/telefonianumero](#get-apisimcardnumero)
        - [GET /simfonia/api/simcard/](#get-apiallsimcard)
        - [GET /simfonia/api/simcard/csv](#get-apiallsimcardcsv)
        - [POST /simfonia/api/simcard/](#post-apicreateosimcard)
        - [POST /simfonia/api/simcard/csv](#post-apicreatesimcardcsv)
        - [PUT /simfonia/api/simcard/:id](#put-apiupdatesimcard)
        - [DELETE /simfonia/api/simcard/:id](#get-apideletesimcardid)
    5. [Estados SimCard](#simCard-estados)
        - [GET /simfonia/api/simcardestado/:id](#get-apisimcardestadoid)
        - [GET /simfonia/api/simcardestado/estado](#get-apisimcardestadoestado)
        - [GET /simfonia/api/simcardestado/](#get-apiallsimcardestado)
        - [GET /simfonia/api/simcardestado/csv](#get-apiallsimcardestadocsv)
        - [POST /simfonia/api/simcardestado/](#post-apicreateosimcardestado)
        - [POST /simfonia/api/simcardestado/csv](#post-apicreatesimcardestadocsv)
        - [PUT /simfonia/api/simcardestado/:id](#put-apiupdatesimcardestado)
        - [DELETE /simfonia/api/simcardestado/:id](#get-apideletesimcardestadoid)
5. [Modelos de Dados](#modelos-de-dados)
6. [Serviços](#serviços)
7. [Controladores](#controladores)
8. [Instruções para Execução](#instruções-para-execução)

## Introdução {#introducao}

Esta documentação descreve uma API CRUD desenvolvida em Golang usando o framework Gin. A API permite criar, ler, atualizar e excluir usuários em um banco de dados PostgreSQL.

## Estrutura do Projeto {#estrutura-do-projeto}

```plaintext
crud-rafael/
├── controller/
│   └── cidadeController.go
│   └── numerosTelefonicosController.go
│   └── operadorasController.go
│   └── simCardController.go
│   └── simCardEstadoController.go
├── db/
│   └── db.go
├── model/
│   └── cidades.go
│   └── numerosTelefonicos.go
│   └── operadoras.go
│   └── simCard.go
│   └── simCardEstado.go
│   └── parametrosDeBusca/
│       └── cidades/
│           └── nome.go
│       └── numerosTelefonicos/
│           └── numero.go
│       └── operadoras/
│           └── abreviacao.go
│           └── nome.go
│       └── simCardEstado/
│           └── nome.go
├── service/
│   └── cidadesService.go
│   └── numerosTelefonicosService.go
│   └── operadorasService.go
│   └── simCardService.go
│   └── simCardEstadoService.go
├── main.go
├── go.sum
└── go.mod

```

## Configuração {#configuracao}

### Banco de Dados {#banco-de-dados}

A configuração do banco de dados está definida no arquivo `db/connection.go`. Um banco de dados PostgreSQL é utilizado, com as seguintes credenciais:

- **Usuário**: postgres
- **Senha**: example
- **Host**: localhost
- **Porta**: 5432
- **Nome do Banco de Dados**: postgres (será criado um banco de dados chamado `usuario`)

### Docker Compose {#docker-compose}

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

### Cidades {#cidades}

#### GET /simfonia/api/cidades/:id {#get-apicidadeid}

##### Descrição

Recupera uma cidade pelo ID.

#### Exemplo de Requisição

```http
GET /simfonia/api/cidades/2
```

#### Exemplo de Resposta

```json
{
    "ID": 2,
    "UUID": "123456789-1234-1234-1234-1234567890ab",
    "Nome": "São Paulo",
    "CodIBGE": 1234567,
    "UF": "RJ",
    "CodArea": 12
}
```

#### GET /simfonia/api/cidades/nome {#get-apicidadenome}

##### Descrição

Recupera uma cidade pelo nome.

##### Exemplo de Requisição

```json
{
    "Nome": "São Paulo"
}
```

#### Exemplo de Resposta

```json
{
    "ID": 2,
    "UUID": "123456789-1234-1234-1234-1234567890ab",
    "Nome": "São Paulo",
    "CodIBGE": 1234567,
    "UF": "RJ",
    "CodArea": 12
}
```

#### GET /simfonia/api/cidades/ {#get-apiallcidades}

##### Descrição

Recupera todos os usuários do banco.

##### Exemplo de Requisição

```http
GET /api/usuario/
```

#### Exemplo de Resposta

```json
[
    {
        "ID": 2,
        "UUID": "123456789-1234-1234-1234-1234567890ab",
        "Nome": "São Paulo",
        "CodIBGE": 1234567,
        "UF": "RJ",
        "CodArea": 12
    },
    {
        "ID": 3,
        "UUID": "123456789-1234-1234-1234-1234567890ab",
        "Nome": "Rio de Janeiro",
        "CodIBGE": 1234567,
        "UF": "SP",
        "CodArea": 12
    }
]
```

#### GET /simfonia/api/cidades/csv {#get-apiallcidadescsv}

##### Descrição

Recupera todas as cidades do banco e exporta em um arquivo `.csv`.

##### Exemplo de Requisição

```http
GET /simfonia/api/cidades/csv
```

##### Exemplo de Resposta

```csv
    ID,UUID,Nome,CodIBGE,UF,CodArea
    2,832696e8-ab2a-40ed-9733-f3e7c86f14a5,Assis Brasil,1200054,AC,68
    3,df26cb7d-46f5-473c-bf59-d967092b8b41,Brasiléia,1200104,AC,68
    4,097c3b36-4757-48ea-aa81-73a448f3dba6,Bujari,1200138,AC,68
    5,4edfc55f-64e0-49de-9d72-3621041cf43a,Capixaba,1200179,AC,68
    6,bce68ddd-1a31-4f38-a3a5-6ea27ac2de5f,Cruzeiro do Sul,1200203,AC,68
    7,1fee47bc-9d10-4092-8b2b-a19590516b9a,Epitaciolândia,1200252,AC,68
    8,f87785f2-06f2-42ad-b85b-271e546a39af,Feijó,1200302,AC,68
    9,f4a6a716-e30c-4fa9-abc1-cac07bc3f53d,Jordão,1200328,AC,68
    10,35abfd59-ce53-4986-a981-efd1f63d8657,Mâncio Lima,1200336,AC,68
```

#### POST /simfonia/api/cidades/ {#post-apicreatecidade}

##### Descrição

Cria uma nova cidade.

##### Exemplo de Requisição

```json

{
    "UUID": "123456789-1234-1234-1234-1234567890ab",
    "Nome": "Rio de Janeiro",
    "CodIBGE": 1234567,
    "UF": "SP",
    "CodArea": 12
}
```

#### Exemplo de Resposta

```json
{
    "ID": 1,
    "UUID": "123456789-1234-1234-1234-1234567890ab",
    "Nome": "Rio de Janeiro",
    "CodIBGE": 1234567,
    "UF": "SP",
    "CodArea": 12
}
```

#### POST /simfonia/api/cidades/csv {#post-apicreatecidadecsv}

##### Descrição

Cria novas cidades em lote, baseados em um arquivo `.csv`.

##### Exemplo de Requisição

```csv
    UUID,Nome,CodIBGE,UF,CodArea
    832696e8-ab2a-40ed-9733-f3e7c86f14a5,Assis Brasil,1200054,AC,68
    df26cb7d-46f5-473c-bf59-d967092b8b41,Brasiléia,1200104,AC,68
    097c3b36-4757-48ea-aa81-73a448f3dba6,Bujari,1200138,AC,68
    4edfc55f-64e0-49de-9d72-3621041cf43a,Capixaba,1200179,AC,68
    bce68ddd-1a31-4f38-a3a5-6ea27ac2de5f,Cruzeiro do Sul,1200203,AC,68
    1fee47bc-9d10-4092-8b2b-a19590516b9a,Epitaciolândia,1200252,AC,68
    f87785f2-06f2-42ad-b85b-271e546a39af,Feijó,1200302,AC,68
    f4a6a716-e30c-4fa9-abc1-cac07bc3f53d,Jordão,1200328,AC,68
    ,35abfd59-ce53-4986-a981-efd1f63d8657,Mâncio Lima,1200336,AC,68
```

```multipart-form
    csv: *arquivo*
```

##### Exemplo de Resposta

```json
{
    "mensagem": "Arquivo carregado e processado com sucesso!"
}
```

#### PUT /simfonia/api/cidades/:id {#put-apiupdatecidade}

##### Descrição

Atualiza uma cidade existente pelo ID.

##### Exemplo de Requisição

```http
    PUT /api/usuario/1
```

```json
{
    "UUID": "123456789-1234-1234-1234-1234567890ab",
    "Nome": "Rio de Janeiro",
    "CodIBGE": 1234567,
    "UF": "SP",
    "CodArea": 12
}
```

##### Exemplo de Resposta

```json
{
    "ID": 1,
    "UUID": "123456789-1234-1234-1234-1234567890ab",
    "Nome": "Rio de Janeiro",
    "CodIBGE": 1234567,
    "UF": "SP",
    "CodArea": 12
}
```

#### DELETE /simfonia/api/cidades/:id {#delete-apideletecidadeid}

##### Descrição

Exclui uma cidade pelo ID.

##### Exemplo de Requisição

```http
DELETE /simfonia/api/cidades/1
```

##### Exemplo de Resposta

```json
{
    "deletado": true
}
```

### Números Telefônicos {#numeros-telefonicos}

#### GET /simfonia/api/numerostelefonicos/:id {#get-apinumerosid}

##### Descrição

Recupera um número telefônico pelo ID.

#### Exemplo de Requisição

```http
GET /simfonia/api/numerostelefonicos/2
```

#### Exemplo de Resposta

```json
{
    "ID": 7,
    "CodArea": 11,
    "Numero": 123456789,
    "Utilizavel": true,
    "PortadoIn": false,
    "PortadoInOperadora": "",
    "PortadoInDate": "2024-07-24T17:21:48.89283-03:00",
    "CodigoCNL": "ABC123",
    "CongeladoAte": null,
    "ExternalID": 1234567890,
    "PortadoOut": false,
    "PortadoOutOperadora": "",
    "PortadoOutDate": "2024-07-24T17:21:48.89283-03:00",
    "DataCriacao": "2024-06-24T10:00:00-03:00",
    "SimCardID": null,
    "SimCard": null,
    "PortadoInOperadoraID": null,
    "PortadoInOperadoraObj": null,
    "PortadoOutOperadoraID": null,
    "PortadoOutOperadoraObj": null
}
```

#### GET /simfonia/api/numerostelefonicos/numero {#get-apinumerosnumero}

##### Descrição

Recupera um número telefônico pelo número.

##### Exemplo de Requisição

```json
{
    "numero": 123456789
}
```

#### Exemplo de Resposta

```json
{
    "ID": 7,
    "CodArea": 11,
    "Numero": 123456789,
    "Utilizavel": true,
    "PortadoIn": false,
    "PortadoInOperadora": "",
    "PortadoInDate": "2024-07-24T17:21:48.89283-03:00",
    "CodigoCNL": "ABC123",
    "CongeladoAte": null,
    "ExternalID": 1234567890,
    "PortadoOut": false,
    "PortadoOutOperadora": "",
    "PortadoOutDate": "2024-07-24T17:21:48.89283-03:00",
    "DataCriacao": "2024-06-24T10:00:00-03:00",
    "SimCardID": null,
    "SimCard": null,
    "PortadoInOperadoraID": null,
    "PortadoInOperadoraObj": null,
    "PortadoOutOperadoraID": null,
    "PortadoOutOperadoraObj": null
}
```

#### GET /simfonia/api/numerostelefonicos/simcard/:id {#get-apinumerossimcardid}

##### Descrição

Recupera um número de telefone pelo id do SimCard correspondente.

##### Exemplo de Requisição

```http
GET /simfonia/api/numerostelefonicos/simcard/235
```

#### Exemplo de Resposta

```json
{
    "ID": 17,
    "CodArea": 12,
    "Numero": 981832641,
    "Utilizavel": true,
    "PortadoIn": false,
    "PortadoInOperadora": "",
    "PortadoInDate": "2024-06-26T16:41:52.571-03:00",
    "CodigoCNL": "0123456789",
    "CongeladoAte": null,
    "ExternalID": 987654321,
    "PortadoOut": false,
    "PortadoOutOperadora": "",
    "PortadoOutDate": "2024-06-26T16:41:52.571-03:00",
    "DataCriacao": "2024-06-26T16:41:52.571-03:00",
    "SimCardID": 235,
    "SimCard": null,
    "PortadoInOperadoraID": null,
    "PortadoInOperadoraObj": null,
    "PortadoOutOperadoraID": null,
    "PortadoOutOperadoraObj": null
}
```

#### GET /simfonia/api/numerostelefonicos/simcard {#get-apinumerossimcard}

##### Descrição

Recupera um número de telefone pelo objeto do SimCard correspondente.

##### Exemplo de Requisição

```http
    TODO
```

##### Exemplo de Resposta

```json
{
    "ID": 17,
    "CodArea": 12,
    "Numero": 981832641,
    "Utilizavel": true,
    "PortadoIn": false,
    "PortadoInOperadora": "",
    "PortadoInDate": "2024-06-26T16:41:52.571-03:00",
    "CodigoCNL": "0123456789",
    "CongeladoAte": null,
    "ExternalID": 987654321,
    "PortadoOut": false,
    "PortadoOutOperadora": "",
    "PortadoOutDate": "2024-06-26T16:41:52.571-03:00",
    "DataCriacao": "2024-06-26T16:41:52.571-03:00",
    "SimCardID": 235,
    "SimCard": null,
    "PortadoInOperadoraID": null,
    "PortadoInOperadoraObj": null,
    "PortadoOutOperadoraID": null,
    "PortadoOutOperadoraObj": null
}
```

#### GET /simfonia/api/numerostelefonicos/ {#get-apiallnumeros}

##### Descrição

Recupera todos os usuários do banco.

##### Exemplo de Requisição

```http
GET /api/usuario/
```

#### Exemplo de Resposta

```json
[
    {
        "ID": 7,
        "CodArea": 11,
        "Numero": 123456789,
        "Utilizavel": true,
        "PortadoIn": false,
        "PortadoInOperadora": "",
        "PortadoInDate": "2024-07-24T17:21:48.89283-03:00",
        "CodigoCNL": "ABC123",
        "CongeladoAte": null,
        "ExternalID": 1234567890,
        "PortadoOut": false,
        "PortadoOutOperadora": "",
        "PortadoOutDate": "2024-07-24T17:21:48.89283-03:00",
        "DataCriacao": "2024-06-24T10:00:00-03:00",
        "SimCardID": null,
        "SimCard": null,
        "PortadoInOperadoraID": null,
        "PortadoInOperadoraObj": null,
        "PortadoOutOperadoraID": null,
        "PortadoOutOperadoraObj": null
    },
    {
        "ID": 8,
        "CodArea": 22,
        "Numero": 234567890,
        "Utilizavel": true,
        "PortadoIn": false,
        "PortadoInOperadora": "",
        "PortadoInDate": "2024-07-24T17:21:48.89283-03:00",
        "CodigoCNL": "DEF456",
        "CongeladoAte": null,
        "ExternalID": 2345678901,
        "PortadoOut": false,
        "PortadoOutOperadora": "",
        "PortadoOutDate": "2024-07-24T17:21:48.89283-03:00",
        "DataCriacao": "2024-06-24T11:00:00-03:00",
        "SimCardID": null,
        "SimCard": null,
        "PortadoInOperadoraID": null,
        "PortadoInOperadoraObj": null,
        "PortadoOutOperadoraID": null,
        "PortadoOutOperadoraObj": null
    }
]
```

#### GET /simfonia/api/numerostelefonicos/csv {#get-apiallnumeroscsv}

##### Descrição

Recupera todos os números telefônicos do banco e exporta em um arquivo `.csv`.

##### Exemplo de Requisição

```http
GET /simfonia/api/numerostelefonicos/csv
```

##### Exemplo de Resposta

```csv
    ID,CodArea,Numero,Utilizavel,PortadoIn,PortadoInOperadora,PortadoInDate,CodigoCNL,CongeladoAte,ExternalID,PortadoOut,PortadoOutOperadora,PortadoOutDate,DataCriacao,SimCardID,PortadoInOperadoraID,PortadoOutOperadoraID
7,11,123456789,true,false,1234567890,2024-07-24 17:21:48.89283 -0300 -03,ABC123,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 10:00:00 -0300 -03,null,null,null
8,22,234567890,true,false,2345678901,2024-07-24 17:21:48.89283 -0300 -03,DEF456,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 11:00:00 -0300 -03,null,null,null
9,33,345678901,true,false,3456789012,2024-07-24 17:21:48.89283 -0300 -03,GHI789,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 12:00:00 -0300 -03,null,null,null
10,44,456789012,true,false,4567890123,2024-07-24 17:21:48.89283 -0300 -03,JKL012,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 13:00:00 -0300 -03,null,null,null
11,55,567890123,true,false,5678901234,2024-07-24 17:21:48.89283 -0300 -03,MNO345,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 14:00:00 -0300 -03,null,null,null
12,66,678901234,true,false,6789012345,2024-07-24 17:21:48.89283 -0300 -03,PQR678,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 15:00:00 -0300 -03,null,null,null
13,77,789012345,true,false,7890123456,2024-07-24 17:21:48.89283 -0300 -03,STU901,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 16:00:00 -0300 -03,null,null,null
14,88,890123456,true,false,8901234567,2024-07-24 17:21:48.89283 -0300 -03,VWX234,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 17:00:00 -0300 -03,null,null,null
15,99,901234567,true,false,9012345678,2024-07-24 17:21:48.89283 -0300 -03,YZA567,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 18:00:00 -0300 -03,null,null,null
17,12,981832641,true,false,987654321,2024-06-26 16:41:52.571 -0300 -03,0123456789,null,,false,,2024-06-26 16:41:52.571 -0300 -03,2024-06-26 16:41:52.571 -0300 -03,null,null,null

```

#### POST /simfonia/api/numerostelefonicos/ {#post-apicreatenumero}

##### Descrição

Cria um novo número telefônico.

##### Exemplo de Requisição

```json
{
    "CodArea": 22,
    "Numero": 234567890,
    "Utilizavel": true,
    "PortadoIn": false,
    "PortadoInOperadora": "",
    "PortadoInDate": "2024-07-24T17:21:48.89283-03:00",
    "CodigoCNL": "DEF456",
    "CongeladoAte": null,
    "ExternalID": 2345678901,
    "PortadoOut": false,
    "PortadoOutOperadora": "",
    "PortadoOutDate": "2024-07-24T17:21:48.89283-03:00",
    "DataCriacao": "2024-06-24T11:00:00-03:00",
    "SimCardID": null,
    "SimCard": null,
    "PortadoInOperadoraID": null,
    "PortadoInOperadoraObj": null,
    "PortadoOutOperadoraID": null,
    "PortadoOutOperadoraObj": null
}
```

#### Exemplo de Resposta

```json
{
    "ID": 8,
    "CodArea": 22,
    "Numero": 234567890,
    "Utilizavel": true,
    "PortadoIn": false,
    "PortadoInOperadora": "",
    "PortadoInDate": "2024-07-24T17:21:48.89283-03:00",
    "CodigoCNL": "DEF456",
    "CongeladoAte": null,
    "ExternalID": 2345678901,
    "PortadoOut": false,
    "PortadoOutOperadora": "",
    "PortadoOutDate": "2024-07-24T17:21:48.89283-03:00",
    "DataCriacao": "2024-06-24T11:00:00-03:00",
    "SimCardID": null,
    "SimCard": null,
    "PortadoInOperadoraID": null,
    "PortadoInOperadoraObj": null,
    "PortadoOutOperadoraID": null,
    "PortadoOutOperadoraObj": null
}
```

#### POST /simfonia/api/numerostelefonicos/csv {#post-apicreatenumerocsv}

##### Descrição

Cria novos números telefônicos em lote, baseados em um arquivo `.csv`.

##### Exemplo de Requisição

```csv
    CodArea,Numero,Utilizavel,PortadoIn,PortadoInOperadora,PortadoInDate,CodigoCNL,CongeladoAte,ExternalID,PortadoOut,PortadoOutOperadora,PortadoOutDate,DataCriacao,SimCardID,PortadoInOperadoraID,PortadoOutOperadoraID
11,123456789,true,false,1234567890,2024-07-24 17:21:48.89283 -0300 -03,ABC123,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 10:00:00 -0300 -03,null,null,null
22,234567890,true,false,2345678901,2024-07-24 17:21:48.89283 -0300 -03,DEF456,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 11:00:00 -0300 -03,null,null,null
33,345678901,true,false,3456789012,2024-07-24 17:21:48.89283 -0300 -03,GHI789,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 12:00:00 -0300 -03,null,null,null
44,456789012,true,false,4567890123,2024-07-24 17:21:48.89283 -0300 -03,JKL012,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 13:00:00 -0300 -03,null,null,null
55,567890123,true,false,5678901234,2024-07-24 17:21:48.89283 -0300 -03,MNO345,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 14:00:00 -0300 -03,null,null,null
66,678901234,true,false,6789012345,2024-07-24 17:21:48.89283 -0300 -03,PQR678,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 15:00:00 -0300 -03,null,null,null
77,789012345,true,false,7890123456,2024-07-24 17:21:48.89283 -0300 -03,STU901,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 16:00:00 -0300 -03,null,null,null
88,890123456,true,false,8901234567,2024-07-24 17:21:48.89283 -0300 -03,VWX234,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 17:00:00 -0300 -03,null,null,null
99,901234567,true,false,9012345678,2024-07-24 17:21:48.89283 -0300 -03,YZA567,null,,false,,2024-07-24 17:21:48.89283 -0300 -03,2024-06-24 18:00:00 -0300 -03,null,null,null
12,981832641,true,false,987654321,2024-06-26 16:41:52.571 -0300 -03,0123456789,null,,false,,2024-06-26 16:41:52.571 -0300 -03,2024-06-26 16:41:52.571 -0300 -03,null,null,null
```

```multipart-form
    csv: *arquivo*
```

##### Exemplo de Resposta

```json
{
    "mensagem": "Arquivo carregado e processado com sucesso!"
}
```

#### PUT /simfonia/api/numerostelefonicos/:id {#put-apiupdatenumero}

##### Descrição

Atualiza um usuário existente pelo ID.

##### Exemplo de Requisição

```http
    PUT /simfonia/api/numerostelefonicos/:id
```

```json
{
    "CodArea": 22,
    "Numero": 234567890,
    "Utilizavel": true,
    "PortadoIn": false,
    "PortadoInOperadora": "",
    "PortadoInDate": "2024-07-24T17:21:48.89283-03:00",
    "CodigoCNL": "DEF456",
    "CongeladoAte": null,
    "ExternalID": 2345678901,
    "PortadoOut": false,
    "PortadoOutOperadora": "",
    "PortadoOutDate": "2024-07-24T17:21:48.89283-03:00",
    "DataCriacao": "2024-06-24T11:00:00-03:00",
    "SimCardID": null,
    "SimCard": null,
    "PortadoInOperadoraID": null,
    "PortadoInOperadoraObj": null,
    "PortadoOutOperadoraID": null,
    "PortadoOutOperadoraObj": null
}
```

##### Exemplo de Resposta

```json
{
    "ID": 8,
    "CodArea": 22,
    "Numero": 234567890,
    "Utilizavel": true,
    "PortadoIn": false,
    "PortadoInOperadora": "",
    "PortadoInDate": "2024-07-24T17:21:48.89283-03:00",
    "CodigoCNL": "DEF456",
    "CongeladoAte": null,
    "ExternalID": 2345678901,
    "PortadoOut": false,
    "PortadoOutOperadora": "",
    "PortadoOutDate": "2024-07-24T17:21:48.89283-03:00",
    "DataCriacao": "2024-06-24T11:00:00-03:00",
    "SimCardID": null,
    "SimCard": null,
    "PortadoInOperadoraID": null,
    "PortadoInOperadoraObj": null,
    "PortadoOutOperadoraID": null,
    "PortadoOutOperadoraObj": null
}
```

#### DELETE /simfonia/api/numerostelefonicos/:id {#delete-apideletenumeroid}

##### Descrição

Exclui um número telefônico pelo ID.

##### Exemplo de Requisição

```http
DELETE /simfonia/api/numerostelefonicos/8
```

##### Exemplo de Resposta

```json
{
    "deletado": true
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
- **FindUsuarioByEmail(email string)**: Busca um usuário pelo e-mail.
- **FindAllUsuarios()**: Busca por todos os usuários da tabela.
- **CreateUsuario(usuario model.Usuario)**: Cria um novo usuário.
- **UpdateUsuario(usuarioRecebido model.Usuario, id uint64)**: Atualiza um usuário existente.
- **DeleteUsuarioById(id uint64)**: Exclui um usuário pelo ID.
- **DeleteAllUsuarios()**: Exclui todos os usuários da tabela.
- **verificarSeEmailEmUso(email string)**: Verifica se o e-mail enviado como parâmetro já está em uso por algum usuário.

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