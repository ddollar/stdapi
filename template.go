package stdapi

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

var (
	helpers   TemplateHelpers
	templates = map[string]*template.Template{}
)

type TemplateHelpers func(r *http.Request) template.FuncMap

func LoadTemplates(dir string, fn TemplateHelpers) error {
	helpers = fn

	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files := []string{}

			files = appendIfExists(files, filepath.Join(dir, "layout.tmpl"))
			files = appendIfExists(files, filepath.Join(filepath.Dir(path), "layout.tmpl"))
			files = append(files, path)

			t, err := template.New("main").ParseFiles(files...)
			if err != nil {
				return err
			}

			rel, err := filepath.Rel(dir, path)
			if err != nil {
				return err
			}

			templates[rel] = t
		}

		return nil
	})
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, path string, params interface{}) error {
	t, ok := templates[fmt.Sprintf("%s.tmpl", path)]
	if !ok {
		return fmt.Errorf("no such template: %s", path)
	}

	if helpers != nil {
		t = t.Funcs(helpers(r))
	}

	var buf bytes.Buffer

	if err := t.Execute(&buf, params); err != nil {
		return errors.WithStack(err)
	}

	if _, err := io.Copy(w, &buf); err != nil {
		return err
	}

	return nil
}

func appendIfExists(files []string, path string) []string {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		files = append(files, path)
	}

	return files
}
