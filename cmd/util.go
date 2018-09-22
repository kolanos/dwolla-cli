package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"text/tabwriter"

	"github.com/kolanos/dwolla-v2-go"
	"github.com/olekukonko/tablewriter"
)

func renderCollection(data [][]string, header []string, footer []string) {
	table := tablewriter.NewWriter(os.Stdout)
	if len(header) > 0 {
		table.SetHeader(header)
	}
	if len(footer) > 0 {
		table.SetFooter(footer)
	}
	table.SetBorder(false)
	table.AppendBulk(data)

	fmt.Println("")
	table.Render()
	fmt.Println("")
}

func renderError(err error) {
	fmt.Println(err)

	validationErr, ok := err.(dwolla.ValidationError)
	if ok {
		data := make([][]string, len(validationErr.Embedded["errors"]))

		for i, e := range validationErr.Embedded["errors"] {
			data[i] = []string{e.Path, e.Message}
		}

		header := []string{"Field", "message"}
		footer := []string{"Total", strconv.Itoa(len(validationErr.Embedded["errors"]))}

		renderCollection(data, header, footer)
	}

	os.Exit(1)
}

func renderResource(res interface{}) {
	s := reflect.ValueOf(res).Elem()
	t := s.Type()
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.AlignRight|tabwriter.Debug)

	for i := 0; i < s.NumField(); i++ {
		if t.Field(i).Name == "Resource" {
			continue
		}

		f := s.Field(i)
		fmt.Fprintln(w, fmt.Sprintf("%s \t %v", t.Field(i).Name, f.Interface()))
	}

	w.Flush()
	fmt.Println("")
}
