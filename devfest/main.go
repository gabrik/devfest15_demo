package devfest

import (

"net/http"


"html/template"

)






//faccio il caching dei template html
var templates = template.Must(template.ParseGlob("template/*"))





func init() {

	//setta gli handler

	http.HandleFunc("/", handler)


	http.HandleFunc("/post",savePostHanlder)
	http.HandleFunc("/get/",getPosts)
	

	http.HandleFunc("/publish", sendPost)


}

