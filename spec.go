package excel

import (
	"github.com/ahamidi/conduit-excel-plugin/config"
	"github.com/ahamidi/conduit-excel-plugin/source"
	sdk "github.com/conduitio/connector-plugin-sdk"
)

type Spec struct{}

// Specification returns the Plugin's Specification.
func Specification() sdk.Specification {
	return sdk.Specification{
		Name:    "excel",
		Summary: "A Microsoft Excel 2007 and later plugin",
		Version: "v0.0.1",
		Author:  "Ali Hamidi",
		DestinationParams: map[string]sdk.Parameter{
			config.ExcelFilepath: {
				Default:     "",
				Required:    true,
				Description: "Path to the Excel file.",
			},
			config.ExcelSheet: {
				Default:     "",
				Required:    false,
				Description: "Worksheet to be used. If none is provided, the first sheet will be used.",
			},
		},
		SourceParams: map[string]sdk.Parameter{
			config.ExcelFilepath: {
				Default:     "",
				Required:    true,
				Description: "Path to the Excel file.",
			},
			config.ExcelSheet: {
				Default:     "",
				Required:    false,
				Description: "Worksheet to be used. If none is provided, the first sheet will be used.",
			},
			source.DefaultPollingPeriod: {
				Default:     "1s",
				Required:    false,
				Description: "Polling interval for checking to see if there are new rows.",
			},
		},
	}
}
