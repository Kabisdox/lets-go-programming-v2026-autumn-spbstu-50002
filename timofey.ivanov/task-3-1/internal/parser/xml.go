package parser

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html/charset"

	"github.com/Kabisdox/task-3/internal/domain"
)

type XMLParser struct{}

func (XMLParser) Parse(path string) ([]domain.Currency, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open xml: %w", err)
	}

	defer func() {
		_ = f.Close()
	}()

	dec := xml.NewDecoder(f)
	dec.CharsetReader = charset.NewReaderLabel

	var doc domain.ValCursXML

	if err := dec.Decode(&doc); err != nil {
		return nil, fmt.Errorf("decode xml: %w", err)
	}

	result := make([]domain.Currency, 0, len(doc.Valutes))

	for _, v := range doc.Valutes {
		normalized := strings.ReplaceAll(v.Value, ",", ".")

		parsed, err := strconv.ParseFloat(normalized, 64)
		if err != nil {
			return nil, fmt.Errorf("parse value %q: %w", v.Value, err)
		}

		result = append(result, domain.Currency{
			NumCode:  v.NumCode,
			CharCode: v.CharCode,
			Value:    parsed,
		})
	}

	return result, nil
}
