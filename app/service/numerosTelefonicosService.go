package service

import (
	"context"
	"errors"
	"pixelnest/app/helpers"
	"pixelnest/app/model/erros"
	"pixelnest/app/model/grpc"
	"pixelnest/app/model/repositories"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/grpc/codes"
)

// Estrutura de serviço para gerenciar operações relacionadas a números telefônicos
type NumeroTelefonicoService struct {
	numeroTelefonicoRepository repositories.NumeroTelefonicoRepository
	cidadeRepository           repositories.CidadeRepository
	simCardRepository          repositories.SimCardRepository
	operadoraRepository        repositories.OperadoraRepository
	simCardEstadoRepository    repositories.SimCardStatusRepository
	clienteRepository          repositories.ClienteRepository
	simCardTelefoneService     *SimCardTelefoneService
}

// Função para criar uma nova instância de NumeroTelefonicoService com o repositório necessário
func NewNumeroTelefonicoService(numeroTelefonicoRepository repositories.NumeroTelefonicoRepository,
	cidadeRepository repositories.CidadeRepository,
	simCardRepository repositories.SimCardRepository,
	operadoraRepository repositories.OperadoraRepository,
	simCardEstadoRepository repositories.SimCardStatusRepository,
	clienteRepository repositories.ClienteRepository,
	simCardTelefoneService *SimCardTelefoneService) *NumeroTelefonicoService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &NumeroTelefonicoService{
		numeroTelefonicoRepository: numeroTelefonicoRepository,
		cidadeRepository:           cidadeRepository,
		simCardRepository:          simCardRepository,
		operadoraRepository:        operadoraRepository,
		simCardEstadoRepository:    simCardEstadoRepository,
		clienteRepository:          clienteRepository,
		simCardTelefoneService:     simCardTelefoneService,
	}
}

// Função para buscar um número telefônico pelo ID
func (numeroTelefonicoService *NumeroTelefonicoService) FindNumeroTelefonicoById(context context.Context, id int32) (*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Busca o número telefônico no repositório pelo ID
	numeroTelefonico, err := numeroTelefonicoService.numeroTelefonicoRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum número telefônico, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Número telefônico não encontrado"),
			}
		}

		return &grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return numeroTelefonicoService.montarObjNumeroCompleto(context, numeroTelefonico)
}

// Função para buscar um número telefônico pelo número
func (numeroTelefonicoService *NumeroTelefonicoService) FindNumeroTelefonicoByNumero(context context.Context, numero int32) (*grpc.NumeroTelefonico, erros.ErroStatus) {
	strNumero := strconv.FormatInt(int64(numero), 10)

	substr := strNumero[2:]

	numeroSemDDD, err := strconv.ParseInt(substr, 10, 32)
	if err != nil {
		return &grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Busca o número telefônico no repositório pelo número
	numerosTelefonicos, err := numeroTelefonicoService.numeroTelefonicoRepository.FindByNumero(context, int32(numeroSemDDD))
	if err != nil {
		// Caso não seja encontrado nenhum número telefônico, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Número telefônico não encontrado"),
			}
		}

		return &grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	for _, numero := range numerosTelefonicos {
		strCodArea := strNumero[0:2]
		codAreaRecebido, err := strconv.ParseInt(strCodArea, 10, 16)
		if err != nil {
			return &grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		if numero.CodArea == int16(codAreaRecebido) {
			return numeroTelefonicoService.montarObjNumeroCompleto(context, numero)
		}
	}

	return &grpc.NumeroTelefonico{}, erros.ErroStatus{
		Status: codes.NotFound,
		Erro:   errors.New("Número telefônico não encontrado"),
	}
}

// Função para buscar um número telefônico pelo ID do SIM card associado
func (numeroTelefonicoService *NumeroTelefonicoService) FindNumeroTelefonicoBySimCardId(context context.Context, id int32) (*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Busca o número telefônico no repositório pelo ID do SIM card
	numeroTelefonico, err := numeroTelefonicoService.numeroTelefonicoRepository.FindBySimCardId(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum número telefônico, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Número telefônico não encontrado"),
			}
		}

		return &grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return numeroTelefonicoService.montarObjNumeroCompleto(context, numeroTelefonico)
}

// Função para buscar um número telefônico pelo ID do SIM card associado
func (numeroTelefonicoService *NumeroTelefonicoService) FindNumeroTelefonicoBySimCardIccid(context context.Context, iccid string) (*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Busca o número telefônico no repositório pelo ID do SIM card
	simCard, err := numeroTelefonicoService.simCardRepository.FindByIccid(context, iccid)
	if err != nil {
		// Caso não seja encontrado nenhum número telefônico, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Número telefônico não encontrado"),
			}
		}

		return &grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	numeroTelefonico, erroService := numeroTelefonicoService.FindNumeroTelefonicoBySimCardId(context, simCard.ID)
	if erroService.Erro != nil {
		return &grpc.NumeroTelefonico{}, erroService
	}

	return numeroTelefonico, erros.ErroStatus{}
}

// Função para buscar todos os números telefônicos
func (numeroTelefonicoService *NumeroTelefonicoService) FindAllNumeroTelefonicos(context context.Context) ([]*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Busca todos os números telefônicos no repositório
	TNumerosTelefonicos, err := numeroTelefonicoService.numeroTelefonicoRepository.FindAll(context)
	if err != nil {
		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum número telefônico, retorna code NotFound
	if len(TNumerosTelefonicos) == 0 {
		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum número telefônico encontrado"),
		}
	}

	var numerosTelefonicos []*grpc.NumeroTelefonico
	for _, numero := range TNumerosTelefonicos {
		numeroTelefonico, erroConversao := numeroTelefonicoService.montarObjNumeroCompleto(context, numero)
		if erroConversao.Erro != nil {
			return []*grpc.NumeroTelefonico{}, erroConversao
		}

		numerosTelefonicos = append(numerosTelefonicos, numeroTelefonico)
	}

	return numerosTelefonicos, erros.ErroStatus{}
}

// Função para buscar números telefônicos disponíveis em certo código de área
func (numeroTelefonicoService *NumeroTelefonicoService) FindNumerosTelefonicosDisponiveisByCodArea(context context.Context, codArea int32) ([]*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Busca todos os números telefônicos utilizáveis ou não no repositório
	tNumerosTelefonicos, err := numeroTelefonicoService.numeroTelefonicoRepository.FindDisponiveis(context, db.FindNumerosTelefonicosDisponiveisParams{Utilizavel: true, CodArea: int16(codArea)})
	if err != nil {
		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum número telefônico, retorna code NotFound
	if len(tNumerosTelefonicos) == 0 {
		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum número telefônico encontrado"),
		}
	}

	var numerosTelefonicos []*grpc.NumeroTelefonico
	for _, numero := range tNumerosTelefonicos {
		if numero.CongeladoAte.Time.After(time.Now()) {
			continue
		}

		numeroTelefonico, erroConversao := numeroTelefonicoService.montarObjNumeroCompleto(context, numero)
		if erroConversao.Erro != nil {
			return []*grpc.NumeroTelefonico{}, erroConversao
		}

		numerosTelefonicos = append(numerosTelefonicos, numeroTelefonico)
	}

	return numerosTelefonicos, erros.ErroStatus{}
}

// Função para buscar números telefônicos disponíveis em certo código de área de certa cidade
func (numeroTelefonicoService *NumeroTelefonicoService) FindNumerosTelefonicosDisponiveisByCidade(context context.Context, codIbge int32) ([]*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Busca a cidade dona do código do IBGE recebido
	tCidade, err := numeroTelefonicoService.cidadeRepository.FindByCodIbge(context, codIbge)
	if err != nil {
		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Busca todos os números telefônicos utilizáveis ou não no código de área da cidade encontrada
	return numeroTelefonicoService.FindNumerosTelefonicosDisponiveisByCodArea(context, tCidade.CodArea)
}

// Função para buscar números telefônicos disponíveis em certo código de área
func (numeroTelefonicoService *NumeroTelefonicoService) ReservarNumerosTelefonicosDisponiveisByCodArea(context context.Context, codArea int32) ([]*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Busca todos os números telefônicos utilizáveis ou não no repositório
	tNumerosTelefonicos, err := numeroTelefonicoService.numeroTelefonicoRepository.FindDisponiveis(context, db.FindNumerosTelefonicosDisponiveisParams{Utilizavel: true, CodArea: int16(codArea)})
	if err != nil {
		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum número telefônico, retorna code NotFound
	if len(tNumerosTelefonicos) == 0 {
		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum número telefônico encontrado"),
		}
	}

	var numerosTelefonicos []*grpc.NumeroTelefonico
	for _, numero := range tNumerosTelefonicos {
		if numero.CongeladoAte.Time.After(time.Now()) {
			continue
		}

		numeroReservaUpdate := db.UpdateNumeroTelefonicoParams{
			CodArea:               numero.CodArea,
			Numero:                numero.Numero,
			Utilizavel:            numero.Utilizavel,
			PortadoIn:             numero.PortadoIn,
			PortadoInOperadoraID:  numero.PortadoInOperadoraID,
			PortadoInDate:         numero.PortadoInDate,
			CodigoCnl:             numero.CodigoCnl,
			CongeladoAte:          pgtype.Date{Time: time.Now().Add(10 * time.Minute), Valid: true},
			ExternalID:            numero.ExternalID,
			PortadoOut:            numero.PortadoOut,
			PortadoOutOperadoraID: numero.PortadoOutOperadoraID,
			PortadoOutDate:        numero.PortadoOutDate,
			DataCriacao:           numero.DataCriacao,
			SimCardID:             numero.SimCardID,
			ID:                    numero.ID,
		}

		// Salva o número telefônico atualizado no repositório
		numeroReservado, err := numeroTelefonicoService.numeroTelefonicoRepository.Update(context, numeroReservaUpdate)
		if err != nil {
			return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		numeroTelefonico, erroConversao := numeroTelefonicoService.montarObjNumeroCompleto(context, numeroReservado)
		if erroConversao.Erro != nil {
			return []*grpc.NumeroTelefonico{}, erroConversao
		}

		numerosTelefonicos = append(numerosTelefonicos, numeroTelefonico)
	}

	return numerosTelefonicos, erros.ErroStatus{}
}

// Função para buscar números telefônicos disponíveis em certo código de área de certa cidade
func (numeroTelefonicoService *NumeroTelefonicoService) ReservarNumerosTelefonicosDisponiveisByCidade(context context.Context, codIbge int32) ([]*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Busca a cidade dona do código do IBGE recebido
	tCidade, err := numeroTelefonicoService.cidadeRepository.FindByCodIbge(context, codIbge)
	if err != nil {
		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Busca todos os números telefônicos utilizáveis ou não no código de área da cidade encontrada
	return numeroTelefonicoService.ReservarNumerosTelefonicosDisponiveisByCodArea(context, tCidade.CodArea)
}

// Função para buscar números telefônicos disponíveis em certo código de área
func (numeroTelefonicoService *NumeroTelefonicoService) FindNumerosTelefonicosByDocumentoCliente(context context.Context, documento string) ([]*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Busca todos os números telefônicos utilizáveis ou não no repositório
	tCliente, err := numeroTelefonicoService.clienteRepository.FindByDocumento(context, documento)
	if err != nil {
		// Caso não seja encontrado nenhum cliente, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Cliente não encontrado"),
			}
		}

		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	tNumerosTelefonicos, err := numeroTelefonicoService.numeroTelefonicoRepository.FindByClienteId(context, tCliente.ID)
	if err != nil {
		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum número telefônico, retorna code NotFound
	if len(tNumerosTelefonicos) == 0 {
		return []*grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum número telefônico encontrado"),
		}
	}

	var numerosTelefonicos []*grpc.NumeroTelefonico
	for _, numero := range tNumerosTelefonicos {
		numeroTelefonico, erroConversao := numeroTelefonicoService.montarObjNumeroCompleto(context, numero)
		if erroConversao.Erro != nil {
			return []*grpc.NumeroTelefonico{}, erroConversao
		}

		numerosTelefonicos = append(numerosTelefonicos, numeroTelefonico)
	}

	return numerosTelefonicos, erros.ErroStatus{}
}

// Função para criar um novo número telefônico
func (numeroTelefonicoService *NumeroTelefonicoService) CreateNumeroTelefonico(context context.Context, numeroTelefonico *grpc.NumeroTelefonico) (*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Busca um número telefônico pelo número enviado para verificar a prévia existência dele
	// Em caso positivo, retorna code AlreadyExists
	mesmaArea, erroBuscaPreExistente := numeroTelefonicoService.numeroTelefonicoRepository.FindByCodArea(context, numeroTelefonico.GetCodArea())
	if erroBuscaPreExistente != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroBuscaPreExistente,
		}
	}

	for _, numero := range mesmaArea {
		if numero.Numero == numeroTelefonico.GetNumero() {
			return nil, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe número telefônico com o número enviado"),
			}
		}
	}

	var err error

	// Faz a conversão das datas string que vêm do modelo de comunicação (grpc) para datas Time utilizadas no modelo de gravação no banco (sqlc)
	var portadoInDate time.Time
	if numeroTelefonico.GetPortadoInDate() != "" {
		portadoInDate, err = time.Parse("2006/01/02", numeroTelefonico.GetPortadoInDate())
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	} else {
		portadoInDate = time.Time{}
	}

	var congeladoAte time.Time
	if numeroTelefonico.GetPortadoInDate() != "" {
		congeladoAte, err = time.Parse("2006/01/02", numeroTelefonico.GetCongeladoAte())
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	} else {
		congeladoAte = time.Time{}
	}

	var portadoOutDate time.Time
	if numeroTelefonico.GetPortadoInDate() != "" {
		portadoOutDate, err = time.Parse("2006/01/02", numeroTelefonico.GetPortadoOutDate())
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	} else {
		portadoOutDate = time.Time{}
	}

	dataCriacao, err := time.Parse("2006/01/02", time.Now().Format("2006/01/02"))
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Faz a conversão dos IDs int32 que vêm do modelo de comunicação (grpc) para int4 do postgres utilizado no modelo de gravação no banco (sqlc)
	var portadoInOperadoraID pgtype.Int4
	if numeroTelefonico.GetPortadoInOperadora().GetID() != 0 {
		portadoInOperadoraID = pgtype.Int4{Int32: numeroTelefonico.GetPortadoInOperadora().GetID(), Valid: true}
	} else {
		portadoInOperadoraID = pgtype.Int4{Int32: numeroTelefonico.GetPortadoInOperadora().GetID(), Valid: false}
	}

	var portadoOutOperadoraID pgtype.Int4
	if numeroTelefonico.GetPortadoOutOperadora().GetID() != 0 {
		portadoOutOperadoraID = pgtype.Int4{Int32: numeroTelefonico.GetPortadoOutOperadora().GetID(), Valid: true}
	} else {
		portadoOutOperadoraID = pgtype.Int4{Int32: numeroTelefonico.GetPortadoOutOperadora().GetID(), Valid: false}
	}

	var simCardID pgtype.Int4
	if numeroTelefonico.GetSimCard().GetID() != 0 {
		simCardID = pgtype.Int4{Int32: numeroTelefonico.GetSimCard().GetID(), Valid: true}
	} else {
		simCardID = pgtype.Int4{Int32: numeroTelefonico.GetSimCard().GetID(), Valid: false}
	}

	// Cria o objeto CreateNumeroTelefonicoParams gerado pelo sqlc para gravação no banco de dados
	numeroTelefonicoCreate := db.CreateNumeroTelefonicoParams{
		CodArea:               int16(numeroTelefonico.GetCodArea()),
		Numero:                numeroTelefonico.GetNumero(),
		Utilizavel:            numeroTelefonico.GetUtilizavel(),
		PortadoIn:             numeroTelefonico.GetPortadoIn(),
		PortadoInOperadoraID:  portadoInOperadoraID,
		PortadoInDate:         pgtype.Timestamptz{Time: portadoInDate, Valid: true},
		CodigoCnl:             numeroTelefonico.GetCodigoCNL(),
		CongeladoAte:          pgtype.Date{Time: congeladoAte, Valid: true},
		ExternalID:            pgtype.Int4{Int32: numeroTelefonico.GetExternalID(), Valid: true},
		PortadoOut:            numeroTelefonico.GetPortadoOut(),
		PortadoOutOperadoraID: portadoOutOperadoraID,
		PortadoOutDate:        pgtype.Timestamptz{Time: portadoOutDate, Valid: true},
		DataCriacao:           pgtype.Timestamptz{Time: dataCriacao, Valid: true},
		SimCardID:             simCardID,
	}

	// Cria o número telefônico no repositório
	numeroTelefonicoCriado, err := numeroTelefonicoService.numeroTelefonicoRepository.Create(context, numeroTelefonicoCreate)
	if err != nil {
		return &grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return numeroTelefonicoService.montarObjNumeroCompleto(context, numeroTelefonicoCriado)
}

// Função para vincular um SimCard a um número telefônico
func (numeroTelefonicoService *NumeroTelefonicoService) VincularSimCard(context context.Context, idSimCard int32, idNumeroTelefonico int32) (*grpc.NumeroTelefonico, erros.ErroStatus) {
	_, numeroTelefonico, err := numeroTelefonicoService.simCardTelefoneService.Vincular(context, idSimCard, idNumeroTelefonico)
	if err.Erro != nil {
		return &grpc.NumeroTelefonico{}, err
	}

	return numeroTelefonicoService.montarObjNumeroCompleto(context, numeroTelefonico)
}

// Função para atualizar um número telefônico existente
func (numeroTelefonicoService *NumeroTelefonicoService) UpdateNumeroTelefonico(context context.Context, numeroTelefonicoRecebido *grpc.NumeroTelefonico, numeroTelefonicoAntigo *grpc.NumeroTelefonico) (*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Verifica se o número foi modificado e, se sim, verifica se já existe outro registro com o mesmo número
	// Em caso positivo, retorna code AlreadyExists
	if numeroTelefonicoAntigo.GetNumero() != numeroTelefonicoRecebido.GetNumero() || numeroTelefonicoAntigo.GetCodArea() != numeroTelefonicoRecebido.GetCodArea() {
		mesmaArea, erroBuscaPreExistente := numeroTelefonicoService.numeroTelefonicoRepository.FindByCodArea(context, numeroTelefonicoRecebido.GetCodArea())
		if erroBuscaPreExistente != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   erroBuscaPreExistente,
			}
		}

		for _, numero := range mesmaArea {
			if numero.Numero == numeroTelefonicoRecebido.GetNumero() {
				return nil, erros.ErroStatus{
					Status: codes.AlreadyExists,
					Erro:   errors.New("Já existe número telefônico com o número enviado"),
				}
			}
		}
	}

	var err error

	// Faz a conversão das datas string que vêm do modelo de comunicação (grpc) para datas Time utilizadas no modelo de gravação no banco (sqlc)
	var portadoInDate time.Time
	if numeroTelefonicoRecebido.GetPortadoInDate() != "" {
		portadoInDate, err = time.Parse("2006/01/02", numeroTelefonicoRecebido.GetPortadoInDate())
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	} else {
		portadoInDate = time.Time{}
	}

	var congeladoAte time.Time
	if numeroTelefonicoRecebido.GetPortadoInDate() != "" {
		congeladoAte, err = time.Parse("2006/01/02", numeroTelefonicoRecebido.GetCongeladoAte())
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	} else {
		congeladoAte = time.Time{}
	}

	var portadoOutDate time.Time
	if numeroTelefonicoRecebido.GetPortadoInDate() != "" {
		portadoOutDate, err = time.Parse("2006/01/02", numeroTelefonicoRecebido.GetPortadoOutDate())
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	} else {
		portadoOutDate = time.Time{}
	}

	dataCriacao, err := time.Parse("2006/01/02", time.Now().Format("2006/01/02"))
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Faz a conversão dos IDs int32 que vêm do modelo de comunicação (grpc) para int4 do postgres utilizado no modelo de gravação no banco (sqlc)
	var portadoInOperadoraID pgtype.Int4
	if numeroTelefonicoRecebido.GetPortadoInOperadora().GetID() != 0 {
		portadoInOperadoraID = pgtype.Int4{Int32: numeroTelefonicoRecebido.GetPortadoInOperadora().GetID(), Valid: true}
	} else {
		portadoInOperadoraID = pgtype.Int4{Int32: numeroTelefonicoRecebido.GetPortadoInOperadora().GetID(), Valid: false}
	}

	var portadoOutOperadoraID pgtype.Int4
	if numeroTelefonicoRecebido.GetPortadoOutOperadora().GetID() != 0 {
		portadoOutOperadoraID = pgtype.Int4{Int32: numeroTelefonicoRecebido.GetPortadoOutOperadora().GetID(), Valid: true}
	} else {
		portadoOutOperadoraID = pgtype.Int4{Int32: numeroTelefonicoRecebido.GetPortadoOutOperadora().GetID(), Valid: false}
	}

	var simCardID pgtype.Int4
	if numeroTelefonicoRecebido.GetSimCard().GetID() != 0 {
		simCardID = pgtype.Int4{Int32: numeroTelefonicoRecebido.GetSimCard().GetID(), Valid: true}
	} else {
		simCardID = pgtype.Int4{Int32: numeroTelefonicoRecebido.GetSimCard().GetID(), Valid: false}
	}

	// Cria o objeto UpdateNumeroTelefonicoParams gerado pelo sqlc para gravação no banco de dados
	numeroTelefonicoUpdate := db.UpdateNumeroTelefonicoParams{
		CodArea:               int16(numeroTelefonicoRecebido.GetCodArea()),
		Numero:                numeroTelefonicoRecebido.GetNumero(),
		Utilizavel:            numeroTelefonicoRecebido.GetUtilizavel(),
		PortadoIn:             numeroTelefonicoRecebido.GetPortadoIn(),
		PortadoInOperadoraID:  portadoInOperadoraID,
		PortadoInDate:         pgtype.Timestamptz{Time: portadoInDate, Valid: true},
		CodigoCnl:             numeroTelefonicoRecebido.GetCodigoCNL(),
		CongeladoAte:          pgtype.Date{Time: congeladoAte, Valid: true},
		ExternalID:            pgtype.Int4{Int32: numeroTelefonicoRecebido.GetExternalID(), Valid: true},
		PortadoOut:            numeroTelefonicoRecebido.GetPortadoOut(),
		PortadoOutOperadoraID: portadoOutOperadoraID,
		PortadoOutDate:        pgtype.Timestamptz{Time: portadoOutDate, Valid: true},
		DataCriacao:           pgtype.Timestamptz{Time: dataCriacao, Valid: true},
		SimCardID:             simCardID,
		ID:                    numeroTelefonicoRecebido.GetID(),
	}

	// Salva o número telefônico atualizado no repositório
	numeroTelefonicoSalvo, err := numeroTelefonicoService.numeroTelefonicoRepository.Update(context, numeroTelefonicoUpdate)
	if err != nil {
		return &grpc.NumeroTelefonico{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return numeroTelefonicoService.montarObjNumeroCompleto(context, numeroTelefonicoSalvo)
}

// Função para deletar um número telefônico pelo ID
func (numeroTelefonicoService *NumeroTelefonicoService) DeleteNumeroTelefonicoById(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Deleta o número telefônico no repositório
	deletado, err := numeroTelefonicoService.numeroTelefonicoRepository.Delete(context, id)
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
			Erro:   errors.New("Nenhum número telefônico encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}

// Essa é uma função privada do serviço para encapsular a montagem do objeto de retorno para o server/controller
func (numeroTelefonicoService *NumeroTelefonicoService) montarObjNumeroCompleto(context context.Context, numeroTelefonico db.TTelefoniaNumero) (*grpc.NumeroTelefonico, erros.ErroStatus) {
	// Faz a montagem do objeto operadora de entrada de portabilidade, caso necessário
	operadoraIn := new(grpc.Operadora)
	if numeroTelefonico.PortadoInOperadoraID.Int32 != 0 {
		tOperadoraIn, err := numeroTelefonicoService.operadoraRepository.FindByID(context, numeroTelefonico.PortadoInOperadoraID.Int32)
		if err != nil {
			return &grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
		operadoraIn = helpers.TOperadoraToPb(tOperadoraIn)
	}

	// Faz a montagem do objeto operadora de saída de portabilidade, caso necessário
	operadoraOut := new(grpc.Operadora)
	if numeroTelefonico.PortadoOutOperadoraID.Int32 != 0 {
		tOperadoraOut, err := numeroTelefonicoService.operadoraRepository.FindByID(context, numeroTelefonico.PortadoInOperadoraID.Int32)
		if err != nil {
			return &grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
		operadoraOut = helpers.TOperadoraToPb(tOperadoraOut)
	}

	// Faz a montagem do objeto SimCard, caso necessário
	simCard := new(grpc.SimCard)
	if numeroTelefonico.SimCardID.Int32 != 0 {
		tSimCard, err := numeroTelefonicoService.simCardRepository.FindByID(context, numeroTelefonico.SimCardID.Int32)
		if err != nil {
			return &grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// Faz a montagem do objeto estado de SimCard para deixar o objeto SimCard completo
		tEstadoSimCard, err := numeroTelefonicoService.simCardEstadoRepository.FindByID(context, int32(tSimCard.StatusID))
		if err != nil {
			return &grpc.NumeroTelefonico{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// O último campo é nulo pois seria o lugar do número telefônico ligado ao simcard, porém é justamente isso que estou montando aqui
		// Passar esse valor para a conversão faria com que isso se tornasse uma função recursiva e entraria em um loop infitino
		simCard = helpers.TSimCardToPb(tSimCard, helpers.TStatusSimCardToPb(tEstadoSimCard), nil)
	}

	// Passando todos os objetos completinhos para o helper que irá converter o objeto completo para mim
	return helpers.TTelefoniaNumeroToPb(numeroTelefonico, operadoraIn, operadoraOut, simCard), erros.ErroStatus{}
}
