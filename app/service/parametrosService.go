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

// Estrutura do serviço de Parâmetro, contendo repositórios necessários
type ParametroService struct {
	parametroRepository repositories.ParametroRepository
}

// Função para criar uma nova instância de ParametroService com os repositórios necessários
func NewParametroService(parametroRepository repositories.ParametroRepository) *ParametroService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &ParametroService{
		parametroRepository: parametroRepository,
	}
}

// Função para buscar todas as permissões
func (parametroService *ParametroService) FindAllParametros(context context.Context) ([]*pb.Parametro, erros.ErroStatus) {
	tParametros, err := parametroService.parametroRepository.FindAll(context)
	if err != nil {
		return []*pb.Parametro{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum parâmetro, retorna code NotFound
	if len(tParametros) == 0 {
		return []*pb.Parametro{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum parâmetro encontrado"),
		}
	}

	var pbParametros []*pb.Parametro
	for _, parametro := range tParametros {
		pbParametros = append(pbParametros, helpers.TParametroToPb(parametro))
	}

	return pbParametros, erros.ErroStatus{}
}

// Função para buscar uma parâmetro pelo nome
func (parametroService *ParametroService) FindParametroByNome(context context.Context, nome string) (*pb.Parametro, erros.ErroStatus) {
	parametro, err := parametroService.parametroRepository.FindByNome(context, nome)
	if err != nil {
		// Caso não seja encontrado nenhum parâmetro, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Parametro{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Parâmetro não encontrado"),
			}
		}

		return &pb.Parametro{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TParametroToPb(parametro), erros.ErroStatus{}
}

// Função para buscar uma parâmetro pelo nome
func (parametroService *ParametroService) FindParametroById(context context.Context, id int32) (*pb.Parametro, erros.ErroStatus) {
	parametro, err := parametroService.parametroRepository.FindById(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum parâmetro, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Parametro{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Parâmetro não encontrado"),
			}
		}

		return &pb.Parametro{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TParametroToPb(parametro), erros.ErroStatus{}
}

// Função para criar uma nova parâmetro
func (parametroService *ParametroService) CreateParametro(context context.Context, parametro *pb.Parametro) (*pb.Parametro, erros.ErroStatus) {
	// Busca uma parâmetro pelo nome enviado para verificar a prévia existência dela
	// Em caso positivo, retorna code AlreadyExists
	if _, err := parametroService.FindParametroByNome(context, parametro.Nome); err.Erro == nil {
		return &pb.Parametro{}, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("Já existe parâmetro com esse nome"),
		}
	}

	// Cria o objeto CreateParametroParams gerado pelo sqlc para gravação no banco de dados
	parametroCreate := db.CreateParametroParams{
		Nome:      pgtype.Text{String: parametro.GetNome(), Valid: true},
		Descricao: pgtype.Text{String: parametro.GetDescricao(), Valid: true},
		Valor:     pgtype.Text{String: parametro.GetValor(), Valid: true},
	}

	// Cria a parâmetro no repositório
	parametroCriada, err := parametroService.parametroRepository.Create(context, parametroCreate)
	if err != nil {
		return &pb.Parametro{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TParametroToPb(parametroCriada), erros.ErroStatus{}
}

// Função para atualizar uma parâmetro existente
func (parametroService *ParametroService) UpdateParametro(context context.Context, parametroRecebido *pb.Parametro, parametroAntigo *pb.Parametro) (*pb.Parametro, erros.ErroStatus) {
	// Verifica se o nome foi modificado e, se sim, verifica se já existe outro registro com o mesmo nome
	// Em caso positivo, retorna code AlreadyExists
	if parametroAntigo.GetNome() != parametroRecebido.GetNome() {
		_, erroBuscaPreExistente := parametroService.FindParametroByNome(context, parametroRecebido.GetNome())
		if erroBuscaPreExistente.Erro == nil {
			return nil, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe parâmetro com o nome enviado"),
			}
		}
	}

	// Cria o objeto UpdateParametroParams gerado pelo sqlc para gravação no banco de dados
	parametroUpdate := db.UpdateParametroParams{
		ID:        parametroAntigo.GetId(),
		Nome:      pgtype.Text{String: parametroRecebido.GetNome(), Valid: true},
		Descricao: pgtype.Text{String: parametroRecebido.GetDescricao(), Valid: true},
		Valor:     pgtype.Text{String: parametroRecebido.GetValor(), Valid: true},
	}

	// Salva a parâmetro atualizada no repositório
	parametroAtualizada, erroUpdate := parametroService.parametroRepository.Update(context, parametroUpdate)
	if erroUpdate != nil {
		return &pb.Parametro{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroUpdate,
		}
	}

	return helpers.TParametroToPb(parametroAtualizada), erros.ErroStatus{}
}

// Função para deletar uma operadora pelo ID
func (parametroService *ParametroService) DeleteParametroById(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Deleta a operadora no repositório
	deletado, err := parametroService.parametroRepository.Delete(context, id)
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
			Erro:   errors.New("Parâmetro não encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}
