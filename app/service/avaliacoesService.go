package service

import (
	"context"
	"errors"
	"pixelnest/app/helpers"
	"pixelnest/app/model/erros"
	pb "pixelnest/app/model/grpc"
	"pixelnest/app/model/repositories"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
)

// Estrutura de serviço para gerenciar operações relacionadas às avaliacoes
type AvaliacaoService struct {
	avaliacaoRepository repositories.AvaliacaoRepository
	usuarioRepository   repositories.UsuarioRepository
	jogoRepository      repositories.JogoRepository
	produtoRepository   repositories.ProdutoRepository
}

// Função para criar uma nova instância de AvaliacaoService com o repositório necessário
func NewAvaliacaoService(avaliacaoRepository repositories.AvaliacaoRepository,
	usuarioRepository repositories.UsuarioRepository,
	jogoRepository repositories.JogoRepository,
	produtoRepository repositories.ProdutoRepository) *AvaliacaoService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &AvaliacaoService{
		avaliacaoRepository: avaliacaoRepository,
		usuarioRepository:   usuarioRepository,
		jogoRepository:      jogoRepository,
		produtoRepository:   produtoRepository,
	}
}

// Função para buscar uma avaliacao pelo ID
func (avaliacaoService *AvaliacaoService) FindAvaliacaoById(context context.Context, id int32) (*pb.Avaliacao, erros.ErroStatus) {
	// Busca a avaliacao no repositório pelo ID
	avaliacao, err := avaliacaoService.avaliacaoRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrada nenhuma avaliacao, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Avaliacao{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Avaliacao não encontrada"),
			}
		}

		return &pb.Avaliacao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return avaliacaoService.montarObjCompleto(context, avaliacao)
}

// Função para buscar uma avaliacao pelo nome
func (avaliacaoService *AvaliacaoService) FindAvaliacaoByUsuario(context context.Context, id int32) ([]*pb.Avaliacao, erros.ErroStatus) {
	// Busca a avaliacoes no repositório pelo nome
	avaliacoes, err := avaliacaoService.avaliacaoRepository.FindByUsuario(context, id)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrada nenhuma avaliacao, retorna code NotFound
	if len(avaliacoes) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma avaliacao encontrada"),
		}
	}

	var pbAvaliacoes []*pb.Avaliacao
	for _, avaliacao := range avaliacoes {
		avaliacaoGRPC, erroMontagem := avaliacaoService.montarObjCompleto(context, avaliacao)
		if erroMontagem.Erro != nil {
			return nil, erroMontagem
		}

		pbAvaliacoes = append(pbAvaliacoes, avaliacaoGRPC)
	}

	return pbAvaliacoes, erros.ErroStatus{}
}

// Função para buscar uma avaliacao pelo nome
func (avaliacaoService *AvaliacaoService) FindAvaliacaoByProduto(context context.Context, id int32) ([]*pb.Avaliacao, erros.ErroStatus) {
	// Busca a avaliacoes no repositório pelo nome
	avaliacoes, err := avaliacaoService.avaliacaoRepository.FindByProduto(context, id)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrada nenhuma avaliacao, retorna code NotFound
	if len(avaliacoes) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma avaliacao encontrada"),
		}
	}

	var pbAvaliacoes []*pb.Avaliacao
	for _, avaliacao := range avaliacoes {
		avaliacaoGRPC, erroMontagem := avaliacaoService.montarObjCompleto(context, avaliacao)
		if erroMontagem.Erro != nil {
			return nil, erroMontagem
		}

		pbAvaliacoes = append(pbAvaliacoes, avaliacaoGRPC)
	}

	return pbAvaliacoes, erros.ErroStatus{}
}

// Função para buscar uma avaliacao pelo nome
func (avaliacaoService *AvaliacaoService) FindAvaliacaoByJogo(context context.Context, id int32) ([]*pb.Avaliacao, erros.ErroStatus) {
	// Busca a avaliacoes no repositório pelo nome
	avaliacoes, err := avaliacaoService.avaliacaoRepository.FindByJogo(context, id)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrada nenhuma avaliacao, retorna code NotFound
	if len(avaliacoes) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma avaliacao encontrada"),
		}
	}

	var pbAvaliacoes []*pb.Avaliacao
	for _, avaliacao := range avaliacoes {
		avaliacaoGRPC, erroMontagem := avaliacaoService.montarObjCompleto(context, avaliacao)
		if erroMontagem.Erro != nil {
			return nil, erroMontagem
		}

		pbAvaliacoes = append(pbAvaliacoes, avaliacaoGRPC)
	}

	return pbAvaliacoes, erros.ErroStatus{}
}

// Função para buscar todas as avaliacoes
func (avaliacaoService *AvaliacaoService) FindAllAvaliacoes(context context.Context) ([]*pb.Avaliacao, erros.ErroStatus) {
	// Busca todas as avaliacoes no repositório
	avaliacoes, err := avaliacaoService.avaliacaoRepository.FindAll(context)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrada nenhuma avaliacao, retorna code NotFound
	if len(avaliacoes) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma avaliacao encontrada"),
		}
	}

	var pbAvaliacoes []*pb.Avaliacao
	for _, avaliacao := range avaliacoes {
		avaliacaoGRPC, erroMontagem := avaliacaoService.montarObjCompleto(context, avaliacao)
		if erroMontagem.Erro != nil {
			return nil, erroMontagem
		}

		pbAvaliacoes = append(pbAvaliacoes, avaliacaoGRPC)
	}

	return pbAvaliacoes, erros.ErroStatus{}
}

// Função para criar uma nova avaliacao
func (avaliacaoService *AvaliacaoService) CreateAvaliacao(context context.Context, avaliacao *pb.Avaliacao) (*pb.Avaliacao, erros.ErroStatus) {
	// Cria o objeto CreateAvaliacaoParams gerado pelo sqlc para gravação no banco de dados
	avaliacaoCreate := db.CreateAvaliacaoParams{
		UsuarioID: avaliacao.GetUsuario().GetID(),
		ProdutoID: pgtype.Int4{Int32: avaliacao.GetProduto().GetID(), Valid: true},
		JogoID:    pgtype.Int4{Int32: avaliacao.GetJogo().GetID(), Valid: true},
		Nota:      pgtype.Int4{Int32: avaliacao.GetNota(), Valid: true},
		Avaliacao: pgtype.Text{String: avaliacao.GetAvaliacao(), Valid: true},
	}

	// Cria a avaliacao no repositório
	avaliacaoCriada, err := avaliacaoService.avaliacaoRepository.Create(context, avaliacaoCreate)
	if err != nil {
		return &pb.Avaliacao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return avaliacaoService.montarObjCompleto(context, avaliacaoCriada)
}

// Função para atualizar uma avaliacao existente
func (avaliacaoService *AvaliacaoService) UpdateAvaliacao(context context.Context, avaliacaoRecebida *pb.Avaliacao, avaliacaoBanco *pb.Avaliacao) (*pb.Avaliacao, erros.ErroStatus) {
	// Cria o objeto UpdateAvaliacaoParams gerado pelo sqlc para gravação no banco de dados
	avaliacaoUpdate := db.UpdateAvaliacaoParams{
		UsuarioID: avaliacaoBanco.GetUsuario().GetID(),
		ProdutoID: pgtype.Int4{Int32: avaliacaoBanco.GetProduto().GetID(), Valid: true},
		JogoID:    pgtype.Int4{Int32: avaliacaoBanco.GetJogo().GetID(), Valid: true},
		Nota:      pgtype.Int4{Int32: avaliacaoRecebida.GetNota(), Valid: true},
		Avaliacao: pgtype.Text{String: avaliacaoRecebida.GetAvaliacao(), Valid: true},
	}

	// Salva a avaliacao atualizada no repositório
	avaliacaoAtualizada, errSalvamento := avaliacaoService.avaliacaoRepository.Update(context, avaliacaoUpdate)
	if errSalvamento != nil {
		return &pb.Avaliacao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errSalvamento,
		}
	}

	return avaliacaoService.montarObjCompleto(context, avaliacaoAtualizada)
}

// Função para deletar uma avaliacao pelo ID
func (avaliacaoService *AvaliacaoService) DeleteAvaliacaoById(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Deleta a avaliacao no repositório pelo ID
	deletados, err := avaliacaoService.avaliacaoRepository.Delete(context, id)

	// Caso ocorra erro na deleção
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo deletados indica o número de linhas deletadas. Se for 0, nenhuma avaliacao foi deletada, pois não existia
	if deletados == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma avaliacao encontrada"),
		}
	}

	return true, erros.ErroStatus{}
}

func (avaliacaoService *AvaliacaoService) montarObjCompleto(context context.Context, avaliacao db.TAvaliaco) (*pb.Avaliacao, erros.ErroStatus) {
	// Faz a montagem do objeto operadora de entrada de portabilidade, caso necessário
	usuario := new(pb.Usuario)
	if avaliacao.UsuarioID != 0 {
		tUsuario, err := avaliacaoService.usuarioRepository.FindByID(context, avaliacao.UsuarioID)
		if err != nil {
			return &pb.Avaliacao{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
		usuario = helpers.TUsuarioToPb(tUsuario)
	}

	// Faz a montagem do objeto operadora de saída de portabilidade, caso necessário
	produto := new(pb.Produto)
	if avaliacao.ProdutoID.Int32 != 0 {
		tProduto, err := avaliacaoService.produtoRepository.FindByID(context, avaliacao.ProdutoID.Int32)
		if err != nil {
			return &pb.Avaliacao{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
		produto = helpers.TProdutoToPb(tProduto)
	}

	// Faz a montagem do objeto SimCard, caso necessário
	jogo := new(pb.Jogo)
	if avaliacao.JogoID.Int32 != 0 {
		tJogo, err := avaliacaoService.jogoRepository.FindByID(context, avaliacao.JogoID.Int32)
		if err != nil {
			return &pb.Avaliacao{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		jogo = helpers.TJogoToPb(tJogo)
	}

	// Passando todos os objetos completinhos para o helper que irá converter o objeto completo para mim
	return helpers.TAvaliacaoToPb(avaliacao, usuario, produto, jogo), erros.ErroStatus{}
}
