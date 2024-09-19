package service

import (
	"context"
	"errors"
	"pixelnest/app/helpers"
	"pixelnest/app/model/erros"
	"pixelnest/app/model/grpc"
	pb "pixelnest/app/model/grpc"
	"pixelnest/app/model/repositories"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
)

// Estrutura do serviço de Jogo, contendo o repositório necessário
type JogoService struct {
	jogoRepository        repositories.JogoRepository
	usuarioJogoRepository repositories.UsuarioJogoRepository
	favoritoRepository    repositories.FavoritoRepository
}

// Função para criar um nova instância de JogoService com o repositório necessário
func NewJogoService(jogoRepository repositories.JogoRepository, usuarioJogoRepository repositories.UsuarioJogoRepository, favoritoRepository repositories.FavoritoRepository) *JogoService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &JogoService{
		jogoRepository:        jogoRepository,
		usuarioJogoRepository: usuarioJogoRepository,
		favoritoRepository:    favoritoRepository,
	}
}

// Função para buscar um jogo pelo ID
func (jogoService *JogoService) FindJogoById(context context.Context, id int32) (*grpc.Jogo, erros.ErroStatus) {
	jogo, err := jogoService.jogoRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum jogo, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.Jogo{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Jogo não encontrado"),
			}
		}

		return &grpc.Jogo{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TJogoToPb(jogo), erros.ErroStatus{}
}

// Função para buscar um jogo pelo nome
func (jogoService *JogoService) FindJogoByNome(context context.Context, nome string) ([]*grpc.Jogo, erros.ErroStatus) {
	tJogos, err := jogoService.jogoRepository.FindByNome(context, nome)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum jogo, retorna code NotFound
	if len(tJogos) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum jogo encontrado"),
		}
	}

	var jogos []*grpc.Jogo
	for _, jogo := range tJogos {
		jogos = append(jogos, helpers.TJogoToPb(jogo))
	}

	return jogos, erros.ErroStatus{}
}

// Função para buscar um jogo pelo nome
func (jogoService *JogoService) FindJogoByGenero(context context.Context, genero string) ([]*grpc.Jogo, erros.ErroStatus) {
	tJogos, err := jogoService.jogoRepository.FindByGenero(context, genero)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum jogo, retorna code NotFound
	if len(tJogos) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum jogo encontrado"),
		}
	}

	var jogos []*grpc.Jogo
	for _, jogo := range tJogos {
		jogos = append(jogos, helpers.TJogoToPb(jogo))
	}

	return jogos, erros.ErroStatus{}
}

// Função para buscar um jogo pelo nome
func (jogoService *JogoService) FindJogoByUsuario(context context.Context, id int32) ([]*grpc.Jogo, erros.ErroStatus) {
	tUsuarioJogos, err := jogoService.usuarioJogoRepository.FindByUsuario(context, id)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum jogo, retorna code NotFound
	if len(tUsuarioJogos) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum jogo encontrado"),
		}
	}

	var jogos []*grpc.Jogo
	for _, usuarioJogo := range tUsuarioJogos {
		tJogo, err := jogoService.jogoRepository.FindByID(context, usuarioJogo.JogoID)
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		jogos = append(jogos, helpers.TJogoToPb(tJogo))
	}

	return jogos, erros.ErroStatus{}
}

// Função para buscar um jogo pelo nome
func (jogoService *JogoService) FindJogoFavoritoByUsuario(context context.Context, id int32) ([]*grpc.Jogo, erros.ErroStatus) {
	tFavoritos, err := jogoService.favoritoRepository.FindByUsuario(context, id)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum jogo, retorna code NotFound
	if len(tFavoritos) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum jogo encontrado"),
		}
	}

	var jogos []*grpc.Jogo
	for _, favorito := range tFavoritos {
		if favorito.JogoID.Int32 != 0 {
			tJogo, err := jogoService.jogoRepository.FindByID(context, favorito.JogoID.Int32)
			if err != nil {
				return nil, erros.ErroStatus{
					Status: codes.Internal,
					Erro:   err,
				}
			}

			jogos = append(jogos, helpers.TJogoToPb(tJogo))
		}
	}

	return jogos, erros.ErroStatus{}
}

// Função para buscar todas as jogos
func (jogoService *JogoService) FindAllJogos(context context.Context) ([]*grpc.Jogo, erros.ErroStatus) {
	tJogos, err := jogoService.jogoRepository.FindAll(context)
	if err != nil {
		return []*grpc.Jogo{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum jogo, retorna code NotFound
	if len(tJogos) == 0 {
		return []*grpc.Jogo{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum jogo encontrado"),
		}
	}

	var jogos []*grpc.Jogo
	for _, jogo := range tJogos {
		jogos = append(jogos, helpers.TJogoToPb(jogo))
	}

	return jogos, erros.ErroStatus{}
}

// Função para criar um nova jogo
func (jogoService *JogoService) CreateJogo(context context.Context, jogo *grpc.Jogo) (*grpc.Jogo, erros.ErroStatus) {
	// Busca um jogo pelo nome enviado para verificar a prévia existência dela
	// Em caso positivo, retorna code AlreadyExists
	tJogos, erroService := jogoService.FindJogoByNome(context, jogo.GetNome())
	if erroService.Erro == nil || len(tJogos) != 0 {
		return &pb.Jogo{}, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("Já existe jogo com esse nome"),
		}
	}

	// Cria o objeto CreateJogoParams gerado pelo sqlc para gravação no banco de dados
	jogoCreate := db.CreateJogoParams{
		Nome:    jogo.GetNome(),
		Sinopse: pgtype.Text{String: jogo.GetSinopse(), Valid: true},
		Genero:  pgtype.Text{String: jogo.GetGenero(), Valid: true},
	}

	// Cria a jogo no repositório
	jogoCriada, err := jogoService.jogoRepository.Create(context, jogoCreate)
	if err != nil {
		return &grpc.Jogo{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TJogoToPb(jogoCriada), erros.ErroStatus{}
}

// Função para atualizar um jogo
func (jogoService *JogoService) UpdateJogo(context context.Context, jogoRecebido *grpc.Jogo, jogoAntiga *grpc.Jogo) (*grpc.Jogo, erros.ErroStatus) {
	// Verifica se o nome foi modificado e, se sim, verifica se já existe outro registro com o mesmo nome
	// Em caso positivo, retorna code AlreadyExists
	if jogoAntiga.GetNome() != jogoRecebido.GetNome() {
		_, erroBuscaPreExistente := jogoService.FindJogoByNome(context, jogoRecebido.GetNome())
		if erroBuscaPreExistente.Erro == nil {
			return nil, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe jogo com o nome enviado"),
			}
		}
	}

	// Cria o objeto UpdateJogoParams gerado pelo sqlc para gravação no banco de dados
	jogoUpdate := db.UpdateJogoParams{
		Nome:    jogoRecebido.GetNome(),
		Sinopse: pgtype.Text{String: jogoRecebido.GetSinopse(), Valid: true},
		Genero:  pgtype.Text{String: jogoRecebido.GetGenero(), Valid: true},
		ID:      jogoRecebido.GetID(),
	}

	// Salva a jogo atualizada no repositório
	jogoAtualizada, erroSalvamento := jogoService.jogoRepository.Update(context, jogoUpdate)
	if erroSalvamento != nil {
		return &grpc.Jogo{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroSalvamento,
		}
	}

	return helpers.TJogoToPb(jogoAtualizada), erros.ErroStatus{}
}

// Função para deletar um jogo pelo ID
func (jogoService *JogoService) DeleteJogoById(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Deleta a jogo no repositório
	deletado, err := jogoService.jogoRepository.Delete(context, id)
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo deletados indica o número de linhas deletadas. Se for 0, nenhum cidade foi deletada, pois não existia
	if deletado == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Jogo não encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}
