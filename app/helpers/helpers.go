package helpers

import (
	"pixelnest/app/model/grpc"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
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

func TCidadeToPb(cidade db.TCidade) *grpc.Cidade {
	return &grpc.Cidade{
		ID:      cidade.ID,
		UUID:    cidade.Uuid.String,
		Nome:    cidade.Nome,
		CodIBGE: cidade.CodIbge,
		UF:      cidade.Uf,
		CodArea: cidade.CodArea,
	}
}

func TClienteToPb(cliente db.TCliente) *grpc.Cliente {
	return &grpc.Cliente{
		ID:         cliente.ID,
		UUID:       cliente.Uuid.String,
		IDExterno:  cliente.IDExterno,
		Nome:       cliente.Nome,
		Documento:  cliente.Documento,
		CodReserva: cliente.CodReserva.String,
	}
}

func TTelefoniaNumeroToPb(numeroTelefonico db.TTelefoniaNumero, operadoraIn *grpc.Operadora, operadoraOut *grpc.Operadora, simCard *grpc.SimCard) *grpc.NumeroTelefonico {
	var portadoInDate string
	if numeroTelefonico.PortadoInDate.Time.IsZero() {
		portadoInDate = ""
	} else {
		portadoInDate = numeroTelefonico.PortadoInDate.Time.Format("2006/01/02")
	}

	var congeladoAte string
	if numeroTelefonico.PortadoInDate.Time.IsZero() {
		congeladoAte = ""
	} else {
		congeladoAte = numeroTelefonico.CongeladoAte.Time.Format("2006/01/02")
	}

	var portadoOutDate string
	if numeroTelefonico.PortadoInDate.Time.IsZero() {
		portadoOutDate = ""
	} else {
		portadoOutDate = numeroTelefonico.PortadoOutDate.Time.Format("2006/01/02")
	}

	return &grpc.NumeroTelefonico{
		ID:                  numeroTelefonico.ID,
		CodArea:             int32(numeroTelefonico.CodArea),
		Numero:              numeroTelefonico.Numero,
		Utilizavel:          numeroTelefonico.Utilizavel,
		PortadoIn:           numeroTelefonico.PortadoIn,
		PortadoInOperadora:  operadoraIn,
		PortadoInDate:       portadoInDate,
		CodigoCNL:           numeroTelefonico.CodigoCnl,
		CongeladoAte:        congeladoAte,
		ExternalID:          numeroTelefonico.ExternalID.Int32,
		PortadoOut:          numeroTelefonico.PortadoOut,
		PortadoOutOperadora: operadoraOut,
		PortadoOutDate:      portadoOutDate,
		DataCriacao:         numeroTelefonico.DataCriacao.Time.Format("2006/01/02"),
		SimCard:             simCard,
	}
}

func TOperadoraToPb(operadora db.TOperadora) *grpc.Operadora {
	return &grpc.Operadora{
		ID:         operadora.ID,
		Nome:       operadora.Nome,
		Abreviacao: operadora.Abreviacao,
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

func TSimCardToPb(simCard db.TSimcard, estado *grpc.SimCardStatus, numeroTelefonico *grpc.NumeroTelefonico) *grpc.SimCard {
	return &grpc.SimCard{
		ID:              simCard.ID,
		ICCID:           simCard.Iccid,
		IMSI:            simCard.Imsi,
		PIN:             simCard.Pin,
		PUK:             simCard.Puk,
		KI:              simCard.Ki,
		OPC:             simCard.Opc,
		Status:          estado,
		TelefoniaNumero: numeroTelefonico,
		DataCriacao:     simCard.DataCriacao.Time.Format("2006/01/02"),
		DataStatus:      simCard.DataStatus.Time.Format("2006/01/02"),
		AtualizadoEm:    simCard.AtualizadoEm.Time.Format("2006/01/02"),
		PUK2:            simCard.Puk2.String,
		PIN2:            simCard.Pin2.String,
	}
}

func TStatusSimCardToPb(simCardStatus db.TSimcardStatus) *grpc.SimCardStatus {
	return &grpc.SimCardStatus{
		ID:        simCardStatus.ID,
		Nome:      simCardStatus.Nome,
		Descricao: simCardStatus.Descricao.String,
	}
}
