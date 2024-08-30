/*
 *
 * Copyright 2024 Semih Tok
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and imitations under the License.
 *
 */

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run(args []string) {
	if len(args) == 0 {
		dir, err := os.ReadDir(".")
		if err != nil {
			panic(err)
		}
		DrawTable(dir, "", 0, 0)
	} else {
		if args[0] == "-t" {
			fileType := "." + strings.Join(args[1:], " ")
			dir, err := os.ReadDir(".")
			if err != nil {
				panic(err)
			}
			DrawTable(dir, fileType, 0, 0)
		}

		if args[0] == "-h" || args[0] == "--help" {
			helpText := `
Usage:	
 lsm [command]

Available Commands:
 lsm                                  : list of files in current directory
 lsm -t {file extension}              : list of files with specific extension in current directory
 lsm -gt {size in MB}                 : list of files greater than size in current directory
 lsm -lt {size in MB}                 : list of files lower than size in current directory
`
			fmt.Println(helpText)
		}

		if args[0] == "-gt" {
			upperLimitArg := strings.Join(args[1:], " ")
			upperLimit, err := strconv.Atoi(upperLimitArg)
			if err != nil {
				panic(err)
			}

			dir, err := os.ReadDir(".")
			if err != nil {
				panic(err)
			}
			DrawTable(dir, "", 0, upperLimit)
		}

		if args[0] == "-lt" {
			lowerLimitArg := strings.Join(args[1:], " ")
			lowerLimit, err := strconv.Atoi(lowerLimitArg)
			if err != nil {
				panic(err)
			}

			dir, err := os.ReadDir(".")
			if err != nil {
				panic(err)
			}
			DrawTable(dir, "", lowerLimit, 0)
		}
	}
}
