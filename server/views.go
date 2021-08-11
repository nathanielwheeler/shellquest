package server

import (
	"html/template"
	"net/http"
)

type view struct {
	Data   *data
	Tpl    *template.Template
	TplErr error
}

type data struct {
	Content     template.HTML
	ContentMeta map[string]interface{}
}

func (s *server) newView(tpls ...string) *view {
	tpl, tplerr := s.parseTemplates(tpls...)
	return &view{
		Tpl:    tpl,
		TplErr: tplerr,
		Data:   &data{},
	}
}

func (s *server) addViewContent(v *view, f string) error {
	d := v.Data
	str, yaml, err := s.parseMarkdown(f)
	if err != nil {
		s.l.Printf("Markdown Parsing Error:\n\t%s\n", err)
		return err
	}
	d.Content = template.HTML(str)
	d.ContentMeta = yaml
	return nil
}

func (s *server) execView(v *view, w http.ResponseWriter, r *http.Request) {
	if v.TplErr != nil {
		s.l.Printf("Template Parsing Error:\n\t%s\n", v.TplErr)
		http.Error(w, errPublic.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")

	v.TplErr = v.Tpl.ExecuteTemplate(w, "default", v.Data)
	if v.TplErr != nil {
		s.l.Printf("Template Execution Error:\n\t%s\n", v.TplErr)
		http.Error(w, errPublic.Error(), http.StatusInternalServerError)
		return
	}
}
