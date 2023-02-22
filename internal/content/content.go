package internal

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v2"
)

type Content struct {
	Metadata map[string]interface{}
	HTML     []byte
}

func loadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func parseMarkdown(data []byte) []byte {
	extensions := blackfriday.CommonExtensions | blackfriday.NoEmptyLineBeforeBlock | blackfriday.FencedCode | blackfriday.Tables | blackfriday.AutoHeadingIDs
	renderer := blackfriday.NewHTMLRenderer(blackfriday.HTMLRendererParameters{
		Flags: blackfriday.CommonHTMLFlags,
	})
	html := blackfriday.Run(data, blackfriday.WithExtensions(extensions), blackfriday.WithRenderer(renderer))
	return html
}

func loadContentFile(path string) (*Content, error) {
	data, err := loadFile(path)
	if err != nil {
		return nil, err
	}
	meta, markdown := splitContent(data)
	html := parseMarkdown(markdown)
	content := &Content{
		Metadata: meta,
		HTML:     html,
	}
	return content, nil
}

func LoadAllContent(dirPath string) ([]*Content, error) {
	var contents []*Content
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !strings.HasSuffix(info.Name(), ".md") {
			return nil
		}
		content, err := loadContentFile(path)
		if err != nil {
			return err
		}
		contents = append(contents, content)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func LoadContent(path string) (*Content, error) {
	content, err := loadContentFile(path)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func splitContent(data []byte) (map[string]interface{}, []byte) {
	parts := bytes.SplitN(data, []byte("\n\n"), 2)
	meta := make(map[string]interface{})
	err := yaml.Unmarshal(parts[0], &meta)
	if err != nil {
		panic(err)
	}
	return meta, parts[1]
}
