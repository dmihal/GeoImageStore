package paths

import (
  "net/http"

  "appengine"
  "appengine/blobstore"

  "lib"
)

func HandleUpload(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  blobs, _, err := blobstore.ParseUpload(r)
  if err != nil {
    lib.ServeError(c, w, err)
    return
  }
  file := blobs["file"]
  if len(file) == 0 {
    c.Errorf("no file uploaded")
    http.Redirect(w, r, "/", http.StatusFound)
    return
  }
  http.Redirect(w, r, "/serve/?blobKey="+string(file[0].BlobKey), http.StatusFound)
}
