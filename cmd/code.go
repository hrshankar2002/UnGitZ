/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sync"
	"ungitz/util"

	"github.com/spf13/cobra"
)

var File_flag []string
var link_flag []string
var File string
var wg = sync.WaitGroup{}
var pattern = `-(.+)`

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "It will open the directory in Visual Studio Code",
	Long: `This command will help to open the unzipped folder to Visual Studio Code.
	In order for this command to work, Visual Studio code should be installed in your system`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(File_flag) < 1 && len(link_flag) < 1 && len(args) < 1 {
			return errors.New("accept(s) 1 argument")
		}
		return nil
	},
	Example: `ungitz code -f <filename>,<branch name>
ungitz code -l <URL>,<filename>,<branch name>`,
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		var fileName string
		var err error
		var argument string
		var link_arg string
		var name_arg string
		var branch_arg string

		// flag check
		if len(File_flag) != 0 {
			argument = File_flag[0]
			branch_arg = File_flag[1]

		} else if len(link_flag) != 0 {
			link_arg = link_flag[0]
			name_arg = link_flag[1]
			branch_arg = link_flag[2]

			// wait period implementation using go-routines for download function
			wg.Add(1)
			go func(name_arg, link_arg string) {
				util.Download(name_arg, link_arg)
				wg.Done()
			}(name_arg, link_arg)
			wg.Wait()
			argument = name_arg
		}

		// file exist check
		FileExists, err := util.FileExists(argument)
		if err != nil {
			fmt.Println(err.Error())
		}
		if FileExists {
			fileName, err = filepath.Abs(argument)
			if err != nil {
				fmt.Println(err.Error())
			}
		} else {
			fmt.Println("File %v does not exist", argument)
			return
		}

		// initialisation of working directory
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		util.Unzip(fileName, wd)

		var testname = util.FilenameWithoutExtension(fileName)

		// regex filter
		re := regexp.MustCompile(pattern)
		submatches := re.FindStringSubmatch(testname)
		if len(submatches) > 1 && submatches[1] != branch_arg {
			var newtestname = testname + "-" + branch_arg
			os.Chdir(newtestname)
		} else {

			os.Chdir(testname)
		}

		// updation of working directory
		wd, err = os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		}

		commandCode := exec.Command("code", wd)
		err = commandCode.Run()

		if err != nil {
			fmt.Println("Visual Studio Code executable not found in %PATH%")
		} else {
			fmt.Println("Unzipping and opening file.")
		}

	},
}

func init() {
	rootCmd.AddCommand(codeCmd)
	codeCmd.PersistentFlags().StringSliceVarP(&File_flag, "file", "f", []string{}, "Arguments:<filename>,<URL>")
	codeCmd.PersistentFlags().StringSliceVarP(&link_flag, "link", "l", []string{}, "Arguments:<URL>,<filename>,<git branch>")
}
