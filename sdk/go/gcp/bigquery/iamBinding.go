// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package bigquery

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} * {{project}}/{{dataset_id}}/{{table_id}} * {{dataset_id}}/{{table_id}} * {{table_id}} Any variables not passed in the import command will be taken from the provider configuration. BigQuery table IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:bigquery/iamBinding:IamBinding editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:bigquery/iamBinding:IamBinding editor "projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}} roles/bigquery.dataOwner"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:bigquery/iamBinding:IamBinding editor projects/{{project}}/datasets/{{dataset_id}}/tables/{{table_id}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type IamBinding struct {
	pulumi.CustomResourceState

	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IamBindingConditionPtrOutput `pulumi:"condition"`
	DatasetId pulumi.StringOutput          `pulumi:"datasetId"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    pulumi.StringOutput `pulumi:"role"`
	TableId pulumi.StringOutput `pulumi:"tableId"`
}

// NewIamBinding registers a new resource with the given unique name, arguments, and options.
func NewIamBinding(ctx *pulumi.Context,
	name string, args *IamBindingArgs, opts ...pulumi.ResourceOption) (*IamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.TableId == nil {
		return nil, errors.New("invalid value for required argument 'TableId'")
	}
	var resource IamBinding
	err := ctx.RegisterResource("gcp:bigquery/iamBinding:IamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIamBinding gets an existing IamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IamBindingState, opts ...pulumi.ResourceOption) (*IamBinding, error) {
	var resource IamBinding
	err := ctx.ReadResource("gcp:bigquery/iamBinding:IamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IamBinding resources.
type iamBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IamBindingCondition `pulumi:"condition"`
	DatasetId *string              `pulumi:"datasetId"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    *string `pulumi:"role"`
	TableId *string `pulumi:"tableId"`
}

type IamBindingState struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IamBindingConditionPtrInput
	DatasetId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    pulumi.StringPtrInput
	TableId pulumi.StringPtrInput
}

func (IamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamBindingState)(nil)).Elem()
}

type iamBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *IamBindingCondition `pulumi:"condition"`
	DatasetId string               `pulumi:"datasetId"`
	Members   []string             `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    string `pulumi:"role"`
	TableId string `pulumi:"tableId"`
}

// The set of arguments for constructing a IamBinding resource.
type IamBindingArgs struct {
	// An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition IamBindingConditionPtrInput
	DatasetId pulumi.StringInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `bigquery.IamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role    pulumi.StringInput
	TableId pulumi.StringInput
}

func (IamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamBindingArgs)(nil)).Elem()
}

type IamBindingInput interface {
	pulumi.Input

	ToIamBindingOutput() IamBindingOutput
	ToIamBindingOutputWithContext(ctx context.Context) IamBindingOutput
}

func (IamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*IamBinding)(nil)).Elem()
}

func (i IamBinding) ToIamBindingOutput() IamBindingOutput {
	return i.ToIamBindingOutputWithContext(context.Background())
}

func (i IamBinding) ToIamBindingOutputWithContext(ctx context.Context) IamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IamBindingOutput)
}

type IamBindingOutput struct {
	*pulumi.OutputState
}

func (IamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IamBindingOutput)(nil)).Elem()
}

func (o IamBindingOutput) ToIamBindingOutput() IamBindingOutput {
	return o
}

func (o IamBindingOutput) ToIamBindingOutputWithContext(ctx context.Context) IamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IamBindingOutput{})
}
