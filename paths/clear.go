package paths

import (
    "fmt"
    "net/http"

    "appengine"
    "appengine/search"

    "lib"
)

func HandleClear(w http.ResponseWriter, r *http.Request){
  c := appengine.NewContext(r)

  index, err := search.Open("photos")
  if err != nil {
    lib.ServeError(c, w, err)
    return
  }

  options := search.SearchOptions{
    IDsOnly: true,
  }

  for t := index.Search(c, "", &options); ; {
    id, err := t.Next(&struct{}{})
    if err == search.Done {
      break
    }
    if err != nil {
      lib.ServeError(c, w, err)
      return
    }
    err = index.Delete(c, id)
  }

  fmt.Fprint(w, "Cleared")
}
