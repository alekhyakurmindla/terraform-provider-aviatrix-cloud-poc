package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	v1 := "1.0"
	provider := New(v1)
	assert.NotNil(t, provider)
}

func TestMetadata(t *testing.T) {
    v1 := "1.0"
	p := New(v1)()

	ctx := context.Background()
	req := provider.MetadataRequest{}
	resp := &provider.MetadataResponse{}
	p.Metadata(ctx, req, resp)
	assert.Equal(t, "aviatrix-cloud-poc", resp.TypeName)
	assert.Equal(t, "1.0", resp.Version)
}

func TestSchema(t *testing.T) {
	v1 := "1.0"
	p := New(v1)()
	ctx := context.Background()
	req := provider.SchemaRequest{}
	resp := &provider.SchemaResponse{}
	p.Schema(ctx, req, resp)
    assert.NotNil(t, resp.Schema.Attributes)
    
        attr, exists := resp.Schema.Attributes["host"]
        assert.True(t, exists, "Key 'host' should exist in the map")
        assert.Equal(t, true, attr.IsOptional(), "Value for key 'resp.Schema.Attributes.IsOptional' should be true")
    
        attr, exists = resp.Schema.Attributes["username"]
        assert.True(t, exists, "Key 'username' should exist in the map")
        assert.Equal(t, true, attr.IsOptional(), "Value for username 'resp.Schema.Attributes.IsOptional' should be true")

        attr, exists = resp.Schema.Attributes["password"]
        assert.True(t, exists, "Key 'password' should exist in the map")	
        assert.Equal(t, true, attr.IsOptional(), "Value for password 'resp.Schema.Attributes.IsOptional' should be true") 
        assert.Equal(t, true, attr.IsSensitive(), "Value for password 'resp.Schema.Attributes.IsSensitive' should be true")    
}

func TestConfigure(t *testing.T){
	avxPro := &aviatrixProvider{}
	ctx := context.Background()
	req := provider.ConfigureRequest{}
	resp := provider.ConfigureResponse{}
	avxPro.Configure(ctx, req, &resp)
	assert.False(t, resp.Diagnostics.HasError())
}

func TestDataSources(t *testing.T){
	p := &aviatrixProvider{}
	ctx := context.Background()
	datasourceArray := p.DataSources(ctx)
	assert.Equal(t, true, len(datasourceArray) > 0)
}

func TestResource(t *testing.T) {
	p := &aviatrixProvider{}
	ctx := context.Background()
	resourceArray := p.Resources(ctx)
	assert.Equal(t, true, len(resourceArray) > 0)
}