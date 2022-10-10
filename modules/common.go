package modules

type MakeArticle struct {
	AuthorId string `json:"author_id" binding:"required" `
	Content
}

type UpadateArticle struct{
	Id      string `json:"id" binding:"required"`
	Content        
}

type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type JSONErrorResponse struct {
	Error string      `json:"error"`
	
}