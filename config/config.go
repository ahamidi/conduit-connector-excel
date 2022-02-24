// Copyright © 2022 Meroxa, Inc.
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

package config

import (
	"fmt"
)

const (
	ExcelFilepath = "filepath"
	ExcelSheet    = "sheet"
)

// Config represents configuration needed for an Excel File
type Config struct {
	Filepath string
	Sheet    string
}

// Parse attempts to parse plugins.Config into a Config struct
func Parse(cfg map[string]string) (Config, error) {
	filepath, ok := cfg[ExcelFilepath]

	if !ok {
		return Config{}, requiredConfigErr(ExcelFilepath)
	}

	sheet, ok := cfg[ExcelSheet]

	if !ok {
		return Config{}, requiredConfigErr(ExcelSheet)
	}

	config := Config{
		Filepath: filepath,
		Sheet:    sheet,
	}

	return config, nil
}

func requiredConfigErr(name string) error {
	return fmt.Errorf("%q config value must be set", name)
}
