package cmd

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"
)

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
