// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package notebooks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Cloud AI Notebooks Instance. Each of these resources serves a different use case:
//
// * `notebooks.InstanceIamPolicy`: Authoritative. Sets the IAM policy for the instance and replaces any existing policy already attached.
// * `notebooks.InstanceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the instance are preserved.
// * `notebooks.InstanceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the instance are preserved.
//
// > **Note:** `notebooks.InstanceIamPolicy` **cannot** be used in conjunction with `notebooks.InstanceIamBinding` and `notebooks.InstanceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `notebooks.InstanceIamBinding` resources **can be** used in conjunction with `notebooks.InstanceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_notebooks\_instance\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		admin, err := organizations.LookupIAMPolicy(ctx, &organizations.LookupIAMPolicyArgs{
// 			Bindings: []organizations.GetIAMPolicyBinding{
// 				organizations.GetIAMPolicyBinding{
// 					Role: "roles/viewer",
// 					Members: []string{
// 						"user:jane@example.com",
// 					},
// 				},
// 			},
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = notebooks.NewInstanceIamPolicy(ctx, "policy", &notebooks.InstanceIamPolicyArgs{
// 			Project:      pulumi.Any(google_notebooks_instance.Instance.Project),
// 			Location:     pulumi.Any(google_notebooks_instance.Instance.Location),
// 			InstanceName: pulumi.Any(google_notebooks_instance.Instance.Name),
// 			PolicyData:   pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_notebooks\_instance\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewInstanceIamBinding(ctx, "binding", &notebooks.InstanceIamBindingArgs{
// 			Project:      pulumi.Any(google_notebooks_instance.Instance.Project),
// 			Location:     pulumi.Any(google_notebooks_instance.Instance.Location),
// 			InstanceName: pulumi.Any(google_notebooks_instance.Instance.Name),
// 			Role:         pulumi.String("roles/viewer"),
// 			Members: pulumi.StringArray{
// 				pulumi.String("user:jane@example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_notebooks\_instance\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/notebooks"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := notebooks.NewInstanceIamMember(ctx, "member", &notebooks.InstanceIamMemberArgs{
// 			Project:      pulumi.Any(google_notebooks_instance.Instance.Project),
// 			Location:     pulumi.Any(google_notebooks_instance.Instance.Location),
// 			InstanceName: pulumi.Any(google_notebooks_instance.Instance.Name),
// 			Role:         pulumi.String("roles/viewer"),
// 			Member:       pulumi.String("user:jane@example.com"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/instances/{{instance_name}} * {{project}}/{{location}}/{{instance_name}} * {{location}}/{{instance_name}} * {{instance_name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud AI Notebooks instance IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/instanceIamMember:InstanceIamMember editor "projects/{{project}}/locations/{{location}}/instances/{{instance_name}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/instanceIamMember:InstanceIamMember editor "projects/{{project}}/locations/{{location}}/instances/{{instance_name}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:notebooks/instanceIamMember:InstanceIamMember editor projects/{{project}}/locations/{{location}}/instances/{{instance_name}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type InstanceIamMember struct {
	pulumi.CustomResourceState

	Condition InstanceIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringOutput `pulumi:"instanceName"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringOutput `pulumi:"location"`
	Member   pulumi.StringOutput `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The role that should be applied. Only one
	// `notebooks.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewInstanceIamMember registers a new resource with the given unique name, arguments, and options.
func NewInstanceIamMember(ctx *pulumi.Context,
	name string, args *InstanceIamMemberArgs, opts ...pulumi.ResourceOption) (*InstanceIamMember, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceName == nil {
		return nil, errors.New("invalid value for required argument 'InstanceName'")
	}
	if args.Member == nil {
		return nil, errors.New("invalid value for required argument 'Member'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource InstanceIamMember
	err := ctx.RegisterResource("gcp:notebooks/instanceIamMember:InstanceIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceIamMember gets an existing InstanceIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceIamMemberState, opts ...pulumi.ResourceOption) (*InstanceIamMember, error) {
	var resource InstanceIamMember
	err := ctx.ReadResource("gcp:notebooks/instanceIamMember:InstanceIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceIamMember resources.
type instanceIamMemberState struct {
	Condition *InstanceIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName *string `pulumi:"instanceName"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   *string `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `notebooks.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type InstanceIamMemberState struct {
	Condition InstanceIamMemberConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringPtrInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `notebooks.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (InstanceIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamMemberState)(nil)).Elem()
}

type instanceIamMemberArgs struct {
	Condition *InstanceIamMemberCondition `pulumi:"condition"`
	// Used to find the parent resource to bind the IAM policy to
	InstanceName string `pulumi:"instanceName"`
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location *string `pulumi:"location"`
	Member   string  `pulumi:"member"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
	// The role that should be applied. Only one
	// `notebooks.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a InstanceIamMember resource.
type InstanceIamMemberArgs struct {
	Condition InstanceIamMemberConditionPtrInput
	// Used to find the parent resource to bind the IAM policy to
	InstanceName pulumi.StringInput
	// A reference to the zone where the machine resides. Used to find the parent resource to bind the IAM policy to
	Location pulumi.StringPtrInput
	Member   pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `notebooks.InstanceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (InstanceIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceIamMemberArgs)(nil)).Elem()
}

type InstanceIamMemberInput interface {
	pulumi.Input

	ToInstanceIamMemberOutput() InstanceIamMemberOutput
	ToInstanceIamMemberOutputWithContext(ctx context.Context) InstanceIamMemberOutput
}

func (InstanceIamMember) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIamMember)(nil)).Elem()
}

func (i InstanceIamMember) ToInstanceIamMemberOutput() InstanceIamMemberOutput {
	return i.ToInstanceIamMemberOutputWithContext(context.Background())
}

func (i InstanceIamMember) ToInstanceIamMemberOutputWithContext(ctx context.Context) InstanceIamMemberOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceIamMemberOutput)
}

type InstanceIamMemberOutput struct {
	*pulumi.OutputState
}

func (InstanceIamMemberOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceIamMemberOutput)(nil)).Elem()
}

func (o InstanceIamMemberOutput) ToInstanceIamMemberOutput() InstanceIamMemberOutput {
	return o
}

func (o InstanceIamMemberOutput) ToInstanceIamMemberOutputWithContext(ctx context.Context) InstanceIamMemberOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(InstanceIamMemberOutput{})
}
