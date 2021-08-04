package category

type CategoryInput struct {
	CategoryName string `json:"categoryname" binding:"required"`
}

type CategoryIDInput struct {
	ID int `uri:"id" binding:"required"`
}