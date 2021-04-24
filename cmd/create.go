package cmd

import (
	"os"
	"path/filepath"
)

func Create(n string, c string) {
	err := os.Mkdir(n, 0o755)
	Check(err)

	f, err := os.Create(filepath.Join(n, n+".vue"))
	Check(err)

	defer f.Close()
	_, err = f.WriteString(c)
	Check(err)

	s, err := os.Create(filepath.Join(n, n+".scss"))
	Check(err)
	defer s.Close()

	_, err = s.WriteString("." + n + " {}")
	Check(err)
}
