package controller

import (
	"context"
	"errors"
	"pixelnest/app/configuration/logger"
	"pixelnest/app/controller/middlewares"
	"pixelnest/app/service"
	"time"

	pb "pixelnest/app/model/grpc" // Importa o pacote gerado pelos arquivos .proto

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Implementação do servidor
type UsuariosServer struct {
	pb.UnimplementedUsuariosServer
	usuarioService      *service.UsuarioService
	permissaoMiddleware *middlewares.PermissoesMiddleware
	chaveSecreta        []byte
}

func NewUsuariosServer(usuarioService *service.UsuarioService, permissaoMiddleware *middlewares.PermissoesMiddleware, chaveSecreta []byte) *UsuariosServer {
	return &UsuariosServer{
		usuarioService:      usuarioService,
		permissaoMiddleware: permissaoMiddleware,
		chaveSecreta:        chaveSecreta,
	}
}

func (usuarioServer *UsuariosServer) mustEmbedUnimplementedUsuariosServer() {}

// Função para buscar por todos os usuários
func (usuarioServer *UsuariosServer) FindAllUsuarios(context context.Context, req *pb.RequestVazio) (*pb.ListaUsuarios, error) {
	usuarioSolicitante, retornoMiddleware := usuarioServer.permissaoMiddleware.PermissaoMiddleware(context, "get-all-usuarios")
	if retornoMiddleware.Erro != nil {
		return &pb.ListaUsuarios{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarios, erroService := usuarioServer.usuarioService.FindAllUsuarios(context)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ListaUsuarios{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Buscados todos os usuários",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ListaUsuarios{Usuarios: usuarios}, nil
}

// Função para buscar por um usuário pelo ID
func (usuarioServer *UsuariosServer) FindUsuarioById(context context.Context, req *pb.RequestId) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioServer.permissaoMiddleware.PermissaoMiddleware(context, "get-usuario-by-id")
	if retornoMiddleware.Erro != nil {
		return &pb.UsuarioPerfis{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.UsuarioPerfis{}, status.Errorf(codes.InvalidArgument, "ID não enviado")
	}

	usuario, erroService := usuarioServer.usuarioService.FindUsuarioById(context, id) // Busca o usuário pelo ID
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.UsuarioPerfis{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	var pbPerfis []*pb.Perfil
	pbPerfis = append(pbPerfis, usuario.Perfis...)

	usuarioRetorno := &pb.UsuarioPerfis{Usuario: usuario.Usuario, Perfis: pbPerfis}

	logger.Logger.Info("Buscado um usuário pelo ID",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("usuarioBuscado", usuarioRetorno),
	)

	return usuarioRetorno, nil
}

// Função para buscar por todos os perfis vinculados àquele usuário
func (usuarioServer *UsuariosServer) GetPerfisVinculados(context context.Context, req *pb.RequestId) (*pb.ResponsePerfisVinculados, error) {
	usuarioSolicitante, retornoMiddleware := usuarioServer.permissaoMiddleware.PermissaoMiddleware(context, "get-perfis-vinculados")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponsePerfisVinculados{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponsePerfisVinculados{}, status.Errorf(codes.InvalidArgument, "ID não enviado")
	}

	perfis, erroService := usuarioServer.usuarioService.GetPerfisVinculados(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponsePerfisVinculados{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	var pbPerfis []*pb.Perfil
	for i := range perfis {
		pbPerfis = append(pbPerfis, &perfis[i])
	}

	logger.Logger.Info("Buscado os perfis vinculados a um usuário pelo ID do usuário",
		zap.Any("usuario", usuarioSolicitante.Usuario),
	)

	return &pb.ResponsePerfisVinculados{Perfis: pbPerfis}, nil
}

// Função para criar um novo usuário
func (usuarioServer *UsuariosServer) CreateUsuario(context context.Context, req *pb.UsuarioPerfis) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioServer.permissaoMiddleware.PermissaoMiddleware(context, "post-create-usuario")
	if retornoMiddleware.Erro != nil {
		return &pb.UsuarioPerfis{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarioCriado, erroService := usuarioServer.usuarioService.CreateUsuario(context, req) // Cria o usuário
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.UsuarioPerfis{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Criado um novo usuário",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("usuarioCriado", usuarioCriado),
	)

	return usuarioCriado, nil
}

// Função para clonar um usuário existente
func (usuarioServer *UsuariosServer) CloneUsuario(context context.Context, req *pb.RequestId) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioServer.permissaoMiddleware.PermissaoMiddleware(context, "post-clone-usuario")
	if retornoMiddleware.Erro != nil {
		return &pb.UsuarioPerfis{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.UsuarioPerfis{}, status.Errorf(codes.InvalidArgument, "ID não enviado")
	}

	usuario, erroService := usuarioServer.usuarioService.CloneUsuario(context, id) // Clona o usuário
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.UsuarioPerfis{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	var pbPerfis []*pb.Perfil
	pbPerfis = append(pbPerfis, usuario.Perfis...)

	logger.Logger.Info("Clonado um usuário existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("usuarioClonado", usuario),
	)

	return &pb.UsuarioPerfis{Usuario: usuario.Usuario, Perfis: pbPerfis}, nil
}

// Função para atualizar um usuário já existente
func (usuarioServer *UsuariosServer) UpdateUsuario(context context.Context, req *pb.UsuarioPerfis) (*pb.UsuarioPerfis, error) {
	usuarioSolicitante, retornoMiddleware := usuarioServer.permissaoMiddleware.PermissaoMiddleware(context, "put-update-usuario")
	if retornoMiddleware.Erro != nil {
		return &pb.UsuarioPerfis{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	usuarioAntigo, erroService := usuarioServer.usuarioService.FindUsuarioById(context, req.GetUsuario().GetID()) // Busca o usuário pelo ID
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.UsuarioPerfis{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	usuarioNovo, erroService := usuarioServer.usuarioService.UpdateUsuario(context, req, usuarioAntigo) // Atualiza o usuário
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.UsuarioPerfis{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	var pbPerfis []*pb.Perfil
	pbPerfis = append(pbPerfis, usuarioNovo.Perfis...)

	logger.Logger.Info("Atualizado um usuário existente",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("usuarioAntigo", usuarioAntigo),
		zap.Any("usuarioAtualizado", usuarioNovo),
	)

	return &pb.UsuarioPerfis{Usuario: usuarioNovo.Usuario, Perfis: pbPerfis}, nil
}

// Função para o administrador do sistema alterar a senha de qualquer usuário baseado no ID
func (usuarioServer *UsuariosServer) AlterarSenhaAdmin(context context.Context, req *pb.RequestAlterarSenhaAdmin) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioServer.permissaoMiddleware.PermissaoMiddleware(context, "put-alterar-senha-admin")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID não enviado")
	}

	usuarioASerAlterado, erroService := usuarioServer.usuarioService.FindUsuarioById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = usuarioServer.usuarioService.AtualizarSenha(context, usuarioASerAlterado.GetUsuario().GetEmail(), req.GetSenhaNova())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Alterada a senha de um usuário existente por meio de permissões de admin",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("usuarioModificado", usuarioASerAlterado),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para o usuário alterar a própria senha
func (usuarioServer *UsuariosServer) AlterarSenhaUsuario(context context.Context, req *pb.RequestAlterarSenhaUsuario) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioServer.permissaoMiddleware.PermissaoMiddleware(context, "put-alterar-propria-senha")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID não enviado")
	}

	usuarioASerAlterado, erroService := usuarioServer.usuarioService.FindUsuarioById(context, id)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	login := pb.LoginUsuario{
		Email: usuarioASerAlterado.Usuario.Email,
		Senha: req.SenhaAntiga,
	}

	usuarioLogado, erroService := usuarioServer.usuarioService.Login(context, &login)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = usuarioServer.usuarioService.AtualizarSenha(context, usuarioLogado.Usuario.Email, req.SenhaNova)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Alterada a senha de um usuário existente por meios próprios",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("usuarioModificado", usuarioASerAlterado),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para ativar um usuário existente
func (usuarioServer *UsuariosServer) AtivarUsuario(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioServer.permissaoMiddleware.PermissaoMiddleware(context, "put-ativar-usuario")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID não enviado")
	}

	usuario, erroService := usuarioServer.usuarioService.FindUsuarioById(context, id) // Busca o usuário pelo ID
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = usuarioServer.usuarioService.AtivarUsuarioById(context, usuario)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Ativou usuário",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("usuarioAtivado", usuario),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para desativar um usuário existente
func (usuarioServer *UsuariosServer) DesativarUsuario(context context.Context, req *pb.RequestId) (*pb.ResponseBool, error) {
	usuarioSolicitante, retornoMiddleware := usuarioServer.permissaoMiddleware.PermissaoMiddleware(context, "put-desativar-usuario")
	if retornoMiddleware.Erro != nil {
		return &pb.ResponseBool{}, status.Errorf(retornoMiddleware.Status, retornoMiddleware.Erro.Error())
	}

	id := req.GetID()
	if id == 0 {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.InvalidArgument, "ID não enviado")
	}

	usuario, erroService := usuarioServer.usuarioService.FindUsuarioById(context, id) // Busca o usuário pelo ID
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	erroService = usuarioServer.usuarioService.DesativarUsuarioById(context, usuario)
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioSolicitante.Usuario))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Desativou usuário",
		zap.Any("usuario", usuarioSolicitante.Usuario),
		zap.Any("usuarioDesativado", usuario),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// Função para realizar o login na aplicação
func (usuarioServer *UsuariosServer) Login(context context.Context, req *pb.LoginUsuario) (*pb.RetornoLoginUsuario, error) {
	usuarioLogado, erroService := usuarioServer.usuarioService.Login(context, req)
	if erroService.Erro != nil {
		logger.Logger.Error("Tentativa de login na API", zap.NamedError("err", erroService.Erro), zap.Any("email", req.GetEmail()))
		return &pb.RetornoLoginUsuario{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	tokenString, err := usuarioServer.createToken(context, usuarioLogado)
	if err != nil {
		logger.Logger.Error(err.Error(), zap.NamedError("err", err), zap.Any("usuario", usuarioLogado.Usuario))
		return &pb.RetornoLoginUsuario{}, status.Errorf(codes.Internal, err.Error())
	}

	logger.Logger.Info("Feito login na API",
		zap.Any("usuario", usuarioLogado.Usuario),
	)

	return &pb.RetornoLoginUsuario{
		ID:    usuarioLogado.GetUsuario().GetID(),
		Nome:  usuarioLogado.GetUsuario().GetNome(),
		Email: usuarioLogado.GetUsuario().Email,
		Token: tokenString,
	}, nil
}

// Função para envio do token de reset de senha
func (usuarioServer *UsuariosServer) TokenResetSenha(context context.Context, req *pb.EmailReset) (*pb.ResponseTokenResetSenha, error) {
	token, err := usuarioServer.createTokenReset(req)
	if err != nil {
		logger.Logger.Error(err.Error(), zap.NamedError("err", err), zap.Any("email", req.GetEmail()))
		return &pb.ResponseTokenResetSenha{}, status.Errorf(codes.Internal, err.Error())
	}

	erroService := usuarioServer.usuarioService.TokenResetSenha(context, token, req.GetEmail())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("email", req.GetEmail()))
		return &pb.ResponseTokenResetSenha{}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Solicitado um token para reset/esquecimento de senha",
		zap.Any("email", req.GetEmail()),
	)

	return &pb.ResponseTokenResetSenha{Token: token}, nil
}

// Função para validar o token de reset de senha e alterar a senha
func (usuarioServer *UsuariosServer) ResetSenha(context context.Context, req *pb.ResetSenhaUsuario) (*pb.ResponseBool, error) {
	email, err := usuarioServer.validarToken(req.Token)
	if err != nil {
		return &pb.ResponseBool{Alterado: false}, status.Errorf(codes.Unauthenticated, "Token inválido")
	}

	erroService := usuarioServer.usuarioService.AtualizarSenha(context, email, req.GetSenhaNova())
	if erroService.Erro != nil {
		logger.Logger.Error(erroService.Erro.Error(), zap.NamedError("err", erroService.Erro), zap.Any("email", email))
		return &pb.ResponseBool{Alterado: false}, status.Errorf(erroService.Status, erroService.Erro.Error())
	}

	logger.Logger.Info("Resetada e alterada a senha de usuário",
		zap.Any("email", email),
	)

	return &pb.ResponseBool{Alterado: true}, nil
}

// createToken cria um token JWT para um usuário autenticado.
func (usuarioServer *UsuariosServer) createToken(context context.Context, usuarioLogado *pb.UsuarioPerfis) (string, error) {
	claims := middlewares.CustomClaims{
		IdUsuario: usuarioLogado.Usuario.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 24), // Define a expiração do token para 24 horas
			},
		},
	}

	logger.Logger.Info("Solicitada a criação de um novo token de acesso",
		zap.Any("usuario", usuarioLogado.Usuario),
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Cria o token JWT com as claims

	tokenString, err := token.SignedString(usuarioServer.chaveSecreta) // Assina o token com a chave secreta
	if err != nil {
		logger.Logger.Error("Falha na criação do token de acesso", zap.NamedError("err", err), zap.Any("usuario", usuarioLogado.Usuario))
		return "", err // Retorna erro se falhar
	}

	erroService := usuarioServer.usuarioService.DefinidoTokenParaNovaSenha(context, usuarioLogado)
	if erroService.Erro != nil {
		logger.Logger.Error("Falha na criação do token de acesso", zap.NamedError("err", erroService.Erro), zap.Any("usuario", usuarioLogado.Usuario))
		return "", erroService.Erro
	}

	logger.Logger.Info("Criado um novo token de acesso",
		zap.Any("usuario", usuarioLogado.Usuario),
	)

	return tokenString, nil // Retorna o token JWT gerado
}

// createToken cria um token JWT para um usuário autenticado para fins de reset da senha do mesmo.
func (usuarioServer *UsuariosServer) createTokenReset(emailReset *pb.EmailReset) (string, error) {
	claims := CustomClaimsResetSenha{
		Email: emailReset.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Minute * 5), // Define a expiração do token para 5 minutos
			},
		},
	}

	logger.Logger.Info("Solicitada a criação de um novo token de reset de senha",
		zap.Any("email", emailReset.GetEmail()),
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Cria o token JWT com as claims

	tokenString, err := token.SignedString(usuarioServer.chaveSecreta) // Assina o token com a chave secreta
	if err != nil {
		logger.Logger.Error("Falha na criação do token de reset de senha", zap.NamedError("err", err), zap.Any("email", emailReset.GetEmail()))
		return "", err // Retorna erro se falhar
	}

	logger.Logger.Info("Criado um novo token de reset de senha",
		zap.Any("email", emailReset.GetEmail()),
	)

	return tokenString, nil // Retorna o token JWT gerado
}

// Struct para capturar as informações do token JWT de reset de senha
type CustomClaimsResetSenha struct {
	Email                string
	jwt.RegisteredClaims // Struct que contém os claims padrão do JWT
}

// Função para verificar a validade de um token JWT
func (usuarioServer *UsuariosServer) validarToken(tokenString string) (string, error) {
	// Analisa o token JWT usando a chave secreta
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsResetSenha{}, func(token *jwt.Token) (interface{}, error) {
		return usuarioServer.chaveSecreta, nil
	})
	if err != nil {
		logger.Logger.Error("Falha na validação de um token de reset de senha", zap.NamedError("err", err))
		return "", err // Retorna erro se o token não puder ser analisado
	}

	if !token.Valid {
		logger.Logger.Error("Enviado um token de reset de senha inválido", zap.NamedError("err", err))
		return "", errors.New("Token inválido") // Retorna erro se o token não for válido
	}

	// Extrai as claims personalizadas do token
	props, ok := token.Claims.(*CustomClaimsResetSenha)
	if !ok {
		logger.Logger.Error("Falha na extração das propriedades de um token de reset de senha", zap.NamedError("err", err))
		return "", errors.New("Propriedades inválidas do token") // Retorna erro se as claims não puderem ser extraídas
	}

	email := props.Email // Obtém o email do usuário a partir das claims

	logger.Logger.Info("Validado um token de reset de senha",
		zap.Any("email", email),
	)

	return email, nil // Retorna o email do usuário
}
