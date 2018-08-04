package xml3

import (
	"fmt"
	"io"
	"strings"
)

func (r *Renderer) out(w io.Writer, d []byte) {
	w.Write(d)
}

func (r *Renderer) outs(w io.Writer, s string) {
	io.WriteString(w, s)
}

func (r *Renderer) cr(w io.Writer) {
	r.outs(w, "\n")
}

func (r *Renderer) outTag(w io.Writer, name string, attrs []string) {
	s := name
	if len(attrs) > 0 {
		s += " " + strings.Join(attrs, " ")
	}
	io.WriteString(w, s+">")
}

func (r *Renderer) outOneOf(w io.Writer, outFirst bool, first string, second string) {
	if outFirst {
		r.outs(w, first)
	} else {
		r.outs(w, second)
	}
}

// outTagContents output the opening tag with possible attributes, then the content
// and then the closing tag.
func (r *Renderer) outTagContent(w io.Writer, name string, attrs []string, content string) {
	s := name
	if len(attrs) > 0 {
		s += " " + strings.Join(attrs, " ")
	}
	io.WriteString(w, s+">")
	io.WriteString(w, content)
	io.WriteString(w, "</"+name[1:]+">\n")
}

func (r *Renderer) sectionClose(w io.Writer) {
	if r.section == nil {
		return
	}

	tag := "</section>"
	if r.section.Special != nil {
		tag = "</note>"
		if isAbstract(r.section.Special) {
			tag = "</abstract>"
		}
	}
	r.outs(w, tag)
	r.cr(w)
}

func attributes(keys, values []string) (s []string) {
	for i, k := range keys {
		if values[i] == "" { // skip entire k=v is value is empty
			continue
		}
		s = append(s, fmt.Sprintf(`%s="%s"`, k, values[i]))
	}
	return s
}

func isAbstract(word []byte) bool {
	return strings.EqualFold(string(word), "abstract")
}
