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

// Estrutura de serviço para gerenciar operações relacionadas às clientes
type ClienteService struct {
	clienteRepository repositories.ClienteRepository
}

// Função para criar uma nova instância de ClienteService com o repositório necessário
func NewClienteService(clienteRepository repositories.ClienteRepository) *ClienteService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &ClienteService{
		clienteRepository: clienteRepository,
	}
}

// Função para buscar um cliente pelo ID
func (clienteService *ClienteService) FindClienteById(context context.Context, id int32) (*pb.Cliente, erros.ErroStatus) {
	// Busca o cliente no repositório pelo ID
	cliente, err := clienteService.clienteRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum cliente, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Cliente{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Cliente não encontrado"),
			}
		}

		return &pb.Cliente{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TClienteToPb(cliente), erros.ErroStatus{}
}

// Função para buscar um cliente pelo ID externo
func (clienteService *ClienteService) FindClienteByIdExterno(context context.Context, id int32) (*pb.Cliente, erros.ErroStatus) {
	// Busca o cliente no repositório pelo ID
	cliente, err := clienteService.clienteRepository.FindByIDExterno(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum cliente, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Cliente{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Cliente não encontrado"),
			}
		}

		return &pb.Cliente{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TClienteToPb(cliente), erros.ErroStatus{}
}

// Função para buscar um cliente pelo nome
func (clienteService *ClienteService) FindClienteByNome(context context.Context, nome string) (*pb.Cliente, erros.ErroStatus) {
	// Busca o cliente no repositório pelo nome
	cliente, err := clienteService.clienteRepository.FindByNome(context, nome)
	if err != nil {
		// Caso não seja encontrado nenhum cliente, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Cliente{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Cliente não encontrado"),
			}
		}

		return &pb.Cliente{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TClienteToPb(cliente), erros.ErroStatus{}
}

// Função para buscar um cliente pelo documento
func (clienteService *ClienteService) FindClienteByDocumento(context context.Context, documento string) (*pb.Cliente, erros.ErroStatus) {
	// Busca o cliente no repositório pelo nome
	cliente, err := clienteService.clienteRepository.FindByDocumento(context, documento)
	if err != nil {
		// Caso não seja encontrado nenhum cliente, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Cliente{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Cliente não encontrado"),
			}
		}

		return &pb.Cliente{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TClienteToPb(cliente), erros.ErroStatus{}
}

// Função para buscar um cliente pelo código de reserva
func (clienteService *ClienteService) FindClienteByCodReserva(context context.Context, codReserva string) (*pb.Cliente, erros.ErroStatus) {
	// Busca o cliente no repositório pelo código de reserva
	cliente, err := clienteService.clienteRepository.FindByCodReserva(context, codReserva)
	if err != nil {
		// Caso não seja encontrado nenhum cliente, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.Cliente{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Cliente não encontrado"),
			}
		}

		return &pb.Cliente{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TClienteToPb(cliente), erros.ErroStatus{}
}

// Função para buscar todas as clientes
func (clienteService *ClienteService) FindAllClientes(context context.Context) ([]*pb.Cliente, erros.ErroStatus) {
	// Busca todas as clientes no repositório
	clientes, err := clienteService.clienteRepository.FindAll(context)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum cliente, retorna code NotFound
	if len(clientes) == 0 {
		return nil, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum cliente encontrado"),
		}
	}

	var pbClientes []*pb.Cliente
	for _, cliente := range clientes {
		pbClientes = append(pbClientes, helpers.TClienteToPb(cliente))
	}

	return pbClientes, erros.ErroStatus{}
}

// Função para criar uma nova cliente
func (clienteService *ClienteService) CreateCliente(context context.Context, cliente *pb.Cliente) (*pb.Cliente, erros.ErroStatus) {
	// Busca um cliente pelo documento enviado para verificar a prévia existência dele
	// Em caso positivo, retorna code AlreadyExists
	_, erroBuscaPreExistente := clienteService.FindClienteByDocumento(context, cliente.GetDocumento())
	if erroBuscaPreExistente.Erro == nil {
		return nil, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("Já existe cliente com o documento enviado"),
		}
	}

	// Cria o objeto CreateClienteParams gerado pelo sqlc para gravação no banco de dados
	clienteCreate := db.CreateClienteParams{
		Uuid:       pgtype.Text{String: cliente.GetUUID(), Valid: true},
		IDExterno:  cliente.GetIDExterno(),
		Nome:       cliente.GetNome(),
		Documento:  cliente.GetDocumento(),
		CodReserva: pgtype.Text{String: cliente.GetCodReserva(), Valid: true},
	}

	// Cria o cliente no repositório
	clienteCriada, err := clienteService.clienteRepository.Create(context, clienteCreate)
	if err != nil {
		return &pb.Cliente{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return helpers.TClienteToPb(clienteCriada), erros.ErroStatus{}
}

// Função para atualizar um cliente existente
func (clienteService *ClienteService) UpdateCliente(context context.Context, clienteRecebido *pb.Cliente, clienteBanco *pb.Cliente) (*pb.Cliente, erros.ErroStatus) {
	// Verifica se o documento foi modificado e, se sim, verifica se já existe outro registro com o mesmo documento
	// Em caso positivo, retorna code AlreadyExists
	if clienteBanco.GetDocumento() != clienteRecebido.GetDocumento() {
		_, erroBuscaPreExistente := clienteService.FindClienteByDocumento(context, clienteRecebido.GetDocumento())
		if erroBuscaPreExistente.Erro == nil {
			return nil, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("Já existe cliente com o documento enviado"),
			}
		}
	}

	// Cria o objeto UpdateClienteParams gerado pelo sqlc para gravação no banco de dados
	clienteUpdate := db.UpdateClienteParams{
		Uuid:       pgtype.Text{String: clienteRecebido.GetUUID(), Valid: true},
		IDExterno:  clienteRecebido.GetIDExterno(),
		Nome:       clienteRecebido.GetNome(),
		Documento:  clienteRecebido.GetDocumento(),
		CodReserva: pgtype.Text{String: clienteRecebido.GetCodReserva(), Valid: true},
		ID:         clienteRecebido.GetID(),
	}

	// Salva o cliente atualizada no repositório
	clienteAtualizada, errSalvamento := clienteService.clienteRepository.Update(context, clienteUpdate)
	if errSalvamento != nil {
		return &pb.Cliente{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errSalvamento,
		}
	}

	return helpers.TClienteToPb(clienteAtualizada), erros.ErroStatus{}
}

// Função para deletar um cliente pelo ID
func (clienteService *ClienteService) DeleteClienteById(context context.Context, id int32) (bool, erros.ErroStatus) {
	// Deleta o cliente no repositório pelo ID
	deletados, err := clienteService.clienteRepository.Delete(context, id)

	// Caso ocorra erro na deleção
	if err != nil {
		return false, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// O atributo deletados indica o número de linhas deletadas. Se for 0, nenhum cliente foi deletada, pois não existia
	if deletados == 0 {
		return false, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum cliente encontrado"),
		}
	}

	return true, erros.ErroStatus{}
}
