package devfest

import (

"time"

)

type Post struct {

	ID int64					`json:"id"`
	Title string				`json:"title"`							
	Content string				`json:"content"`	
	Excerpt string				`json:"excerpt"`	
	CreateDate time.Time 		`json:"create_time"`	
	UpdateTime time.Time 		`json:"update_time"`	
	Image string				`json:"image"`	
} 


type Result struct {
	Result int64
}


type HeaderData struct {
	Email string
	Url string
}