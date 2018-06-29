// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create project from assets",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		assets := parseStringFlag(cmd, "assets")
		pname := parseStringFlag(cmd, "name")
		if err := os.Mkdir(pname, 0755); err != nil && !os.IsExist(err) {
			log.Fatal(fmt.Sprintf("fail to make directory %s", pname), err)
		}
		box := packr.NewBox("../assets/" + assets)
		list := box.List()
		for _, fname := range list {
			b := box.Bytes(fname)
			writeFile(pname, fname, b)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("assets", "a", "", "Which assets to use")
	createCmd.Flags().StringP("name", "n", "", "Project name")
}
