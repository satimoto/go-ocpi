package deise

import (
	"encoding/csv"
	"io"

	"github.com/gocarina/gocsv"
)

func UnmarshalDto(body io.ReadCloser) ([]*DeIseTariffDto, error) {
	response := []*DeIseTariffDto{}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r
	})

	if err := gocsv.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return response, nil
}
