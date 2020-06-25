// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package projects

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Retrieve information about a set of projects based on a filter. See the
// [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list)
// for more details.
//
// ## Example Usage
func GetProject(ctx *pulumi.Context, args *GetProjectArgs, opts ...pulumi.InvokeOption) (*GetProjectResult, error) {
	var rv GetProjectResult
	err := ctx.Invoke("gcp:projects/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type GetProjectArgs struct {
	// A string filter as defined in the [REST API](https://cloud.google.com/resource-manager/reference/rest/v1/projects/list#query-parameters).
	Filter string `pulumi:"filter"`
}

// A collection of values returned by getProject.
type GetProjectResult struct {
	Filter string `pulumi:"filter"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// A list of projects matching the provided filter. Structure is defined below.
	Projects []GetProjectProject `pulumi:"projects"`
}
