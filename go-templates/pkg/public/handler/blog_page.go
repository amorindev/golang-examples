package handler

import (
	"html/template"
	"net/http"
)

type Article struct {
	ID      string
	Title   string
	Content string
}

var articles = []*Article{
	{
		ID:      "1",
		Title:   "first article",
		Content: "first article with content",
	},
	{
		ID:      "2",
		Title:   "second article",
		Content: "second article with content",
	},
}

func BlogPage(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"web/public/templates/base.html",
		"web/public/templates/blog.html",
		"web/public/templates/components/header.html",		
		"web/public/templates/components/footer.html",		
		"web/public/templates/components/articles_card.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := struct {
		ActivePage string
		Articles   []*Article
	}{
		ActivePage: "blog",
		Articles:   articles,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
