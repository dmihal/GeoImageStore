package geoimagestore

import (
  "net/http"

  "paths"
)

func init() {
  http.HandleFunc("/", paths.HandleRoot)
  http.HandleFunc("/serve/", paths.HandleServe)
  http.HandleFunc("/upload", paths.HandleUpload)
  http.HandleFunc("/upload_url", paths.HandleUploadURL)
  http.HandleFunc("/list", paths.HandleList)
  http.HandleFunc("/nearby", paths.HandleNearby)
  http.HandleFunc("/nearby.html", paths.HandleNearbyList)
  http.HandleFunc("/clear", paths.HandleClear)
}
