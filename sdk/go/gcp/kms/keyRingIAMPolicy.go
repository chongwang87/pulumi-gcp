// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package kms

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for KMS key ring. Each of these resources serves a different use case:
//
// * `kms.KeyRingIAMPolicy`: Authoritative. Sets the IAM policy for the key ring and replaces any existing policy already attached.
// * `kms.KeyRingIAMBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the key ring are preserved.
// * `kms.KeyRingIAMMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the key ring are preserved.
//
// > **Note:** `kms.KeyRingIAMPolicy` **cannot** be used in conjunction with `kms.KeyRingIAMBinding` and `kms.KeyRingIAMMember` or they will fight over what your policy should be.
//
// > **Note:** `kms.KeyRingIAMBinding` resources **can be** used in conjunction with `kms.KeyRingIAMMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_kms_key_ring_iam.html.markdown.
type KeyRingIAMPolicy struct {
	pulumi.CustomResourceState

	// (Computed) The etag of the key ring's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringOutput `pulumi:"keyRingId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
}

// NewKeyRingIAMPolicy registers a new resource with the given unique name, arguments, and options.
func NewKeyRingIAMPolicy(ctx *pulumi.Context,
	name string, args *KeyRingIAMPolicyArgs, opts ...pulumi.ResourceOption) (*KeyRingIAMPolicy, error) {
	if args == nil || args.KeyRingId == nil {
		return nil, errors.New("missing required argument 'KeyRingId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &KeyRingIAMPolicyArgs{}
	}
	var resource KeyRingIAMPolicy
	err := ctx.RegisterResource("gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKeyRingIAMPolicy gets an existing KeyRingIAMPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKeyRingIAMPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KeyRingIAMPolicyState, opts ...pulumi.ResourceOption) (*KeyRingIAMPolicy, error) {
	var resource KeyRingIAMPolicy
	err := ctx.ReadResource("gcp:kms/keyRingIAMPolicy:KeyRingIAMPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering KeyRingIAMPolicy resources.
type keyRingIAMPolicyState struct {
	// (Computed) The etag of the key ring's IAM policy.
	Etag *string `pulumi:"etag"`
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId *string `pulumi:"keyRingId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
}

type KeyRingIAMPolicyState struct {
	// (Computed) The etag of the key ring's IAM policy.
	Etag pulumi.StringPtrInput
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
}

func (KeyRingIAMPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingIAMPolicyState)(nil)).Elem()
}

type keyRingIAMPolicyArgs struct {
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId string `pulumi:"keyRingId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
}

// The set of arguments for constructing a KeyRingIAMPolicy resource.
type KeyRingIAMPolicyArgs struct {
	// The key ring ID, in the form
	// `{project_id}/{location_name}/{key_ring_name}` or
	// `{location_name}/{key_ring_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	KeyRingId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
}

func (KeyRingIAMPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*keyRingIAMPolicyArgs)(nil)).Elem()
}
