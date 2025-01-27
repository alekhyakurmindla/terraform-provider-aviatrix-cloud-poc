package aviatrixresource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/stretchr/testify/assert"
)

func TestMetadata(t *testing.T) {
	avxRc := AviatrixAccountResource{}
	ctx := context.Background()
	req := resource.MetadataRequest{
		ProviderTypeName: "test-provider-name",
	}
	resp := resource.MetadataResponse{}
	avxRc.Metadata(ctx, req, &resp)
	assert.Equal(t, "test-provider-name_account_user", resp.TypeName)
}

func TestSchema(t *testing.T) {
	avxRc := AviatrixAccountResource{}
	ctx := context.Background()
	req := resource.SchemaRequest{}
	resp := resource.SchemaResponse{}

	avxRc.Schema(ctx, req, &resp)
	assert.NotNil(t, resp.Schema.Attributes)

	emailAttr, exists := resp.Schema.Attributes["email"]
	assert.True(t, exists, "Key 'email' should exist in the map")
	assert.False(t, emailAttr.IsComputed(), "isComputed should be false")
	assert.True(t, true, emailAttr.IsRequired(), "isComputed should be true")

	usernameAttr, exists := resp.Schema.Attributes["username"]
	assert.True(t, exists, "Key 'attrUsername' should exist in the map")
	assert.False(t, usernameAttr.IsComputed(), "isComputed should be false")
	assert.True(t, true, usernameAttr.IsRequired(), "isComputed should be true")

	passwordAttr, exists := resp.Schema.Attributes["password"]
	assert.True(t, exists, "Key 'attrUsername' should exist in the map")
	assert.False(t, passwordAttr.IsComputed(), "isComputed should be false")
	assert.True(t, true, passwordAttr.IsRequired(), "isComputed should be true")
}

func TestConfigure(t *testing.T) {
	avxRc := AviatrixAccountResource{}
	ctx := context.Background()
	req := resource.ConfigureRequest{}
	resp := resource.ConfigureResponse{}

	avxRc.Configure(ctx, req, &resp)
	assert.False(t, resp.Diagnostics.HasError())
}

func TestRead(t *testing.T) {
	// Create the mock object
	//mockHttpHandler := new(mocks.HttpHandler)

	avxRc := &AviatrixAccountResource{}
	ctx := context.Background()
	req := resource.ReadRequest{}
	resp := resource.ReadResponse{}

	
	avxRc.Read(ctx, req, &resp)
	assert.False(t, resp.Diagnostics.HasError())
}


