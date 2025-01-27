package controllerdatasource

import (
	"context"

	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-provider-aviatrix-cloud-poc/client"
	pb "github.com/terraform-provider-aviatrix-cloud-poc/gen-protogo/aviatrix"
)

// @TODO : This  data source needs update schema properly.
// Will revisit
type ControllerDataSource struct {
	client client.HttpHandler
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

	client, ok := req.ProviderData.(client.HttpHandler)
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
	// Connect to the gRPC server
	conn, err := client.NewGRPCClient(ctx)
	if err != nil {
		tflog.Error(ctx, fmt.Sprintf("ControllerDataSource : Failed to connect gRPC: %v", err))
	}
	defer conn.ClientConn.Close()

	client := pb.NewAviatrixControllerserviceClient(conn.ClientConn)

	apiResp, err := client.GetAviatrixControllers(ctx, &pb.GetAviatrixControllersRequest{})
	if err != nil {
		//log.Fatalf("Error calling GetAviatrixControllers: %v", err)
		tflog.Error(ctx, fmt.Sprintf("Error calling GetAviatrixControllers: %v", err))
	}

	//tflog.Error(ctx, fmt.Sprintf("=============> gRPCCCCCCCCCCCCCCC :  AviatrixControllers: %s", apiResp.AviatrixControllers))

	state := &ControllerListModel{
		ControllerListModel: []*ControllerModel{},
	}

	for _, d := range apiResp.AviatrixControllers {
		obj := &ControllerModel{
			Region:     types.StringValue(d.Region),
			CloudType:  types.StringValue(d.CloudType),
			InstanceId: types.StringValue(d.InstanceId),
			VpcId:      types.StringValue(d.VpcId),
		}
		state.ControllerListModel = append(state.ControllerListModel, obj)
	}

	// Set state
	if len(state.ControllerListModel) > 0 {
		diags := resp.State.Set(ctx, state)
		resp.Diagnostics.Append(diags...)
	}

	if resp.Diagnostics.HasError() {
		return
	}

}
