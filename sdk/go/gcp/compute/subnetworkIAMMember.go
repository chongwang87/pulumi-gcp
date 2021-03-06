// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/regions/{{region}}/subnetworks/{{name}} * {{project}}/{{region}}/{{name}} * {{region}}/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Compute Engine subnetwork IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/subnetworkIAMMember:SubnetworkIAMMember editor "projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/subnetworkIAMMember:SubnetworkIAMMember editor "projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}} roles/compute.networkUser"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:compute/subnetworkIAMMember:SubnetworkIAMMember editor projects/{{project}}/regions/{{region}}/subnetworks/{{subnetwork}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type SubnetworkIAMMember struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition SubnetworkIAMMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringOutput `pulumi:"etag"`
	Member pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringOutput `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringOutput `pulumi:"subnetwork"`
}

// NewSubnetworkIAMMember registers a new resource with the given unique name, arguments, and options.
func NewSubnetworkIAMMember(ctx *pulumi.Context,
	name string, args *SubnetworkIAMMemberArgs, opts ...pulumi.ResourceOption) (*SubnetworkIAMMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	if args.Subnetwork == nil {
		return nil, errors.New("invalid value for required argument 'Subnetwork'")
	}
	var resource SubnetworkIAMMember
	err := ctx.RegisterResource("gcp:compute/subnetworkIAMMember:SubnetworkIAMMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnetworkIAMMember gets an existing SubnetworkIAMMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnetworkIAMMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetworkIAMMemberState, opts ...pulumi.ResourceOption) (*SubnetworkIAMMember, error) {
	var resource SubnetworkIAMMember
	err := ctx.ReadResource("gcp:compute/subnetworkIAMMember:SubnetworkIAMMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubnetworkIAMMember resources.
type subnetworkIAMMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *SubnetworkIAMMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag   *string `pulumi:"etag"`
	Member *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork *string `pulumi:"subnetwork"`
}

type SubnetworkIAMMemberState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition SubnetworkIAMMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag   pulumi.StringPtrInput
	Member pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringPtrInput
}

func (SubnetworkIAMMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkIAMMemberState)(nil)).Elem()
}

type subnetworkIAMMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *SubnetworkIAMMemberCondition `pulumi:"condition"`
	Member    string                        `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region *string `pulumi:"region"`
	// The role that should be applied. Only one
	// `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork string `pulumi:"subnetwork"`
}

// The set of arguments for constructing a SubnetworkIAMMember resource.
type SubnetworkIAMMemberArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition SubnetworkIAMMemberConditionPtrInput
	Member    pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The GCP region for this subnetwork.
	// Used to find the parent resource to bind the IAM policy to. If not specified,
	// the value will be parsed from the identifier of the parent resource. If no region is provided in the parent identifier and no
	// region is specified, it is taken from the provider configuration.
	Region pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `compute.SubnetworkIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
	// Used to find the parent resource to bind the IAM policy to
	Subnetwork pulumi.StringInput
}

func (SubnetworkIAMMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetworkIAMMemberArgs)(nil)).Elem()
}

type SubnetworkIAMMemberInput interface {
	pulumi.Input

	ToSubnetworkIAMMemberOutput() SubnetworkIAMMemberOutput
	ToSubnetworkIAMMemberOutputWithContext(ctx context.Context) SubnetworkIAMMemberOutput
}

func (SubnetworkIAMMember) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetworkIAMMember)(nil)).Elem()
}

func (i SubnetworkIAMMember) ToSubnetworkIAMMemberOutput() SubnetworkIAMMemberOutput {
	return i.ToSubnetworkIAMMemberOutputWithContext(context.Background())
}

func (i SubnetworkIAMMember) ToSubnetworkIAMMemberOutputWithContext(ctx context.Context) SubnetworkIAMMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetworkIAMMemberOutput)
}

type SubnetworkIAMMemberOutput struct {
	*pulumi.OutputState
}

func (SubnetworkIAMMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetworkIAMMemberOutput)(nil)).Elem()
}

func (o SubnetworkIAMMemberOutput) ToSubnetworkIAMMemberOutput() SubnetworkIAMMemberOutput {
	return o
}

func (o SubnetworkIAMMemberOutput) ToSubnetworkIAMMemberOutputWithContext(ctx context.Context) SubnetworkIAMMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubnetworkIAMMemberOutput{})
}
