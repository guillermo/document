package document

import (
	"github.com/guillermo/terminal/assert"
	"strings"
	"testing"
)

var docTests = []struct {
	in  string
	out []string
}{
	{`<P>Paragraph</p><p>Paragraph 2</p>`, []string{"Paragraph ", "Paragraph ", "2         "}},
	{`<P style="display:inline;">P1</p><p style="display:inline;">P2</p>`, []string{"P1P2"}},
}

func TestDocument(t *testing.T) {
	for _, tCase := range docTests {
		t.Run(tCase.in, func(t *testing.T) {

			a := &assert.TArea{
				T: t,
			}
			doc := &Document{
				Width: 10,
				Area:  a,
			}

			doc.Parse(strings.NewReader(tCase.in))

			a.AssertLines(tCase.out...)

		})
	}

}
