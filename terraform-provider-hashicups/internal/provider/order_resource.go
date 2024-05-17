package provider

import (
	"context"

	"github.com/hashicope-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &orderResource{}
	resource.ResourceWithConfigure = &orderResource{}
)

// NewOrderResource is a helper function to simplify the privider implementation.
func NewOrderResource() resource.Resource {
	return &orderResource{}
}

// orderResource is the resource implementation.
type orderResource struct{}

// Metadata return the resource type name.
func (r *orderResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponese) {
	resp.TypeName = req.prividerTypeName + "_order"
}

// Schema defines the schema for the resource.
func (r *orderResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Create creates the resource and sets the initial Trraform state.
func (r *orderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

// Read refreshes the  Terraform state with the latest data.
func (r *orderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on sucess.
func (r *orderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateRespense) {
}

// Delete deletes the resource and removes the Terraform state on succes.
func (r *orderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

// Resources defines the resources implemented in the provider.
func (p *hashicupsProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewOrderResource,
	}
}

// orderResource is the resource implementation.

type orderResource struct{
	client *hashicups.Client
}

// Configure adds the provider configured client to the resource.
func (r *orderResource) Configure(_ context.Configure, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*hashicups.Client)

	if !ok{
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *hashicups.Client, fot: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return 
	}

	r.client = client
}
