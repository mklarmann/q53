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

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Greeting").Order("-Date").Limit(10)
	greetings := make([]Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := guestbookTemplate.Execute(w, greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var guestbookTemplate = template.Must(template.New("book").Parse(guestbookTemplateHTML))

const guestbookTemplateHTML = `
<html>
  <body>
    {{range .}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`

func sign(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	c.Infof("New entry from URL: %v", r.URL)
	g := Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Greeting", nil), &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
