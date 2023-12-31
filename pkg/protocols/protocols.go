package protocols


type GenericRequest struct {
	Command string `json:"command"`
}

type GenericResponse struct {
	Command string `json:"command"`
	Status string `json:"status"`
}


type CreateSiteRequest struct {
	Command string `json:"command"`
}

type CreateSiteResponse struct {
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

type UpdateSiteResponse struct {
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
