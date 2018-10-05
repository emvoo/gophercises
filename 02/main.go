package main

import (
	"net/http"
	"fmt"
	"github.com/boltdb/bolt"
)

func main()  {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("paths"))
		if err != nil {
			return err
		}
		bucket.Put([]byte("/urlshort"), []byte("https://godoc.org/github.com/gophercises/urlshort"))
		bucket.Put([]byte("/urlshort-final"), []byte("https://github.com/gophercises/urlshort/tree/solution"))
		return nil
	})

	// yamlFile := flag.String("yaml", "paths.yaml", "makes it possible to load yaml files")
	// jsonFile := flag.String("json", "paths.json", "makes it possible to load yaml files")
	// flag.Parse()

	// yaml, err := ioutil.ReadFile(*yamlFile)
	// if err != nil {
	// 	panic(err)
	// }
	//
	// jsons, err := ioutil.ReadFile(*jsonFile)
	// if err != nil {
	// 	panic(err)
	// }

	mux := newMuxer()

	paths := map[string]string{
		// "/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		// "/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	mapHandler := MapHandler(paths, mux)
	//
	// yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	// if err != nil {
	// 	panic(err)
	// }

	// jsonHandler, err := JSONHandler(jsons, mapHandler)
	// if err != nil {
	// 	panic(err)
	// }

	boltHandler, err := BoltHandler(db, mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", boltHandler)
}

func newMuxer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w,"Hello world")
}