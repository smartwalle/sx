package sx

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type FileStock struct {
	words []string
}

func NewFileStock(p string) (*FileStock, error) {
	var s = &FileStock{}

	var rFile, err = os.OpenFile(p, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer rFile.Close()
	var reader = bufio.NewReader(rFile)

	var line []byte

	for {
		if line, _, err = reader.ReadLine(); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		var sLine = strings.TrimSpace(string(line))

		if sLine == "" {
			continue
		}

		s.words = append(s.words, sLine)
	}

	return s, nil
}

func (this *FileStock) ReadAll() []string {
	return this.words
}
