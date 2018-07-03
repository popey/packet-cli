// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// operatingSystemCmd represents the operatingSystem command
var retrieveOperatingSystemCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieves a list of available operating systems.",
	Long: `Example:
  packet operating-systems get`,
	Run: func(cmd *cobra.Command, args []string) {
		oss, _, err := PacknGo.OperatingSystems.List()
		if err != nil {
			fmt.Println("Client error:", err)
			return
		}

		data := make([][]string, len(oss))

		for i, os := range oss {
			data[i] = []string{os.Name, os.Slug, os.Distro, os.Version}
		}
		header := []string{"Name", "Slug", "Distro", "Version"}

		output(oss, header, &data)
	},
}

func init() {
	retrieveOperatingSystemCmd.PersistentFlags().BoolVarP(&isJSON, "json", "j", false, "JSON output")
	retrieveOperatingSystemCmd.PersistentFlags().BoolVarP(&isYaml, "yaml", "y", false, "YAML output")
}
