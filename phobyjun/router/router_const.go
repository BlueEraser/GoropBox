package router

const (
	API           = "/api"
	APIAuth       = API + "/auth"
	APIAuthSignup = APIAuth + "/new"
)

const (
	APIFile    = API + "/file"
	APIFilesID = APIFile + "/:id"
)
