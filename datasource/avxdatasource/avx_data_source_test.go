package avxdatasource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stretchr/testify/assert"
)

func TestMetadata(t *testing.T){
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

func TestSchema(t *testing.T){
	cnt := &AvxDatasource{}
	ctx := context.Background()
	req := datasource.SchemaRequest{}
	resp := datasource.SchemaResponse{}
	cnt.Schema(ctx, req, &resp)
	assert.NotNil(t, resp.Schema.Attributes)
	
}

