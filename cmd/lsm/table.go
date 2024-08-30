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
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"os"
	"path/filepath"
	"sync"
)

func DrawTable(dir []os.DirEntry, fileType string, lowerLimit int, upperLimit int) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"#", "Name", "Modified At", "Size"})
	var sum int64
	var index = 1

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, entry := range dir {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		wg.Add(1)
		go func(entry os.DirEntry) {
			defer wg.Done()

			// Calculate the size of the entry, including directory contents
			entrySize := CalculateSize(entry)

			// Filter by file type, if specified
			if len(fileType) > 0 {
				fileExtension := filepath.Ext(info.Name())
				if fileExtension != fileType {
					return
				}
			}

			// Apply size limits
			if upperLimit > 0 && int64(upperLimit*1024*1024) > entrySize {
				return
			}

			if lowerLimit > 0 && int64(lowerLimit*1024*1024) < entrySize {
				return
			}

			formattedTime := info.ModTime().Format("2006-01-02 15:04:05")

			mu.Lock()
			t.AppendRows([]table.Row{
				{index, entry.Name(), formattedTime, FormatSize(entrySize)},
			})
			index++
			sum += entrySize
			mu.Unlock()
		}(entry)
	}

	wg.Wait()

	t.SetStyle(table.StyleColoredBlackOnBlueWhite)
	t.Style().Color = table.ColorOptions{
		IndexColumn:  text.Colors{text.BgHiBlue, text.FgHiWhite, text.Bold},
		Header:       text.Colors{text.BgHiBlue, text.FgHiWhite, text.Bold},
		Footer:       text.Colors{text.BgHiBlue, text.FgHiWhite, text.Bold},
		Row:          text.Colors{text.BgHiWhite, text.FgBlack},
		RowAlternate: text.Colors{text.BgHiWhite, text.FgBlack},
	}

	t.Style().Format.Footer = text.FormatLower
	t.Style().Options.DrawBorder = false

	t.AppendFooter(table.Row{"", "", "Total", FormatSize(sum)})
	t.Render()
}
