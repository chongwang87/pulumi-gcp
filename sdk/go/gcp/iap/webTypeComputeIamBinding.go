// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iap

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/iap_web/compute * {{project}} Any variables not passed in the import command will be taken from the provider configuration. Identity-Aware Proxy webtypecompute IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeComputeIamBinding:WebTypeComputeIamBinding editor "projects/{{project}}/iap_web/compute roles/iap.httpsResourceAccessor user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeComputeIamBinding:WebTypeComputeIamBinding editor "projects/{{project}}/iap_web/compute roles/iap.httpsResourceAccessor"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:iap/webTypeComputeIamBinding:WebTypeComputeIamBinding editor projects/{{project}}/iap_web/compute
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type WebTypeComputeIamBinding struct {
	pulumi.CustomResourceState

	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeComputeIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewWebTypeComputeIamBinding registers a new resource with the given unique name, arguments, and options.
func NewWebTypeComputeIamBinding(ctx *pulumi.Context,
	name string, args *WebTypeComputeIamBindingArgs, opts ...pulumi.ResourceOption) (*WebTypeComputeIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource WebTypeComputeIamBinding
	err := ctx.RegisterResource("gcp:iap/webTypeComputeIamBinding:WebTypeComputeIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTypeComputeIamBinding gets an existing WebTypeComputeIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTypeComputeIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTypeComputeIamBindingState, opts ...pulumi.ResourceOption) (*WebTypeComputeIamBinding, error) {
	var resource WebTypeComputeIamBinding
	err := ctx.ReadResource("gcp:iap/webTypeComputeIamBinding:WebTypeComputeIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTypeComputeIamBinding resources.
type webTypeComputeIamBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebTypeComputeIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type WebTypeComputeIamBindingState struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeComputeIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (WebTypeComputeIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeComputeIamBindingState)(nil)).Elem()
}

type webTypeComputeIamBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *WebTypeComputeIamBindingCondition `pulumi:"condition"`
	Members   []string                           `pulumi:"members"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a WebTypeComputeIamBinding resource.
type WebTypeComputeIamBindingArgs struct {
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition WebTypeComputeIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `iap.WebTypeComputeIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (WebTypeComputeIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeComputeIamBindingArgs)(nil)).Elem()
}

type WebTypeComputeIamBindingInput interface {
	pulumi.Input

	ToWebTypeComputeIamBindingOutput() WebTypeComputeIamBindingOutput
	ToWebTypeComputeIamBindingOutputWithContext(ctx context.Context) WebTypeComputeIamBindingOutput
}

func (WebTypeComputeIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTypeComputeIamBinding)(nil)).Elem()
}

func (i WebTypeComputeIamBinding) ToWebTypeComputeIamBindingOutput() WebTypeComputeIamBindingOutput {
	return i.ToWebTypeComputeIamBindingOutputWithContext(context.Background())
}

func (i WebTypeComputeIamBinding) ToWebTypeComputeIamBindingOutputWithContext(ctx context.Context) WebTypeComputeIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebTypeComputeIamBindingOutput)
}

type WebTypeComputeIamBindingOutput struct {
	*pulumi.OutputState
}

func (WebTypeComputeIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTypeComputeIamBindingOutput)(nil)).Elem()
}

func (o WebTypeComputeIamBindingOutput) ToWebTypeComputeIamBindingOutput() WebTypeComputeIamBindingOutput {
	return o
}

func (o WebTypeComputeIamBindingOutput) ToWebTypeComputeIamBindingOutputWithContext(ctx context.Context) WebTypeComputeIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebTypeComputeIamBindingOutput{})
}
