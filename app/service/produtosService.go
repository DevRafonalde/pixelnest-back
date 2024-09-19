package service

import (
	"context"
	"errors"
	"net/http"
	"pixelnest/app/helpers"
	"pixelnest/app/model/erros"
	"pixelnest/app/model/grpc"
	pb "pixelnest/app/model/grpc"
	"pixelnest/app/model/repositories"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
)

// Estrutura do serviço de Produto
type ProdutoService struct {
	produtoRepository repositories.ProdutoRepository
	layout            string
}

// Função para criar uma nova instância de ProdutoService
func NewProdutoService(produtoRepository repositories.ProdutoRepository) *ProdutoService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &ProdutoService{
		produtoRepository: produtoRepository,
		layout:            "2006-01-02", // Define o formato de data/hora
	}
}

// Função para buscar um Produto pelo ID
func (produtoService *ProdutoService) FindProdutoById(context context.Context, id int32) (*grpc.Produto, erros.ErroStatus) {
	produto, err := produtoService.produtoRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum Produto, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.Produto{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Produto não encontrado"),
			}
		}

		return &grpc.Produto{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TProdutoToPb(produto), erros.ErroStatus{}
}

// Função para buscar um Produto pelo ICCID
func (produtoService *ProdutoService) FindProdutoByNome(context context.Context, nome string) ([]*grpc.Produto, erros.ErroStatus) {
	produtos, err := produtoService.produtoRepository.FindByNome(context, nome)
	if err != nil {
		return []*grpc.Produto{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum Produto, retorna code NotFound
	if len(produtos) == 0 {
		return []*grpc.Produto{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum Produto encontrado"),
		}
	}

	pbProdutos := []*grpc.Produto{}

	for _, produto := range produtos {
		pbProdutos = append(pbProdutos, helpers.TProdutoToPb(produto))
	}

	return pbProdutos, erros.ErroStatus{}
}

// Função para buscar um Produto pelo ICCID
func (produtoService *ProdutoService) FindProdutoByGenero(context context.Context, nome string) ([]*grpc.Produto, erros.ErroStatus) {
	produtos, err := produtoService.produtoRepository.FindByGenero(context, nome)
	if err != nil {
		return []*grpc.Produto{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum Produto, retorna code NotFound
	if len(produtos) == 0 {
		return []*grpc.Produto{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum Produto encontrado"),
		}
	}

	pbProdutos := []*grpc.Produto{}

	for _, produto := range produtos {
		pbProdutos = append(pbProdutos, helpers.TProdutoToPb(produto))
	}

	return pbProdutos, erros.ErroStatus{}
}

// Função para buscar todos os Produtos
func (produtoService *ProdutoService) FindAllProdutos(context context.Context) ([]*grpc.Produto, erros.ErroStatus) {
	produtos, err := produtoService.produtoRepository.FindAll(context)
	if err != nil {
		return []*grpc.Produto{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum Produto, retorna code NotFound
	if len(produtos) == 0 {
		return []*grpc.Produto{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum Produto encontrado"),
		}
	}

	pbProdutos := []*grpc.Produto{}

	for _, produto := range produtos {
		pbProdutos = append(pbProdutos, helpers.TProdutoToPb(produto))
	}

	return pbProdutos, erros.ErroStatus{}
}

// Função para criar um novo Produto
func (produtoService *ProdutoService) CreateProduto(context context.Context, produtoRecebido *grpc.Produto) (*grpc.Produto, erros.ErroStatus) {
	// Busca um jogo pelo nome enviado para verificar a prévia existência dela
	// Em caso positivo, retorna code AlreadyExists
	tProdutos, erroService := produtoService.FindProdutoByNome(context, produtoRecebido.GetNome())
	if erroService.Erro == nil || len(tProdutos) != 0 {
		return &pb.Produto{}, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("Já existe produto com esse nome"),
		}
	}

	// Cria o objeto CreateProdutoParams gerado pelo sqlc para gravação no banco de dados
	produtoCreate := db.CreateProdutoParams{
		Nome:      produtoRecebido.GetNome(),
		Descricao: pgtype.Text{String: produtoRecebido.GetDescricao(), Valid: true},
		Genero:    pgtype.Text{String: produtoRecebido.GetGenero(), Valid: true},
	}

	produtoCriado, err := produtoService.produtoRepository.Create(context, produtoCreate)
	if err != nil {
		return &grpc.Produto{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TProdutoToPb(produtoCriado), erros.ErroStatus{}
}

// Função para atualizar um Produto existente
func (produtoService *ProdutoService) UpdateProduto(context context.Context, produtoRecebido *grpc.Produto) (*grpc.Produto, erros.ErroStatus) {
	produtoBanco, err := produtoService.produtoRepository.FindByID(context, produtoRecebido.ID)
	if err != nil {
		return &grpc.Produto{}, erros.ErroStatus{
			Status: http.StatusBadRequest,
			Erro:   err,
		}
	}

	if produtoBanco.Nome != produtoRecebido.GetNome() {
		// Busca um jogo pelo nome enviado para verificar a prévia existência dela
		// Em caso positivo, retorna code AlreadyExists
		tProdutos, erroService := produtoService.FindProdutoByNome(context, produtoRecebido.GetNome())
		if erroService.Erro == nil || len(tProdutos) != 0 {
			return &pb.Produto{}, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe produto com esse nome"),
			}
		}
	}

	// Cria o objeto UpdateProdutoParams gerado pelo sqlc para gravação no banco de dados
	produtoUpdate := db.UpdateProdutoParams{
		Nome:      produtoRecebido.GetNome(),
		Descricao: pgtype.Text{String: produtoRecebido.GetDescricao(), Valid: true},
		Genero:    pgtype.Text{String: produtoRecebido.GetGenero(), Valid: true},
	}

	// Salva o Produto atualizado no repositório
	produtoAtualizado, erroSalvamento := produtoService.produtoRepository.Update(context, produtoUpdate)
	if erroSalvamento != nil {
		// Mesmo caso do create
		if strings.Contains(erroSalvamento.Error(), "fk_t_simcard_estado_sim_cards") {
			return &grpc.Produto{}, erros.ErroStatus{
				Status: codes.InvalidArgument,
				Erro:   errors.New("Não existe estado de Produto com o id enviado"),
			}
		}

		return &grpc.Produto{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroSalvamento,
		}
	}

	return helpers.TProdutoToPb(produtoAtualizado), erros.ErroStatus{}
}

// Função para deletar um Produto pelo ID
func (produtoService *ProdutoService) DeleteProdutoById(context context.Context, id int32) (bool, erros.ErroStatus) {
	deletados, err := produtoService.produtoRepository.Delete(context, id)
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo deletados indica o número de linhas deletadas. Se for 0, nenhuma cidade foi deletada, pois não existia
	if deletados == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Status de Produto não encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}
