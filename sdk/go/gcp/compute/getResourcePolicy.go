// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provide access to a Resource Policy's attributes. For more information see [the official documentation](https://cloud.google.com/compute/docs/disks/scheduled-snapshots) or the [API](https://cloud.google.com/compute/docs/reference/rest/beta/resourcePolicies).
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-gcp/sdk/v4/go/gcp/compute"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "us-central1"
// 		_, err := compute.LookupResourcePolicy(ctx, &compute.LookupResourcePolicyArgs{
// 			Name:   "daily",
// 			Region: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupResourcePolicy(ctx *pulumi.Context, args *LookupResourcePolicyArgs, opts ...pulumi.InvokeOption) (*LookupResourcePolicyResult, error) {
	var rv LookupResourcePolicyResult
	err := ctx.Invoke("gcp:compute/getResourcePolicy:getResourcePolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourcePolicy.
type LookupResourcePolicyArgs struct {
	// The name of the Resource Policy.
	Name string `pulumi:"name"`
	// Project from which to list the Resource Policy. Defaults to project declared in the provider.
	Project *string `pulumi:"project"`
	// Region where the Resource Policy resides.
	Region *string `pulumi:"region"`
}

// A collection of values returned by getResourcePolicy.
type LookupResourcePolicyResult struct {
	GroupPlacementPolicies []GetResourcePolicyGroupPlacementPolicy `pulumi:"groupPlacementPolicies"`
	// The provider-assigned unique ID for this managed resource.
	Id      string  `pulumi:"id"`
	Name    string  `pulumi:"name"`
	Project *string `pulumi:"project"`
	Region  *string `pulumi:"region"`
	// The URI of the resource.
	SelfLink                 string                                    `pulumi:"selfLink"`
	SnapshotSchedulePolicies []GetResourcePolicySnapshotSchedulePolicy `pulumi:"snapshotSchedulePolicies"`
}
