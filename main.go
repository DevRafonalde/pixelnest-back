package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"pixelnest/app/controller"
	"pixelnest/app/controller/middlewares"
	pb "pixelnest/app/model/grpc"
	"pixelnest/app/model/repositories"
	"pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	"pixelnest/app/service"
	"pixelnest/db"

	"google.golang.org/grpc"
)

func main() {
	// Configuração de chave secreta do token JWT
	if os.Getenv("GENERATE_KEY") == "true" {
		if err := GeraChaveSecreta(); err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao gerar a chave secreta: %v\n", err)
			os.Exit(1)
		}
	}

	chaveSecreta, err := os.ReadFile("./jwt/jwt_secret_key.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro ao ler a chave secreta: %v\n", err)
		os.Exit(1)
	}

	db := db.CreateConnection()
	defer db.Close()

	sqlcQueries := repositoryIMPL.New(db)

	context := context.Background()

	// Repositórios
	// Aqui são feitas as implementações das funções que realmente interagirão com o banco de dados
	avaliacaoRepository := repositories.NewAvaliacaoRepository(sqlcQueries)
	favoritoRepository := repositories.NewFavoritoRepository(sqlcQueries)
	jogoRepository := repositories.NewJogoRepository(sqlcQueries)
	parametroRepository := repositories.NewParametroRepository(sqlcQueries)
	perfilPermissaoRepository := repositories.NewPerfilPermissaoRepository(sqlcQueries)
	perfilRepository := repositories.NewPerfilRepository(sqlcQueries)
	permissaoRepository := repositories.NewPermissaoRepository(sqlcQueries)
	produtoRepository := repositories.NewProdutoRepository(sqlcQueries)
	usuarioJogoRepository := repositories.NewUsuarioJogoRepository(sqlcQueries)
	usuarioPerfilRepository := repositories.NewUsuarioPerfilRepository(sqlcQueries)
	usuarioRepository := repositories.NewUsuarioRepository(sqlcQueries)

	// Serviços
	// Aqui é a camada da minha regra de negócio, todas as validações, modificações e adaptações dos dados são feitas aqui
	avaliacaoService := service.NewAvaliacaoService(avaliacaoRepository, usuarioRepository, jogoRepository, produtoRepository)
	favoritoService := service.NewFavoritoService(favoritoRepository, usuarioRepository, jogoRepository, produtoRepository)
	jogoService := service.NewJogoService(jogoRepository, usuarioJogoRepository, favoritoRepository)
	parametroService := service.NewParametroService(parametroRepository)
	perfilService := service.NewPerfilService(perfilPermissaoRepository, permissaoRepository, usuarioPerfilRepository, perfilRepository, usuarioRepository)
	permissaoService := service.NewPermissaoService(perfilPermissaoRepository, permissaoRepository)
	produtoService := service.NewProdutoService(produtoRepository)
	usuarioService := service.NewUsuarioService(perfilRepository, usuarioPerfilRepository, usuarioRepository)

	// Esse trecho de código abaixo serve para que sejam definidos, nas variáveis de ambiente, todos os parâmetros da tabela de parâmetros
	parametros, erroService := parametroService.FindAllParametros(context)
	if erroService.Erro != nil { // Se houver erro na busca, retorna erro
		log.Fatalf("Falha ao buscar por todos os parâmetros %v", err.Error())
		os.Exit(1)
	}

	for _, parametro := range parametros {
		os.Setenv(parametro.GetNome(), parametro.GetValor())
	}

	// Middleware de acesso
	permissaoMiddleware := middlewares.NewPermissoesMiddleware(usuarioService, perfilService)
	// O "middleware" acima está sendo utilizado atualmente como uma chamada no início da execução de cada uma das funções dos servers/controllers
	// A função chama a função passando o contexto e a permissão específica daquele comando
	// Dessa forma, caso o usuário tenha as permissões necessárias, é retornado um erro nulo e continada a requisição
	// Caso o usuário não tenha as permissões necessárias, é abortada a requisição passando o erro que impediu esse acesso

	// Servers gRPC
	// Aqui é a camada de controle, aqui não temos regras de negócio, apenas a comunicação efetiva com o favorito
	avaliacoesServer := controller.NewAvaliacaosServer(avaliacaoService, permissaoMiddleware)
	favoritosServer := controller.NewFavoritosServer(favoritoService, permissaoMiddleware)
	jogosServer := controller.NewJogosServer(jogoService, permissaoMiddleware)
	parametrosServer := controller.NewParametrosServer(parametroService, permissaoMiddleware)
	perfisServer := controller.NewPerfisServer(perfilService, permissaoMiddleware)
	permissoesServer := controller.NewPermissoesServer(permissaoService, permissaoMiddleware)
	produtoServer := controller.NewProdutosServer(produtoService, permissaoMiddleware)
	usuarioServer := controller.NewUsuariosServer(usuarioService, permissaoMiddleware, chaveSecreta)

	// Criado o listener para a porta da aplicação
	lis, err := net.Listen("tcp", ":8601")
	if err != nil {
		log.Fatalf("Falha ao escutar: %v", err)
		return
	}

	// Carrega o certificado e a chave para TLS
	// creds, err := credentials.NewServerTLSFromFile("cert.pem", "key.pem")
	// if err != nil {
	// 	log.Fatalf("Falha ao carregar credenciais de TLS: %v", err)
	// 	return
	// }

	// Configurações do servidor gRPC com TLS
	serverGrpc := grpc.NewServer()

	// Registro dos servidores/controllers
	pb.RegisterAvaliacoesServer(serverGrpc, avaliacoesServer)
	pb.RegisterFavoritosServer(serverGrpc, favoritosServer)
	pb.RegisterJogosServer(serverGrpc, jogosServer)
	pb.RegisterParametrosServer(serverGrpc, parametrosServer)
	pb.RegisterPerfisServer(serverGrpc, perfisServer)
	pb.RegisterPermissoesServer(serverGrpc, permissoesServer)
	pb.RegisterProdutosServer(serverGrpc, produtoServer)
	pb.RegisterUsuariosServer(serverGrpc, usuarioServer)

	// Feitas as configurações e inicializações, vamos dar início ao processo de inicialização da aplicação

	// Criação do primeiro usuário (admin) baseado no json abaixo
	_, erroService = usuarioService.FindUsuarioById(context, 1)
	if erroService.Erro != nil {
		jsonFile, err := os.Open("./seedUsuarioAdmin.json")
		if err != nil {
			log.Fatalf("Falha ao criar o usuário admin: %v", err.Error())
			os.Exit(1)
		}

		defer jsonFile.Close()
		byteValueJson, _ := io.ReadAll(jsonFile)
		objUsuario := pb.UsuarioPerfis{}
		json.Unmarshal(byteValueJson, &objUsuario)
		_, erroCriacaoUsuario := usuarioService.CreateUsuario(context, &objUsuario)
		if erroCriacaoUsuario.Erro != nil {
			if erroService.Erro != nil { // Se houver erro na criação, retorna erro
				log.Fatalf("Falha ao criar o usuário admin: %v", err.Error())
				os.Exit(1)
			}
		}
	}

	log.Println("Iniciando servidor gRPC na porta 8601...")

	if err := serverGrpc.Serve(lis); err != nil {
		log.Fatalf("Falha ao iniciar o servidor gRPC: %v", err)
	}
}
