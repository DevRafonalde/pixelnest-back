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

// Estrutura do serviço de Perfil, contendo repositórios necessários
type PerfilService struct {
	perfilPermissaoRepository repositories.PerfilPermissaoRepository
	permissaoRepository       repositories.PermissaoRepository
	usuarioPerfilRepository   repositories.UsuarioPerfilRepository
	perfilRepository          repositories.PerfilRepository
	usuarioRepository         repositories.UsuarioRepository
}

// Função para criar uma nova instância de PerfilService com os repositórios necessários
func NewPerfilService(perfilPermissaoRepository repositories.PerfilPermissaoRepository,
	permissaoRepository repositories.PermissaoRepository,
	usuarioPerfilRepository repositories.UsuarioPerfilRepository,
	perfilRepository repositories.PerfilRepository, usuarioRepository repositories.UsuarioRepository) *PerfilService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &PerfilService{
		perfilPermissaoRepository: perfilPermissaoRepository,
		permissaoRepository:       permissaoRepository,
		usuarioPerfilRepository:   usuarioPerfilRepository,
		perfilRepository:          perfilRepository,
		usuarioRepository:         usuarioRepository,
	}
}

// Função para buscar todos os perfis
// Como um perfil pode estar vinculado a várias permissoes,
// Nessa função, os perfis são retornados sem as permissões vinculadas
func (perfilService *PerfilService) FindAllPerfis(context context.Context) ([]*pb.Perfil, erros.ErroStatus) {
	perfis, err := perfilService.perfilRepository.FindAll(context)
	if err != nil {
		return []*pb.Perfil{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum perfil, retorna code NotFound
	if len(perfis) == 0 {
		return []*pb.Perfil{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum perfil encontrado"),
		}
	}

	var pbPermissoes []*pb.Perfil
	for _, perfil := range perfis {
		pbPermissoes = append(pbPermissoes, helpers.TPerfToPb(perfil))
	}

	return pbPermissoes, erros.ErroStatus{}
}

// Função para buscar um perfil pelo ID
// Diferente da busca por todos, aqui a aplicação retorna o perfil com todas as permissões vinculadas
func (perfilService *PerfilService) FindPerfilById(context context.Context, id int32) (*pb.PerfilPermissoes, erros.ErroStatus) {
	perfilEncontrado, err := perfilService.perfilRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum perfil, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.PerfilPermissoes{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Perfil não encontrado"),
			}
		}

		return &pb.PerfilPermissoes{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	permissoes, erroPermissoes := perfilService.GetPermissoesVinculadas(context, perfilEncontrado.ID)
	if erroPermissoes.Erro != nil {
		return &pb.PerfilPermissoes{}, erroPermissoes
	}

	pbPerfilPermissao := pb.PerfilPermissoes{
		Perfil:     helpers.TPerfToPb(perfilEncontrado),
		Permissoes: permissoes,
	}

	return &pbPerfilPermissao, erros.ErroStatus{}
}

// Função para obter as permissões vinculadas a um perfil, filtrando apenas as ativas
func (perfilService *PerfilService) GetPermissoesVinculadas(context context.Context, id int32) ([]*pb.Permissao, erros.ErroStatus) {
	perfilPermissoes, errBuscaPermissoes := perfilService.perfilPermissaoRepository.FindByPerfil(context, id)
	if errBuscaPermissoes != nil {
		// Caso não seja encontrado nenhum perfil, retorna code NotFound
		if errBuscaPermissoes.Error() == "no rows in result set" {
			return []*pb.Permissao{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Perfil não encontrado"),
			}
		}

		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errBuscaPermissoes,
		}
	}

	// Filtra apenas as permissões que estão vinculadas e ativas
	var permissoes []*pb.Permissao
	for _, perfilPermissao := range perfilPermissoes {
		permissaoEncontrada, err := perfilService.permissaoRepository.FindByID(context, perfilPermissao.PermissaoID)
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		if permissaoEncontrada.Ativo.Bool {
			permissoes = append(permissoes, helpers.TPermissaoToPb(permissaoEncontrada))
		}
	}

	return permissoes, erros.ErroStatus{}
}

// Função para obter os usuários vinculados a um perfil, filtrando apenas os ativos
func (perfilService *PerfilService) GetUsuariosVinculados(context context.Context, id int32) ([]*pb.Usuario, erros.ErroStatus) {
	perfilUsuarios, errBuscaUsuarios := perfilService.usuarioPerfilRepository.FindByPerfil(context, id)
	if errBuscaUsuarios != nil {
		// Caso não seja encontrado nenhum perfil, retorna code NotFound
		if errBuscaUsuarios.Error() == "no rows in result set" {
			return []*pb.Usuario{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Perfil não encontrado"),
			}
		}

		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errBuscaUsuarios,
		}
	}

	// Filtra apenas os usuários que estão vinculados e ativos
	var usuarios []*pb.Usuario
	for _, perfilUsuario := range perfilUsuarios {
		usuarioEncontrado, err := perfilService.usuarioRepository.FindByID(context, perfilUsuario.UsuarioID)
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		if usuarioEncontrado.Ativo.Bool {
			usuarios = append(usuarios, helpers.TUsuarioToPb(usuarioEncontrado))
		}
	}

	return usuarios, erros.ErroStatus{}
}

// Função para clonar um perfil, mantendo as permissões, mas retornando um novo perfil vazio
func (perfilService *PerfilService) ClonePerfil(context context.Context, id int32) (*pb.PerfilPermissoes, erros.ErroStatus) {
	perfil, err := perfilService.FindPerfilById(context, id)
	if err.Erro != nil {
		return &pb.PerfilPermissoes{}, err
	}

	// Define o perfil clonado como vazio
	perfil.Perfil = &pb.Perfil{}

	return perfil, erros.ErroStatus{}
}

// Função para criar um novo perfil com permissões
func (perfilService *PerfilService) CreatePerfil(context context.Context, perfilPermissaoRecebido *pb.PerfilPermissoes) (*pb.PerfilPermissoes, erros.ErroStatus) {
	// Busca um perfil pelo nome enviado para verificar a prévia existência dele
	// Em caso positivo, retorna code AlreadyExists
	_, err := perfilService.perfilRepository.FindByNome(context, perfilPermissaoRecebido.Perfil.Nome)
	if err == nil {
		return &pb.PerfilPermissoes{}, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("Nome já está sendo utilizado"),
		}
	}

	// Cria o objeto CreatePerfilParams gerado pelo sqlc para gravação no banco de dados
	perfilCreate := db.CreatePerfilParams{
		Nome:                  perfilPermissaoRecebido.Perfil.Nome,
		Descricao:             perfilPermissaoRecebido.Perfil.Descricao,
		Ativo:                 pgtype.Bool{Bool: true, Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
	}

	// Cria o perfil no repositório
	perfil, err := perfilService.perfilRepository.Create(context, perfilCreate)
	perfilPermissaoRecebido.Perfil = helpers.TPerfToPb(perfil)
	if err != nil {
		return &pb.PerfilPermissoes{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Salva a relação entre o perfil e cada permissão
	for i, permissao := range perfilPermissaoRecebido.Permissoes {
		tPermissao, err := perfilService.permissaoRepository.FindByID(context, permissao.ID)
		if err != nil {
			return &pb.PerfilPermissoes{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// Coloca o objeto completo da permissão na lista de permissões do objeto recebido
		// Como esse é o objeto retornado, as permissões retornam completas e prontas para serem usadas
		perfilPermissaoRecebido.Permissoes[i] = helpers.TPermissaoToPb(tPermissao)

		// Cria o objeto CreatePerfilPermissaoParams gerado pelo sqlc para gravação no banco de dados
		perfilPermissao := db.CreatePerfilPermissaoParams{
			PerfilID:    perfil.ID,
			PermissaoID: permissao.ID,
			DataHora:    pgtype.Timestamp{Time: time.Now(), Valid: true},
		}

		perfilService.perfilPermissaoRepository.Create(context, perfilPermissao)
	}

	return perfilPermissaoRecebido, erros.ErroStatus{}
}

// Função para atualizar um perfil e suas permissões
func (perfilService *PerfilService) UpdatePerfil(context context.Context, perfilPermissao *pb.PerfilPermissoes, perfilPermissaoAntigo *pb.PerfilPermissoes) (*pb.PerfilPermissoes, erros.ErroStatus) {
	perfilRecebido := perfilPermissao.Perfil

	// Verifica se o nome foi modificado e, se sim, verifica se já existe outro registro com o mesmo nome
	// Em caso positivo, retorna code AlreadyExists
	if perfilPermissaoAntigo.GetPerfil().GetNome() != perfilRecebido.Nome {
		_, err := perfilService.perfilRepository.FindByNome(context, perfilRecebido.Nome)
		if err == nil {
			return &pb.PerfilPermissoes{}, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Nome já está sendo utilizado"),
			}
		}
	}

	// Cria o objeto UpdatePerfilParams gerado pelo sqlc para gravação no banco de dados
	perfilUpdate := db.UpdatePerfilParams{
		Nome:                  perfilRecebido.Nome,
		Descricao:             perfilRecebido.Descricao,
		Ativo:                 pgtype.Bool{Bool: perfilPermissaoAntigo.GetPerfil().GetAtivo(), Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		ID:                    perfilRecebido.ID,
	}

	// Salva o perfil atualizado no repositório
	perfilAtualizado, err := perfilService.perfilRepository.Update(context, perfilUpdate)
	if err != nil {
		return &pb.PerfilPermissoes{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Deleta todas as relações antigas entre o perfil e suas permissões
	registrosExistentes, errBuscaPermissoes := perfilService.perfilPermissaoRepository.FindByPerfil(context, perfilAtualizado.ID)
	if errBuscaPermissoes != nil {
		return &pb.PerfilPermissoes{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	for i := 0; i < len(registrosExistentes); i++ {
		erroDelete := perfilService.perfilPermissaoRepository.Delete(context, registrosExistentes[i].ID)
		if erroDelete != nil {
			return &pb.PerfilPermissoes{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   erroDelete,
			}
		}
	}

	// Cria novas relações entre o perfil atualizado e as novas permissões
	permissoes := perfilPermissao.Permissoes

	for i, permissao := range permissoes {
		tPermissao, err := perfilService.permissaoRepository.FindByID(context, permissao.ID)
		if err != nil {
			return &pb.PerfilPermissoes{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// Coloca o objeto completo da permissão na lista de permissões do objeto recebido
		// Como esse é o objeto retornado, as permissões retornam completas e prontas para serem usadas
		permissoes[i] = helpers.TPermissaoToPb(tPermissao)

		// Cria o objeto CreatePerfilPermissaoParams gerado pelo sqlc para gravação no banco de dados
		perfilPermissao := db.CreatePerfilPermissaoParams{
			PerfilID:    perfilAtualizado.ID,
			PermissaoID: permissao.ID,
			DataHora:    pgtype.Timestamp{Time: time.Now(), Valid: true},
		}

		perfilService.perfilPermissaoRepository.Create(context, perfilPermissao)
	}

	perfilPermissaoRetorno := pb.PerfilPermissoes{
		Perfil:     helpers.TPerfToPb(perfilAtualizado),
		Permissoes: permissoes,
	}

	return &perfilPermissaoRetorno, erros.ErroStatus{}
}

// Função para desativar um perfil pelo ID
func (perfilService *PerfilService) DesativarPerfilById(context context.Context, perfil *pb.PerfilPermissoes) erros.ErroStatus {
	// Cria o objeto UpdatePerfilParams gerado pelo sqlc para gravação no banco de dados
	perfilUpdate := db.UpdatePerfilParams{
		Nome:                  perfil.Perfil.Nome,
		Descricao:             perfil.Perfil.Descricao,
		Ativo:                 pgtype.Bool{Bool: false, Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		ID:                    perfil.Perfil.ID,
	}

	// Salva o perfil atualizado no repositório
	_, errAtt := perfilService.perfilRepository.Update(context, perfilUpdate)

	return erros.ErroStatus{
		Status: codes.Internal,
		Erro:   errAtt,
	}
}

// Função para ativar um perfil pelo ID
func (perfilService *PerfilService) AtivarPerfilById(context context.Context, perfil *pb.PerfilPermissoes) erros.ErroStatus {
	// Cria o objeto UpdatePerfilParams gerado pelo sqlc para gravação no banco de dados
	perfilUpdate := db.UpdatePerfilParams{
		Nome:                  perfil.Perfil.Nome,
		Descricao:             perfil.Perfil.Descricao,
		Ativo:                 pgtype.Bool{Bool: true, Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		ID:                    perfil.Perfil.ID,
	}

	// Salva o perfil atualizado no repositório
	_, errAtt := perfilService.perfilRepository.Update(context, perfilUpdate)

	return erros.ErroStatus{
		Status: codes.Internal,
		Erro:   errAtt,
	}
}
