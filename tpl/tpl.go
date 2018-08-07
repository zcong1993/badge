package tpl

import "html/template"

func add(a, b float64) float64 {
	return a + b
}

// TPL is struct of tpl engine
type TPL struct {
	DefaultTpl *template.Template
	FlatTpl    *template.Template
}

// NewTPL is constructor for TPL
func NewTPL() *TPL {
	tpl := &TPL{}
	funcs := template.FuncMap{
		"add": add,
	}
	dft := template.Must(template.New("default").Funcs(funcs).Parse(DefaultTpl))
	flat := template.Must(template.New("flat").Funcs(funcs).Parse(FlatTpl))

	tpl.DefaultTpl = dft
	tpl.FlatTpl = flat

	return tpl
}
