package server

import (
	"html/template"
)

func (s *server) parseTemplates(tpls ...string) (*template.Template, error) {
	const (
		tplRoot       = "templates/"
		fileExt       = ".html"
		defaultLayout = "layouts/default"
	)
	var defaultComponents = []string{
		"components/navbar",
	}
	tpls = append(tpls, defaultLayout)
	tpls = append(tpls, defaultComponents...)
	for i, t := range tpls {
		tpls[i] = tplRoot + t + fileExt
	}

	return template.New("protoTpl").
  Funcs(template.FuncMap{
    // getMetaField retrieves a top-level field from Data.ContentMeta map and returns it as an HTML string.
    // ! this will crash template if `d` is not of data type
    "getMetaField": func(d data, k string) template.HTML {
      v, ok := d.ContentMeta[k].(string)
      if ok {
        switch k {
        case "Title":
          v = `<h1 class="col-12 mt-4 text-center">` + v + "</h1>"
          return template.HTML(v)
        default:
          return template.HTML(v)
        }
      }
      s.l.Printf(`tpl exe func err: {{getMetaField %v %s}}
d.ContentMeta[k].(string) !ok`, d, k)
      return ""
    },
  }).ParseFS(s.fs, tpls...)
}
