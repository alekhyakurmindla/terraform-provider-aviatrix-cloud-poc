package aviatrixresource

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/terraform-provider-aviatrix-cloud-poc/client"
)

type AviatrixAccountResourceModel struct {
	Email    types.String `tfsdk:"id"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

type AviatrixAccountResource struct {
	client *client.Client
}

func NewAviatrixAccountUser() resource.Resource {
	return &AviatrixAccountResource{}
}

func (r *AviatrixAccountResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_account_user"
}

// Schema defines the schema for the resource.
func (r *AviatrixAccountResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				Computed: false,
				Required: true,
			},
			"username": schema.StringAttribute{
				Computed: false,
				Required: true,
			},
			"password": schema.StringAttribute{
				Computed:  false,
				Required: true,
			},
		},
	}
}

// Configure adds the provider configured client to the resource.
func (r *AviatrixAccountResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hashicups.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *AviatrixAccountResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	resource := &AviatrixAccountResourceModel{}
	// Set state to fully populated data
	//diags := req.Plan.Get(ctx, resource)
	diags := resp.State.Set(ctx, resource)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *AviatrixAccountResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

func (r *AviatrixAccountResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *AviatrixAccountResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {

}
