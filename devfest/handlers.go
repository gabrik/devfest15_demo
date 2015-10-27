
package devfest

import (
"fmt"
"net/http"
"time"
"encoding/json"
"strconv"


"appengine"
"appengine/datastore"

"appengine/user"


"html"
"html/template"

)




func handler(w http.ResponseWriter, r *http.Request) {



	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r)




	//carico il template angularjs per la single-page app
	indexTmpl := template.New("index.html").Delims("<<<", ">>>") //cambio i delimiter usati da golang per i template
	indexTmpl, _ = indexTmpl.ParseFiles("static_files/html/index.html")
	err := indexTmpl.Execute(w, nil) 

	if err != nil {
		c.Errorf("load: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


}




//crea la vista per l'invio del post, controlla se l'utente è autenticato
func sendPost(w http.ResponseWriter, r *http.Request) {


	w.Header().Set("Content-type", "text/html; charset=utf-8")
	c := appengine.NewContext(r)
	u := user.Current(c)
	
	var err error

	if u == nil { 
		//se utente non autenticato s
		url, _ := user.LoginURL(c, "/publish")
		err = templates.ExecuteTemplate(w, "navbar_nu", url)
		if err != nil {
			c.Errorf("load: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = templates.ExecuteTemplate(w, "noauth", url)
		if err != nil {
			c.Errorf("load: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	

	//se utente autenticato
	
	url, _ := user.LogoutURL(c, "/")
	err = templates.ExecuteTemplate(w, "navbar", &HeaderData{ Email:u.Email,Url:url, })
	if err != nil {
		c.Errorf("load: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = templates.ExecuteTemplate(w, "publishpage", u)
	if err != nil {
		c.Errorf("load: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	


	//url, _ := user.LogoutURL(c, "/")
	//fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)




}







//salva i post sul db e stampa a vista il risultato
func savePostHanlder(w http.ResponseWriter, r *http.Request){


	c := appengine.NewContext(r)
	u := user.Current(c)

	w.Header().Set("Content-type", "text/html; charset=utf-8")


	//fmt.Fprintln(w,r)

	if u == nil {
		//_, _ := user.LoginURL(c, "/")
		fmt.Fprintf(w, `{"Result":-2}`)
		

		return
	}

	if (r.Method == "POST"){

		
		
		p:=Post{
			//ID : 0,		
			Title : html.EscapeString(r.FormValue("title")),					
			Content : html.EscapeString(r.FormValue("content")),
			Excerpt : html.EscapeString(r.FormValue("excerpt")),
			Image : html.EscapeString(r.FormValue("image")),
		} 

		p.CreateDate=time.Now()
		p.UpdateTime=time.Now()

		key := datastore.NewIncompleteKey(c, "Post", postKey(c))

		k,err := datastore.Put(c, key, &p)
		if err != nil {

			/*r:=Result{
				Result:-1,
				}*/

			/*b, err := json.Marshal(r)
			if err != nil {
				fmt.Println("error:", err)
				}*/

			//fmt.Fprint(w, string(b))

			//var err error


				url, _ := user.LogoutURL(c, "/")
				err = templates.ExecuteTemplate(w, "navbar", &HeaderData{ Email:u.Email,Url:url, })
				if err != nil {
					c.Errorf("load: %v", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				err = templates.ExecuteTemplate(w, "postko", u)
				if err != nil {
					c.Errorf("load: %v", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}



		} else {

			/*r:=Result{
				Result:k.IntID() ,
				}*/



				p.ID=k.IntID() 

				datastore.Put(c, k, &p)

				datastore.Delete(c, key)

			/*b, err := json.Marshal(r)
			if err != nil {
				fmt.Println("error:", err)
				}*/


			//fmt.Fprint(w, string(b))
				url, _ := user.LogoutURL(c, "/")
				err = templates.ExecuteTemplate(w, "navbar", &HeaderData{ Email:u.Email,Url:url, })
				if err != nil {
					c.Errorf("load: %v", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

				err = templates.ExecuteTemplate(w, "postok", u)
				if err != nil {
					c.Errorf("load: %v", err)
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}

			}






		} else {

			var err error

			url, _ := user.LoginURL(c, "/publish")
		//fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
			err = templates.ExecuteTemplate(w, "navbar_nu", url)
			if err != nil {
				c.Errorf("load: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			err = templates.ExecuteTemplate(w, "noauth", url)
			if err != nil {
				c.Errorf("load: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}


			return
		}

	}



//restituisce i post in JSON
	func getPosts(w http.ResponseWriter, r *http.Request) {







		if (r.Method == "GET"){


			//fmt.Fprintln(w,r.URL.Path[len("/get/"):] )
			c := appengine.NewContext(r)




			if(r.URL.Path[len("/get/"):]!=""){

				id,_:=strconv.ParseInt(r.URL.Path[len("/get/"):], 10, 64)
				
				//controllare errore

				q := datastore.NewQuery("Post").Filter("ID=",id)
				posts := make([]Post, 0, 10)
				if _, err := q.GetAll(c, &posts); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				p:=posts[0]
				b, err := json.Marshal(p)
				if err != nil {
					fmt.Println("error:", err)
				}

				fmt.Fprint(w, string(b))

				

			} else {

				

				q := datastore.NewQuery("Post")


				posts := make([]Post, 0, 10)
				if _, err := q.GetAll(c, &posts); err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}


		//fmt.Fprintln(w, posts)
				b, err := json.Marshal(posts)
				if err != nil {
					fmt.Println("error:", err)
				}

				fmt.Fprint(w, string(b))


			}



		} else {



			http.Error(w, "Wrong Method", http.StatusMethodNotAllowed)
			return

		}

	}

