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

package source

import (
	"context"
	"encoding/json"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"

	sdk "github.com/conduitio/connector-plugin-sdk"
)

// Source connector
type Source struct {
	sdk.UnimplementedSource
	config   Config
	file     *excelize.File
	iterator *excelize.Rows
	p        int
}

func NewSource() sdk.Source {
	return &Source{}
}

// Configure parses and stores the configurations
// returns an error in case of invalid config
func (s *Source) Configure(ctx context.Context, cfg map[string]string) error {
	log.Println("plugin configure")
	config2, err := Parse(cfg)
	if err != nil {
		log.Println("configure error")
		return err
	}

	s.config = config2

	return nil
}

// Open prepare the plugin to start sending records from the given position
func (s *Source) Open(ctx context.Context, rp sdk.Position) error {
	sdk.Logger(ctx).Debug().Msg("open")
	f, err := excelize.OpenFile(s.config.Filepath)
	if err != nil {
		log.Println("open error")
		return err
	}

	s.file = f
	if rp != nil {
		pos, err := strconv.Atoi(string(rp))
		if err != nil {
			log.Println("open error")
			return err
		}
		s.p = pos
	}

	// TODO: handle case where sheet is not provided
	i, err := f.Rows(s.config.Sheet)
	if err != nil {
		log.Println("open error")
		return err
	}

	s.iterator = i
	return nil
}

// Read gets the next row from the Excel file
func (s *Source) Read(ctx context.Context) (sdk.Record, error) {
	sdk.Logger(ctx).Debug().Msg("read")
	if !s.iterator.Next() {
		return sdk.Record{}, sdk.ErrBackoffRetry
	}
	row, err := s.iterator.Columns()
	if err != nil {
		return sdk.Record{}, sdk.ErrBackoffRetry
	}
	recordRow := make([]interface{}, len(row))
	for k, v := range row {
		recordRow[k] = v
	}

	jsonPayload, err := json.Marshal(recordRow)
	if err != nil {
		return sdk.Record{}, err
	}
	pos := s.iterator.CurrentRow()
	key := strconv.Itoa(int(pos))
	return sdk.Record{
		Key:      sdk.RawData(key),
		Position: sdk.Position(key),
		Payload:  sdk.RawData(jsonPayload),
	}, nil
}

func (s *Source) Teardown(ctx context.Context) error {
	var err error
	if s.iterator != nil {
		err = s.iterator.Close()
		if err != nil {
			return err
		}
	}
	if s.file != nil {
		err = s.file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Source) Ack(ctx context.Context, position sdk.Position) error {
	sdk.Logger(ctx).Debug().Str("position", string(position)).Msg("got ack")
	return nil // no ack needed
}
