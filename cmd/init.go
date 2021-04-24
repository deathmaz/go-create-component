package cmd

import (
	"flag"
)

func Init() {
	kind := flag.String("kind", "", "Kind of component")
	ts := flag.Bool("ts", true, "Should work with typesript?")

	flag.Parse()

	var name string
	if len(flag.Args()) > 0 {
		name = flag.Args()[0]
	} else {
		panic("Please provide component name!")
	}

	vue := NewVue(Options{
		Name:       name,
		Typescript: *ts,
		Kind:       *kind,
	})
	Create(name, vue.GetComponent())
}
