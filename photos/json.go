package photos

import (
  "encoding/json"
)

func EncodeJSON(photos []Photo) (string, error) {
  bytes, err := json.Marshal(photos)
  return string(bytes), err
}
