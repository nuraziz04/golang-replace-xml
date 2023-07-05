package controller

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"github.com/nuraziz04/golang-esta/exception"
	"github.com/nuraziz04/golang-esta/model"
	"github.com/tealeg/xlsx"
)

type EstaControllerImpl struct {
}

func NewEstaController() EstaController {
	return &EstaControllerImpl{}
}

func (controller *EstaControllerImpl) Generate(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(request.Body)

	rb := model.RequestBody{}
	decoder.Decode(&rb)

	fileExcel := rb.FileExcel + ".xlsx"
	fileDss := rb.FileDss
	fileDssResult := fileDss + "-modified"

	// Get the absolute path of the file
	filePathExcel := filepath.Join(".", "file", "excel", fileExcel)

	// Open the Excel file
	xlFile, err := xlsx.OpenFile(filePathExcel)
	if err != nil {
		err = errors.New("File Excel Not Found")
		panic(exception.NewNotFoundError(err.Error()))
	}

	// Get the sheet containing the ID replacements
	sheet := xlFile.Sheets[0]

	// Read the ID replacements from the Excel sheet
	idReplacements := make(map[string]string)
	for _, row := range sheet.Rows {
		idCell := row.Cells[0]
		configCell := row.Cells[1]
		id := idCell.String()
		config := configCell.String()
		idReplacements[id] = config
	}

	// Get the absolute path of the file
	filePathDss := filepath.Join(".", "file", "original", fileDss+".dbs")

	// Open the .dbs file
	dbsFile, err := os.Open(filePathDss)
	if err != nil {
		err = errors.New("File DSS Not Found")
		panic(exception.NewNotFoundError(err.Error()))
	}
	defer dbsFile.Close()

	// Parse the .dbs file
	decoderXml := xml.NewDecoder(dbsFile)
	queries := model.Queries{}
	err = decoderXml.Decode(&queries)
	if err != nil {
		panic(err)
	}

	// Replace the IDs in the queries
	for i := range queries.Data.Queries {
		query := &queries.Data.Queries[i]
		if newConfig, ok := idReplacements[query.ID]; ok {
			query.UseConfig = newConfig
		}
	}

	// Get the absolute path of the file
	filePathDssResult := filepath.Join(".", "file", "modified", fileDssResult+".dbs")

	// Write the modified queries back to the .dbs file
	dbsOutput, err := os.Create(filePathDssResult)
	if err != nil {
		panic(err)
	}
	defer dbsOutput.Close()

	encoderXml := xml.NewEncoder(dbsOutput)
	encoderXml.Indent("", "   ")
	err = encoderXml.Encode(queries)
	if err != nil {
		panic(err)
	}

	// response
	response := model.Response{
		Message: "Success",
	}

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}
