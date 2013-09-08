// Copyright 2013 of Manuel Klarmann (aka Q52 or mklarmann). All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package q53

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"html/template"
	"net/http"
	"time"
)


type Answers struct {
	Phrases  []string
	Probabilities []float64
	Transitions []float64
	Votes [2]int // first for Probabilities, seccond for Transitions
}

type Question struct {
	Phrase  string
	Responses Answers
	
	
	Account string
	Date    time.Time
}



var (
	// initialization variables / questions
	
	q_alpha = Question{
		Phrase:  "Are you seeking for answers?",
		Responses: Answers{
				Phrases: []string{"yes","no"},
				Probabilities: []float64{0.5,0.5},
				Transitions: []float64{0.5,0.5},
				Votes: [2]int{0,1},
			},
	}

	q_omega = Question{
		Phrase:  "Are you satisfied?",
		Responses: Answers{
				Phrases: []string{"yes","no"},
				Probabilities: []float64{0.5,0.5},
				Transitions: []float64{0.5,0.5},
				Votes: [2]int{0,1},
			},

	}

)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)


	q := datastore.NewQuery("Question").Limit(10)
	questions_query := make([]Question, 0, 10)
	if _, err := q.GetAll(c, &questions_query); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	questions := append(questions_query,q_alpha,q_omega)
	if err := guestbookTemplate.Execute(w, questions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var guestbookTemplate = template.Must(template.New("book").Parse(guestbookTemplateHTML))

const guestbookTemplateHTML = `
<html>
  <body>
    {{range .}}
	  {{with .Account}}
        <p><b>{{.}}</b> wanted to know:</p>
      {{else}}
        <p>The following generic questions are stored:</p>
      {{end}}
      <pre>{{.Phrase}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="phrase" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Ask a new Question"></div>
    </form>
  </body>
</html>
`

func sign(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("New entry from URL: %v", r.URL)
	q := Question{
		Phrase: r.FormValue("phrase"),
		Date:   time.Now(),
	}
	if u := user.Current(c); u != nil {
		q.Account = u.String()
	}
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Question", nil), &q)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
