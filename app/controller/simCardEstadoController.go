package controller

import (
	"encoding/csv"
	"net/http"
	"os"
	"simfonia-golang-back/app/model"
	parametrosdebusca "simfonia-golang-back/app/model/parametrosDeBusca/simCardEstado"
	"simfonia-golang-back/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SimCardEstadoController struct {
	simCardEstadoService *service.SimCardEstadoService
}

func NewSimCardEstadoController(simCardEstadoService *service.SimCardEstadoService) *SimCardEstadoController {
	return &SimCardEstadoController{
		simCardEstadoService: simCardEstadoService,
	}
}

func (simCardEstadoController *SimCardEstadoController) InitRoutes(app *gin.Engine) {
	api := app.Group("/simfonia/api/simcardestado")

	api.GET("/:id", simCardEstadoController.findSimCardEstadoById)
	api.GET("/estado", simCardEstadoController.FindSimCardEstadoByEstado)
	api.GET("/", simCardEstadoController.findAllSimCardEstados)
	api.GET("/csv", simCardEstadoController.findAllSimCardEstadosExportCSV)
	api.POST("/", simCardEstadoController.createSimCardEstado)
	api.POST("/csv", simCardEstadoController.createSimCardEstadoByCSV)
	api.PUT("/:id", simCardEstadoController.updateSimCardEstado)
	api.DELETE("/:id", simCardEstadoController.deleteSimCardEstadoById)
}

func (simCardEstadoController *SimCardEstadoController) findSimCardEstadoById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	simCardEstado, err := simCardEstadoController.simCardEstadoService.FindSimCardEstadoById(id)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "SimCardEstado não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, simCardEstado)
}

func (simCardEstadoController *SimCardEstadoController) FindSimCardEstadoByEstado(context *gin.Context) {
	estado := new(parametrosdebusca.Estado)

	erroJson := context.ShouldBindJSON(&estado)
	if erroJson != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Não foi enviado o Json apenas com o estado, por favor siga a documentação"},
		)
		return
	}

	cidade, err := simCardEstadoController.simCardEstadoService.FindSimCardEstadoByEstado(estado.Estado)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Estado não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, cidade)
}

func (simCardEstadoController *SimCardEstadoController) findAllSimCardEstados(context *gin.Context) {
	simCards, err := simCardEstadoController.simCardEstadoService.FindAllSimCardEstados()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		simCards,
	)
}

func (simCardEstadoController *SimCardEstadoController) findAllSimCardEstadosExportCSV(context *gin.Context) {
	simCards, err := simCardEstadoController.simCardEstadoService.FindAllSimCardEstados()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na busca pelas simCards"})
		return
	}

	filePath := "./downloads/simCards.csv"
	file, err := os.Create(filePath)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao criar o arquivo .csv"})
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	headers := []string{
		"ID",
		"ICCID",
		"IMSI",
		"PIN",
		"PUK",
		"KI",
		"OPC",
		"EstadoID",
		"TelefoniaNumeroID",
		"DataCriacao",
		"DataEstado",
		"AtualizadoEm",
		"PUK2",
		"PIN2",
	}
	if err := csvWriter.Write(headers); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao escrever no arquivo .csv"})
		return
	}

	for _, simCardEstado := range simCards {
		record := []string{
			strconv.FormatUint(simCardEstado.ID, 10),
			simCardEstado.Estado,
			simCardEstado.Descricao,
		}
		if err := csvWriter.Write(record); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao escrever no arquivo .csv"})
			return
		}
	}

	context.File(filePath)
}

func (simCardEstadoController *SimCardEstadoController) createSimCardEstado(context *gin.Context) {
	simCardEstado := new(model.SimCardEstado)
	isNotSimCardEstado := context.ShouldBindJSON(&simCardEstado)

	if isNotSimCardEstado != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é um SimCardEstado"},
		)
		return
	}

	simCardCriado, err := simCardEstadoController.simCardEstadoService.CreateSimCardEstado(*simCardEstado)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		simCardCriado,
	)
}

func (simCardEstadoController *SimCardEstadoController) createSimCardEstadoByCSV(context *gin.Context) {
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

		simCardEstado := model.SimCardEstado{
			Estado:    record[0],
			Descricao: record[1],
		}

		_, err := simCardEstadoController.simCardEstadoService.CreateSimCardEstado(simCardEstado)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na criação da simCardEstado na linha " + iString})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"mensagem": "Arquivo carregado e processado com sucesso!"})
}

func (simCardEstadoController *SimCardEstadoController) updateSimCardEstado(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	simCardEstado := new(model.SimCardEstado)
	isNotSimCardEstado := context.ShouldBindJSON(&simCardEstado)
	if isNotSimCardEstado != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é um SimCardEstado"},
		)
	}

	simCardNovo, erroUpdate := simCardEstadoController.simCardEstadoService.UpdateSimCardEstado(*simCardEstado, id)
	if erroUpdate != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": erroUpdate},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		simCardNovo,
	)
}

func (simCardEstadoController *SimCardEstadoController) deleteSimCardEstadoById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	deletado, erroDelete := simCardEstadoController.simCardEstadoService.DeleteSimCardEstadoById(id)
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
