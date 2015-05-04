package paths

import (
  "net/http"

  "appengine"
  "appengine/image"

  "lib"
)

func HandleServe(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)

  options := image.ServingURLOptions {
    Secure: false, // whether the URL should use HTTPS
    Size: 400,
    Crop: false,
  }
  key := appengine.BlobKey(r.FormValue("blobKey"))
  url, err := image.ServingURL(c, key, &options)

  if err != nil {
    lib.ServeError(c, w, err)
    return
  }

  http.Redirect(w, r, url.String(), http.StatusFound)
}
