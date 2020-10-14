// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package compute

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Google Compute Engine Instance Group Manager API creates and manages pools
// of homogeneous Compute Engine virtual machine instances from a common instance
// template. For more information, see [the official documentation](https://cloud.google.com/compute/docs/instance-groups/manager)
// and [API](https://cloud.google.com/compute/docs/reference/latest/instanceGroupManagers)
//
// > **Note:** Use [compute.RegionInstanceGroupManager](https://www.terraform.io/docs/providers/google/r/compute_region_instance_group_manager.html) to create a regional (multi-zone) instance group manager.
//
// ## Example Usage
type InstanceGroupManager struct {
	pulumi.CustomResourceState

	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies InstanceGroupManagerAutoHealingPoliciesPtrOutput `pulumi:"autoHealingPolicies"`
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName pulumi.StringOutput `pulumi:"baseInstanceName"`
	// An optional textual description of the instance
	// group manager.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The fingerprint of the instance group manager.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// The full URL of the instance group created by the manager.
	InstanceGroup pulumi.StringOutput `pulumi:"instanceGroup"`
	// - Version name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupManagerNamedPortArrayOutput `pulumi:"namedPorts"`
	Operation  pulumi.StringOutput                      `pulumi:"operation"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// The URL of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks InstanceGroupManagerStatefulDiskArrayOutput `pulumi:"statefulDisks"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayOutput `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntOutput `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy InstanceGroupManagerUpdatePolicyOutput `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions InstanceGroupManagerVersionArrayOutput `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrOutput `pulumi:"waitForInstances"`
	// The zone that instances in this group should be created
	// in.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstanceGroupManager registers a new resource with the given unique name, arguments, and options.
func NewInstanceGroupManager(ctx *pulumi.Context,
	name string, args *InstanceGroupManagerArgs, opts ...pulumi.ResourceOption) (*InstanceGroupManager, error) {
	if args == nil || args.BaseInstanceName == nil {
		return nil, errors.New("missing required argument 'BaseInstanceName'")
	}
	if args == nil || args.Versions == nil {
		return nil, errors.New("missing required argument 'Versions'")
	}
	if args == nil {
		args = &InstanceGroupManagerArgs{}
	}
	var resource InstanceGroupManager
	err := ctx.RegisterResource("gcp:compute/instanceGroupManager:InstanceGroupManager", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstanceGroupManager gets an existing InstanceGroupManager resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstanceGroupManager(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceGroupManagerState, opts ...pulumi.ResourceOption) (*InstanceGroupManager, error) {
	var resource InstanceGroupManager
	err := ctx.ReadResource("gcp:compute/instanceGroupManager:InstanceGroupManager", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InstanceGroupManager resources.
type instanceGroupManagerState struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies *InstanceGroupManagerAutoHealingPolicies `pulumi:"autoHealingPolicies"`
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName *string `pulumi:"baseInstanceName"`
	// An optional textual description of the instance
	// group manager.
	Description *string `pulumi:"description"`
	// The fingerprint of the instance group manager.
	Fingerprint *string `pulumi:"fingerprint"`
	// The full URL of the instance group created by the manager.
	InstanceGroup *string `pulumi:"instanceGroup"`
	// - Version name.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []InstanceGroupManagerNamedPort `pulumi:"namedPorts"`
	Operation  *string                         `pulumi:"operation"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// The URL of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks []InstanceGroupManagerStatefulDisk `pulumi:"statefulDisks"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools []string `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize *int `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy *InstanceGroupManagerUpdatePolicy `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions []InstanceGroupManagerVersion `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances *bool `pulumi:"waitForInstances"`
	// The zone that instances in this group should be created
	// in.
	Zone *string `pulumi:"zone"`
}

type InstanceGroupManagerState struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies InstanceGroupManagerAutoHealingPoliciesPtrInput
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName pulumi.StringPtrInput
	// An optional textual description of the instance
	// group manager.
	Description pulumi.StringPtrInput
	// The fingerprint of the instance group manager.
	Fingerprint pulumi.StringPtrInput
	// The full URL of the instance group created by the manager.
	InstanceGroup pulumi.StringPtrInput
	// - Version name.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupManagerNamedPortArrayInput
	Operation  pulumi.StringPtrInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// The URL of the created resource.
	SelfLink pulumi.StringPtrInput
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks InstanceGroupManagerStatefulDiskArrayInput
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayInput
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntPtrInput
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy InstanceGroupManagerUpdatePolicyPtrInput
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions InstanceGroupManagerVersionArrayInput
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrInput
	// The zone that instances in this group should be created
	// in.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupManagerState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupManagerState)(nil)).Elem()
}

type instanceGroupManagerArgs struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies *InstanceGroupManagerAutoHealingPolicies `pulumi:"autoHealingPolicies"`
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName string `pulumi:"baseInstanceName"`
	// An optional textual description of the instance
	// group manager.
	Description *string `pulumi:"description"`
	// - Version name.
	Name *string `pulumi:"name"`
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts []InstanceGroupManagerNamedPort `pulumi:"namedPorts"`
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks []InstanceGroupManagerStatefulDisk `pulumi:"statefulDisks"`
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools []string `pulumi:"targetPools"`
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize *int `pulumi:"targetSize"`
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy *InstanceGroupManagerUpdatePolicy `pulumi:"updatePolicy"`
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions []InstanceGroupManagerVersion `pulumi:"versions"`
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances *bool `pulumi:"waitForInstances"`
	// The zone that instances in this group should be created
	// in.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a InstanceGroupManager resource.
type InstanceGroupManagerArgs struct {
	// The autohealing policies for this managed instance
	// group. You can specify only one value. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/creating-groups-of-managed-instances#monitoring_groups).
	AutoHealingPolicies InstanceGroupManagerAutoHealingPoliciesPtrInput
	// The base instance name to use for
	// instances in this group. The value must be a valid
	// [RFC1035](https://www.ietf.org/rfc/rfc1035.txt) name. Supported characters
	// are lowercase letters, numbers, and hyphens (-). Instances are named by
	// appending a hyphen and a random four-character string to the base instance
	// name.
	BaseInstanceName pulumi.StringInput
	// An optional textual description of the instance
	// group manager.
	Description pulumi.StringPtrInput
	// - Version name.
	Name pulumi.StringPtrInput
	// The named port configuration. See the section below
	// for details on configuration.
	NamedPorts InstanceGroupManagerNamedPortArrayInput
	// The ID of the project in which the resource belongs. If it
	// is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Disks created on the instances that will be preserved on instance delete, update, etc. Structure is documented below. For more information see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/configuring-stateful-disks-in-migs).
	StatefulDisks InstanceGroupManagerStatefulDiskArrayInput
	// The full URL of all target pools to which new
	// instances in the group are added. Updating the target pools attribute does
	// not affect existing instances.
	TargetPools pulumi.StringArrayInput
	// - The number of instances calculated as a fixed number or a percentage depending on the settings. Structure is documented below.
	TargetSize pulumi.IntPtrInput
	// The update policy for this managed instance group. Structure is documented below. For more information, see the [official documentation](https://cloud.google.com/compute/docs/instance-groups/updating-managed-instance-groups) and [API](https://cloud.google.com/compute/docs/reference/rest/v1/instanceGroupManagers/patch)
	UpdatePolicy InstanceGroupManagerUpdatePolicyPtrInput
	// Application versions managed by this instance group. Each
	// version deals with a specific instance template, allowing canary release scenarios.
	// Structure is documented below.
	Versions InstanceGroupManagerVersionArrayInput
	// Whether to wait for all instances to be created/updated before
	// returning. Note that if this is set to true and the operation does not succeed, this provider will
	// continue trying until it times out.
	WaitForInstances pulumi.BoolPtrInput
	// The zone that instances in this group should be created
	// in.
	Zone pulumi.StringPtrInput
}

func (InstanceGroupManagerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceGroupManagerArgs)(nil)).Elem()
}
