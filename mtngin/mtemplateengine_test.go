package mtngin

import "testing"

const (
	PASSMARK          = "\u2713"
	FAILMARK          = "\u2717"
	DIR_TEMPLATE_ROOT = "./template"
	DIR_LAYOUT        = "layout"
	DIR_STATIC        = "static"
)

func TestParseTemplateLayoutFiles(t *testing.T) {
	t.Logf("executing test: parse template layout files %s", PASSMARK)

	te := NewTemplateEngine(DIR_TEMPLATE_ROOT, DIR_LAYOUT, DIR_STATIC)

	layoutDir := te.Layouts
	layouts, _ := te.ParseDirTemplates(layoutDir)

	for k, v := range layouts {
		t.Logf("layout cache key: %s %s\n", k, PASSMARK)
		t.Logf("layout cached file: %s %s\n", v, PASSMARK)
	}

	t.Logf("%s", PASSMARK)
}

func TestParseTemplateStaticFiles(t *testing.T) {
	t.Logf("executing test: parse template static files %s", PASSMARK)

	te := NewTemplateEngine(DIR_TEMPLATE_ROOT, DIR_LAYOUT, DIR_STATIC)

	staticDir := te.Statics
	statics, _ := te.ParseDirTemplates(staticDir)

	for k, v := range statics {
		t.Logf("static cache key: %s %s\n", k, PASSMARK)
		t.Logf("static cached file: %s %s\n", v, PASSMARK)
	}

	t.Logf("%s", PASSMARK)
}

func TestCacheTemplates(t *testing.T) {
	t.Logf("executing test: cache templates %s", PASSMARK)

	te := NewTemplateEngine(DIR_TEMPLATE_ROOT, DIR_LAYOUT, DIR_STATIC)

	layouts, _ := te.ParseDirTemplates(te.Layouts)
	statics, _ := te.ParseDirTemplates(te.Statics)

	var key string = "base"
	base, ok := layouts[key]

	if !ok {
		t.Errorf("could not find layout file with key: %s %s\n", key, FAILMARK)
	}

	te.CacheTemplates(base, statics)

	t.Logf("%s", PASSMARK)
}
