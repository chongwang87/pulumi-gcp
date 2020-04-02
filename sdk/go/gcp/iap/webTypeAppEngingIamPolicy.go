// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package iap

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Three different resources help you manage your IAM policy for Identity-Aware Proxy WebTypeAppEngine. Each of these resources serves a different use case:
//
// * `iap.WebTypeAppEngingIamPolicy`: Authoritative. Sets the IAM policy for the webtypeappengine and replaces any existing policy already attached.
// * `iap.WebTypeAppEngingIamBinding`: Authoritative for a given role. Updates the IAM policy to grant a role to a list of members. Other roles within the IAM policy for the webtypeappengine are preserved.
// * `iap.WebTypeAppEngingIamMember`: Non-authoritative. Updates the IAM policy to grant a role to a new member. Other members for the role for the webtypeappengine are preserved.
//
// > **Note:** `iap.WebTypeAppEngingIamPolicy` **cannot** be used in conjunction with `iap.WebTypeAppEngingIamBinding` and `iap.WebTypeAppEngingIamMember` or they will fight over what your policy should be.
//
// > **Note:** `iap.WebTypeAppEngingIamBinding` resources **can be** used in conjunction with `iap.WebTypeAppEngingIamMember` resources **only if** they do not grant privilege to the same role.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/iap_web_type_app_engine_iam.html.markdown.
type WebTypeAppEngingIamPolicy struct {
	pulumi.CustomResourceState

	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringOutput `pulumi:"appId"`
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringOutput `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
}

// NewWebTypeAppEngingIamPolicy registers a new resource with the given unique name, arguments, and options.
func NewWebTypeAppEngingIamPolicy(ctx *pulumi.Context,
	name string, args *WebTypeAppEngingIamPolicyArgs, opts ...pulumi.ResourceOption) (*WebTypeAppEngingIamPolicy, error) {
	if args == nil || args.AppId == nil {
		return nil, errors.New("missing required argument 'AppId'")
	}
	if args == nil || args.PolicyData == nil {
		return nil, errors.New("missing required argument 'PolicyData'")
	}
	if args == nil {
		args = &WebTypeAppEngingIamPolicyArgs{}
	}
	var resource WebTypeAppEngingIamPolicy
	err := ctx.RegisterResource("gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebTypeAppEngingIamPolicy gets an existing WebTypeAppEngingIamPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebTypeAppEngingIamPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebTypeAppEngingIamPolicyState, opts ...pulumi.ResourceOption) (*WebTypeAppEngingIamPolicy, error) {
	var resource WebTypeAppEngingIamPolicy
	err := ctx.ReadResource("gcp:iap/webTypeAppEngingIamPolicy:WebTypeAppEngingIamPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebTypeAppEngingIamPolicy resources.
type webTypeAppEngingIamPolicyState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId *string `pulumi:"appId"`
	// (Computed) The etag of the IAM policy.
	Etag *string `pulumi:"etag"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData *string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

type WebTypeAppEngingIamPolicyState struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringPtrInput
	// (Computed) The etag of the IAM policy.
	Etag pulumi.StringPtrInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebTypeAppEngingIamPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeAppEngingIamPolicyState)(nil)).Elem()
}

type webTypeAppEngingIamPolicyArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId string `pulumi:"appId"`
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData string `pulumi:"policyData"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a WebTypeAppEngingIamPolicy resource.
type WebTypeAppEngingIamPolicyArgs struct {
	// Id of the App Engine application. Used to find the parent resource to bind the IAM policy to
	AppId pulumi.StringInput
	// The policy data generated by
	// a `organizations.getIAMPolicy` data source.
	PolicyData pulumi.StringInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the project will be parsed from the identifier of the parent resource. If no project is provided in the parent identifier and no project is specified, the provider project is used.
	Project pulumi.StringPtrInput
}

func (WebTypeAppEngingIamPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webTypeAppEngingIamPolicyArgs)(nil)).Elem()
}
