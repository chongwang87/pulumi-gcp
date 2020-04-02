// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package composer

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides access to available Cloud Composer versions in a region for a given project.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_google_composer_image_versions.html.markdown.
func GetImageVersions(ctx *pulumi.Context, args *GetImageVersionsArgs, opts ...pulumi.InvokeOption) (*GetImageVersionsResult, error) {
	var rv GetImageVersionsResult
	err := ctx.Invoke("gcp:composer/getImageVersions:getImageVersions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getImageVersions.
type GetImageVersionsArgs struct {
	// The ID of the project to list versions in.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The location to list versions in.
	// If it is not provider, the provider region is used.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getImageVersions.
type GetImageVersionsResult struct {
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of composer image versions available in the given project and location. Each `imageVersion` contains:
	ImageVersions []GetImageVersionsImageVersion `pulumi:"imageVersions"`
	Project       string                         `pulumi:"project"`
	Region        string                         `pulumi:"region"`
}
