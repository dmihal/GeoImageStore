package paths

import (
  "net/http"

  "appengine"
  "appengine/blobstore"
)

func HandleServe(w http.ResponseWriter, r *http.Request) {
  blobstore.Send(w, appengine.BlobKey(r.FormValue("blobKey")))
}
