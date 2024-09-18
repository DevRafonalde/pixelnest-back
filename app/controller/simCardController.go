package controller

import (
	"encoding/csv"
	"net/http"
	"os"
	"simfonia-golang-back/app/model"
	"simfonia-golang-back/app/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SimCardController struct {
	simCardService *service.SimCardService
}

func NewSimCardController(simCardService *service.SimCardService) *SimCardController {
	return &SimCardController{
		simCardService: simCardService,
	}
}

func (simCardController *SimCardController) InitRoutes(app *gin.Engine) {
	api := app.Group("/simfonia/api/simcard")

	api.GET("/:id", simCardController.findSimCardById)
	api.GET("/telefonianumero/:id", simCardController.findSimCardByTelefoniaNumeroID)
	api.GET("/telefonianumero", simCardController.findSimCardByTelefoniaNumero)
	api.GET("/", simCardController.findAllSimCards)
	api.GET("/csv", simCardController.findAllSimCardsExportCSV)
	api.POST("/", simCardController.createSimCard)
	api.POST("/csv", simCardController.createSimCardByCSV)
	api.PUT("/:id", simCardController.updateSimCard)
	api.DELETE("/:id", simCardController.deleteSimCardById)
}

func (simCardController *SimCardController) findSimCardById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	simCard, err := simCardController.simCardService.FindSimCardById(id)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "SimCard não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, simCard)
}

func (simCardController *SimCardController) findSimCardByTelefoniaNumeroID(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	simCard, err := simCardController.simCardService.FindSimCardByTelefoniaNumeroID(int(id))
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "SimCard não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, simCard)
}

// TODO Verificar como esse Numero de Telefone seria enviado, seria o objeto completo, apenas algum parâmetro? Esse acaba resolvendo o de baixo também

func (simCardController *SimCardController) findSimCardByTelefoniaNumero(context *gin.Context) {
	numeroTelefone := new(model.NumeroTelefonico)
	isNotSimCard := context.ShouldBindJSON(&numeroTelefone)

	if isNotSimCard != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é um SimCard"},
		)
		return
	}

	// TODO Verificar como esse Numero de Telefone seria enviado (com ou sem ID, se for sem ID, o que eu posso usar pra pesquisar?)

	simCard, err := simCardController.simCardService.FindSimCardByTelefoniaNumeroID(int(numeroTelefone.ID))
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "SimCard não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, simCard)
}

func (simCardController *SimCardController) findAllSimCards(context *gin.Context) {
	simCards, err := simCardController.simCardService.FindAllSimCards()
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

func (simCardController *SimCardController) findAllSimCardsExportCSV(context *gin.Context) {
	simCards, err := simCardController.simCardService.FindAllSimCards()
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

	for _, simCard := range simCards {
		record := []string{
			strconv.FormatUint(simCard.ID, 10),
			simCard.ICCID,
			simCard.IMSI,
			simCard.PIN,
			simCard.PUK,
			simCard.KI,
			simCard.OPC,
			strconv.FormatUint(uint64(*simCard.EstadoID), 10),
			strconv.FormatUint(*simCard.TelefoniaNumeroID, 10),
			simCard.DataCriacao.String(),
			simCard.DataEstado.String(),
			simCard.AtualizadoEm.String(),
			simCard.PUK2,
			simCard.PIN2,
		}
		if err := csvWriter.Write(record); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao escrever no arquivo .csv"})
			return
		}
	}

	context.File(filePath)
}

func (simCardController *SimCardController) createSimCard(context *gin.Context) {
	simCard := new(model.SimCard)
	isNotSimCard := context.ShouldBindJSON(&simCard)

	if isNotSimCard != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é um SimCard"},
		)
		return
	}

	simCardCriado, err := simCardController.simCardService.CreateSimCard(*simCard)
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

func (simCardController *SimCardController) createSimCardByCSV(context *gin.Context) {
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

		estadoID, errConv := strconv.ParseUint(record[6], 10, 64)
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter o número da linha " + iString})
			return
		}

		telefoniaNumeroId, errConv := strconv.ParseUint(record[7], 10, 64)
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter o número da linha " + iString})
			return
		}

		layoutData := "02/01/2006 15:04:05"
		dataCriacao, errConv := time.Parse(layoutData, record[8])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter a flag \"utilizável\" da linha " + iString})
			return
		}

		dataEstado, errConv := time.Parse(layoutData, record[9])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter a flag \"utilizável\" da linha " + iString})
			return
		}

		AtualizadoEm, errConv := time.Parse(layoutData, record[10])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter a flag \"utilizável\" da linha " + iString})
			return
		}

		// TODO Verificar, nos casos de cadastro por CSV, se de só enviar o ID das cahves estrangeiras já resolve o problema

		simCard := model.SimCard{
			ICCID:             record[0],
			IMSI:              record[1],
			PIN:               record[2],
			PUK:               record[3],
			KI:                record[4],
			OPC:               record[5],
			EstadoID:          &estadoID,
			TelefoniaNumeroID: &telefoniaNumeroId,
			DataCriacao:       dataCriacao,
			DataEstado:        dataEstado,
			AtualizadoEm:      AtualizadoEm,
			PUK2:              record[11],
			PIN2:              record[12],
		}

		_, err := simCardController.simCardService.CreateSimCard(simCard)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na criação da simCard na linha " + iString})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"mensagem": "Arquivo carregado e processado com sucesso!"})
}

func (simCardController *SimCardController) updateSimCard(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	simCard := new(model.SimCard)
	isNotSimCard := context.ShouldBindJSON(&simCard)
	if isNotSimCard != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é um SimCard"},
		)
	}

	simCardNovo, erroUpdate := simCardController.simCardService.UpdateSimCard(*simCard, id)
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

func (simCardController *SimCardController) deleteSimCardById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	deletado, erroDelete := simCardController.simCardService.DeleteSimCardById(id)
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
