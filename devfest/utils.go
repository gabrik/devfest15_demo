package devfest

import (
	"io/ioutil"
	//"appengine/log"
	"appengine"
)



func loadPage(title string, c appengine.Context) string {
    filename := "/views/" + title + ".html"
    body, err := ioutil.ReadFile(filename)
    if err != nil {
        c.Errorf("Failed to retrieve file: %v", err)
        
    }


    header,err:= ioutil.ReadFile("/views/include/header.html")
    if err != nil {
        c.Errorf("Failed to retrieve file: %v", err)
        
    }

    footer,err:=ioutil.ReadFile("/views/include/footer.html")
    if err != nil {
        c.Errorf("Failed to retrieve file: %v", err)
    }


    bodys:=string(header)+string(body)+string(footer)

    return bodys
}

