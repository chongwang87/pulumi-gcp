// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages a Target Pool within GCE. This is a collection of instances used as
// target of a network load balancer (Forwarding Rule). For more information see
// [the official
// documentation](https://cloud.google.com/compute/docs/load-balancing/network/target-pools)
// and [API](https://cloud.google.com/compute/docs/reference/latest/targetPools).
//
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_target_pool.html.markdown.
type TargetPool struct {
	pulumi.CustomResourceState

	// URL to the backup target pool. Must also set
	// failover\_ratio.
	BackupPool pulumi.StringPtrOutput `pulumi:"backupPool"`
	// Textual description field.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio pulumi.Float64PtrOutput `pulumi:"failoverRatio"`
	// List of zero or one health check name or self_link. Only
	// legacy `compute.HttpHealthCheck` is supported.
	HealthChecks pulumi.StringPtrOutput `pulumi:"healthChecks"`
	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name". Note that the instances need not exist
	// at the time of target pool creation, so there is no need to use the
	// interpolation to create a dependency on the instances from the
	// target pool.
	Instances pulumi.StringArrayOutput `pulumi:"instances"`
	// A unique name for the resource, required by GCE. Changing
	// this forces a new resource to be created.
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Where the target pool resides. Defaults to project
	// region.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	SessionAffinity pulumi.StringPtrOutput `pulumi:"sessionAffinity"`
}

// NewTargetPool registers a new resource with the given unique name, arguments, and options.
func NewTargetPool(ctx *pulumi.Context,
	name string, args *TargetPoolArgs, opts ...pulumi.ResourceOption) (*TargetPool, error) {
	if args == nil {
		args = &TargetPoolArgs{}
	}
	var resource TargetPool
	err := ctx.RegisterResource("gcp:compute/targetPool:TargetPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTargetPool gets an existing TargetPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTargetPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TargetPoolState, opts ...pulumi.ResourceOption) (*TargetPool, error) {
	var resource TargetPool
	err := ctx.ReadResource("gcp:compute/targetPool:TargetPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TargetPool resources.
type targetPoolState struct {
	// URL to the backup target pool. Must also set
	// failover\_ratio.
	BackupPool *string `pulumi:"backupPool"`
	// Textual description field.
	Description *string `pulumi:"description"`
	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio *float64 `pulumi:"failoverRatio"`
	// List of zero or one health check name or self_link. Only
	// legacy `compute.HttpHealthCheck` is supported.
	HealthChecks *string `pulumi:"healthChecks"`
	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name". Note that the instances need not exist
	// at the time of target pool creation, so there is no need to use the
	// interpolation to create a dependency on the instances from the
	// target pool.
	Instances []string `pulumi:"instances"`
	// A unique name for the resource, required by GCE. Changing
	// this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Where the target pool resides. Defaults to project
	// region.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	SessionAffinity *string `pulumi:"sessionAffinity"`
}

type TargetPoolState struct {
	// URL to the backup target pool. Must also set
	// failover\_ratio.
	BackupPool pulumi.StringPtrInput
	// Textual description field.
	Description pulumi.StringPtrInput
	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio pulumi.Float64PtrInput
	// List of zero or one health check name or self_link. Only
	// legacy `compute.HttpHealthCheck` is supported.
	HealthChecks pulumi.StringPtrInput
	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name". Note that the instances need not exist
	// at the time of target pool creation, so there is no need to use the
	// interpolation to create a dependency on the instances from the
	// target pool.
	Instances pulumi.StringArrayInput
	// A unique name for the resource, required by GCE. Changing
	// this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Where the target pool resides. Defaults to project
	// region.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	SessionAffinity pulumi.StringPtrInput
}

func (TargetPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*targetPoolState)(nil)).Elem()
}

type targetPoolArgs struct {
	// URL to the backup target pool. Must also set
	// failover\_ratio.
	BackupPool *string `pulumi:"backupPool"`
	// Textual description field.
	Description *string `pulumi:"description"`
	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio *float64 `pulumi:"failoverRatio"`
	// List of zero or one health check name or self_link. Only
	// legacy `compute.HttpHealthCheck` is supported.
	HealthChecks *string `pulumi:"healthChecks"`
	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name". Note that the instances need not exist
	// at the time of target pool creation, so there is no need to use the
	// interpolation to create a dependency on the instances from the
	// target pool.
	Instances []string `pulumi:"instances"`
	// A unique name for the resource, required by GCE. Changing
	// this forces a new resource to be created.
	Name *string `pulumi:"name"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Where the target pool resides. Defaults to project
	// region.
	Region *string `pulumi:"region"`
	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	SessionAffinity *string `pulumi:"sessionAffinity"`
}

// The set of arguments for constructing a TargetPool resource.
type TargetPoolArgs struct {
	// URL to the backup target pool. Must also set
	// failover\_ratio.
	BackupPool pulumi.StringPtrInput
	// Textual description field.
	Description pulumi.StringPtrInput
	// Ratio (0 to 1) of failed nodes before using the
	// backup pool (which must also be set).
	FailoverRatio pulumi.Float64PtrInput
	// List of zero or one health check name or self_link. Only
	// legacy `compute.HttpHealthCheck` is supported.
	HealthChecks pulumi.StringPtrInput
	// List of instances in the pool. They can be given as
	// URLs, or in the form of "zone/name". Note that the instances need not exist
	// at the time of target pool creation, so there is no need to use the
	// interpolation to create a dependency on the instances from the
	// target pool.
	Instances pulumi.StringArrayInput
	// A unique name for the resource, required by GCE. Changing
	// this forces a new resource to be created.
	Name pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Where the target pool resides. Defaults to project
	// region.
	Region pulumi.StringPtrInput
	// How to distribute load. Options are "NONE" (no
	// affinity). "CLIENT\_IP" (hash of the source/dest addresses / ports), and
	// "CLIENT\_IP\_PROTO" also includes the protocol (default "NONE").
	SessionAffinity pulumi.StringPtrInput
}

func (TargetPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*targetPoolArgs)(nil)).Elem()
}
