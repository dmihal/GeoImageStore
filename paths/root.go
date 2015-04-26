package paths

import (
  "html/template"
  "net/http"

  "appengine"
  "appengine/blobstore"

  "lib"
)

var rootTemplate = template.Must(template.New("root").Parse(rootTemplateHTML))

const rootTemplateHTML = `
<html><body>
<form action="{{.}}" method="POST" enctype="multipart/form-data">
Upload File: <input type="file" name="file"><br>
<input type="submit" name="submit" value="Submit">
</form></body></html>
`

func HandleRoot(w http.ResponseWriter, r *http.Request) {
  c := appengine.NewContext(r)
  uploadURL, err := blobstore.UploadURL(c, "/upload", nil)
  if err != nil {
    lib.ServeError(c, w, err)
    return
  }
  w.Header().Set("Content-Type", "text/html")
  err = rootTemplate.Execute(w, uploadURL)
  if err != nil {
    c.Errorf("%v", err)
  }
}
