package controlletdatasource

import (
	"context"

	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-provider-aviatrix-cloud-poc/client"
)

// @TODO : This  data source needs update schema properly.
// Will revisit
type ControllerDataSource struct {
	client *client.Client
}

type ControllerModel struct {
	Region     types.String `tfsdk:"region"`
	VpcId      types.String `tfsdk:"vpc_id"`
	InstanceId types.String `tfsdk:"instance_id"`
	CloudType  types.String `tfsdk:"cloud_type"`
}
type ControllerListModel struct {
	ControllerListModel []*ControllerModel `tfsdk:"controller"`
}

type ControllerJsonModel struct {
	Region     string `json:"region"`
	VpcId      string `json:"vpc_id"`
	InstanceId string `json:"instance_id"`
	CloudType  string `json:"cloud_type"`
}
type ControllerListJsonModel struct {
	ControllerListJsonModel []*ControllerJsonModel `json:"controller"`
}

func NewControllerDataSource() datasource.DataSource {
	return &ControllerDataSource{}
}

// Metadata returns the data source type name.
func (d *ControllerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_controller"
}

// Schema defines the schema for the data source.
func (d *ControllerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"controller": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"region": schema.StringAttribute{
							Computed: true,
						},
						"vpc_id": schema.StringAttribute{
							Computed: true,
						},
						"instance_id": schema.StringAttribute{
							Computed: true,
						},
						"cloud_type": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *ControllerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *ControllerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	//state := &ControllerListModel{}
	output := &ControllerListJsonModel{}
	httpResponse, err := d.client.Get(ctx, "/get-controller-data", "", output)
	if err != nil {
		errMsg := fmt.Sprintf("got an errorin controller datasource. Error : %v , response code : %v", err.Error(), httpResponse)
		tflog.Error(ctx, errMsg)

		resp.Diagnostics.AddAttributeError(
			path.Root("controller"),
			errMsg,
			"The provider cannot get the controller data",
		)
		return
	}
	

	state := &ControllerListModel{
		ControllerListModel: []*ControllerModel{},
	}

	for _, d := range output.ControllerListJsonModel {
		obj := &ControllerModel{
			Region:     types.StringValue(d.Region),
			CloudType:  types.StringValue(d.CloudType),
			InstanceId: types.StringValue(d.InstanceId),
			VpcId:      types.StringValue(d.VpcId),
		}
		state.ControllerListModel = append(state.ControllerListModel, obj)
	}

	// Set state
	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

}
