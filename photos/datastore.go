package photos

import (
  "fmt"

  "appengine"
  "appengine/datastore"
  "appengine/search"
)

const PHOTO_KIND = "Photo"
const KILOMETER = 1000

func Add(photo Photo, c appengine.Context) error {
  sPhoto := makeSearchPhoto(photo)

  index, err := search.Open("photos")
  if err != nil {
    return err
  }
  indexKey, err := index.Put(c, "", &sPhoto)
  if err != nil {
    return err
  }

  key := datastore.NewKey(c, PHOTO_KIND, indexKey, 0, nil)
  datastore.Put(c, key, &photo)

  return nil
}
func GetAll(c appengine.Context) []Photo {
  query := datastore.NewQuery(PHOTO_KIND)

  return runQuery(query, c)
}
func GetNearest(point appengine.GeoPoint, c appengine.Context) ([]Photo, error) {
  index, err := search.Open("photos")
  if err != nil {
    return nil, err
  }
  distanceStr := fmt.Sprintf("distance(Location, geopoint(%f, %f))", point.Lat, point.Lng)
  query := fmt.Sprintf("%s < %d", distanceStr, 500*KILOMETER)

  sort := search.SortOptions {
    Expressions: []search.SortExpression{
      search.SortExpression {
        Expr: distanceStr,
        Reverse: true,
      },
    },
  }
  options := search.SearchOptions{
    Limit: 20,
    IDsOnly: true,
    Sort: &sort,
  }

  var keys []*datastore.Key
  for t := index.Search(c, query, &options); ; {
    var item searchPhoto
    id, err := t.Next(&item)
    if err == search.Done {
      break
    }
    if err != nil {
      return nil, err
    }
    key := datastore.NewKey(c, PHOTO_KIND, id, 0, nil)
    c.Infof(key.Encode())
    keys = append(keys, key)
  }

  photos := make([]Photo, len(keys))
  err = datastore.GetMulti(c, keys, photos)
  if err != nil {
    return nil, err
  }

  return photos, nil
}
func runQuery(query *datastore.Query, c appengine.Context) []Photo {
  var photos []Photo

  for t := query.Run(c); ; {
    var photo Photo
    _, err := t.Next(&photo)
    if err == datastore.Done {
      break
    }
    if err != nil {
      return nil
    }
    photos = append(photos, photo)
  }
  return photos
}
