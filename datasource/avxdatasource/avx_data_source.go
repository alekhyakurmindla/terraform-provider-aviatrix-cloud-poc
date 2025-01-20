package avxdatasource

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-provider-aviatrix-cloud-poc/client"
)

//@TODO : This is dummy data source. Will remove after testing
type AvxDatasource struct {
	client *client.Client
}

// avxDataSourceModel maps the data source schema data.
type AvxDataSourceModel struct {
	AvxData []*AvxModel `tfsdk:"avxdata"`
}

// avxModel maps avx schema data.
type AvxModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
}

func NewAviatrixDataSource() datasource.DataSource {
	return &AvxDatasource{}
}

func (d *AvxDatasource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {

	host := os.Getenv("AVIATRIX_HOST")
	username := os.Getenv("AVIATRIX_USERNAME")
	password := os.Getenv("AVIATRIX_PASSWORD")
	hostDeatails := fmt.Sprintf("==========>  AvxDatasource host : %v, username : %v, password : %v ", host, username, password)
	tflog.Debug(ctx, hostDeatails)
	fmt.Printf("Hello terraform debug : req.ProviderTypeName: %v ", req.ProviderTypeName)

	resp.TypeName = req.ProviderTypeName + "_avx"
}

// Schema defines the schema for the data source.
func (d *AvxDatasource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"avxdata": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *AvxDatasource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *AvxDatasource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	tflog.Debug(ctx, "===> avx data source read")
	//d.client.Get(ctx, "/", "test-data")

	// if d.client == nil {
	// 		panic("=============> PANIC the client is nil")
	// }else {
	// 	panic("=============> ELSE PANIC the client is not nil")
	// }
	//i := interface{}
	state := &AvxDataSourceModel{
		AvxData: []*AvxModel{
			{
				ID:          types.Int64Value(10),
				Name:        types.StringValue("Test name"),
				Description: types.StringValue("Test model description"),
			},
		},
	}
	// Set state
	diags := resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
