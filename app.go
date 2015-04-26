package geoimagestore

import (
  "net/http"

  "paths"
)

func init() {
  http.HandleFunc("/", paths.HandleRoot)
  http.HandleFunc("/serve/", paths.HandleServe)
  http.HandleFunc("/upload", paths.HandleUpload)
}
