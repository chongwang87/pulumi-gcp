// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package storage

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// For all import syntaxes, the "resource in question" can take any of the following forms* b/{{name}} * {{name}} Any variables not passed in the import command will be taken from the provider configuration. Cloud Storage bucket IAM resources can be imported using the resource identifiers, role, and member. IAM member imports use space-delimited identifiersthe resource in question, the role, and the member identity, e.g.
//
// ```sh
//  $ pulumi import gcp:storage/bucketIAMBinding:BucketIAMBinding editor "b/{{bucket}} roles/storage.objectViewer user:jane@example.com"
// ```
//
//  IAM binding imports use space-delimited identifiersthe resource in question and the role, e.g.
//
// ```sh
//  $ pulumi import gcp:storage/bucketIAMBinding:BucketIAMBinding editor "b/{{bucket}} roles/storage.objectViewer"
// ```
//
//  IAM policy imports use the identifier of the resource in question, e.g.
//
// ```sh
//  $ pulumi import gcp:storage/bucketIAMBinding:BucketIAMBinding editor b/{{bucket}}
// ```
//
//  -> **Custom Roles**If you're importing a IAM resource with a custom role, make sure to use the
//
// full name of the custom role, e.g. `[projects/my-project|organizations/my-org]/roles/my-custom-role`.
type BucketIAMBinding struct {
	pulumi.CustomResourceState

	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringOutput `pulumi:"bucket"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BucketIAMBindingConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members pulumi.StringArrayOutput `pulumi:"members"`
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewBucketIAMBinding registers a new resource with the given unique name, arguments, and options.
func NewBucketIAMBinding(ctx *pulumi.Context,
	name string, args *BucketIAMBindingArgs, opts ...pulumi.ResourceOption) (*BucketIAMBinding, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Bucket == nil {
		return nil, errors.New("invalid value for required argument 'Bucket'")
	}
	if args.Members == nil {
		return nil, errors.New("invalid value for required argument 'Members'")
	}
	if args.Role == nil {
		return nil, errors.New("invalid value for required argument 'Role'")
	}
	var resource BucketIAMBinding
	err := ctx.RegisterResource("gcp:storage/bucketIAMBinding:BucketIAMBinding", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBucketIAMBinding gets an existing BucketIAMBinding resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBucketIAMBinding(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BucketIAMBindingState, opts ...pulumi.ResourceOption) (*BucketIAMBinding, error) {
	var resource BucketIAMBinding
	err := ctx.ReadResource("gcp:storage/bucketIAMBinding:BucketIAMBinding", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BucketIAMBinding resources.
type bucketIAMBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket *string `pulumi:"bucket"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *BucketIAMBindingCondition `pulumi:"condition"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type BucketIAMBindingState struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringPtrInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BucketIAMBindingConditionPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (BucketIAMBindingState) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIAMBindingState)(nil)).Elem()
}

type bucketIAMBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket string `pulumi:"bucket"`
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition *BucketIAMBindingCondition `pulumi:"condition"`
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members []string `pulumi:"members"`
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a BucketIAMBinding resource.
type BucketIAMBindingArgs struct {
	// Used to find the parent resource to bind the IAM policy to
	Bucket pulumi.StringInput
	// ) An [IAM Condition](https://cloud.google.com/iam/docs/conditions-overview) for a given binding.
	// Structure is documented below.
	Condition BucketIAMBindingConditionPtrInput
	// Identities that will be granted the privilege in `role`.
	// Each entry can have one of the following values:
	// * **allUsers**: A special identifier that represents anyone who is on the internet; with or without a Google account.
	// * **allAuthenticatedUsers**: A special identifier that represents anyone who is authenticated with a Google account or a service account.
	// * **user:{emailid}**: An email address that represents a specific Google account. For example, alice@gmail.com or joe@example.com.
	// * **serviceAccount:{emailid}**: An email address that represents a service account. For example, my-other-app@appspot.gserviceaccount.com.
	// * **group:{emailid}**: An email address that represents a Google group. For example, admins@example.com.
	// * **domain:{domain}**: A G Suite domain (primary, instead of alias) name that represents all the users of that domain. For example, google.com or example.com.
	Members pulumi.StringArrayInput
	// The role that should be applied. Only one
	// `storage.BucketIAMBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (BucketIAMBindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bucketIAMBindingArgs)(nil)).Elem()
}

type BucketIAMBindingInput interface {
	pulumi.Input

	ToBucketIAMBindingOutput() BucketIAMBindingOutput
	ToBucketIAMBindingOutputWithContext(ctx context.Context) BucketIAMBindingOutput
}

func (BucketIAMBinding) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketIAMBinding)(nil)).Elem()
}

func (i BucketIAMBinding) ToBucketIAMBindingOutput() BucketIAMBindingOutput {
	return i.ToBucketIAMBindingOutputWithContext(context.Background())
}

func (i BucketIAMBinding) ToBucketIAMBindingOutputWithContext(ctx context.Context) BucketIAMBindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BucketIAMBindingOutput)
}

type BucketIAMBindingOutput struct {
	*pulumi.OutputState
}

func (BucketIAMBindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BucketIAMBindingOutput)(nil)).Elem()
}

func (o BucketIAMBindingOutput) ToBucketIAMBindingOutput() BucketIAMBindingOutput {
	return o
}

func (o BucketIAMBindingOutput) ToBucketIAMBindingOutputWithContext(ctx context.Context) BucketIAMBindingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BucketIAMBindingOutput{})
}
