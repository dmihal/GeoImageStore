package paths

import (
    "fmt"
    "net/http"

    "appengine"
    "appengine/blobstore"

    "lib"
)

func HandleUploadURL(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  uploadURL, err := blobstore.UploadURL(c, "/upload", nil)
  if err != nil {
    lib.ServeError(c, w, err)
    return
  }

  fmt.Fprint(w, uploadURL)
}
