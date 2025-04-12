package Controller

import (
	"MarkDownAPI/Helper"
	"MarkDownAPI/Helper/StructStore"
	"MarkDownAPI/Package/Utility"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func AddFile(writer http.ResponseWriter, Req *http.Request) {
	responseInstance := StructStore.AddFileResponse{}

	// File Processing :-
	Req.ParseMultipartForm(1000 << 20)

	file, handler, err := Req.FormFile(Helper.FileHeaderName)

	if err != nil {

		responseInstance.IsAnyError = true

		response, _ := json.Marshal(responseInstance)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
		return
	}

	fmt.Print(handler.Filename)

	fileBytes, err := ioutil.ReadAll(file)

	if err != nil {

		responseInstance.IsAnyError = true

		response, _ := json.Marshal(responseInstance)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
		return
	}

	uuiID, err := uuid.NewRandom()

	Helper.FilesStored[uuiID.String()] = handler.Filename

	if err != nil {

		responseInstance.IsAnyError = true

		response, _ := json.Marshal(responseInstance)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
		return
	}

	fileLoc := Helper.FileStoreLocation + "/" + handler.Filename

	err = os.WriteFile(fileLoc, fileBytes, 0666)
	if err != nil {

		responseInstance.IsAnyError = true

		response, _ := json.Marshal(responseInstance)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
		return
	}

	//--
	responseInstance.FileName = uuiID.String()
	responseInstance.IsAnyError = true

	res, _ := json.Marshal(responseInstance)

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(res)
	return

}

func GetRenderedFileByID(writer http.ResponseWriter, Req *http.Request) {
	responseInstance := StructStore.GetFileResponse{}

	// Validating Input:
	requestGetRenderedFile := &StructStore.GetFileRequest{}

	Utility.ParseBody(Req, requestGetRenderedFile)

	IsBadRequest := ValidateReuqestBody(requestGetRenderedFile)
	if IsBadRequest != true {

		responseInstance.IsAnyError = true

		response, _ := json.Marshal(responseInstance)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
		return
	}

	//--

	_, exists := Helper.FilesStored[requestGetRenderedFile.FileName]
	if exists != true {
		responseInstance.IsAnyError = true

		response, _ := json.Marshal(responseInstance)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
		return
	}

	fileLoc := Helper.FileStoreLocation + "/" + Helper.FilesStored[requestGetRenderedFile.FileName]

	data, err := os.ReadFile(fileLoc)
	if err != nil {
		responseInstance.IsAnyError = true

		response, _ := json.Marshal(responseInstance)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "text/html; charset=utf-8")
	writer.Write(data)
	return

}

func GetFileByID(writer http.ResponseWriter, Req *http.Request) {
	responseInstance := StructStore.GetFileResponse{}

	// Validating Input:
	requestGetRenderedFile := &StructStore.GetFileRequest{}

	Utility.ParseBody(Req, requestGetRenderedFile)

	IsBadRequest := ValidateReuqestBody(requestGetRenderedFile)
	if IsBadRequest != true {

		responseInstance.IsAnyError = true

		response, _ := json.Marshal(responseInstance)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
		return
	}

	//--

	_, exists := Helper.FilesStored[requestGetRenderedFile.FileName]
	if exists != true {
		responseInstance.IsAnyError = true

		response, _ := json.Marshal(responseInstance)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
		return
	}

	fileLoc := Helper.FileStoreLocation + "/" + Helper.FilesStored[requestGetRenderedFile.FileName]

	data, err := os.ReadFile(fileLoc)
	if err != nil {
		responseInstance.IsAnyError = true

		response, _ := json.Marshal(responseInstance)

		writer.WriteHeader(http.StatusInternalServerError)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(response)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Disposition", "attachment; filename="+requestGetRenderedFile.FileName)
	writer.Header().Set("Content-Type", "application/octet-stream") // Generic binary type
	writer.Header().Set("Content-Length", string(len(data)))
	writer.Write(data)
	return
}

func GetAllFileByID(writer http.ResponseWriter, Req *http.Request) {

	var allFiles = StructStore.GetAllFileResponse{}

	if len(Helper.FilesStored) <= 0 {

		allFiles.FilesMetaData = nil
		allFiles.IsAnyError = true
		writer.WriteHeader(http.StatusInternalServerError)

	} else {

		allFiles.FilesMetaData = Helper.FilesStored
		allFiles.IsAnyError = false
		writer.WriteHeader(http.StatusOK)
	}

	response, _ := json.Marshal(allFiles)

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(response)

	return
}

//  ===========   VALIDATOR ==============

func ValidateReuqestBody[GenericType interface{}](ReqBody GenericType) bool {
	validatorInstance := validator.New()

	err := validatorInstance.Struct(ReqBody)

	if err != nil {
		return false
	}
	return true
}

//--
