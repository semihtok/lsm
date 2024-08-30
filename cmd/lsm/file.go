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
	"path/filepath"
)

func FormatSize(size int64) string {
	const unit = 1024
	if size < unit {
		return fmt.Sprintf("%d B", size)
	}
	div, exp := int64(unit), 0
	for n := size / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(size)/float64(div), "KMGTPE"[exp])
}

// CalculateSize concurrently calculates the size of a directory or a file.
func CalculateSize(entry os.DirEntry) int64 {
	if entry.IsDir() {
		var dirSize int64
		err := filepath.WalkDir(filepath.Join(".", entry.Name()), func(_ string, d os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			info, err := d.Info()
			if err != nil {
				return err
			}
			dirSize += info.Size()
			return nil
		})
		if err != nil {
			return 0
		}
		return dirSize
	}
	info, err := entry.Info()
	if err != nil {
		return 0
	}
	return info.Size()
}
