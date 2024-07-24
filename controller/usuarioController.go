package controller

import (
	"crud-rafael/model"
	parametrosdebusca "crud-rafael/model/parametrosDeBusca"
	"crud-rafael/service"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	usuarioService *service.UsuarioService
}

func NewusuarioController(usuarioService *service.UsuarioService) *UsuarioController {
	return &UsuarioController{
		usuarioService: usuarioService,
	}
}

func (usuarioController *UsuarioController) InitRoutes() {
	app := gin.Default()
	api := app.Group("/api/usuario")

	api.GET("/:id", usuarioController.findUsuarioById)
	api.GET("/email", usuarioController.findUsuarioByEmail)
	api.GET("/", usuarioController.findAllUsuarios)
	api.GET("/csv", usuarioController.findAllUsuariosExportCSV)
	api.POST("/", usuarioController.createUsuario)
	api.POST("/csv", usuarioController.createUsuarioByCSV)
	api.PUT("/:id", usuarioController.updateUsuario)
	api.DELETE("/:id", usuarioController.deleteUsuarioById)
	api.DELETE("/all", usuarioController.deleteAllUsuarios)

	app.Run(":8601")
}

func (usuarioController *UsuarioController) createUsuarioByCSV(context *gin.Context) {
	file, err := context.FormFile("csv")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Nenhum arquivo foi recebido"})
		return
	}

	filePath := "./uploads/" + file.Filename
	if err := context.SaveUploadedFile(file, filePath); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao salvar o arquivo"})
		return
	}

	f, err := os.Open(filePath)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao abrir o arquivo"})
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao ler o arquivo"})
		return
	}

	for i, record := range records {
		if len(record) < 2 {
			continue
		}
		if i == 0 {
			continue
		}
		usuario := model.Usuario{
			Nome:  record[0],
			Email: record[1],
		}

		fmt.Println(usuario)
		_, err := usuarioController.usuarioService.CreateUsuario(usuario)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na criação de um usuário"})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"mensagem": "Arquivo carregado e processado com sucesso!"})
}

func (usuarioController *UsuarioController) findAllUsuariosExportCSV(context *gin.Context) {
	usuarios, err := usuarioController.usuarioService.FindAllUsuarios()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na busca pelos usuários"})
		return
	}

	filePath := "./downloads/usuarios.csv"
	file, err := os.Create(filePath)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao criar o arquivo .csv"})
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	headers := []string{"ID", "Nome", "Email"}
	if err := csvWriter.Write(headers); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao escrever no arquivo .csv"})
		return
	}

	for _, usuario := range usuarios {
		record := []string{strconv.FormatUint(usuario.ID, 10), usuario.Nome, usuario.Email}
		if err := csvWriter.Write(record); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao escrever no arquivo .csv"})
			return
		}
	}

	context.File(filePath)
}

func (usuarioController *UsuarioController) findUsuarioById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	usuario, err := usuarioController.usuarioService.FindUsuarioById(id)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Usuário não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, usuario)
}

func (usuarioController *UsuarioController) findUsuarioByEmail(context *gin.Context) {
	email := new(parametrosdebusca.Email)

	erroJson := context.ShouldBindJSON(&email)
	if erroJson != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Não foi enviado o Json apenas com o e-mail, por favor siga a documentação"},
		)
		return
	}
	fmt.Println(email.Email)

	isEmail := strings.ContainsAny(email.Email, "@")

	if !isEmail {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "E-mail enviado não é válido"},
		)
		return
	}

	usuario, err := usuarioController.usuarioService.FindUsuarioByEmail(email.Email)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Usuário não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, usuario)
}

func (usuarioController *UsuarioController) findAllUsuarios(context *gin.Context) {
	usuarios, err := usuarioController.usuarioService.FindAllUsuarios()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		usuarios,
	)
}

func (usuarioController *UsuarioController) createUsuario(context *gin.Context) {
	usuario := new(model.Usuario)
	isNotUsuario := context.ShouldBindJSON(&usuario)

	if isNotUsuario != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é um usuário"},
		)
		return
	}

	usuarioCriado, err := usuarioController.usuarioService.CreateUsuario(*usuario)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	} else if (usuarioCriado == model.Usuario{}) {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "O e-mail enviado já foi utilizado"},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		usuarioCriado,
	)
}

func (usuarioController *UsuarioController) updateUsuario(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	usuario := new(model.Usuario)
	isNotUsuario := context.ShouldBindJSON(&usuario)
	if isNotUsuario != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é um usuário"},
		)
	}

	usuarioNovo, erroUpdate := usuarioController.usuarioService.UpdateUsuario(*usuario, id)
	if erroUpdate != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": erroUpdate},
		)
		return
	} else if (usuarioNovo == model.Usuario{}) {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "O e-mail enviado já foi utilizado"},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		usuarioNovo,
	)
}

func (controller *UsuarioController) deleteUsuarioById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	deletado, erroDelete := controller.usuarioService.DeleteUsuarioById(id)
	if erroDelete != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": erroDelete},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"deletado": deletado},
	)
}

func (usuarioController *UsuarioController) deleteAllUsuarios(context *gin.Context) {
	err := usuarioController.usuarioService.DeleteAllUsuarios()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao deletar os usuários"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"mensagem": "Todos os usuários foram deletados com sucesso!"})
}
