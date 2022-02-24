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
	"testing"
)

var exampleConfig = map[string]string{
	"excel.filepath": "test.xlsx",
	"excel.sheet":    "Sheet1",
}

func configWith(pairs ...string) map[string]string {
	cfg := make(map[string]string)

	for key, value := range exampleConfig {
		cfg[key] = value
	}

	for i := 0; i < len(pairs); i += 2 {
		key := pairs[i]
		value := pairs[i+1]
		cfg[key] = value
	}

	return cfg
}

func configWithout(keys ...string) map[string]string {
	cfg := make(map[string]string)

	for key, value := range exampleConfig {
		cfg[key] = value
	}

	for _, key := range keys {
		delete(cfg, key)
	}

	return cfg
}

func TestFilepath(t *testing.T) {
	t.Run("Successful", func(t *testing.T) {
		c, err := Parse(configWith("excel.filepath", "some-value"))

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if c.Filepath != "some-value" {
			t.Fatalf("expected Filepath to be %q, got %q", "some-value", c.Filepath)
		}
	})

	t.Run("Missing", func(t *testing.T) {
		_, err := Parse(configWithout("excel.filepath"))

		if err == nil {
			t.Fatal("expected error, got nothing")
		}

		expectedErrMsg := `"excel.filepath" config value must be set`

		if err.Error() != expectedErrMsg {
			t.Fatalf("expected error msg to be %q, got %q", expectedErrMsg, err.Error())
		}
	})
}

func TestSheet(t *testing.T) {
	t.Run("Successful", func(t *testing.T) {
		c, err := Parse(configWith("excel.sheet", "some-value"))

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if c.Sheet != "some-value" {
			t.Fatalf("expected Sheet to be %q, got %q", "some-value", c.Sheet)
		}
	})

	t.Run("Missing", func(t *testing.T) {
		_, err := Parse(configWithout("excel.sheet"))

		if err == nil {
			t.Fatal("expected error, got nothing")
		}

		expectedErrMsg := `"excel.sheet" config value must be set`

		if err.Error() != expectedErrMsg {
			t.Fatalf("expected error msg to be %q, got %q", expectedErrMsg, err.Error())
		}
	})
}
