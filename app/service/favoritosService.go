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

// Estrutura de serviço para gerenciar operações relacionadas às favoritos
type FavoritoService struct {
	favoritoRepository repositories.FavoritoRepository
	usuarioRepository  repositories.UsuarioRepository
	jogoRepository     repositories.JogoRepository
	produtoRepository  repositories.ProdutoRepository
}

// Função para criar uma nova instância de FavoritoService com o repositório necessário
func NewFavoritoService(favoritoRepository repositories.FavoritoRepository,
	usuarioRepository repositories.UsuarioRepository,
	jogoRepository repositories.JogoRepository,
	produtoRepository repositories.ProdutoRepository) *FavoritoService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &FavoritoService{
		favoritoRepository: favoritoRepository,
		usuarioRepository:  usuarioRepository,
		jogoRepository:     jogoRepository,
		produtoRepository:  produtoRepository,
	}
}

// Função para buscar um favorito pelo ID
func (favoritoService *FavoritoService) FindFavoritoById(context context.Context, id int32) (*pb.Favorito, erros.ErroStatus) {
	// Busca o favorito no repositório pelo ID
	favorito, err := favoritoService.favoritoRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum favorito, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Favorito{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Favorito não encontrado"),
			}
		}

		return &pb.Favorito{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return favoritoService.montarObjCompleto(context, favorito)
}

// Função para buscar um favorito pelo ID externo
func (favoritoService *FavoritoService) FindJogosFavoritosByUsuario(context context.Context, id int32) ([]*pb.Favorito, erros.ErroStatus) {
	// Busca o favoritos no repositório pelo ID
	favoritos, err := favoritoService.favoritoRepository.FindByUsuario(context, id)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum favorito, retorna code NotFound
	if len(favoritos) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum favorito encontrado"),
		}
	}

	var pbFavoritos []*pb.Favorito
	for _, favorito := range favoritos {
		if favorito.JogoID.Int32 != 0 {
			favoritoGRPC, erroMontagem := favoritoService.montarObjCompleto(context, favorito)
			if erroMontagem.Erro != nil {
				return nil, erroMontagem
			}

			pbFavoritos = append(pbFavoritos, favoritoGRPC)
		}
	}

	return pbFavoritos, erros.ErroStatus{}
}

// Função para buscar um favorito pelo ID externo
func (favoritoService *FavoritoService) FindProdutosFavoritosByUsuario(context context.Context, id int32) ([]*pb.Favorito, erros.ErroStatus) {
	// Busca o favoritos no repositório pelo ID
	favoritos, err := favoritoService.favoritoRepository.FindByUsuario(context, id)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum favorito, retorna code NotFound
	if len(favoritos) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum favorito encontrado"),
		}
	}

	var pbFavoritos []*pb.Favorito
	for _, favorito := range favoritos {
		if favorito.JogoID.Int32 != 0 {
			favoritoGRPC, erroMontagem := favoritoService.montarObjCompleto(context, favorito)
			if erroMontagem.Erro != nil {
				return nil, erroMontagem
			}

			pbFavoritos = append(pbFavoritos, favoritoGRPC)
		}
	}

	return pbFavoritos, erros.ErroStatus{}
}

// Função para criar uma nova favorito
func (favoritoService *FavoritoService) CreateFavorito(context context.Context, favorito *pb.Favorito) (*pb.Favorito, erros.ErroStatus) {
	// Cria o objeto CreateFavoritoParams gerado pelo sqlc para gravação no banco de dados
	favoritoCreate := db.CreateFavoritoParams{
		UsuarioID: favorito.GetUsuario().GetID(),
		ProdutoID: pgtype.Int4{Int32: favorito.GetProduto().GetID(), Valid: true},
		JogoID:    pgtype.Int4{Int32: favorito.GetJogo().GetID(), Valid: true},
	}

	// Cria o favorito no repositório
	favoritoCriada, err := favoritoService.favoritoRepository.Create(context, favoritoCreate)
	if err != nil {
		return &pb.Favorito{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return favoritoService.montarObjCompleto(context, favoritoCriada)
}

// Função para deletar um favorito pelo ID
func (favoritoService *FavoritoService) DeleteFavoritoById(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Deleta o favorito no repositório pelo ID
	deletados, err := favoritoService.favoritoRepository.Delete(context, id)

	// Caso ocorra erro na deleção
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo deletados indica o número de linhas deletadas. Se for 0, nenhum favorito foi deletada, pois não existia
	if deletados == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum favorito encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}

func (favoritoService *FavoritoService) montarObjCompleto(context context.Context, favorito db.TFavorito) (*pb.Favorito, erros.ErroStatus) {
	// Faz a montagem do objeto operadora de entrada de portabilidade, caso necessário
	usuario := new(pb.Usuario)
	if favorito.UsuarioID != 0 {
		tUsuario, err := favoritoService.usuarioRepository.FindByID(context, favorito.UsuarioID)
		if err != nil {
			return &pb.Favorito{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
		usuario = helpers.TUsuarioToPb(tUsuario)
	}

	// Faz a montagem do objeto operadora de saída de portabilidade, caso necessário
	produto := new(pb.Produto)
	if favorito.ProdutoID.Int32 != 0 {
		tProduto, err := favoritoService.produtoRepository.FindByID(context, favorito.ProdutoID.Int32)
		if err != nil {
			return &pb.Favorito{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
		produto = helpers.TProdutoToPb(tProduto)
	}

	// Faz a montagem do objeto SimCard, caso necessário
	jogo := new(pb.Jogo)
	if favorito.JogoID.Int32 != 0 {
		tJogo, err := favoritoService.jogoRepository.FindByID(context, favorito.JogoID.Int32)
		if err != nil {
			return &pb.Favorito{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		jogo = helpers.TJogoToPb(tJogo)
	}

	// Passando todos os objetos completinhos para o helper que irá converter o objeto completo para mim
	return helpers.TFavoritoToPb(favorito, usuario, produto, jogo), erros.ErroStatus{}
}
