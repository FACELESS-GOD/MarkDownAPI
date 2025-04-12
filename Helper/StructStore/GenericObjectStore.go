package StructStore

type GenericMessageReponse struct {
	Message string
}

type GenericCountReponse struct {
	Count int
}

// ==========    Request Struct    ==========
type AddFileRequest struct {
	FileName string `json:"filename" validate:"required,min=2,max=50"`
}

type GetFileRequest struct {
	FileID   int64
	FileName string `json:"filename" validate:"required,min=2,max=50"`
}

//--

// ==========    Response Struct    ==========

type AddFileResponse struct {
	FileName   string
	IsAnyError bool
}

type GetFileResponse struct {
	FileID     int64
	FileName   string
	IsAnyError bool
}

type GetAllFileResponse struct {
	FilesMetaData map[string]string
	IsAnyError    bool
}
