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

// Estrutura de serviço para gerenciar operações relacionadas às cidades
type CidadeService struct {
	cidadeRepository repositories.CidadeRepository
}

// Função para criar uma nova instância de CidadeService com o repositório necessário
func NewCidadeService(cidadeRepository repositories.CidadeRepository) *CidadeService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &CidadeService{
		cidadeRepository: cidadeRepository,
	}
}

// Função para buscar uma cidade pelo ID
func (cidadeService *CidadeService) FindCidadeById(context context.Context, id int32) (*pb.Cidade, erros.ErroStatus) {
	// Busca a cidade no repositório pelo ID
	cidade, err := cidadeService.cidadeRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrada nenhuma cidade, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Cidade{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Cidade não encontrada"),
			}
		}

		return &pb.Cidade{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TCidadeToPb(cidade), erros.ErroStatus{}
}

// Função para buscar uma cidade pelo nome
func (cidadeService *CidadeService) FindCidadeByNome(context context.Context, nome string) (*pb.Cidade, erros.ErroStatus) {
	// Busca a cidade no repositório pelo nome
	cidade, err := cidadeService.cidadeRepository.FindByNome(context, nome)
	if err != nil {
		// Caso não seja encontrada nenhuma cidade, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Cidade{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Cidade não encontrada"),
			}
		}

		return &pb.Cidade{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TCidadeToPb(cidade), erros.ErroStatus{}
}

// Função para buscar uma cidade pela UF
func (cidadeService *CidadeService) FindCidadeByUF(context context.Context, nome string) ([]*pb.Cidade, erros.ErroStatus) {
	// Busca as cidades no repositório pela UF
	cidades, err := cidadeService.cidadeRepository.FindByUF(context, nome)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrada nenhuma cidade, retorna code NotFound
	if len(cidades) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma cidade encontrada"),
		}
	}

	var pbCidades []*pb.Cidade
	for _, cidade := range cidades {
		pbCidades = append(pbCidades, helpers.TCidadeToPb(cidade))
	}

	return pbCidades, erros.ErroStatus{}
}

// Função para buscar uma cidade pelo nome
func (cidadeService *CidadeService) FindCidadeByCodIbge(context context.Context, codIbge int32) (*pb.Cidade, erros.ErroStatus) {
	// Busca a cidade no repositório pelo nome
	cidade, err := cidadeService.cidadeRepository.FindByCodIbge(context, codIbge)
	if err != nil {
		// Caso não seja encontrada nenhuma cidade, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Cidade{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Cidade não encontrada"),
			}
		}

		return &pb.Cidade{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TCidadeToPb(cidade), erros.ErroStatus{}
}

// Função para buscar todas as cidades
func (cidadeService *CidadeService) FindAllCidades(context context.Context) ([]*pb.Cidade, erros.ErroStatus) {
	// Busca todas as cidades no repositório
	cidades, err := cidadeService.cidadeRepository.FindAll(context)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrada nenhuma cidade, retorna code NotFound
	if len(cidades) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma cidade encontrada"),
		}
	}

	var pbCidades []*pb.Cidade
	for _, cidade := range cidades {
		pbCidades = append(pbCidades, helpers.TCidadeToPb(cidade))
	}

	return pbCidades, erros.ErroStatus{}
}

// Função para criar uma nova cidade
func (cidadeService *CidadeService) CreateCidade(context context.Context, cidade *pb.Cidade) (*pb.Cidade, erros.ErroStatus) {
	// Busca uma cidade pelo nome enviado para verificar a prévia existência dela
	// Em caso positivo, retorna code AlreadyExists
	_, erroBuscaPreExistente := cidadeService.FindCidadeByNome(context, cidade.GetNome())
	if erroBuscaPreExistente.Erro == nil {
		return nil, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("Já existe cidade com o nome enviado"),
		}
	}

	// Cria o objeto CreateCidadeParams gerado pelo sqlc para gravação no banco de dados
	cidadeCreate := db.CreateCidadeParams{
		Uuid:    pgtype.Text{String: cidade.GetUUID(), Valid: true},
		Nome:    cidade.GetNome(),
		CodIbge: cidade.GetCodIBGE(),
		Uf:      cidade.GetUF(),
		CodArea: cidade.GetCodArea(),
	}

	// Cria a cidade no repositório
	cidadeCriada, err := cidadeService.cidadeRepository.Create(context, cidadeCreate)
	if err != nil {
		return &pb.Cidade{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TCidadeToPb(cidadeCriada), erros.ErroStatus{}
}

// Função para atualizar uma cidade existente
func (cidadeService *CidadeService) UpdateCidade(context context.Context, cidadeRecebida *pb.Cidade, cidadeBanco *pb.Cidade) (*pb.Cidade, erros.ErroStatus) {
	// Verifica se o nome foi modificado e, se sim, verifica se já existe outro registro com o mesmo nome
	// Em caso positivo, retorna code AlreadyExists
	if cidadeBanco.GetNome() != cidadeRecebida.GetNome() {
		_, erroBuscaPreExistente := cidadeService.FindCidadeByNome(context, cidadeRecebida.GetNome())
		if erroBuscaPreExistente.Erro == nil {
			return nil, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe cidade com o nome enviado"),
			}
		}
	}

	// Cria o objeto UpdateCidadeParams gerado pelo sqlc para gravação no banco de dados
	cidadeUpdate := db.UpdateCidadeParams{
		Uuid:    pgtype.Text{String: cidadeRecebida.GetUUID(), Valid: true},
		Nome:    cidadeRecebida.GetNome(),
		CodIbge: cidadeRecebida.GetCodIBGE(),
		Uf:      cidadeRecebida.GetUF(),
		CodArea: cidadeRecebida.GetCodArea(),
		ID:      cidadeRecebida.GetID(),
	}

	// Salva a cidade atualizada no repositório
	cidadeAtualizada, errSalvamento := cidadeService.cidadeRepository.Update(context, cidadeUpdate)
	if errSalvamento != nil {
		return &pb.Cidade{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errSalvamento,
		}
	}

	return helpers.TCidadeToPb(cidadeAtualizada), erros.ErroStatus{}
}

// Função para deletar uma cidade pelo ID
func (cidadeService *CidadeService) DeleteCidadeById(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Deleta a cidade no repositório pelo ID
	deletados, err := cidadeService.cidadeRepository.Delete(context, id)

	// Caso ocorra erro na deleção
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
			Erro:   errors.New("Nenhuma cidade encontrada"),
		}
	}

	return true, erros.ErroStatus{}
}
