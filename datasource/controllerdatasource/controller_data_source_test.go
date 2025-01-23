package controllerdatasource

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/stretchr/testify/assert"
	"github.com/terraform-provider-aviatrix-cloud-poc/mocks"
)

func TestMetaData(t *testing.T) {
	cnt := &ControllerDataSource{}
	ctx := context.Background()
	req := datasource.MetadataRequest{
		ProviderTypeName: "test-provider-name",
	}
	resp := datasource.MetadataResponse{}
	cnt.Metadata(ctx, req, &resp) 
	assert.Equal(t, "test-provider-name_controller", resp.TypeName)
	
}

func TestSchema(t *testing.T) {
	cnt := &ControllerDataSource{}
	ctx := context.Background()
	req := datasource.SchemaRequest{}
	resp := datasource.SchemaResponse{}
	cnt.Schema(ctx, req, &resp)
	assert.NotNil(t, resp.Schema.Attributes)

	controllerAttr := resp.Schema.Attributes["controller"]
	assert.NotNil(t, controllerAttr)
	assert.Equal(t, true, controllerAttr.IsComputed(), "Value for region 'resp.Schema.Attributes.Computed' should be true")

}

func TestRead(t *testing.T) {

	// Create the mock object
	mockHttpHandler := new(mocks.HttpHandler)

	cnt := &ControllerDataSource{
		client: mockHttpHandler,
	}
	ctx := context.Background()
	req := datasource.ReadRequest{}
	resp := datasource.ReadResponse{}

	apiOutput := &ControllerListJsonModel{}

	mockHttpHandler.On("Get", ctx, "/get-controller-data", "", apiOutput).Return(200, nil)
	cnt.Read(ctx, req, &resp)
}
