package controller

import (
	"encoding/csv"
	"net/http"
	"os"
	"simfonia-golang-back/app/model"
	parametrosdebusca "simfonia-golang-back/app/model/parametrosDeBusca/cidades"
	"simfonia-golang-back/app/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CidadeController struct {
	cidadeService *service.CidadeService
}

func NewCidadeController(cidadeService *service.CidadeService) *CidadeController {
	return &CidadeController{
		cidadeService: cidadeService,
	}
}

func (cidadeController *CidadeController) InitRoutes(app *gin.Engine) {
	api := app.Group("/simfonia/api/cidades")

	api.GET("/:id", cidadeController.findCidadeById)
	api.GET("/nome", cidadeController.findCidadeByNome)
	api.GET("/", cidadeController.findAllCidades)
	api.GET("/csv", cidadeController.findAllCidadesExportCSV)
	api.POST("/", cidadeController.createCidade)
	api.POST("/csv", cidadeController.createCidadeByCSV)
	api.PUT("/:id", cidadeController.updateCidade)
	api.DELETE("/:id", cidadeController.deleteCidadeById)
}

func (cidadeController *CidadeController) findCidadeById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	cidade, err := cidadeController.cidadeService.FindCidadeById(id)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Cidade não encontrada"},
		)
		return
	}

	context.JSON(http.StatusOK, cidade)
}

func (cidadeController *CidadeController) findCidadeByNome(context *gin.Context) {
	nome := new(parametrosdebusca.Nome)

	erroJson := context.ShouldBindJSON(&nome)
	if erroJson != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Não foi enviado o Json apenas com o nome, por favor siga a documentação"},
		)
		return
	}

	cidade, err := cidadeController.cidadeService.FindCidadeByNome(nome.Nome)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Cidade não encontrada"},
		)
		return
	}

	context.JSON(http.StatusOK, cidade)
}

func (cidadeController *CidadeController) findAllCidades(context *gin.Context) {
	cidades, err := cidadeController.cidadeService.FindAllCidades()
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		cidades,
	)
}

func (cidadeController *CidadeController) findAllCidadesExportCSV(context *gin.Context) {
	cidades, err := cidadeController.cidadeService.FindAllCidades()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na busca pelas cidades"})
		return
	}

	filePath := "./downloads/cidades.csv"
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

	for _, cidade := range cidades {
		record := []string{strconv.FormatUint(cidade.ID, 10), cidade.UUID.String(), cidade.Nome, strconv.Itoa(cidade.CodIBGE), cidade.UF, strconv.Itoa(cidade.CodArea)}
		if err := csvWriter.Write(record); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao escrever no arquivo .csv"})
			return
		}
	}

	context.File(filePath)
}

func (cidadeController *CidadeController) createCidade(context *gin.Context) {
	cidade := new(model.Cidade)
	isNotCidade := context.ShouldBindJSON(&cidade)

	if isNotCidade != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é uma cidade"},
		)
		return
	}

	cidadeCriado, err := cidadeController.cidadeService.CreateCidade(*cidade)
	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err},
		)
		return
	}

	context.JSON(
		http.StatusCreated,
		cidadeCriado,
	)
}

func (cidadeController *CidadeController) createCidadeByCSV(context *gin.Context) {
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
		codIbge, errConv1 := strconv.Atoi(record[2])
		if errConv1 != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter o código IBGE da linha " + iString})
			return
		}

		codArea, errConv2 := strconv.Atoi(record[4])
		if errConv2 != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro ao converter o código de área da linha " + iString})
			return
		}
		cidade := model.Cidade{
			UUID:    uuid.MustParse(record[0]),
			Nome:    record[1],
			CodIBGE: codIbge,
			UF:      record[3],
			CodArea: codArea,
		}

		_, err := cidadeController.cidadeService.CreateCidade(cidade)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Ocorreu um erro na criação da cidade na linha " + iString})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{"mensagem": "Arquivo carregado e processado com sucesso!"})
}

func (cidadeController *CidadeController) updateCidade(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	cidade := new(model.Cidade)
	isNotCidade := context.ShouldBindJSON(&cidade)
	if isNotCidade != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "Objeto enviado não é uma cidade"},
		)
	}

	cidadeNova, erroUpdate := cidadeController.cidadeService.UpdateCidade(*cidade, id)
	if erroUpdate != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"error": erroUpdate},
		)
		return
	}

	context.JSON(
		http.StatusOK,
		cidadeNova,
	)
}

func (cidadeController *CidadeController) deleteCidadeById(context *gin.Context) {
	idString := context.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"error": "ID enviado não é válido"},
		)
		return
	}

	deletado, erroDelete := cidadeController.cidadeService.DeleteCidadeById(id)
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
