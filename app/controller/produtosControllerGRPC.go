package controller

import (
	"context"
	"pixelnest/app/configuration/logger"
	"pixelnest/app/controller/middlewares"
	"pixelnest/app/service"

	pb "pixelnest/app/model/grpc" // Importa o pacote gerado pelos arquivos .proto

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Implementação do servidor
type ProdutosServer struct {
	pb.UnimplementedProdutosServer
	produtoService      *service.ProdutoService
	permissaoMiddleware *middlewares.PermissoesMiddleware
}

func NewProdutosServer(produtoService *service.ProdutoService, permissaoMiddleware *middlewares.PermissoesMiddleware) *ProdutosServer {
	return &ProdutosServer{
		produtoService:      produtoService,
		permissaoMiddleware: permissaoMiddleware,
	}
}

func (produtoServer *ProdutosServer) mustEmbedUnimplementedProdutosServer() {}

// Função para buscar por todos os produtos existentes no banco de dados
func (produtoServer *ProdutosServer) FindAllProdutos(context context.Context, req *pb.RequestVazio) (*pb.ListaProdutos, error) {
	usuarioSolicitante, retornoMiddleware := produtoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-produtos")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaProdutos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	produtos, erroService := produtoServer.produtoService.FindAllProdutos(context)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os produtos",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaProdutos{Produtos: produtos}, nil
}

// Função para buscar por um produto pelo ID
func (produtoServer *ProdutosServer) FindProdutoById(context context.Context, req *pb.RequestId) (*pb.Produto, error) {
	usuarioSolicitante, retornoMiddleware := produtoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-produto-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.Produto{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	produto, erroService := produtoServer.produtoService.FindProdutoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada um produto pelo ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("produto", produto),
	)

	return produto, nil
}

// Função para buscar por um produto pelo nome
func (produtoServer *ProdutosServer) FindProdutoByNome(context context.Context, req *pb.RequestNome) (*pb.ListaProdutos, error) {
	usuarioSolicitante, retornoMiddleware := produtoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-produto-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaProdutos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	produto, erroService := produtoServer.produtoService.FindProdutoByNome(context, req.GetNome())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada um produto pelo nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("produto", produto),
	)

	return &pb.ListaProdutos{Produtos: produto}, nil
}

// Função para buscar por um produto pelo nome
func (produtoServer *ProdutosServer) FindProdutoByGenero(context context.Context, req *pb.RequestNome) (*pb.ListaProdutos, error) {
	usuarioSolicitante, retornoMiddleware := produtoServer.permissaoMiddleware.PermissaoMiddleware(context, "get-produto-by-nome")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaProdutos{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	produto, erroService := produtoServer.produtoService.FindProdutoByGenero(context, req.GetNome())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscada um produto pelo nome",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("produto", produto),
	)

	return &pb.ListaProdutos{Produtos: produto}, nil
}

// Função para criar um nova produto
func (produtoServer *ProdutosServer) CreateProduto(context context.Context, req *pb.Produto) (*pb.Produto, error) {
	usuarioSolicitante, retornoMiddleware := produtoServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-produto")
	if retornoMiddleware.Erro != nil {
		return &pb.Produto{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	produtoCriada, erroService := produtoServer.produtoService.CreateProduto(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criada um nova produto",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("produto", produtoCriada),
	)

	return produtoCriada, nil
}

// Função para atualizar um produto já existente no banco
func (produtoServer *ProdutosServer) UpdateProduto(context context.Context, req *pb.Produto) (*pb.Produto, error) {
	usuarioSolicitante, retornoMiddleware := produtoServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-produto")
	if retornoMiddleware.Erro != nil {
		return &pb.Produto{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	// produtoAntigo, erroService := produtoServer.produtoService.FindProdutoById(context, req.GetID())
	// if erroService.Erro != nil {
	// 	logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
	// 	return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	// }

	produtoCriada, erroService := produtoServer.produtoService.UpdateProduto(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Atualizado um produto existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		// zap.Any("produto", produtoAntigo),
		zap.Any("produtoAtualizada", produtoCriada),
	)

	return produtoCriada, nil
}

// Função para deletar um produto existente no banco
func (produtoServer *ProdutosServer) DeleteProduto(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := produtoServer.permissaoMiddleware.PermissaoMiddleware(context, "delete-produto-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID enviado não é válido ou não foi enviado")
	}

	produto, erroService := produtoServer.produtoService.FindProdutoById(context, req.GetID())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return nil, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	deletado, erroService := produtoServer.produtoService.DeleteProdutoById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	if !deletado {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, "Não existe produto com o ID enviado")
	}

	logger.Logger.Info("Deletada um produto",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("produto", produto),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}
