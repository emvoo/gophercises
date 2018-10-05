package main

import (
	"net/http"
	"gopkg.in/yaml.v2"
	"encoding/json"
	"github.com/boltdb/bolt"
	"fmt"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls := []PathUrl{}
	if err := yaml.Unmarshal(yml, &pathUrls); err != nil {
		return nil, err
	}

	pathsToUrl := make(map[string]string)
	for _, item := range pathUrls {
		pathsToUrl[item.Path] = item.Url
	}

	return MapHandler(pathsToUrl, fallback), nil
}

func JSONHandler(jSon []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls := []PathUrl{}
	if err := json.Unmarshal(jSon, &pathUrls); err != nil {
		return nil, err
	}

	pathsToUrl := make(map[string]string)
	for _, item := range pathUrls {
		pathsToUrl[item.Path] = item.Url
	}

	return MapHandler(pathsToUrl, fallback), nil
}

func BoltHandler(db *bolt.DB, fallback http.Handler) (http.HandlerFunc, error) {
	pathsToUrl := make(map[string]string)
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("paths"))
		cursor := bucket.Cursor()

		for k, v := cursor.First(); k != nil; k, v = cursor.Next() {
			key := fmt.Sprintf("%s", k)
			value := fmt.Sprintf("%s", v)
			pathsToUrl[key] = value
		}
		return nil
	})

	return MapHandler(pathsToUrl, fallback), nil
}

type PathUrl struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}
