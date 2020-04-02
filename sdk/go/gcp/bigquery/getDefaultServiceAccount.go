// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package bigquery

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Use this data source to retrieve default service account for this project
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/google_bigquery_default_service_account.html.markdown.
func GetDefaultServiceAccount(ctx *pulumi.Context, args *GetDefaultServiceAccountArgs, opts ...pulumi.InvokeOption) (*GetDefaultServiceAccountResult, error) {
	var rv GetDefaultServiceAccountResult
	err := ctx.Invoke("gcp:bigquery/getDefaultServiceAccount:getDefaultServiceAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDefaultServiceAccount.
type GetDefaultServiceAccountArgs struct {
	// The project ID. If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
}

// A collection of values returned by getDefaultServiceAccount.
type GetDefaultServiceAccountResult struct {
	// Email address of the default service account used by bigquery encryption in this project
	Email string `pulumi:"email"`
	// id is the provider-assigned unique ID for this managed resource.
	Id      string `pulumi:"id"`
	Project string `pulumi:"project"`
}
