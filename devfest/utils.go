package devfest

import (

	"appengine"
    "appengine/datastore"

)




//genera la chiave per i post nel datastore
func postKey(c appengine.Context) *datastore.Key {

    return datastore.NewKey(c, "Post", "default_post", 0, nil)
}

