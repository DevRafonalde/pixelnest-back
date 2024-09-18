package controller

import (
	"encoding/csv"
	"net/http"
	"os"
	"simfonia-golang-back/app/model"
	parametrosdebusca "simfonia-golang-back/app/model/parametrosDeBusca/operadoras"
	"simfonia-golang-back/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OperadoraController struct {
	operadoraService *service.OperadoraService
}

func NewOperadoraController(operadoraService *service.OperadoraService) *OperadoraController {
	return &OperadoraController{
		operadoraService: operadoraService,
	}
}

func (operadoraController *OperadoraController) InitRoutes(app *gin.Engine) {
	api := app.Group("/simfonia/api/operadoras")

	api.GET("/:id", operadoraController.findOperadoraById)
	api.GET("/nome", operadoraController.findOperadoraByNome)
	api.GET("/abreviacao", operadoraController.findOperadoraByAbreviacao)
	api.GET("/", operadoraController.findAllOperadoras)
	api.GET("/csv", operadoraController.findAllOperadorasExportCSV)
	api.POST("/", operadoraController.createOperadora)
	api.POST("/csv", operadoraController.createOperadoraByCSV)
	api.PUT("/:id", operadoraController.updateOperadora)
	api.DELETE("/:id", operadoraController.deleteOperadoraById)
}

func (operadoraController *OperadoraController) findOperadoraById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	operadora, err := operadoraController.operadoraService.FindOperadoraById(id)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Operadora não encontrada"},
		)
		return
	}

	context.JSON(http.StatusOK, operadora)
}

func (operadoraController *OperadoraController) findOperadoraByNome(context *gin.Context) {
	nome := new(parametrosdebusca.Nome)

	erroJson := context.ShouldBindJSON(&nome)
	if erroJson != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Não foi enviado o Json apenas com o nome, por favor siga a documentação"},
		)
		return
	}

	operadora, err := operadoraController.operadoraService.FindOperadoraByNome(nome.Nome)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Operadora não encontrada"},
		)
		return
	}

	context.JSON(http.StatusOK, operadora)
}

func (operadoraController *OperadoraController) findOperadoraByAbreviacao(context *gin.Context) {
	abreviacao := new(parametrosdebusca.Abreviacao)

	erroJson := context.ShouldBindJSON(&abreviacao)
	if erroJson != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Não foi enviado o Json apenas com a Abreviação, por favor siga a documentação"},
		)
		return
	}

	operadora, err := operadoraController.operadoraService.FindOperadoraByAbreviacao(abreviacao.Abreviacao)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Operadora não encontrada"},
		)
		return
	}

	context.JSON(http.StatusOK, operadora)
}

func (operadoraController *OperadoraController) findAllOperadoras(context *gin.Context) {
	operadoras, err := operadoraController.operadoraService.FindAllOperadoras()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		operadoras,
	)
}

func (operadoraController *OperadoraController) findAllOperadorasExportCSV(context *gin.Context) {
	operadoras, err := operadoraController.operadoraService.FindAllOperadoras()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na busca pelas operadoras"})
		return
	}

	filePath := "./downloads/operadoras.csv"
	file, err := os.Create(filePath)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao criar o arquivo .csv"})
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	headers := []string{"ID", "UUID", "Nome", "CodIBGE", "UF", "CodArea"}
	if err := csvWriter.Write(headers); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao escrever no arquivo .csv"})
		return
	}

	for _, operadora := range operadoras {
		record := []string{
			strconv.FormatUint(operadora.ID, 10),
			operadora.Nome,
			operadora.Abreviacao,
		}
		if err := csvWriter.Write(record); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao escrever no arquivo .csv"})
			return
		}
	}

	context.File(filePath)
}

func (operadoraController *OperadoraController) createOperadora(context *gin.Context) {
	operadora := new(model.Operadora)
	isNotOperadora := context.ShouldBindJSON(&operadora)

	if isNotOperadora != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é uma operadora"},
		)
		return
	}

	operadoraCriado, err := operadoraController.operadoraService.CreateOperadora(*operadora)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		operadoraCriado,
	)
}

func (operadoraController *OperadoraController) createOperadoraByCSV(context *gin.Context) {
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
		iString := strconv.Itoa(i)

		if len(record) < 2 {
			continue
		}

		if i == 0 {
			continue
		}

		operadora := model.Operadora{
			Nome:       record[0],
			Abreviacao: record[1],
		}

		_, err := operadoraController.operadoraService.CreateOperadora(operadora)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na criação da operadora na linha " + iString})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"mensagem": "Arquivo carregado e processado com sucesso!"})
}

func (operadoraController *OperadoraController) updateOperadora(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	operadora := new(model.Operadora)
	isNotOperadora := context.ShouldBindJSON(&operadora)
	if isNotOperadora != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é uma operadora"},
		)
	}

	operadoraNova, erroUpdate := operadoraController.operadoraService.UpdateOperadora(*operadora, id)
	if erroUpdate != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": erroUpdate},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		operadoraNova,
	)
}

func (operadoraController *OperadoraController) deleteOperadoraById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	deletado, erroDelete := operadoraController.operadoraService.DeleteOperadoraById(id)
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
