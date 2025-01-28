package avxdatasource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stretchr/testify/assert"
)

func TestMetadata(t *testing.T) {
	cnt := &AvxDatasource{}
	ctx := context.Background()
	req := datasource.MetadataRequest{
		ProviderTypeName: "test-provider-name",
	}
	resp := datasource.MetadataResponse{}
	cnt.Metadata(ctx, req, &resp)

	expectedName := "test-provider-name_avx"
	assert.Equal(t, expectedName, resp.TypeName)
}

func TestSchema(t *testing.T) {
	cnt := &AvxDatasource{}
	ctx := context.Background()
	req := datasource.SchemaRequest{}
	resp := datasource.SchemaResponse{}
	cnt.Schema(ctx, req, &resp)
	assert.NotNil(t, resp.Schema.Attributes)

	avxdataAttr := resp.Schema.Attributes["avxdata"]
	assert.NotNil(t, avxdataAttr)

	assert.Equal(t, true, avxdataAttr.IsComputed(), "Value for id 'resp.Schema.Attributes.Computed' should be true")
	assert.Equal(t, true, avxdataAttr.IsComputed(), "Value for name 'resp.Schema.Attributes.Computed' should be true")
	assert.Equal(t, true, avxdataAttr.IsComputed(), "Value for description 'resp.Schema.Attributes.Computed' should be true")
}

func TestConfigure(t *testing.T) {
	ds := &AvxDatasource{}
	ctx := context.Background()
	req := datasource.ConfigureRequest{}
	resp := datasource.ConfigureResponse{}
	ds.Configure(ctx, req, &resp)

	assert.False(t, resp.Diagnostics.HasError())
}

func TestRead(t *testing.T) {

	avx := &AvxDatasource{}
	ctx := context.Background()
	req := datasource.ReadRequest{}
	resp := datasource.ReadResponse{}

	resp.Diagnostics.HasError()
	avx.Read(ctx, req, &resp)
}
