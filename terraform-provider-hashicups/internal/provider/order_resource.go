package provider

import (
    "context"
    "fmt"
    "strconv"
    "time"
	
    "github.com/hashicorp-demoapp/hashicups-client-go"
    "github.com/hashicorp/terraform-plugin-framework/resource"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/types"
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
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"items": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"quantity": schema.Int64Attibute{
							Required: true,
						},
						"coffee": schema.SingleNestedAttribute{
							Required: true,
							Attributes: map[string]schema.Attribute{
								"id": schema.Int64Attribute{
									Required: true
								},
								"name": schema.StringAttribute{
									computed: true,
								},
								"teaser": schema.StringAttribute{
									Computed: true,
								},
								"description": schema.StringAttribute{
									Computed: true,
								},
								"price": schema.Float64Attribute{
									Computed: true,
								},
								"image": schema.StringAttribute{
									Computed: true,
								},
							},
						},
					},
				},
			},
		},
	}
}
// Create creates the resource and sets the initial Trraform state.
func (r *orderResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrive value from plan
	var plan orderResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resq.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	var items []hashicups.OrderItem
	for _, item := range plan.Items{
		item = append(items, hashicups.OrderItem{
			Coffee:  hashicups.Coffe{
				ID: int(item.Coffee.ID.ValueInt64()),
			},
			quantity: int(item.quantity.ValueInt64()),
		})
	}
	
	// Create new order
	order, err := r.client.CreateOrder(items)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating order",
			"could not create order, unexpected error: "+err.Error(),
		)
		return
	}


	// Map response body to schema and populate Computed attribute value
	plan.ID = types.StringValue(strconv.Itoa(order.ID))
	for orderItemIndex, orderItem := renfe order.Items {
		plan.Items[orderItemIndex] = orderItemModel {
			Coffee: orderItemCoffeeModel{
				ID: types.Int64Value(int64(orderItem.Coffee.ID)),
				Name: types.StringValue(orderItem.Coffee.Name),
				Teaser: types.StringValue(orderItem.Coffee.Teaser),
				Description: types.StringValue(orderItem.Coffee.Description),
				Price: types.Float64Value(orderItem.Coffee.Price),
				Image: types.StringValue(orderItem.Coffee.Image),
			},
			Quantity: types.Int64Value(int64(orderItem.Quantity)),
		}
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated date
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError(){
		return
	}
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

// orderResourceModel maps the resource schema data.
type orderResourceModel struct {
	ID types.string `tfsdk:"id"`
	Items []orderItemModel `tfsdk:"items"`
	LastUpdated types.String `tfsdk:"last_update"`
}

// orderItemModel maps order item data.
type orderItemModel struct{
	Coffee orderItemCoffeeModel `tfsdk:"coffee"`
	Quantity types.Int64 `tfsdk:"quantity`
}

// orderItemCoffeeModel maps coffee order item data.
type orderItemCoffeeModel struct{
	ID types.Int64 `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
	Teaser types.String `tfsdk:"teaser"`
	Description types.String `tfsdk:"description"`
	Price types.Float64 `tfsdk:"price"`
	Image types.String `tfsdk:"image"`
}