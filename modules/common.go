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

type UpdateAuthor struct{
	Id      string `json:"id" binding:"required"`
	Firstname string     `json:"firstname"   binding:"required"  minLength:"2" maxLength:"30" example:"Jack"`
	Lastname  string     `json:"lastname"   binding:"required" minLength:"2" maxLength:"30" example:"Haldson"`
	     
}
