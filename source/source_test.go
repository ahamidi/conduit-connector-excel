package source

import (
	"context"
	"github.com/ahamidi/conduit-excel-plugin/config"
	sdk "github.com/conduitio/connector-plugin-sdk"
	"log"
	"testing"
)

func TestSource_ReadOne(t *testing.T) {
	cfg := map[string]string{
		config.ExcelFilepath:   "test/test.xlsx",
		config.ExcelSheet:      "Sheet1",
		ConfigKeyPollingPeriod: "100ms",
	}

	ctx := context.Background()
	source := &Source{}
	err := source.Configure(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	err = source.Open(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	r, err := source.Read(ctx)
	if err != nil && err.Error() != sdk.ErrBackoffRetry.Error() {
		t.Fatalf("expected a BackoffRetry error, got: %v", err)
	}

	log.Printf("r Key: %+v, Position: %+v, Payload: %+v", string(r.Key.Bytes()), string(r.Position), string(r.Payload.Bytes()))

	err = source.Teardown(ctx)
	if err != nil {
		t.Fatalf("expected a no error, got: %v", err)
	}
}
