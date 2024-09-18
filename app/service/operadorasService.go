package service

import (
	"context"
	"errors"
	"pixelnest/app/helpers"
	"pixelnest/app/model/erros"
	"pixelnest/app/model/grpc"
	"pixelnest/app/model/repositories"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	"strings"
	"time"

	"google.golang.org/grpc/codes"
)

// Estrutura do serviço de Operadora, contendo o repositório necessário
type OperadoraService struct {
	operadorarepository repositories.OperadoraRepository
}

// Função para criar uma nova instância de OperadoraService com o repositório necessário
func NewOperadoraService(operadorarepository repositories.OperadoraRepository) *OperadoraService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &OperadoraService{
		operadorarepository: operadorarepository,
	}
}

// Função para buscar uma operadora pelo ID
func (operadoraService *OperadoraService) FindOperadoraById(context context.Context, id int32) (*grpc.Operadora, erros.ErroStatus) {
	operadora, err := operadoraService.operadorarepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrada nenhuma operadora, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.Operadora{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Operadora não encontrada"),
			}
		}

		return &grpc.Operadora{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TOperadoraToPb(operadora), erros.ErroStatus{}
}

// Função para buscar uma operadora pelo nome
func (operadoraService *OperadoraService) FindOperadoraByNome(context context.Context, nome string) (*grpc.Operadora, erros.ErroStatus) {
	operadora, err := operadoraService.operadorarepository.FindByNome(context, nome)
	if err != nil {
		// Caso não seja encontrada nenhuma operadora, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.Operadora{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Operadora não encontrada"),
			}
		}

		return &grpc.Operadora{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TOperadoraToPb(operadora), erros.ErroStatus{}
}

// Função para buscar uma operadora pela abreviação
func (operadoraService *OperadoraService) FindOperadoraByAbreviacao(context context.Context, abreviacao string) (*grpc.Operadora, erros.ErroStatus) {
	operadora, err := operadoraService.operadorarepository.FindByAbreviacao(context, abreviacao)
	if err != nil {
		// Caso não seja encontrada nenhuma operadora, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.Operadora{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Operadora não encontrada"),
			}
		}

		return &grpc.Operadora{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TOperadoraToPb(operadora), erros.ErroStatus{}
}

// Função para buscar todas as operadoras
func (operadoraService *OperadoraService) FindAllOperadoras(context context.Context) ([]*grpc.Operadora, erros.ErroStatus) {
	tOperadoras, err := operadoraService.operadorarepository.FindAll(context)
	if err != nil {
		return []*grpc.Operadora{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrada nenhuma operadora, retorna code NotFound
	if len(tOperadoras) == 0 {
		return []*grpc.Operadora{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma operadora encontrada"),
		}
	}

	var operadoras []*grpc.Operadora
	for _, operadora := range tOperadoras {
		operadoras = append(operadoras, helpers.TOperadoraToPb(operadora))
	}

	return operadoras, erros.ErroStatus{}
}

// Função para criar uma nova operadora
func (operadoraService *OperadoraService) CreateOperadora(context context.Context, operadora *grpc.Operadora) (*grpc.Operadora, erros.ErroStatus) {
	// Cria o objeto CreateOperadoraParams gerado pelo sqlc para gravação no banco de dados
	operadoraCreate := db.CreateOperadoraParams{
		Nome:       operadora.GetNome(),
		Abreviacao: operadora.GetAbreviacao(),
	}

	// Cria a operadora no repositório
	operadoraCriada, err := operadoraService.operadorarepository.Create(context, operadoraCreate)
	if err != nil {
		// Se ocorrer algum dos erros contendo as chaves especificadas abaixo, significa que ela já existe no banco
		// Portanto, retorna code AlreadyExists
		if strings.Contains(err.Error(), "t_operadoras_nome_key") {
			return &grpc.Operadora{}, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe operadora com esse nome"),
			}
		}

		if strings.Contains(err.Error(), "t_operadoras_abreviacao_key") {
			return &grpc.Operadora{}, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe operadora com essa abreviação"),
			}
		}

		return &grpc.Operadora{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TOperadoraToPb(operadoraCriada), erros.ErroStatus{}
}

// Função para atualizar uma operadora
func (operadoraService *OperadoraService) UpdateOperadora(context context.Context, operadoraRecebida *grpc.Operadora, operadoraAntiga *grpc.Operadora) (*grpc.Operadora, erros.ErroStatus) {
	// Verifica se o nome foi modificado e, se sim, verifica se já existe outro registro com o mesmo nome
	// Em caso positivo, retorna code AlreadyExists
	if operadoraAntiga.GetNome() != operadoraRecebida.GetNome() {
		_, erroBuscaPreExistente := operadoraService.FindOperadoraByNome(context, operadoraRecebida.GetNome())
		if erroBuscaPreExistente.Erro == nil {
			return nil, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe operadora com o nome enviado"),
			}
		}
	}

	// Verifica se a abreviação foi modificada e, se sim, verifica se já existe outro registro com a mesma abreviação
	// Em caso positivo, retorna code AlreadyExists
	if operadoraAntiga.GetAbreviacao() != operadoraRecebida.GetAbreviacao() {
		_, erroBuscaPreExistente := operadoraService.FindOperadoraByAbreviacao(context, operadoraRecebida.GetAbreviacao())
		if erroBuscaPreExistente.Erro == nil {
			return nil, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe operadora com a abreviação enviada"),
			}
		}
	}

	// Cria o objeto UpdateOperadoraParams gerado pelo sqlc para gravação no banco de dados
	operadoraUpdate := db.UpdateOperadoraParams{
		Nome:       operadoraRecebida.GetNome(),
		Abreviacao: operadoraRecebida.GetAbreviacao(),
		ID:         operadoraRecebida.GetID(),
	}

	// Salva a operadora atualizada no repositório
	operadoraAtualizada, erroSalvamento := operadoraService.operadorarepository.Update(context, operadoraUpdate)
	if erroSalvamento != nil {
		return &grpc.Operadora{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroSalvamento,
		}
	}

	return helpers.TOperadoraToPb(operadoraAtualizada), erros.ErroStatus{}
}

// Função para deletar uma operadora pelo ID
func (operadoraService *OperadoraService) DeleteOperadoraById(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Deleta a operadora no repositório
	deletado, err := operadoraService.operadorarepository.Delete(context, id)
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo deletados indica o número de linhas deletadas. Se for 0, nenhuma cidade foi deletada, pois não existia
	if deletado == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Operadora não encontrada"),
		}
	}

	return true, erros.ErroStatus{}
}
