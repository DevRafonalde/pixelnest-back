package controller

import (
	"encoding/csv"
	"net/http"
	"os"
	"simfonia-golang-back/app/model"
	parametrosdebusca "simfonia-golang-back/app/model/parametrosDeBusca/numerosTelefonicos"
	"simfonia-golang-back/app/service"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type NumeroTelefonicoController struct {
	numeroTelefonicoService *service.NumeroTelefonicoService
}

func NewNumeroTelefonicoController(numeroTelefonicoService *service.NumeroTelefonicoService) *NumeroTelefonicoController {
	return &NumeroTelefonicoController{
		numeroTelefonicoService: numeroTelefonicoService,
	}
}

func (numeroTelefonicoController *NumeroTelefonicoController) InitRoutes(app *gin.Engine) {
	api := app.Group("/simfonia/api/numerostelefonicos")

	api.GET("/:id", numeroTelefonicoController.findNumeroTelefonicoById)
	api.GET("/numero", numeroTelefonicoController.findNumeroTelefonicoByNumero)
	api.GET("/simcard/:id", numeroTelefonicoController.findNumeroTelefonicoBySimCardId)
	api.GET("/simcard", numeroTelefonicoController.findNumeroTelefonicoBySimCard)
	api.GET("/", numeroTelefonicoController.findAllNumeroTelefonicos)
	api.GET("/csv", numeroTelefonicoController.findAllNumeroTelefonicosExportCSV)
	api.POST("/", numeroTelefonicoController.createNumeroTelefonico)
	api.POST("/csv", numeroTelefonicoController.createNumeroTelefonicoByCSV)
	api.PUT("/:id", numeroTelefonicoController.updateNumeroTelefonico)
	api.DELETE("/:id", numeroTelefonicoController.deleteNumeroTelefonicoById)
}

func (numeroTelefonicoController *NumeroTelefonicoController) findNumeroTelefonicoById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	numeroTelefonico, err := numeroTelefonicoController.numeroTelefonicoService.FindNumeroTelefonicoById(id)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Número Telefônico não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, numeroTelefonico)
}

func (numeroTelefonicoController *NumeroTelefonicoController) findNumeroTelefonicoByNumero(context *gin.Context) {
	numero := new(parametrosdebusca.Numero)

	erroJson := context.ShouldBindJSON(&numero)
	if erroJson != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Não foi enviado o Json apenas com o número, por favor siga a documentação"},
		)
		return
	}

	numeroTelefonico, err := numeroTelefonicoController.numeroTelefonicoService.FindNumeroTelefonicoByNumero(int(numero.Numero))
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Número Telefônico não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, numeroTelefonico)
}

func (numeroTelefonicoController *NumeroTelefonicoController) findNumeroTelefonicoBySimCardId(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	numeroTelefonico, err := numeroTelefonicoController.numeroTelefonicoService.FindNumeroTelefonicoBySimCardId(int(id))
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Número Telefônico não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, numeroTelefonico)
}

// TODO Verificar como esse SimCard seria enviado, seria o objeto completo, apenas algum parâmetro? Esse acaba resolvendo o de baixo também

func (numeroTelefonicoController *NumeroTelefonicoController) findNumeroTelefonicoBySimCard(context *gin.Context) {
	simCard := new(model.SimCard)
	isNotSimCard := context.ShouldBindJSON(&simCard)

	if isNotSimCard != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é um SimCard"},
		)
		return
	}

	// TODO Verificar como esse SimCard seria enviado (com ou sem ID, se for sem ID, o que eu posso usar pra pesquisar?)

	numeroTelefonico, err := numeroTelefonicoController.numeroTelefonicoService.FindNumeroTelefonicoBySimCardId(int(simCard.ID))
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Número Telefônico não encontrado"},
		)
		return
	}

	context.JSON(http.StatusOK, numeroTelefonico)
}

func (numeroTelefonicoController *NumeroTelefonicoController) findAllNumeroTelefonicos(context *gin.Context) {
	numerosTelefonicos, err := numeroTelefonicoController.numeroTelefonicoService.FindAllNumeroTelefonicos()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		numerosTelefonicos,
	)
}

func (numeroTelefonicoController *NumeroTelefonicoController) findAllNumeroTelefonicosExportCSV(context *gin.Context) {
	numerosTelefonicos, err := numeroTelefonicoController.numeroTelefonicoService.FindAllNumeroTelefonicos()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na busca pelas numerosTelefonicos"})
		return
	}

	filePath := "./downloads/numerosTelefonicos.csv"
	file, err := os.Create(filePath)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao criar o arquivo .csv"})
		return
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	defer csvWriter.Flush()

	headers := []string{"ID",
		"CodArea",
		"Numero",
		"Utilizavel",
		"PortadoIn",
		"PortadoInOperadora",
		"PortadoInDate",
		"CodigoCNL",
		"CongeladoAte",
		"ExternalID",
		"PortadoOut",
		"PortadoOutOperadora",
		"PortadoOutDate",
		"DataCriacao",
		"SimCardID",
		"PortadoInOperadoraID",
		"PortadoOutOperadoraID",
	}

	if err := csvWriter.Write(headers); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao escrever no arquivo .csv"})
		return
	}

	for _, numeroTelefonico := range numerosTelefonicos {

		var portadoInDate string
		if numeroTelefonico.PortadoInDate == nil {
			portadoInDate = "null"
		} else {
			portadoInDate = numeroTelefonico.PortadoInDate.String()
		}

		var congeladoAte string
		if numeroTelefonico.CongeladoAte == nil {
			congeladoAte = "null"
		} else {
			congeladoAte = numeroTelefonico.CongeladoAte.String()
		}

		var portadoOutDate string
		if numeroTelefonico.PortadoOutDate == nil {
			portadoOutDate = "null"
		} else {
			portadoOutDate = numeroTelefonico.PortadoOutDate.String()
		}

		var dataCriacao string
		if numeroTelefonico.DataCriacao == nil {
			dataCriacao = "null"
		} else {
			dataCriacao = numeroTelefonico.DataCriacao.String()
		}

		var codArea string
		if numeroTelefonico.CodArea == nil {
			codArea = "null"
		} else {
			codArea = strconv.Itoa(*numeroTelefonico.CodArea)
		}

		var portadoInOperadora string
		if numeroTelefonico.PortadoInOperadora == nil {
			portadoInOperadora = "null"
		} else {
			portadoInOperadora = *numeroTelefonico.PortadoInOperadora
		}

		var externalID string
		if numeroTelefonico.ExternalID == nil {
			portadoInOperadora = "null"
		} else {
			portadoInOperadora = strconv.FormatUint(*numeroTelefonico.ExternalID, 10)
		}

		var portadoOutOperadora string
		if numeroTelefonico.PortadoOutOperadora == nil {
			portadoOutOperadora = "null"
		} else {
			portadoOutOperadora = *numeroTelefonico.PortadoOutOperadora
		}

		var simCardID string
		if numeroTelefonico.SimCardID == nil {
			simCardID = "null"
		} else {
			simCardID = strconv.FormatUint(*numeroTelefonico.SimCardID, 10)
		}

		var portadoInOperadoraID string
		if numeroTelefonico.PortadoInOperadoraID == nil {
			portadoInOperadoraID = "null"
		} else {
			portadoInOperadoraID = strconv.FormatUint(*numeroTelefonico.PortadoInOperadoraID, 10)
		}

		var portadoOutOperadoraID string
		if numeroTelefonico.PortadoOutOperadoraID == nil {
			portadoOutOperadoraID = "null"
		} else {
			portadoOutOperadoraID = strconv.FormatUint(*numeroTelefonico.PortadoOutOperadoraID, 10)
		}

		record := []string{
			strconv.FormatUint(numeroTelefonico.ID, 10),
			codArea,
			strconv.FormatUint(numeroTelefonico.Numero, 10),
			strconv.FormatBool(numeroTelefonico.Utilizavel),
			strconv.FormatBool(numeroTelefonico.PortadoIn),
			portadoInOperadora,
			portadoInDate,
			numeroTelefonico.CodigoCNL,
			congeladoAte,
			externalID,
			strconv.FormatBool(numeroTelefonico.PortadoOut),
			portadoOutOperadora,
			portadoOutDate,
			dataCriacao,
			simCardID,
			portadoInOperadoraID,
			portadoOutOperadoraID,
		}
		if err := csvWriter.Write(record); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao escrever no arquivo .csv"})
			return
		}
	}

	context.File(filePath)
}

func (numeroTelefonicoController *NumeroTelefonicoController) createNumeroTelefonico(context *gin.Context) {
	numeroTelefonico := new(model.NumeroTelefonico)
	isNotNumeroTelefonico := context.ShouldBindJSON(&numeroTelefonico)

	if isNotNumeroTelefonico != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é um Número Telefônico"},
		)
		return
	}

	numeroTelefonicoCriado, err := numeroTelefonicoController.numeroTelefonicoService.CreateNumeroTelefonico(*numeroTelefonico)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		numeroTelefonicoCriado,
	)
}

func (numeroTelefonicoController *NumeroTelefonicoController) createNumeroTelefonicoByCSV(context *gin.Context) {
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

		codArea, errConv := strconv.Atoi(record[0])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter o código de área da linha " + iString})
			return
		}

		numero, errConv := strconv.ParseUint(record[1], 10, 64)
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter o número da linha " + iString})
			return
		}

		utilizavel, errConv := strconv.ParseBool(record[2])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter a flag \"utilizável\" da linha " + iString})
			return
		}

		portadoIn, errConv := strconv.ParseBool(record[3])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter a flag \"portadoIn\" da linha " + iString})
			return
		}

		layoutData := "02/01/2006 15:04:05"
		portadoInDate, errConv := time.Parse(layoutData, record[5])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter a data portadoInDate da linha " + iString})
			return
		}

		congeladoAte, errConv := time.Parse(layoutData, record[7])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter a data congeladoAte da linha " + iString})
			return
		}

		externalId, errConv := strconv.ParseUint(record[8], 10, 64)
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter o externalId da linha " + iString})
			return
		}

		portadoOut, errConv := strconv.ParseBool(record[9])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter a flag \"portadoOut\" da linha " + iString})
			return
		}

		portadoOutDate, errConv := time.Parse(layoutData, record[11])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter a data portadoOutDate da linha " + iString})
			return
		}

		dataCriacao, errConv := time.Parse(layoutData, record[12])
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter a data de criação da linha " + iString})
			return
		}

		simCardId, errConv := strconv.ParseUint(record[13], 10, 64)
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter o simCardId da linha " + iString})
			return
		}

		portadoInOperadoraId, errConv := strconv.ParseUint(record[14], 10, 64)
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter o portadoInOperadoraId da linha " + iString})
			return
		}

		portadoOutOperadoraId, errConv := strconv.ParseUint(record[15], 10, 64)
		if errConv != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter o portadoOutOperadoraId da linha " + iString})
			return
		}

		// TODO Verificar, nos casos de cadastro por CSV, se de só enviar o ID das cahves estrangeiras já resolve o problema

		numeroTelefonico := model.NumeroTelefonico{
			CodArea:               &codArea,
			Numero:                numero,
			Utilizavel:            utilizavel,
			PortadoIn:             portadoIn,
			PortadoInOperadora:    &record[4],
			PortadoInDate:         &portadoInDate,
			CodigoCNL:             record[6],
			CongeladoAte:          &congeladoAte,
			ExternalID:            &externalId,
			PortadoOut:            portadoOut,
			PortadoOutOperadora:   &record[10],
			PortadoOutDate:        &portadoOutDate,
			DataCriacao:           &dataCriacao,
			SimCardID:             &simCardId,
			PortadoInOperadoraID:  &portadoInOperadoraId,
			PortadoOutOperadoraID: &portadoOutOperadoraId,
		}

		_, err := numeroTelefonicoController.numeroTelefonicoService.CreateNumeroTelefonico(numeroTelefonico)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na criação da numeroTelefonico na linha " + iString})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"mensagem": "Arquivo carregado e processado com sucesso!"})
}

func (numeroTelefonicoController *NumeroTelefonicoController) updateNumeroTelefonico(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	numeroTelefonico := new(model.NumeroTelefonico)
	isNotNumeroTelefonico := context.ShouldBindJSON(&numeroTelefonico)
	if isNotNumeroTelefonico != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é um Número Telefônico"},
		)
	}

	numeroTelefonicoNovo, erroUpdate := numeroTelefonicoController.numeroTelefonicoService.UpdateNumeroTelefonico(*numeroTelefonico, id)
	if erroUpdate != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": erroUpdate},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		numeroTelefonicoNovo,
	)
}

func (numeroTelefonicoController *NumeroTelefonicoController) deleteNumeroTelefonicoById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	deletado, erroDelete := numeroTelefonicoController.numeroTelefonicoService.DeleteNumeroTelefonicoById(id)
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
