// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package servicedirectory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Service Directory Namespace. Each of these resources serves a different use case:
//
// * `servicedirectory.NamespaceIamPolicy`: Authoritative. Sets the IAM policy for the namespace and replaces any existing policy already attached.
// * `servicedirectory.NamespaceIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the namespace are preserved.
// * `servicedirectory.NamespaceIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the namespace are preserved.
//
// > **Note:** `servicedirectory.NamespaceIamPolicy` **cannot** be used in conjunction with `servicedirectory.NamespaceIamBinding` and `servicedirectory.NamespaceIamMember` or they will fight over what your policy should be.
//
// > **Note:** `servicedirectory.NamespaceIamBinding` resources **can be** used in conjunction with `servicedirectory.NamespaceIamMember` resources **only if** they do not grant privilege to the same role.
//
// ## google\_service\_directory\_namespace\_iam\_policy
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/organizations"
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/servicedirectory"
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
// 		_, err = servicedirectory.NewNamespaceIamPolicy(ctx, "policy", &servicedirectory.NamespaceIamPolicyArgs{
// 			PolicyData: pulumi.String(admin.PolicyData),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## google\_service\_directory\_namespace\_iam\_binding
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicedirectory.NewNamespaceIamBinding(ctx, "binding", &servicedirectory.NamespaceIamBindingArgs{
// 			Role: pulumi.String("roles/viewer"),
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
// ## google\_service\_directory\_namespace\_iam\_member
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/servicedirectory"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := servicedirectory.NewNamespaceIamMember(ctx, "member", &servicedirectory.NamespaceIamMemberArgs{
// 			Role:   pulumi.String("roles/viewer"),
// 			Member: pulumi.String("user:jane@example.com"),
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
// For all import syntaxes, the "resource in question" can take any of the following forms* projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} * {{project}}/{{location}}/{{namespace_id}} * {{location}}/{{namespace_id}} Any variables not passed in the import command will be taken from the provider configuration. Service Directory namespace IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} roles/viewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding editor "projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}} roles/viewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding editor projects/{{project}}/locations/{{location}}/namespaces/{{namespace_id}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type NamespaceIamBinding struct {
	pulumi.CustomResourceState

	Condition NamespaceIamBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringOutput      `pulumi:"etag"`
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringOutput `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewNamespaceIamBinding registers a new resource with the given unique name, arguments, and options.
func NewNamespaceIamBinding(ctx *pulumi.Context,
	name string, args *NamespaceIamBindingArgs, opts ...pulumi.ResourceOption) (*NamespaceIamBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource NamespaceIamBinding
	err := ctx.RegisterResource("gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNamespaceIamBinding gets an existing NamespaceIamBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNamespaceIamBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NamespaceIamBindingState, opts ...pulumi.ResourceOption) (*NamespaceIamBinding, error) {
	var resource NamespaceIamBinding
	err := ctx.ReadResource("gcp:servicedirectory/namespaceIamBinding:NamespaceIamBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NamespaceIamBinding resources.
type namespaceIamBindingState struct {
	Condition *NamespaceIamBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag    *string  `pulumi:"etag"`
	Members []string `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type NamespaceIamBindingState struct {
	Condition NamespaceIamBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag    pulumi.StringPtrInput
	Members pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (NamespaceIamBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIamBindingState)(nil)).Elem()
}

type namespaceIamBindingArgs struct {
	Condition *NamespaceIamBindingCondition `pulumi:"condition"`
	Members   []string                      `pulumi:"members"`
	// Used to find the parent resource to bind the IAM policy to
	Name *string `pulumi:"name"`
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a NamespaceIamBinding resource.
type NamespaceIamBindingArgs struct {
	Condition NamespaceIamBindingConditionPtrInput
	Members   pulumi.StringArrayInput
	// Used to find the parent resource to bind the IAM policy to
	Name pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `servicedirectory.NamespaceIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (NamespaceIamBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*namespaceIamBindingArgs)(nil)).Elem()
}

type NamespaceIamBindingInput interface {
	pulumi.Input

	ToNamespaceIamBindingOutput() NamespaceIamBindingOutput
	ToNamespaceIamBindingOutputWithContext(ctx context.Context) NamespaceIamBindingOutput
}

func (NamespaceIamBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceIamBinding)(nil)).Elem()
}

func (i NamespaceIamBinding) ToNamespaceIamBindingOutput() NamespaceIamBindingOutput {
	return i.ToNamespaceIamBindingOutputWithContext(context.Background())
}

func (i NamespaceIamBinding) ToNamespaceIamBindingOutputWithContext(ctx context.Context) NamespaceIamBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NamespaceIamBindingOutput)
}

type NamespaceIamBindingOutput struct {
	*pulumi.OutputState
}

func (NamespaceIamBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NamespaceIamBindingOutput)(nil)).Elem()
}

func (o NamespaceIamBindingOutput) ToNamespaceIamBindingOutput() NamespaceIamBindingOutput {
	return o
}

func (o NamespaceIamBindingOutput) ToNamespaceIamBindingOutputWithContext(ctx context.Context) NamespaceIamBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(NamespaceIamBindingOutput{})
}
