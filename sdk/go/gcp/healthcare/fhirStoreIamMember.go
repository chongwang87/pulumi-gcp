// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package healthcare

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Healthcare FHIR store. Each of these resources serves a different use case:
//
// * `healthcare.FhirStoreIamPolicy`: Authoritative. Sets the IAM policy for the FHIR store and replaces any existing policy already attached.
// * `healthcare.FhirStoreIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the FHIR store are preserved.
// * `healthcare.FhirStoreIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the FHIR store are preserved.
//
// > **Note:** `healthcare.FhirStoreIamPolicy` **cannot** be used in conjunction with `healthcare.FhirStoreIamBinding` and `healthcare.FhirStoreIamMember` or they will fight over what your policy should be.
//
// > **Note:** `healthcare.FhirStoreIamBinding` resources **can be** used in conjunction with `healthcare.FhirStoreIamMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/healthcare_fhir_store_iam.html.markdown.
type FhirStoreIamMember struct {
	pulumi.CustomResourceState

	Condition FhirStoreIamMemberConditionPtrOutput `pulumi:"condition"`
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringOutput `pulumi:"fhirStoreId"`
	Member      pulumi.StringOutput `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringOutput `pulumi:"role"`
}

// NewFhirStoreIamMember registers a new resource with the given unique name, arguments, and options.
func NewFhirStoreIamMember(ctx *pulumi.Context,
	name string, args *FhirStoreIamMemberArgs, opts ...pulumi.ResourceOption) (*FhirStoreIamMember, error) {
	if args == nil || args.FhirStoreId == nil {
		return nil, errors.New("missing required argument 'FhirStoreId'")
	}
	if args == nil || args.Member == nil {
		return nil, errors.New("missing required argument 'Member'")
	}
	if args == nil || args.Role == nil {
		return nil, errors.New("missing required argument 'Role'")
	}
	if args == nil {
		args = &FhirStoreIamMemberArgs{}
	}
	var resource FhirStoreIamMember
	err := ctx.RegisterResource("gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetFhirStoreIamMember gets an existing FhirStoreIamMember resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetFhirStoreIamMember(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FhirStoreIamMemberState, opts ...pulumi.ResourceOption) (*FhirStoreIamMember, error) {
	var resource FhirStoreIamMember
	err := ctx.ReadResource("gcp:healthcare/fhirStoreIamMember:FhirStoreIamMember", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering FhirStoreIamMember resources.
type fhirStoreIamMemberState struct {
	Condition *FhirStoreIamMemberCondition `pulumi:"condition"`
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag *string `pulumi:"etag"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId *string `pulumi:"fhirStoreId"`
	Member      *string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role *string `pulumi:"role"`
}

type FhirStoreIamMemberState struct {
	Condition FhirStoreIamMemberConditionPtrInput
	// (Computed) The etag of the FHIR store's IAM policy.
	Etag pulumi.StringPtrInput
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringPtrInput
	Member      pulumi.StringPtrInput
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringPtrInput
}

func (FhirStoreIamMemberState) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreIamMemberState)(nil)).Elem()
}

type fhirStoreIamMemberArgs struct {
	Condition *FhirStoreIamMemberCondition `pulumi:"condition"`
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId string `pulumi:"fhirStoreId"`
	Member      string `pulumi:"member"`
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role string `pulumi:"role"`
}

// The set of arguments for constructing a FhirStoreIamMember resource.
type FhirStoreIamMemberArgs struct {
	Condition FhirStoreIamMemberConditionPtrInput
	// The FHIR store ID, in the form
	// `{project_id}/{location_name}/{dataset_name}/{fhir_store_name}` or
	// `{location_name}/{dataset_name}/{fhir_store_name}`. In the second form, the provider's
	// project setting will be used as a fallback.
	FhirStoreId pulumi.StringInput
	Member      pulumi.StringInput
	// The role that should be applied. Only one
	// `healthcare.FhirStoreIamBinding` can be used per role. Note that custom roles must be of the format
	// `[projects|organizations]/{parent-name}/roles/{role-name}`.
	Role pulumi.StringInput
}

func (FhirStoreIamMemberArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fhirStoreIamMemberArgs)(nil)).Elem()
}
