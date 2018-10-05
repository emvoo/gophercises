package controllers

import (
	"gophercises/03/app"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"gophercises/03/types"
	"bytes"
	"html/template"
)

func LoadStory(s *app.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		byt, err := ioutil.ReadFile("story.json")
		if err != nil {
			panic(err)
		}
		var stories types.Stories

		if err = json.Unmarshal(byt, &stories); err != nil {
			panic(err)
		}

		storyStr := r.URL.Query().Get("story")
		var story types.Story
		for _, v := range stories {
			switch storyStr {
			case "new-york":
				story = v.NewYork.Story
				break
			case "debate":
				story = v.Debate.Story
				break
			case "sean-kelly":
				story = v.SeanKelly.Story
				break
			case "mark-bates":
				story = v.MarkBates.Story
				break
			case "denver":
				story = v.Denver.Story
				break
			case "home":
				story = v.Home.Story
				break
			default:
				story = v.Intro.Story
			}
		}

		tpl := template.Must(template.ParseFiles("views/index.html"))
		var writer bytes.Buffer
		if err = tpl.ExecuteTemplate(&writer, "main", story); err != nil {
			panic(err)
		}

		if _, err = w.Write(writer.Bytes()); err != nil {
			panic(err)
		}
	}
}
