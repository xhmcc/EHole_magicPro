/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"ehole/module/finger/source"
	"ehole/module/fofaext"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// quakeextCmd represents the quakeext command
var quakeextCmd = &cobra.Command{
	Use:   "quakeext",
	Short: "ehole的quake提取模块",
	Long:  `从quake api提取资产并保存成xlsx，支持大批量ip提取,支持quake所有语法。`,
	Run: func(cmd *cobra.Command, args []string) {
		file := strings.Split(ext_quake_output, ".")

		if len(file) == 2 {
			if file[1] == "xlsx" {
				if ext_quakeip != "" {
					results := source.Quakeips_out(ext_quakeip)
					fofaext.Fofaext(results, ext_quake_output)
					os.Exit(1)
				}
				if ext_quakesearche != "" {
					fmt.Println(ext_quakesearche)
					results := source.Quakeall_out(ext_quakesearche)
					fofaext.Fofaext(results, ext_quake_output)
					os.Exit(1)
				}
			} else {
				log.Println("文件名错误！！！")
			}
		} else {
			log.Println("文件名错误！！！")
		}
	},
}

var (
	ext_quakeip      string
	ext_quakesearche string
	ext_quake_output string
)

func init() {
	rootCmd.AddCommand(quakeextCmd)
	quakeextCmd.Flags().StringVarP(&ext_quakeip, "ipfile", "l", "", "从文本获取IP，在quake搜索，支持大量ip，默认保存所有结果。")
	quakeextCmd.Flags().StringVarP(&ext_quakesearche, "quake", "s", "", "从quake提取资产，支持quake所有语法，默认保存所有结果。")
	quakeextCmd.Flags().StringVarP(&ext_quake_output, "output", "o", "results.xlsx", "指定输出文件名和位置，当前仅支持xlsx后缀的文件。")
} 