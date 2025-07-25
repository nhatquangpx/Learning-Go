package post

type Post struct {
	ID      int    `gorm:"primaryKey" json:"id"`
	Title   string `json:"title" validate:"required,min=3"`	
	Content string `json:"content" validate:"required"`
	AuthorID int   `json:"author_id"` 
}

type updateTitleRequest struct {
	Title string `json: "title" validate:"required, min = 3"`
}


