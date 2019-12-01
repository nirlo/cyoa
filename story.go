package cyoa

import (
  "encoding/json"
  "io"
  "net/http"
  "http/template"
)

func init() {
  tpl = template.Must(template.New("").Parse(defaultHandlerTmpl))
}

var tpl *template.template

var defaultHandlerTmpl = `
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <title>Choose Your Own Adventure</title>
</head>
<body>
  <h1>{{.Title}}</h1>
  {{ range .Paragraphs }}
  <p>{{ . }}</p>
  {{ end }}
  <ul>
  {{ range .Options }}
    <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
  {{ end }}
</body>
</html>
`

func NewHandler(s Story) http.Handler {
  return handler{s}
}

type handler struct {
  s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  err := tpl.Execute(w, h.s["intro"])
  if err != nil {
    panic(err)
  }
}

func JsonStory(r io.Reder) (Story, error) {
  d := json NewDecoder(r)
  var story cyoa.Story
  if err := d.Decode(&story); err != nil {
    return nil, err
  }
  return story, nill
}

type Story map[string]Chapter

type Chapter struct {
  Title string `json:"title"`
  Paragraphs []string `json:"story"`
  Options []option `json:"options"`
}

type Option struct {
  Text string `json:"text"`
  Chapter string `json:"arc"`
}


