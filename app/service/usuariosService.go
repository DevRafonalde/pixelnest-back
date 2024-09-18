package service

import (
	"context"
	"encoding/base64"
	"errors"
	"pixelnest/app/helpers"
	"pixelnest/app/model/erros"
	pb "pixelnest/app/model/grpc"
	"pixelnest/app/model/repositories"
	db "pixelnest/app/model/repositories/sqlc/repositoryIMPL"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
)

// Estrutura que define o serviço de usuário, contendo os repositórios necessários
type UsuarioService struct {
	perfilRepository        repositories.PerfilRepository
	usuarioPerfilRepository repositories.UsuarioPerfilRepository
	usuarioRepository       repositories.UsuarioRepository
}

// Função para criar uma nova instância de UsuarioService com os repositórios fornecidos
func NewUsuarioService(perfilRepository repositories.PerfilRepository, usuarioPerfilRepository repositories.UsuarioPerfilRepository, usuarioRepository repositories.UsuarioRepository) *UsuarioService {
	time.Local, _ = time.LoadLocation("America/Sao_Paulo") // Define o fuso horário local
	return &UsuarioService{
		perfilRepository:        perfilRepository,
		usuarioPerfilRepository: usuarioPerfilRepository,
		usuarioRepository:       usuarioRepository,
	}
}

// Função para buscar todos os usuários
// Como um usuário pode estar vinculado a vários perfis,
// Nessa função, os usuários são retornados sem as perfis vinculados
func (usuarioService *UsuarioService) FindAllUsuarios(context context.Context) ([]*pb.Usuario, erros.ErroStatus) {
	usuarios, err := usuarioService.usuarioRepository.FindAll(context)
	if err != nil {
		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Caso não seja encontrado nenhum usuário, retorna code NotFound
	if len(usuarios) == 0 {
		return []*pb.Usuario{}, erros.ErroStatus{
			Status: codes.NotFound,
			Erro:   errors.New("Nenhum usuário encontrado"),
		}
	}

	var pbUsuarios []*pb.Usuario
	for _, usuario := range usuarios {
		pbUsuarios = append(pbUsuarios, helpers.TUsuarioToPb(usuario))
	}

	return pbUsuarios, erros.ErroStatus{}
}

// Função para buscar um usuário por ID
// Diferente da busca por todos, aqui a aplicação retorna o usuário com todos os perfis vinculados
func (usuarioService *UsuarioService) FindUsuarioById(context context.Context, id int32) (*pb.UsuarioPerfis, erros.ErroStatus) {
	usuario, err := usuarioService.usuarioRepository.FindByID(context, id)
	if err != nil {
		// Caso não seja encontrado nenhum usuário, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário não encontrado"),
			}
		}

		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	perfis, erroPerfis := usuarioService.GetPerfisVinculados(context, usuario.ID)
	if erroPerfis.Erro != nil {
		return &pb.UsuarioPerfis{}, erroPerfis
	}

	var pbPerfis []*pb.Perfil
	for i := range perfis {
		pbPerfis = append(pbPerfis, &perfis[i])
	}

	retorno := pb.UsuarioPerfis{
		Usuario: helpers.TUsuarioToPb(usuario),
		Perfis:  pbPerfis,
	}

	return &retorno, erros.ErroStatus{}
}

// Função para buscar perfis vinculados a um usuário, filtrando apenas os perfis ativos
func (usuarioService *UsuarioService) GetPerfisVinculados(context context.Context, id int32) ([]pb.Perfil, erros.ErroStatus) {
	usuarioPerfis, errBuscaPerfis := usuarioService.usuarioPerfilRepository.FindByUsuario(context, id)
	if errBuscaPerfis != nil {
		// Caso não seja encontrado nenhum usuário, retorna code NotFound
		if errBuscaPerfis.Error() == "no rows in result set" {
			return []pb.Perfil{}, erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário não encontrado"),
			}
		}

		return nil, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errBuscaPerfis,
		}
	}

	// Filtra apenas os perfis que estão vinculados e ativos
	var perfis []pb.Perfil
	for _, usuarioPerfil := range usuarioPerfis {
		perfilEncontrado, err := usuarioService.perfilRepository.FindByID(context, usuarioPerfil.PerfilID)
		if err != nil {
			return nil, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		if perfilEncontrado.Ativo.Bool {
			perfis = append(perfis, *helpers.TPerfToPb(perfilEncontrado))
		}
	}

	return perfis, erros.ErroStatus{}
}

// Função de login, que verifica o email, a senha do usuário e se ele está ativo ou não
// Retorna o usuário com seus perfis vinculados (UsuarioPerfis)
func (usuarioService *UsuarioService) Login(context context.Context, loginUsuario *pb.LoginUsuario) (*pb.UsuarioPerfis, erros.ErroStatus) {
	usuarioBanco, err := usuarioService.usuarioRepository.FindByEmail(context, loginUsuario.Email)
	if err != nil {
		// Caso não seja encontrado nenhum usuário, retorna code Unauthenticated
		// Dessa forma, o usuário que está tentando acessar não consegue saber se o que ele errou foi o e-mail ou a senha
		// Assim acredito que fique mais seguro contra tentativas maliciosas de acesso
		if err.Error() == "no rows in result set" {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.Unauthenticated,
				Erro:   errors.New("E-mail e/ou senha incorretos"),
			}
		}

		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	if !usuarioBanco.Ativo.Bool {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.InvalidArgument,
			Erro:   errors.New("Usuário está desativado"),
		}
	}

	perfisVinculados, erroGetPerfis := usuarioService.GetPerfisVinculados(context, usuarioBanco.ID)
	if erroGetPerfis.Erro != nil {
		return &pb.UsuarioPerfis{}, erroGetPerfis
	}

	var pbPerfis []*pb.Perfil
	for i := range perfisVinculados {
		pbPerfis = append(pbPerfis, &perfisVinculados[i])
	}

	usuarioPerfil := new(pb.UsuarioPerfis)
	usuarioPerfil.Usuario = helpers.TUsuarioToPb(usuarioBanco)
	usuarioPerfil.Perfis = pbPerfis

	// Validação da senha codificada do banco
	senhaHash, err := base64.StdEncoding.DecodeString(usuarioBanco.Senha)
	if err != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	erroValidacao := bcrypt.CompareHashAndPassword(senhaHash, []byte(loginUsuario.Senha))
	if erroValidacao != nil {
		// Caso a senha esteja incorreta, não informa que é a senha o problema, mas sim um erro genérico de login
		// Dessa forma, o usuário que está tentando acessar não consegue saber se o que ele errou foi o e-mail ou a senha
		// Assim acredito que fique mais seguro contra tentativas maliciosas de acesso
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Unauthenticated,
			Erro:   errors.New("E-mail e/ou senha incorretos"),
		}
	} else {
		return usuarioPerfil, erros.ErroStatus{}
	}
}

// Essa função insere na base de dados o token criado para o reset da senha do usuário em caso de esquecimento
// Retorna apenas um erro caso aconteça algum
func (usuarioService *UsuarioService) TokenResetSenha(context context.Context, token string, email string) erros.ErroStatus {
	usuarioBanco, err := usuarioService.usuarioRepository.FindByEmail(context, email)
	if err != nil {
		// Caso não seja encontrado nenhum usuário, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário não encontrado"),
			}
		}

		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Cria o objeto UpdateUsuarioParams gerado pelo sqlc para gravação no banco de dados
	usuarioUpdate := db.UpdateUsuarioParams{
		ID:                    usuarioBanco.ID,
		Nome:                  usuarioBanco.Nome,
		Email:                 usuarioBanco.Email,
		Senha:                 usuarioBanco.Senha,
		Ativo:                 usuarioBanco.Ativo,
		TokenResetSenha:       pgtype.Text{String: token, Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		SenhaAtualizada:       usuarioBanco.SenhaAtualizada,
	}

	// Salva o usuário atualizado no repositório
	_, err = usuarioService.usuarioRepository.Update(context, usuarioUpdate)
	if err != nil {
		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return erros.ErroStatus{}
}

// Função para atualizar a senha de um usuário
// Retorna apenas um erro caso aconteça algum
func (usuarioService *UsuarioService) AtualizarSenha(context context.Context, email string, senhaNova string) erros.ErroStatus {
	usuarioEncontrado, err := usuarioService.usuarioRepository.FindByEmail(context, email)
	if err != nil {
		// Caso não seja encontrado nenhum usuário, retorna code NotFound
		if err.Error() == "no rows in result set" {
			return erros.ErroStatus{
				Status: codes.NotFound,
				Erro:   errors.New("Usuário não encontrado"),
			}
		}

		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Criptografia da senha
	senhaEmBytes := []byte(senhaNova)
	senhaCriptografada, erroCriptografia := bcrypt.GenerateFromPassword(senhaEmBytes, 10)
	if erroCriptografia != nil {
		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroCriptografia,
		}
	}

	// Cria o objeto UpdateUsuarioParams gerado pelo sqlc para gravação no banco de dados
	// Aqui eu defino a flag `SenhaAtualizada` como true para que, caso o usuário tente acessar a API novamente
	// Ele seja obrigado a fazer o login com a senha nova
	usuarioUpdate := db.UpdateUsuarioParams{
		ID:                    usuarioEncontrado.ID,
		Nome:                  usuarioEncontrado.Nome,
		Email:                 usuarioEncontrado.Email,
		Senha:                 base64.StdEncoding.EncodeToString(senhaCriptografada),
		Ativo:                 usuarioEncontrado.Ativo,
		TokenResetSenha:       usuarioEncontrado.TokenResetSenha,
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		SenhaAtualizada:       pgtype.Bool{Bool: true, Valid: true},
	}

	// Salva o usuário atualizado no repositório
	_, err = usuarioService.usuarioRepository.Update(context, usuarioUpdate)
	if err != nil {
		return erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	return erros.ErroStatus{}
}

// Função para clonar um usuário, mantendo as permissões, mas retornando um novo usuário vazio
func (usuarioService *UsuarioService) CloneUsuario(context context.Context, id int32) (*pb.UsuarioPerfis, erros.ErroStatus) {
	usuarioPerfil, err := usuarioService.FindUsuarioById(context, id)
	if err.Erro != nil {
		return &pb.UsuarioPerfis{}, err
	}

	// Define o usuário clonado como vazio
	usuarioPerfil.Usuario = &pb.Usuario{}

	return usuarioPerfil, erros.ErroStatus{}
}

// Função para criar um novo usuário, verificando se o e-mail já existe e criptografando a senha
func (usuarioService *UsuarioService) CreateUsuario(context context.Context, usuarioPerfilRecebido *pb.UsuarioPerfis) (*pb.UsuarioPerfis, erros.ErroStatus) {
	// Busca um usuário pelo e-mail enviado para verificar a prévia existência dele
	// Em caso positivo, retorna code AlreadyExists
	_, err := usuarioService.usuarioRepository.FindByEmail(context, usuarioPerfilRecebido.GetUsuario().GetEmail())
	if err == nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.AlreadyExists,
			Erro:   errors.New("E-mail já está sendo utilizado"),
		}
	}

	// Criptografia da senha
	senhaEmBytes := []byte(usuarioPerfilRecebido.GetUsuario().GetSenha())
	senhaCriptografada, erroCriptografia := bcrypt.GenerateFromPassword(senhaEmBytes, 10)
	if erroCriptografia != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   erroCriptografia,
		}
	}

	// Cria o objeto CreateUsuarioParams gerado pelo sqlc para gravação no banco de dados
	usuarioCreate := db.CreateUsuarioParams{
		Nome:                  usuarioPerfilRecebido.GetUsuario().GetNome(),
		Email:                 usuarioPerfilRecebido.GetUsuario().GetEmail(),
		Senha:                 base64.StdEncoding.EncodeToString(senhaCriptografada),
		Ativo:                 pgtype.Bool{Bool: true, Valid: true},
		TokenResetSenha:       pgtype.Text{String: usuarioPerfilRecebido.GetUsuario().GetTokenResetSenha(), Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		SenhaAtualizada:       pgtype.Bool{Bool: true, Valid: true},
	}

	// Cria o usuário no repositório
	usuario, err := usuarioService.usuarioRepository.Create(context, usuarioCreate)
	usuarioPerfilRecebido.Usuario = helpers.TUsuarioToPb(usuario)
	if err != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Salva a relação entre o usuário e cada perfil
	for i, perfil := range usuarioPerfilRecebido.GetPerfis() {
		perfil, err := usuarioService.perfilRepository.FindByID(context, perfil.ID)
		if err != nil {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// Coloca o objeto completo do perfil na lista de perfis do objeto recebido
		// Como esse é o objeto retornado, os perfis retornam completos e prontos para serem usados
		usuarioPerfilRecebido.Perfis[i] = helpers.TPerfToPb(perfil)

		// Cria o objeto CreateUsuarioPerfilParams gerado pelo sqlc para gravação no banco de dados
		usuarioPerfil := db.CreateUsuarioPerfilParams{
			UsuarioID: usuario.ID,
			PerfilID:  perfil.ID,
			DataHora:  pgtype.Timestamp{Time: time.Now(), Valid: true},
		}

		usuarioService.usuarioPerfilRepository.Create(context, usuarioPerfil)
	}

	return usuarioPerfilRecebido, erros.ErroStatus{}
}

// Função para atualizar um usuário, verificando o email e atualizando perfis relacionados
func (usuarioService *UsuarioService) UpdateUsuario(context context.Context, usuarioPerfil *pb.UsuarioPerfis, usuarioAntigo *pb.UsuarioPerfis) (*pb.UsuarioPerfis, erros.ErroStatus) {
	usuarioRecebido := usuarioPerfil.Usuario
	usuarioBanco := usuarioAntigo.GetUsuario()
	// Verifica se o e-mail foi modificado e, se sim, verifica se já existe outro registro com o mesmo e-mail
	// Em caso positivo, retorna code AlreadyExists
	if usuarioBanco.Email != usuarioRecebido.Email {
		_, err := usuarioService.usuarioRepository.FindByEmail(context, usuarioRecebido.Email)
		if err == nil {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.AlreadyExists,
				Erro:   errors.New("E-mail já está sendo utilizado"),
			}
		}
	}

	// Cria o objeto UpdateUsuarioParams gerado pelo sqlc para gravação no banco de dados
	usuarioUpdate := db.UpdateUsuarioParams{
		ID:                    usuarioBanco.ID,
		Nome:                  usuarioRecebido.Nome,
		Email:                 usuarioRecebido.Email,
		Senha:                 usuarioBanco.Senha,
		Ativo:                 pgtype.Bool{Bool: usuarioBanco.Ativo, Valid: true},
		TokenResetSenha:       pgtype.Text{String: usuarioBanco.TokenResetSenha, Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		SenhaAtualizada:       pgtype.Bool{Bool: usuarioBanco.SenhaAtualizada, Valid: true},
	}

	// Salva o usuário atualizado no repositório
	usuarioAtualizado, err := usuarioService.usuarioRepository.Update(context, usuarioUpdate)
	if err != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   err,
		}
	}

	// Deleta todas as relações antigas entre o usuário e seus perfis
	registrosExistentes, errBuscaPerfis := usuarioService.usuarioPerfilRepository.FindByUsuario(context, usuarioAtualizado.ID)
	if errBuscaPerfis != nil {
		return &pb.UsuarioPerfis{}, erros.ErroStatus{
			Status: codes.Internal,
			Erro:   errBuscaPerfis,
		}
	}

	for i := 0; i < len(registrosExistentes); i++ {
		err := usuarioService.usuarioPerfilRepository.Delete(context, registrosExistentes[i].ID)
		if err != nil {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}
	}

	// Cria novas relações entre o perfil atualizado e as novas permissões
	perfis := usuarioPerfil.Perfis

	for i, perfil := range perfis {
		tPerfil, err := usuarioService.perfilRepository.FindByID(context, perfil.ID)
		if err != nil {
			return &pb.UsuarioPerfis{}, erros.ErroStatus{
				Status: codes.Internal,
				Erro:   err,
			}
		}

		// Coloca o objeto completo do perfil na lista de perfis do objeto recebido
		// Como esse é o objeto retornado, os perfis retornam completos e prontos para serem usados
		perfis[i] = helpers.TPerfToPb(tPerfil)

		// Cria o objeto CreateUsuarioPerfilParams gerado pelo sqlc para gravação no banco de dados
		usuarioPerfil := db.CreateUsuarioPerfilParams{
			UsuarioID: usuarioAtualizado.ID,
			PerfilID:  perfil.ID,
			DataHora:  pgtype.Timestamp{Time: time.Now(), Valid: true},
		}

		usuarioService.usuarioPerfilRepository.Create(context, usuarioPerfil)
	}

	usuarioPerfilRetorno := pb.UsuarioPerfis{
		Usuario: helpers.TUsuarioToPb(usuarioAtualizado),
		Perfis:  perfis,
	}

	return &usuarioPerfilRetorno, erros.ErroStatus{}
}

// Função para desativar um usuário pelo ID
func (usuarioService *UsuarioService) DesativarUsuarioById(context context.Context, usuario *pb.UsuarioPerfis) erros.ErroStatus {
	// Cria o objeto UpdateUsuarioParams gerado pelo sqlc para gravação no banco de dados
	usuarioUpdate := db.UpdateUsuarioParams{
		ID:                    usuario.GetUsuario().GetID(),
		Nome:                  usuario.GetUsuario().GetNome(),
		Email:                 usuario.GetUsuario().GetEmail(),
		Senha:                 usuario.GetUsuario().GetSenha(),
		Ativo:                 pgtype.Bool{Bool: false, Valid: true},
		TokenResetSenha:       pgtype.Text{String: usuario.GetUsuario().GetTokenResetSenha(), Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		SenhaAtualizada:       pgtype.Bool{Bool: usuario.GetUsuario().GetSenhaAtualizada(), Valid: true},
	}

	// Salva o usuário atualizado no repositório
	_, errAtt := usuarioService.usuarioRepository.Update(context, usuarioUpdate)

	return erros.ErroStatus{
		Status: codes.Internal,
		Erro:   errAtt,
	}
}

// Função para ativar um usuário, alterando seu status ativo
func (usuarioService *UsuarioService) AtivarUsuarioById(context context.Context, usuario *pb.UsuarioPerfis) erros.ErroStatus {
	// Cria o objeto UpdateUsuarioParams gerado pelo sqlc para gravação no banco de dados
	usuarioUpdate := db.UpdateUsuarioParams{
		ID:                    usuario.GetUsuario().GetID(),
		Nome:                  usuario.GetUsuario().GetNome(),
		Email:                 usuario.GetUsuario().GetEmail(),
		Senha:                 usuario.GetUsuario().GetSenha(),
		Ativo:                 pgtype.Bool{Bool: true, Valid: true},
		TokenResetSenha:       pgtype.Text{String: usuario.GetUsuario().GetTokenResetSenha(), Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		SenhaAtualizada:       pgtype.Bool{Bool: usuario.GetUsuario().GetSenhaAtualizada(), Valid: true},
	}

	// Salva o usuário atualizado no repositório
	_, errAtt := usuarioService.usuarioRepository.Update(context, usuarioUpdate)

	return erros.ErroStatus{
		Status: codes.Internal,
		Erro:   errAtt,
	}
}

// Essa função serve apenas para definir a flag `SenhaAtualizada` como false
// Dessa forma o usuário consegue passar pelo middleware e acessar a API
// Essa função é chamada sempre que um usuário faz login
func (usuarioService *UsuarioService) DefinidoTokenParaNovaSenha(context context.Context, usuario *pb.UsuarioPerfis) erros.ErroStatus {
	// Cria o objeto UpdateUsuarioParams gerado pelo sqlc para gravação no banco de dados
	usuarioUpdate := db.UpdateUsuarioParams{
		ID:                    usuario.Usuario.ID,
		Nome:                  usuario.Usuario.Nome,
		Email:                 usuario.Usuario.Email,
		Senha:                 usuario.Usuario.Senha,
		Ativo:                 pgtype.Bool{Bool: true, Valid: true},
		TokenResetSenha:       pgtype.Text{String: usuario.Usuario.TokenResetSenha, Valid: true},
		DataUltimaAtualizacao: pgtype.Timestamp{Time: time.Now(), Valid: true},
		SenhaAtualizada:       pgtype.Bool{Bool: false, Valid: true},
	}

	// Salva o usuário atualizado no repositório
	_, err := usuarioService.usuarioRepository.Update(context, usuarioUpdate)

	return erros.ErroStatus{
		Status: codes.Internal,
		Erro:   err,
	}
}
