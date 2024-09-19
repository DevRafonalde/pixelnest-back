package helpers

import (
	"fmt"
	"pixelnest/app/model/grpc"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

// Conversão de TUsuario para pb.Usuario
func TUsuarioToPb(usuario db.TUsuario) *grpc.Usuario {
	return &grpc.Usuario{
		ID:                    usuario.ID,
		Nome:                  usuario.Nome,
		Email:                 usuario.Email,
		Senha:                 usuario.Senha,
		Ativo:                 usuario.Ativo.Bool,
		TokenResetSenha:       usuario.TokenResetSenha.String,
		DataUltimaAtualizacao: usuario.DataUltimaAtualizacao.Time.Format("2006/01/02"),
		SenhaAtualizada:       usuario.SenhaAtualizada.Bool,
	}
}

// Conversão de TPerfi para pb.Perfil
func TPerfToPb(perfil db.TPerfi) *grpc.Perfil {
	return &grpc.Perfil{
		ID:                    perfil.ID,
		Nome:                  perfil.Nome,
		Descricao:             perfil.Descricao,
		Ativo:                 perfil.Ativo.Bool,
		DataUltimaAtualizacao: perfil.DataUltimaAtualizacao.Time.Format("2006/01/02"),
	}
}

func TPermissaoToPb(permissao db.TPermisso) *grpc.Permissao {
	return &grpc.Permissao{
		ID:                    permissao.ID,
		Nome:                  permissao.Nome,
		Descricao:             permissao.Descricao,
		Ativo:                 permissao.Ativo.Bool,
		DataUltimaAtualizacao: permissao.DataUltimaAtualizacao.Time.Format("2006/01/02"),
	}
}

func TAvaliacaoToPb(cidade db.TAvaliaco, usuario *grpc.Usuario, produto *grpc.Produto, jogo *grpc.Jogo) *grpc.Avaliacao {
	return &grpc.Avaliacao{
		ID:        cidade.ID,
		Usuario:   usuario,
		Produto:   produto,
		Jogo:      jogo,
		Nota:      cidade.Nota.Int32,
		Avaliacao: cidade.Avaliacao.String,
	}
}

func TFavoritoToPb(cliente db.TFavorito, usuario *grpc.Usuario, produto *grpc.Produto, jogo *grpc.Jogo) *grpc.Favorito {
	return &grpc.Favorito{
		ID:      cliente.ID,
		Usuario: usuario,
		Produto: produto,
		Jogo:    jogo,
	}
}

func TJogoToPb(jogo db.TJogo) *grpc.Jogo {
	avaliacao, err := jogo.Avaliacao.Float64Value()
	if err != nil {
		fmt.Println("Deu ruim na conversão da avaliação do jogo aqui no helper")
		avaliacao = pgtype.Float8{Float64: 0.0}
	}

	return &grpc.Jogo{
		ID:        jogo.ID,
		Nome:      jogo.Nome,
		Sinopse:   jogo.Sinopse.String,
		Genero:    jogo.Genero.String,
		Avaliacao: avaliacao.Float64,
	}
}

func TParametroToPb(parametro db.TParametro) *grpc.Parametro {
	return &grpc.Parametro{
		Id:        parametro.ID,
		Nome:      parametro.Nome.String,
		Descricao: parametro.Descricao.String,
		Valor:     parametro.Valor.String,
	}
}

func TProdutoToPb(produto db.TProduto) *grpc.Produto {
	avaliacao, err := produto.Avaliacao.Float64Value()
	if err != nil {
		fmt.Println("Deu ruim na conversão da avaliação do jogo aqui no helper")
		avaliacao = pgtype.Float8{Float64: 0.0}
	}

	return &grpc.Produto{
		ID:        produto.ID,
		Nome:      produto.Nome,
		Descricao: produto.Descricao.String,
		Genero:    produto.Genero.String,
		Avaliacao: avaliacao.Float64,
	}
}
