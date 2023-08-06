package domain

import (
	"errors"
	"io"
	"strings"
	"testing"
)

type testReadCloser struct {
	Reader io.Reader
}

func (testReadCloser) Close() error                       { return nil }
func (t testReadCloser) Read(p []byte) (n int, err error) { return t.Reader.Read(p) }

type errReader struct{}

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("Some error!")
}

func TestExtractPrice(t *testing.T) {
	ext := PriceExtractor{}
	s := `
		<html>
			<head>
				<title>test</title>
			</head>
			<body>
				<table>
					<td> покупка: 3232.00</td>
				</table>
			</body>
		</html>`
	r := strings.NewReader(s)
	rc := testReadCloser{r}
	var want float32 = 3232.00
	got, err := ext.ExtractPrice(rc)
	if err != nil {
		t.Errorf("got error: %v, wanted %v", err, want)
	}
	if got != want && err == nil {
		t.Errorf("got %v, wanted %v", got, want)
	}

	s = `
		<html>
			<head>
				<title>test</title>
			</head>
			<body>
				<table>
					<td> покупка: 3232.00wrong value</td>
				</table>
			</body>
		</html>`
	rc.Reader = strings.NewReader(s)
	want = 0.00
	errTemplt := "cannot parse the string data:"
	got, err = ext.ExtractPrice(rc)
	if err != nil && !strings.Contains(err.Error(), errTemplt) {
		t.Errorf("got not wanted error: %v, wanted error template: %v", err, errTemplt)
	}
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	s = `
		<html>
			<head>
				<title>test</title>
			</head>
			<body>
				<table>
					<td> покупка: wrong value</td>
				</table>
			</body>
		</html>`
	rc.Reader = strings.NewReader(s)
	errTemplt = "the document does not have a price value with the tag:"
	got, err = ext.ExtractPrice(rc)
	if err != nil && !strings.Contains(err.Error(), errTemplt) {
		t.Errorf("got not wanted error: %v, wanted error template: %v", err, errTemplt)
	}
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	errTemplt = "cannot parse the body to an HTML document:"
	rc.Reader = errReader{}
	got, err = ext.ExtractPrice(rc)
	if err != nil && !strings.Contains(err.Error(), errTemplt) {
		t.Errorf("got not wanted error: %v, wanted error template: %v", err, errTemplt)
	}
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
