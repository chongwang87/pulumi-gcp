// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Get the IP address from a static address reserved for a Global Forwarding Rule which are only used for HTTP load balancing. For more information see
// the official [API](https://cloud.google.com/compute/docs/reference/latest/globalAddresses) documentation.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/datasource_compute_global_address.html.markdown.
func LookupGlobalAddress(ctx *pulumi.Context, args *LookupGlobalAddressArgs, opts ...pulumi.InvokeOption) (*LookupGlobalAddressResult, error) {
	var rv LookupGlobalAddressResult
	err := ctx.Invoke("gcp:compute/getGlobalAddress:getGlobalAddress", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getGlobalAddress.
type LookupGlobalAddressArgs struct {
	// A unique name for the resource, required by GCE.
	Name string `pulumi:"name"`
	// The project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getGlobalAddress.
type LookupGlobalAddressResult struct {
	// The IP of the created resource.
	Address string `pulumi:"address"`
	// id is the provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	Project string `pulumi:"project"`
	// The URI of the created resource.
	SelfLink string `pulumi:"selfLink"`
	// Indicates if the address is used. Possible values are: RESERVED or IN_USE.
	Status string `pulumi:"status"`
}
