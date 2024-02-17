// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package imagebuilder

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/imagebuilder"
	awstypes "github.com/aws/aws-sdk-go-v2/service/imagebuilder/types"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/id"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/enum"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// @FrameworkResource(name="Lifecycle Policy")
func newResourceLifecyclePolicy(_ context.Context) (resource.ResourceWithConfigure, error) {
	return &resourceLifecyclePolicy{}, nil
}

const (
	ResNameLifecyclePolicy = "Lifecycle Policy"
)

type resourceLifecyclePolicy struct {
	framework.ResourceWithConfigure
}

func (r *resourceLifecyclePolicy) Metadata(_ context.Context, request resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "aws_imagebuilder_lifecycle_policy"
}

func (r *resourceLifecyclePolicy) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"arn": framework.ARNAttributeComputedOnly(),
			"description": schema.StringAttribute{
				Optional: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"execution_role": schema.StringAttribute{
				Required: true,
			},
			"id": framework.IDAttribute(),
			"name": schema.StringAttribute{
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"resource_type": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					enum.FrameworkValidate[awstypes.LifecyclePolicyResourceType](),
				},
			},
			"status": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					enum.FrameworkValidate[awstypes.LifecyclePolicyStatus](),
				},
			},
			names.AttrTags:    tftags.TagsAttribute(),
			names.AttrTagsAll: tftags.TagsAttributeComputedOnly(),
		},
		Blocks: map[string]schema.Block{
			"policy_details": schema.SetNestedBlock{
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
					setvalidator.SizeAtMost(3),
				},
				NestedObject: schema.NestedBlockObject{
					Blocks: map[string]schema.Block{
						"action": schema.ListNestedBlock{
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
								listvalidator.IsRequired(),
							},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											enum.FrameworkValidate[awstypes.LifecyclePolicyDetailActionType](),
										},
									},
								},
								Blocks: map[string]schema.Block{
									"include_resources": schema.ListNestedBlock{
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"amis": schema.BoolAttribute{
													Optional: true,
												},
												"containers": schema.BoolAttribute{
													Optional: true,
												},
												"snapshots": schema.BoolAttribute{
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
						"filter": schema.ListNestedBlock{
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
								listvalidator.IsRequired(),
							},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"type": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											enum.FrameworkValidate[awstypes.LifecyclePolicyDetailFilterType](),
										},
									},
									"value": schema.Int64Attribute{
										Required: true,
									},
									"retain_at_least": schema.Int64Attribute{
										Optional: true,
									},
									"unit": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											enum.FrameworkValidate[awstypes.LifecyclePolicyTimeUnit](),
										},
									},
								},
							},
						},
						"exclusion_rules": schema.ListNestedBlock{
							Validators: []validator.List{
								listvalidator.SizeAtMost(1),
							},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"tag_map": schema.MapAttribute{
										ElementType: types.StringType,
										Optional:    true,
									},
								},
								Blocks: map[string]schema.Block{
									"amis": schema.ListNestedBlock{
										Validators: []validator.List{
											listvalidator.SizeAtMost(1),
										},
										NestedObject: schema.NestedBlockObject{
											Attributes: map[string]schema.Attribute{
												"is_public": schema.BoolAttribute{
													Optional: true,
												},
												"regions": schema.ListAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
												"shared_accounts": schema.ListAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
												"tag_map": schema.MapAttribute{
													ElementType: types.StringType,
													Optional:    true,
												},
											},
											Blocks: map[string]schema.Block{
												"last_launched": schema.ListNestedBlock{
													Validators: []validator.List{
														listvalidator.SizeAtMost(1),
													},
													NestedObject: schema.NestedBlockObject{
														Attributes: map[string]schema.Attribute{
															"unit": schema.StringAttribute{
																Required: true,
																Validators: []validator.String{
																	enum.FrameworkValidate[awstypes.LifecyclePolicyTimeUnit](),
																},
															},
															"value": schema.Int64Attribute{
																Required: true,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"resource_selection": schema.ListNestedBlock{
				Validators: []validator.List{
					listvalidator.SizeAtMost(1),
				},
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"tag_map": schema.MapAttribute{
							ElementType: types.StringType,
							Optional:    true,
						},
					},
					Blocks: map[string]schema.Block{
						"recipes": schema.SetNestedBlock{
							Validators: []validator.Set{
								setvalidator.SizeAtLeast(1),
								setvalidator.SizeAtMost(50),
							},
							NestedObject: schema.NestedBlockObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Required: true,
									},
									"semantic_version": schema.StringAttribute{
										Required: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func (r *resourceLifecyclePolicy) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan resourceLifecyclePolicyData

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	conn := r.Meta().ImageBuilderClient(ctx)

	in := &imagebuilder.CreateLifecyclePolicyInput{
		ClientToken:   aws.String(id.UniqueId()),
		ExecutionRole: aws.String(plan.ExecutionRole.ValueString()),
		Name:          aws.String(plan.Name.ValueString()),
		ResourceType:  awstypes.LifecyclePolicyResourceType(plan.ResourceType.ValueString()),
		Status:        awstypes.LifecyclePolicyStatus(plan.Status.ValueString()),
		//Tags:          getTagsIn(ctx),
	}

	if !plan.Description.IsNull() {
		in.Description = aws.String(plan.Description.ValueString())
	}

	if !plan.PolicyDetails.IsNull() {
		var tfList []resourcePolicyDetailsData
		resp.Diagnostics.Append(plan.PolicyDetails.ElementsAs(ctx, &tfList, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		policyDetails, d := expandPolicyDetails(ctx, tfList)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
		in.PolicyDetails = policyDetails
	}

	if !plan.ResourceSelection.IsNull() {
		var tfList []resourceResourceSelectionData
		resp.Diagnostics.Append(plan.ResourceSelection.ElementsAs(ctx, &tfList, false)...)
		if resp.Diagnostics.HasError() {
			return
		}

		resourceSelection, d := expandResourceSelection(ctx, tfList)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}
		in.ResourceSelection = resourceSelection
	}

	out, err := conn.CreateLifecyclePolicy(ctx, in)
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.ImageBuilder, create.ErrActionCreating, ResNameLifecyclePolicy, plan.Name.String(), err),
			err.Error(),
		)
		return
	}
	if out == nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.ImageBuilder, create.ErrActionCreating, ResNameLifecyclePolicy, plan.Name.String(), nil),
			errors.New("empty output").Error(),
		)
		return
	}

	plan.ID = flex.StringToFramework(ctx, out.LifecyclePolicyArn)
	plan.ARN = flex.StringToFrameworkARN(ctx, out.LifecyclePolicyArn)

	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *resourceLifecyclePolicy) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	conn := r.Meta().ImageBuilderClient(ctx)

	var state resourceLifecyclePolicyData
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	out, err := conn.GetLifecyclePolicy(ctx, &imagebuilder.GetLifecyclePolicyInput{
		LifecyclePolicyArn: aws.String(state.ID.ValueString()),
	})
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.ImageBuilder, create.ErrActionReading, ResNameLifecyclePolicy, state.Name.String(), nil),
			err.Error(),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *resourceLifecyclePolicy) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	conn := r.Meta().ImageBuilderClient(ctx)

	var plan, state resourceLifecyclePolicyData
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !plan.Description.Equal(state.Description) ||
		!plan.ExecutionRole.Equal(state.ExecutionRole) ||
		!plan.PolicyDetails.Equal(state.PolicyDetails) ||
		!plan.ResourceSelection.Equal(state.ResourceSelection) ||
		!plan.ResourceType.Equal(state.ResourceType) ||
		!plan.Status.Equal(state.Status) {

		in := &imagebuilder.UpdateLifecyclePolicyInput{
			ExecutionRole: aws.String(plan.ExecutionRole.ValueString()),
			ResourceType:  awstypes.LifecyclePolicyResourceType(plan.ResourceType.ValueString()),
			Status:        awstypes.LifecyclePolicyStatus(plan.Status.ValueString()),
		}

		if !plan.Description.Equal(state.Description) {
			in.Description = aws.String(plan.Description.ValueString())
		}

		if !plan.PolicyDetails.Equal(state.PolicyDetails) {
			var tfList []resourcePolicyDetailsData
			resp.Diagnostics.Append(plan.PolicyDetails.ElementsAs(ctx, &tfList, false)...)
			if resp.Diagnostics.HasError() {
				return
			}

			policyDetails, d := expandPolicyDetails(ctx, tfList)
			resp.Diagnostics.Append(d...)
			if resp.Diagnostics.HasError() {
				return
			}
			in.PolicyDetails = policyDetails
		}

		if !plan.ResourceSelection.Equal(state.ResourceSelection) {
			var tfList []resourceResourceSelectionData
			resp.Diagnostics.Append(plan.ResourceSelection.ElementsAs(ctx, &tfList, false)...)
			if resp.Diagnostics.HasError() {
				return
			}

			resourceSelection, d := expandResourceSelection(ctx, tfList)
			resp.Diagnostics.Append(d...)
			if resp.Diagnostics.HasError() {
				return
			}
			in.ResourceSelection = resourceSelection
		}

		out, err := conn.UpdateLifecyclePolicy(ctx, in)
		if err != nil {
			resp.Diagnostics.AddError(
				create.ProblemStandardMessage(names.ImageBuilder, create.ErrActionUpdating, ResNameLifecyclePolicy, plan.ID.String(), err),
				err.Error(),
			)
			return
		}
		if out == nil {
			resp.Diagnostics.AddError(
				create.ProblemStandardMessage(names.ImageBuilder, create.ErrActionUpdating, ResNameLifecyclePolicy, plan.ID.String(), nil),
				errors.New("empty output").Error(),
			)
			return
		}
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourceLifecyclePolicy) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	conn := r.Meta().ImageBuilderClient(ctx)

	var state resourceLifecyclePolicyData
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := conn.DeleteLifecyclePolicy(ctx, &imagebuilder.DeleteLifecyclePolicyInput{
		LifecyclePolicyArn: aws.String(state.ID.ValueString()),
	})
	if err != nil {
		var nfe *awstypes.ResourceNotFoundException
		if errors.As(err, &nfe) {
			return
		}
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.ImageBuilder, create.ErrActionDeleting, ResNameLifecyclePolicy, state.Name.String(), nil),
			err.Error(),
		)
	}
}

func (r *resourceLifecyclePolicy) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}

func (r *resourceLifecyclePolicy) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	r.SetTagsAll(ctx, req, resp)
}

func expandPolicyDetails(ctx context.Context, tfList []resourcePolicyDetailsData) ([]awstypes.LifecyclePolicyDetail, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(tfList) == 0 {
		return nil, diags
	}

	apiResult := []awstypes.LifecyclePolicyDetail{}

	for _, policyDetail := range tfList {
		apiObject := awstypes.LifecyclePolicyDetail{}

		if !policyDetail.Action.IsNull() {
			var tfList []resourceActionData
			diags.Append(policyDetail.Action.ElementsAs(ctx, &tfList, false)...)
			if diags.HasError() {
				return nil, diags
			}

			action, d := expandPolicyDetailAction(ctx, tfList)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			apiObject.Action = action
		}

		if !policyDetail.Filter.IsNull() {
			var tfList []resourceFilterData
			diags.Append(policyDetail.Filter.ElementsAs(ctx, &tfList, false)...)
			if diags.HasError() {
				return nil, diags
			}

			filter, d := expandPolicyDetailFilter(ctx, tfList)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			apiObject.Filter = filter
		}

		if !policyDetail.ExclusionRules.IsNull() {
			var tfList []resourceExclusionRulesData
			diags.Append(policyDetail.ExclusionRules.ElementsAs(ctx, &tfList, false)...)
			if diags.HasError() {
				return nil, diags
			}

			exclusionRules, d := expandPolicyDetailExclusionRules(ctx, tfList)
			diags.Append(d...)
			if diags.HasError() {
				return nil, diags
			}
			apiObject.ExclusionRules = exclusionRules
		}

		apiResult = append(apiResult, apiObject)
	}

	return apiResult, diags
}

func expandPolicyDetailAction(ctx context.Context, tfList []resourceActionData) (*awstypes.LifecyclePolicyDetailAction, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(tfList) == 0 {
		return nil, diags
	}

	tfObj := tfList[0]

	apiObject := awstypes.LifecyclePolicyDetailAction{}

	if !tfObj.IncludeResources.IsNull() {
		var tfList []resourceIncludeResourcesData
		diags.Append(tfObj.IncludeResources.ElementsAs(ctx, &tfList, false)...)
		if diags.HasError() {
			return nil, diags
		}

		apiObject.IncludeResources = expandPolicyDetailActionIncludeResources(tfList)
	}

	if !tfObj.Type.IsNull() {
		apiObject.Type = awstypes.LifecyclePolicyDetailActionType(tfObj.Type.ValueString())
	}

	return &apiObject, diags
}

func expandPolicyDetailActionIncludeResources(tfList []resourceIncludeResourcesData) *awstypes.LifecyclePolicyDetailActionIncludeResources {
	tfObj := tfList[0]

	apiObject := awstypes.LifecyclePolicyDetailActionIncludeResources{}

	if !tfObj.Amis.IsNull() {
		apiObject.Amis = aws.ToBool(tfObj.Amis.ValueBoolPointer())
	}

	if !tfObj.Containers.IsNull() {
		apiObject.Containers = aws.ToBool(tfObj.Containers.ValueBoolPointer())
	}

	if !tfObj.Snapshots.IsNull() {
		apiObject.Snapshots = aws.ToBool(tfObj.Snapshots.ValueBoolPointer())
	}

	return &apiObject
}

func expandPolicyDetailFilter(ctx context.Context, tfList []resourceFilterData) (*awstypes.LifecyclePolicyDetailFilter, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(tfList) == 0 {
		return nil, diags
	}

	tfObj := tfList[0]

	apiObject := awstypes.LifecyclePolicyDetailFilter{}

	if !tfObj.Type.IsNull() {
		apiObject.Type = awstypes.LifecyclePolicyDetailFilterType(tfObj.Type.ValueString())
	}

	if !tfObj.Value.IsNull() {
		apiObject.Value = aws.Int32(int32(tfObj.Value.ValueInt64()))
	}

	if !tfObj.RetainAtLeast.IsNull() {
		apiObject.RetainAtLeast = aws.Int32(int32(tfObj.RetainAtLeast.ValueInt64()))
	}

	if !tfObj.Unit.IsNull() {
		apiObject.Unit = awstypes.LifecyclePolicyTimeUnit(tfObj.Type.ValueString())
	}

	return &apiObject, diags
}

func expandPolicyDetailExclusionRules(ctx context.Context, tfList []resourceExclusionRulesData) (*awstypes.LifecyclePolicyDetailExclusionRules, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(tfList) == 0 {
		return nil, diags
	}

	tfObj := tfList[0]

	apiObject := awstypes.LifecyclePolicyDetailExclusionRules{}

	if !tfObj.AMIs.IsNull() {
		var tfList []resourceAMIsData
		diags.Append(tfObj.AMIs.ElementsAs(ctx, &tfList, false)...)
		if diags.HasError() {
			return nil, diags
		}

		Amis, d := expandPolicyDetailExclusionRulesAmis(ctx, tfList)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}
		apiObject.Amis = Amis
	}

	if !tfObj.TagMap.IsNull() {
		apiObject.TagMap = flex.ExpandFrameworkStringValueMap(ctx, tfObj.TagMap)
	}

	return &apiObject, diags
}

func expandPolicyDetailExclusionRulesAmis(ctx context.Context, tfList []resourceAMIsData) (*awstypes.LifecyclePolicyDetailExclusionRulesAmis, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(tfList) == 0 {
		return nil, diags
	}

	tfObj := tfList[0]

	apiObject := awstypes.LifecyclePolicyDetailExclusionRulesAmis{}

	if !tfObj.IsPublic.IsNull() {
		apiObject.IsPublic = aws.ToBool(tfObj.IsPublic.ValueBoolPointer())
	}

	if !tfObj.LastLaunched.IsNull() {
		var tfList []resourceLastLaunchedData
		diags.Append(tfObj.LastLaunched.ElementsAs(ctx, &tfList, false)...)
		if diags.HasError() {
			return nil, diags
		}

		apiObject.LastLaunched = expandPolicyDetailExclusionRulesAmisLastLaunched(tfList)
	}

	if !tfObj.Regions.IsNull() {
		apiObject.Regions = flex.ExpandFrameworkStringValueList(ctx, tfObj.LastLaunched)
	}

	if !tfObj.SharedAccounts.IsNull() {
		apiObject.Regions = flex.ExpandFrameworkStringValueList(ctx, tfObj.SharedAccounts)
	}

	if !tfObj.TagMap.IsNull() {
		apiObject.TagMap = flex.ExpandFrameworkStringValueMap(ctx, tfObj.TagMap)
	}

	return &apiObject, diags
}

func expandPolicyDetailExclusionRulesAmisLastLaunched(tfList []resourceLastLaunchedData) *awstypes.LifecyclePolicyDetailExclusionRulesAmisLastLaunched {
	tfObj := tfList[0]

	apiObject := awstypes.LifecyclePolicyDetailExclusionRulesAmisLastLaunched{}

	if !tfObj.Unit.IsNull() {
		apiObject.Unit = awstypes.LifecyclePolicyTimeUnit(tfObj.Unit.ValueString())
	}

	if !tfObj.Value.IsNull() {
		apiObject.Value = aws.Int32(int32(tfObj.Value.ValueInt64()))
	}

	return &apiObject
}

func expandResourceSelection(ctx context.Context, tfList []resourceResourceSelectionData) (*awstypes.LifecyclePolicyResourceSelection, diag.Diagnostics) {
	var diags diag.Diagnostics

	if len(tfList) == 0 {
		return nil, diags
	}

	tfObj := tfList[0]

	apiObject := awstypes.LifecyclePolicyResourceSelection{}

	if !tfObj.Recipes.IsNull() {
		var tfList []resourceRecipesData
		diags.Append(tfObj.Recipes.ElementsAs(ctx, &tfList, false)...)
		if diags.HasError() {
			return nil, diags
		}

		apiObject.Recipes = expandResourceSelectionRecipes(tfList)
	}

	if !tfObj.TagMap.IsNull() {
		apiObject.TagMap = flex.ExpandFrameworkStringValueMap(ctx, tfObj.TagMap)
	}

	return &apiObject, diags
}

func expandResourceSelectionRecipes(tfList []resourceRecipesData) []awstypes.LifecyclePolicyResourceSelectionRecipe {
	apiResult := []awstypes.LifecyclePolicyResourceSelectionRecipe{}

	for _, tfObj := range tfList {
		apiObject := awstypes.LifecyclePolicyResourceSelectionRecipe{}

		if !tfObj.Name.IsNull() {
			apiObject.SemanticVersion = aws.String(tfObj.Name.ValueString())
		}
		if !tfObj.SemanticVersion.IsNull() {
			apiObject.SemanticVersion = aws.String(tfObj.SemanticVersion.ValueString())
		}

		apiResult = append(apiResult, apiObject)

	}
	return apiResult
}

type resourceLifecyclePolicyData struct {
	ID                types.String `tfsdk:"id"`
	ARN               fwtypes.ARN  `tfsdk:"arn"`
	Description       types.String `tfsdk:"description"`
	Name              types.String `tfsdk:"name"`
	ExecutionRole     types.String `tfsdk:"execution_role"`
	ResourceType      types.String `tfsdk:"resource_type"`
	Status            types.String `tfsdk:"status"`
	PolicyDetails     types.Set    `tfsdk:"policy_details"`
	ResourceSelection types.List   `tfsdk:"resource_selection"`
}

type resourcePolicyDetailsData struct {
	Action         types.List `tfsdk:"action"`
	Filter         types.List `tfsdk:"filter"`
	ExclusionRules types.List `tfsdk:"exclusion_rules"`
}

type resourceResourceSelectionData struct {
	TagMap  types.Map `tfsdk:"tag_map"`
	Recipes types.Set `tfsdk:"recipes"`
}

type resourceRecipesData struct {
	Name            types.String `tfsdk:"name"`
	SemanticVersion types.String `tfsdk:"semantic_version"`
}

type resourceActionData struct {
	Type             types.String `tfsdk:"type"`
	IncludeResources types.List   `tfsdk:"include_resources"`
}

type resourceIncludeResourcesData struct {
	Amis       types.Bool `tfsdk:"amis"`
	Containers types.Bool `tfsdk:"containers"`
	Snapshots  types.Bool `tfsdk:"snapshots"`
}

type resourceFilterData struct {
	Type          types.String `tfsdk:"type"`
	Value         types.Int64  `tfsdk:"value"`
	RetainAtLeast types.Int64  `tfsdk:"retain_at_least"`
	Unit          types.String `tfsdk:"unit"`
}

type resourceExclusionRulesData struct {
	AMIs   types.List `tfsdk:"ami"`
	TagMap types.Map  `tfsdk:"tag_map"`
}

type resourceAMIsData struct {
	IsPublic       types.Bool `tfsdk:"is_public"`
	LastLaunched   types.List `tfsdk:"last_launched"`
	Regions        types.List `tfsdk:"regions"`
	SharedAccounts types.List `tfsdk:"shared_accounts"`
	TagMap         types.Map  `tfsdk:"tag_map"`
}

type resourceLastLaunchedData struct {
	Unit  types.String `tfsdk:"unit"`
	Value types.Int64  `tfsdk:"value"`
}
