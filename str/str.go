package str

import (
	"io"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

func Tmpl2Str(tmpl string, args interface{}) (string, error) {
	uuidV4, _ := uuid.NewRandom()
	t, err := template.New(uuidV4.String()).Parse(tmpl)
	if err != nil {
		return "", err
	}
	buf := &strings.Builder{}
	if err = t.Execute(buf, args); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func Tmpl2writer(tmpl string, args interface{}, writer io.Writer) error {
	uuidV4, _ := uuid.NewRandom()
	t, err := template.New(uuidV4.String()).Parse(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(writer, args)
}
