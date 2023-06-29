package protocols

type CreateSiteRequest struct {
	Command string `json:"command"`
}

type CreateSiteReponse struct {
	Command string `json:"command"`
	Status string `json:"status"`
}


type GetSiteRequest struct {
	Command string `json:"command"`
}

type GetSiteResponse struct {
	Command string `json:"command"`
	Status string `json:"status"`
}


type UpdateSiteRequest struct {
	Command string `json:"command"`
}

type UpdateSiteReponse struct {
	Command string `json:"command"`
	Status string `json:"status"`
}


type CreateFileRequest struct {
	Command string `json:"command"`
}

type CreateFileResponse struct {
	Command string `json:"command"`
	Status string `json:"status"`
}


type GetFileRequest struct {
	Command string `json:"command"`
}

type GetFileResponse struct {
	Command string `json:"command"`
	Status string `json:"status"`
}


type UpdateFileRequest struct {
	Command string `json:"command"`
}

type UpdateFileResponse struct {
	Command string `json:"command"`
	Status string `json:"status"`
}
