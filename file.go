package xcel

import (
	"encoding/json"

	"github.com/360EntSecGroup-Skylar/excelize"
)

// File represents an xlsx file
type File struct {
	*excelize.File
	Name   string
	Sheets []*Sheet
}

// NewFile creates a new empty file, with no sheets
func NewFile(name string) *File {
	f := File{
		excelize.NewFile(),
		name,
		[]*Sheet{},
	}

	return &f
}

// AddSheet adds a sheet with the given name to the file
func (f *File) AddSheet(name string) *Sheet {
	if len(f.Sheets) == 0 && f.SheetCount == 1 {
		f.SetSheetName(f.GetSheetName(1), name)
	} else {
		f.NewSheet(name)
	}
	sheet := Sheet{
		Name: name,
		File: f,
	}

	f.Sheets = append(f.Sheets, &sheet)
	return &sheet
}

// NewStyle creates a style in the file and assigns it an id
func (f *File) NewStyle(style Style) (int, error) {
	bts, err := json.Marshal(style)
	if err != nil {
		return 0, err
	}

	return f.File.NewStyle(string(bts))
}
