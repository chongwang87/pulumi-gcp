// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package accesscontextmanager

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An AccessLevel is a label that can be applied to requests to GCP services,
// along with a list of requirements necessary for the label to be applied.
//
//
// To get more information about AccessLevel, see:
//
// * [API documentation](https://cloud.google.com/access-context-manager/docs/reference/rest/v1/accessPolicies.accessLevels)
// * How-to Guides
//     * [Access Policy Quickstart](https://cloud.google.com/access-context-manager/docs/quickstart)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/access_context_manager_access_level.html.markdown.
type AccessLevel struct {
	pulumi.CustomResourceState

	// A set of predefined conditions for the access level and a combining function.
	Basic AccessLevelBasicPtrOutput `pulumi:"basic"`
	// Description of the AccessLevel and its use. Does not affect behavior.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
	// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name pulumi.StringOutput `pulumi:"name"`
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
	Parent pulumi.StringOutput `pulumi:"parent"`
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringOutput `pulumi:"title"`
}

// NewAccessLevel registers a new resource with the given unique name, arguments, and options.
func NewAccessLevel(ctx *pulumi.Context,
	name string, args *AccessLevelArgs, opts ...pulumi.ResourceOption) (*AccessLevel, error) {
	if args == nil || args.Parent == nil {
		return nil, errors.New("missing required argument 'Parent'")
	}
	if args == nil || args.Title == nil {
		return nil, errors.New("missing required argument 'Title'")
	}
	if args == nil {
		args = &AccessLevelArgs{}
	}
	var resource AccessLevel
	err := ctx.RegisterResource("gcp:accesscontextmanager/accessLevel:AccessLevel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAccessLevel gets an existing AccessLevel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAccessLevel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AccessLevelState, opts ...pulumi.ResourceOption) (*AccessLevel, error) {
	var resource AccessLevel
	err := ctx.ReadResource("gcp:accesscontextmanager/accessLevel:AccessLevel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AccessLevel resources.
type accessLevelState struct {
	// A set of predefined conditions for the access level and a combining function.
	Basic *AccessLevelBasic `pulumi:"basic"`
	// Description of the AccessLevel and its use. Does not affect behavior.
	Description *string `pulumi:"description"`
	// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
	// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name *string `pulumi:"name"`
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
	Parent *string `pulumi:"parent"`
	// Human readable title. Must be unique within the Policy.
	Title *string `pulumi:"title"`
}

type AccessLevelState struct {
	// A set of predefined conditions for the access level and a combining function.
	Basic AccessLevelBasicPtrInput
	// Description of the AccessLevel and its use. Does not affect behavior.
	Description pulumi.StringPtrInput
	// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
	// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name pulumi.StringPtrInput
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
	Parent pulumi.StringPtrInput
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringPtrInput
}

func (AccessLevelState) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelState)(nil)).Elem()
}

type accessLevelArgs struct {
	// A set of predefined conditions for the access level and a combining function.
	Basic *AccessLevelBasic `pulumi:"basic"`
	// Description of the AccessLevel and its use. Does not affect behavior.
	Description *string `pulumi:"description"`
	// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
	// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name *string `pulumi:"name"`
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
	Parent string `pulumi:"parent"`
	// Human readable title. Must be unique within the Policy.
	Title string `pulumi:"title"`
}

// The set of arguments for constructing a AccessLevel resource.
type AccessLevelArgs struct {
	// A set of predefined conditions for the access level and a combining function.
	Basic AccessLevelBasicPtrInput
	// Description of the AccessLevel and its use. Does not affect behavior.
	Description pulumi.StringPtrInput
	// Resource name for the Access Level. The short_name component must begin with a letter and only include alphanumeric and
	// '_'. Format: accessPolicies/{policy_id}/accessLevels/{short_name}
	Name pulumi.StringPtrInput
	// The AccessPolicy this AccessLevel lives in. Format: accessPolicies/{policy_id}
	Parent pulumi.StringInput
	// Human readable title. Must be unique within the Policy.
	Title pulumi.StringInput
}

func (AccessLevelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*accessLevelArgs)(nil)).Elem()
}
