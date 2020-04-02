// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package projects

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Allows management of a single API service for an existing Google Cloud Platform project.
//
// For a list of services available, visit the
// [API library page](https://console.cloud.google.com/apis/library) or run `gcloud services list`.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/google_project_service.html.markdown.
type Service struct {
	pulumi.CustomResourceState

	// If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
	// If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
	DisableDependentServices pulumi.BoolPtrOutput `pulumi:"disableDependentServices"`
	// If true, disable the service when the resource is destroyed.  Defaults to true.  May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy pulumi.BoolPtrOutput `pulumi:"disableOnDestroy"`
	// The project ID. If not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The service to enable.
	Service pulumi.StringOutput `pulumi:"service"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil || args.Service == nil {
		return nil, errors.New("missing required argument 'Service'")
	}
	if args == nil {
		args = &ServiceArgs{}
	}
	var resource Service
	err := ctx.RegisterResource("gcp:projects/service:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("gcp:projects/service:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
	// If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
	// If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
	DisableDependentServices *bool `pulumi:"disableDependentServices"`
	// If true, disable the service when the resource is destroyed.  Defaults to true.  May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy *bool `pulumi:"disableOnDestroy"`
	// The project ID. If not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The service to enable.
	Service *string `pulumi:"service"`
}

type ServiceState struct {
	// If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
	// If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
	DisableDependentServices pulumi.BoolPtrInput
	// If true, disable the service when the resource is destroyed.  Defaults to true.  May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy pulumi.BoolPtrInput
	// The project ID. If not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The service to enable.
	Service pulumi.StringPtrInput
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
	// If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
	DisableDependentServices *bool `pulumi:"disableDependentServices"`
	// If true, disable the service when the resource is destroyed.  Defaults to true.  May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy *bool `pulumi:"disableOnDestroy"`
	// The project ID. If not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The service to enable.
	Service string `pulumi:"service"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// If `true`, services that are enabled and which depend on this service should also be disabled when this service is destroyed.
	// If `false` or unset, an error will be generated if any enabled services depend on this service when destroying it.
	DisableDependentServices pulumi.BoolPtrInput
	// If true, disable the service when the resource is destroyed.  Defaults to true.  May be useful in the event that a project is long-lived but the infrastructure running in that project changes frequently.
	DisableOnDestroy pulumi.BoolPtrInput
	// The project ID. If not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The service to enable.
	Service pulumi.StringInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}
