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

// Estrutura do serviço de Permissão, contendo repositórios necessários
type PermissaoService struct {
	perfilPermissaoRepository repositories.PerfilPermissaoRepository
	permissaoRepository       repositories.PermissaoRepository
}

// Função para criar uma nova instância de PermissaoService com os repositórios necessários
func NewPermissaoService(perfilPermissaoRepository repositories.PerfilPermissaoRepository, permissaoRepository repositories.PermissaoRepository) *PermissaoService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &PermissaoService{
		perfilPermissaoRepository: perfilPermissaoRepository,
		permissaoRepository:       permissaoRepository,
	}
}

// Função para buscar todas as permissões
func (permissaoService *PermissaoService) FindAllPermissoes(context context.Context) ([]*pb.Permissao, erros.ErroStatus) {
	tPermissoes, err := permissaoService.permissaoRepository.FindAll(context)
	if err != nil {
		return []*pb.Permissao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrada nenhuma permissão, retorna code NotFound
	if len(tPermissoes) == 0 {
		return []*pb.Permissao{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhuma permissão encontrada"),
		}
	}

	var pbPermissoes []*pb.Permissao
	for _, permissao := range tPermissoes {
		pbPermissoes = append(pbPermissoes, helpers.TPermissaoToPb(permissao))
	}

	return pbPermissoes, erros.ErroStatus{}
}

// Função para buscar uma permissão pelo ID
func (permissaoService *PermissaoService) FindPermissaoById(context context.Context, id int32) (*pb.Permissao, erros.ErroStatus) {
	permissao, err := permissaoService.permissaoRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrada nenhuma permissão, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Permissao{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Permissão não encontrada"),
			}
		}

		return &pb.Permissao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TPermissaoToPb(permissao), erros.ErroStatus{}
}

// Função para buscar uma permissão pelo nome
func (permissaoService *PermissaoService) FindPermissaoByNome(context context.Context, nome string) (*pb.Permissao, erros.ErroStatus) {
	permissao, err := permissaoService.permissaoRepository.FindByNome(context, nome)
	if err != nil {
		// Caso não seja encontrada nenhuma permissão, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Permissao{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Permissão não encontrada"),
			}
		}

		return &pb.Permissao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TPermissaoToPb(permissao), erros.ErroStatus{}
}

// Função para criar uma nova permissão
func (permissaoService *PermissaoService) CreatePermissao(context context.Context, permissao *pb.Permissao) (*pb.Permissao, erros.ErroStatus) {
	// Busca uma permissão pelo nome enviado para verificar a prévia existência dela
	// Em caso positivo, retorna code AlreadyExists
	if _, err := permissaoService.FindPermissaoByNome(context, permissao.Nome); err.Erro == nil {
		return &pb.Permissao{}, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("Já existe permissão com esse nome"),
		}
	}

	// Cria o objeto CreatePermissaoParams gerado pelo sqlc para gravação no banco de dados
	permissaoCreate := db.CreatePermissaoParams{
		Nome:                  permissao.Nome,
		Descricao:             permissao.Descricao,
		Ativo:                 pgtype.Bool{Bool: true, Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	// Cria a permissão no repositório
	permissaoCriada, err := permissaoService.permissaoRepository.Create(context, permissaoCreate)
	if err != nil {
		return &pb.Permissao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TPermissaoToPb(permissaoCriada), erros.ErroStatus{}
}

// Função para atualizar uma permissão existente
func (permissaoService *PermissaoService) UpdatePermissao(context context.Context, permissaoRecebida *pb.Permissao) (*pb.Permissao, erros.ErroStatus) {
	permissaoBanco, err := permissaoService.FindPermissaoById(context, permissaoRecebida.GetID())
	if err.Erro != nil {
		return &pb.Permissao{}, err
	}

	// Verifica se o nome foi modificado e, se sim, verifica se já existe outro registro com o mesmo nome
	// Em caso positivo, retorna code AlreadyExists
	if permissaoBanco.GetNome() != permissaoRecebida.GetNome() {
		_, err := permissaoService.FindPermissaoByNome(context, permissaoRecebida.GetNome())
		if err.Erro == nil {
			return &pb.Permissao{}, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Nome já está sendo utilizado"),
			}
		}
	}

	// Cria o objeto UpdatePermissaoParams gerado pelo sqlc para gravação no banco de dados
	permissaoUpdate := db.UpdatePermissaoParams{
		Nome:                  permissaoRecebida.GetNome(),
		Descricao:             permissaoRecebida.GetDescricao(),
		Ativo:                 pgtype.Bool{Bool: true, Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		ID:                    permissaoBanco.GetID(),
	}

	// Salva a permissão atualizada no repositório
	permissaoAtualizada, erroUpdate := permissaoService.permissaoRepository.Update(context, permissaoUpdate)
	if erroUpdate != nil {
		return &pb.Permissao{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroUpdate,
		}
	}

	return helpers.TPermissaoToPb(permissaoAtualizada), erros.ErroStatus{}
}

// Função para desativar uma permissão pelo ID
func (permissaoService *PermissaoService) DesativarPermissaoById(context context.Context, permissao *pb.Permissao) erros.ErroStatus {
	// Cria o objeto UpdatePermissaoParams gerado pelo sqlc para gravação no banco de dados
	permissaoUpdate := db.UpdatePermissaoParams{
		Nome:                  permissao.Nome,
		Descricao:             permissao.Descricao,
		Ativo:                 pgtype.Bool{Bool: false, Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		ID:                    permissao.ID,
	}

	// Salva a permissão atualizada no repositório
	_, errAtt := permissaoService.permissaoRepository.Update(context, permissaoUpdate)

	return erros.ErroStatus{
		Status: codes.Internal,
		Erro:   errAtt,
	}
}

// Função para ativar uma permissão pelo ID
func (permissaoService *PermissaoService) AtivarPermissaoById(context context.Context, permissao *pb.Permissao) erros.ErroStatus {
	// Cria o objeto UpdatePermissaoParams gerado pelo sqlc para gravação no banco de dados
	permissaoUpdate := db.UpdatePermissaoParams{
		Nome:                  permissao.Nome,
		Descricao:             permissao.Descricao,
		Ativo:                 pgtype.Bool{Bool: true, Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		ID:                    permissao.ID,
	}

	// Salva a permissão atualizada no repositório
	_, errAtt := permissaoService.permissaoRepository.Update(context, permissaoUpdate)

	return erros.ErroStatus{
		Status: codes.Internal,
		Erro:   errAtt,
	}
}
