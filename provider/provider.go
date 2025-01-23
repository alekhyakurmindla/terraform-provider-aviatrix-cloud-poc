package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/terraform-provider-aviatrix-cloud-poc/client"
	"github.com/terraform-provider-aviatrix-cloud-poc/datasource/avxdatasource"
	"github.com/terraform-provider-aviatrix-cloud-poc/datasource/controllerdatasource"
	"github.com/terraform-provider-aviatrix-cloud-poc/resource/aviatrixresource"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &aviatrixProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &aviatrixProvider{
			version: version,
		}
	}
}

// aviatrixProvider is the provider implementation.
type aviatrixProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *aviatrixProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "aviatrix-cloud-poc"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *aviatrixProvider) Schema(ctx context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	tflog.Trace(ctx, "Debug provider schema")
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

// hashicupsProviderModel maps provider schema data to a Go type.
type aviatrixProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

// Configure prepares a Aviatrix API client for data sources and resources.
func (p *aviatrixProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Aviatrix client")
	fmt.Println("Hello World")

	// Retrieve provider data from configuration
	config := aviatrixProviderModel{}
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	
	host := os.Getenv("AVIATRIX_HOST")
	username := os.Getenv("AVIATRIX_USERNAME")
	password := os.Getenv("AVIATRIX_PASSWORD")

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Host",
			"The provider cannot create the Aviatrix API client as there is an unknown configuration value for the host.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown Username",
			"The provider cannot create the Aviatrix API client as there is an unknown configuration value for the username.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Password",
			"The provider cannot create the Aviatrix API client as there is an unknown configuration value for the password.",
		)
	}

	hostDeatails := fmt.Sprintf("==========>  host : %v, username : %v, password : %v ", host, username, password)
	tflog.Debug(ctx, hostDeatails)
	fmt.Println("Hello terraform debug")

	client, err := client.NewClient(host, username, password)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Aviatrix API Client",
			"An unexpected error occurred when creating the Aviatrix API client. "+
				"If the error is not clear, please contact the provider.\n\n"+
				"Aviatrix Client Error: "+err.Error(),
		)
		return
	}
	resp.DataSourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *aviatrixProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		avxdatasource.NewAviatrixDataSource,
        controllerdatasource.NewControllerDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *aviatrixProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		aviatrixresource.NewAviatrixAccountUser,
	}
}
