package repositories

import (
	"context"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"

	"github.com/jackc/pgx/v5/pgtype"
)

type ClienteRepository interface {
	FindAll(context context.Context) ([]db.TCliente, error)
	FindByID(context context.Context, id int32) (db.TCliente, error)
	FindByIDExterno(context context.Context, id int32) (db.TCliente, error)
	FindByNome(context context.Context, nome string) (db.TCliente, error)
	FindByDocumento(context context.Context, documento string) (db.TCliente, error)
	FindByCodReserva(context context.Context, codIbge string) (db.TCliente, error)
	Create(context context.Context, cliente db.CreateClienteParams) (db.TCliente, error)
	Update(context context.Context, cliente db.UpdateClienteParams) (db.TCliente, error)
	Delete(context context.Context, id int32) (int64, error)
}

type clienteRepository struct {
	*db.Queries
}

func NewClienteRepository(queries *db.Queries) ClienteRepository {
	return &clienteRepository{
		Queries: queries,
	}
}

func (clienteRepository *clienteRepository) FindAll(ctx context.Context) ([]db.TCliente, error) {
	return clienteRepository.FindAllClientes(ctx)

}

func (clienteRepository *clienteRepository) FindByID(context context.Context, id int32) (db.TCliente, error) {
	cliente, err := clienteRepository.FindClienteByID(context, id)
	if err != nil {
		return db.TCliente{}, err
	}

	return cliente, nil
}

func (clienteRepository *clienteRepository) FindByIDExterno(context context.Context, id int32) (db.TCliente, error) {
	cliente, err := clienteRepository.FindClienteByIDExterno(context, id)
	if err != nil {
		return db.TCliente{}, err
	}

	return cliente, nil
}

func (clienteRepository *clienteRepository) FindByNome(context context.Context, nome string) (db.TCliente, error) {
	cliente, err := clienteRepository.FindClienteByNome(context, nome)
	if err != nil {
		return db.TCliente{}, err
	}

	return cliente, nil
}

func (clienteRepository *clienteRepository) FindByDocumento(context context.Context, documento string) (db.TCliente, error) {
	clientes, err := clienteRepository.FindClienteByDocumento(context, documento)
	if err != nil {
		return db.TCliente{}, err
	}

	return clientes, nil
}

func (clienteRepository *clienteRepository) FindByCodReserva(context context.Context, codIbge string) (db.TCliente, error) {
	cliente, err := clienteRepository.FindClienteByCodReserva(context, pgtype.Text{String: codIbge, Valid: true})
	if err != nil {
		return db.TCliente{}, err
	}

	return cliente, nil
}

func (clienteRepository *clienteRepository) Create(context context.Context, cliente db.CreateClienteParams) (db.TCliente, error) {
	clienteCriada, err := clienteRepository.CreateCliente(context, cliente)
	if err != nil {
		return db.TCliente{}, err
	}

	return clienteCriada, nil
}

func (clienteRepository *clienteRepository) Update(context context.Context, cliente db.UpdateClienteParams) (db.TCliente, error) {
	clienteAtualizada, err := clienteRepository.UpdateCliente(context, cliente)
	if err != nil {
		return db.TCliente{}, err
	}

	return clienteAtualizada, nil
}

func (clienteRepository *clienteRepository) Delete(context context.Context, id int32) (int64, error) {
	linhasAfetadas, err := clienteRepository.DeleteClienteById(context, id)
	if err != nil {
		return 0, err
	}

	return linhasAfetadas, nil
}
