package mtngin

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
)

type templates map[string]*template.Template

type templateDir struct {
	Root    string
	Layouts string
	Statics string
}

type TemplateEngine struct {
	TemplateCache templates
	*templateDir
}

func (te *TemplateEngine) ParseDirTemplates(dir string) (map[string]string, error) {
	var staticHTML map[string]string = make(map[string]string)
	var file string
	var key string
	var errstring string = "error loading static HTML:"

	d, err := os.Open(dir)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %v\n", errstring, err))
	}

	defer d.Close()

	filenames, err := d.Readdirnames(0)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%s %v\n", errstring, err))
	}

	var offset = 0

	for _, v := range filenames {
		file = path.Base(v)
		offset = len(file) - len(filepath.Ext(file)) // could also use TrimSuffix ..
		key = file[0:offset]
		staticHTML[key] = file
	}

	return staticHTML, nil
}

func (te *TemplateEngine) CacheTemplates(layout string, statics map[string]string) error {
	t := path.Join(te.Layouts, layout)
	for k, v := range statics {
		s := path.Join(te.Statics, v)
		te.TemplateCache[k] = te.fromLayout(t, s)
	}
	return nil
}

func (te *TemplateEngine) fromLayout(layout, static string) *template.Template {
	return template.Must(
		template.ParseFiles(
			layout,
			static,
		),
	)
}

func NewTemplateEngine(dirRoot, dirLayout, dirStatic string) *TemplateEngine {
	return &TemplateEngine{
		TemplateCache: make(templates),
		templateDir: &templateDir{
			Root:    dirRoot,
			Layouts: path.Join(dirRoot, dirLayout),
			Statics: path.Join(dirRoot, dirStatic),
		},
	}
}
