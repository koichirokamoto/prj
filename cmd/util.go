package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func parseStringFlag(cmd *cobra.Command, name string) string {
	s, err := cmd.Flags().GetString(name)
	if err != nil {
		log.Fatal(fmt.Sprintf("fail to get %s from flag", name), err)
	}
	if s == "" {
		cmd.Help()
		log.Fatal(fmt.Sprintf("%s is required", name))
	}
	return s
}

func writeFile(pname, fname string, b []byte) {
	dst := filepath.Join(pname, fname)
	dir := filepath.Dir(dst)
	if err := os.MkdirAll(dir, 0755); err != nil && !os.IsExist(err) {
		log.Fatal(fmt.Sprintf("fail to make directory %s", dir), err)
	}
	if err := ioutil.WriteFile(dst, b, 0755); err != nil {
		log.Fatal("write asset data to file", err)
	}
}
