// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package compute

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Represents a NodeTemplate resource. Node templates specify properties
// for creating sole-tenant nodes, such as node type, vCPU and memory
// requirements, node affinity labels, and region.
//
//
// To get more information about NodeTemplate, see:
//
// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeTemplates)
// * How-to Guides
//     * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_node_template.html.markdown.
type NodeTemplate struct {
	pulumi.CustomResourceState

	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// An optional textual description of the resource.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels pulumi.StringMapOutput `pulumi:"nodeAffinityLabels"`
	// Node type to use for nodes group that are created from this template. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	NodeType pulumi.StringPtrOutput `pulumi:"nodeType"`
	// Flexible properties for the desired node type. Node groups that use this node template will create nodes of a type that
	// matches these properties. Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityPtrOutput `pulumi:"nodeTypeFlexibility"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringOutput `pulumi:"project"`
	// Region where nodes using the node template will be created. If it is not provided, the provider region is used.
	Region pulumi.StringOutput `pulumi:"region"`
	// The URI of the created resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// The server binding policy for nodes using this template. Determines where the nodes should restart following a
	// maintenance event.
	ServerBinding NodeTemplateServerBindingOutput `pulumi:"serverBinding"`
}

// NewNodeTemplate registers a new resource with the given unique name, arguments, and options.
func NewNodeTemplate(ctx *pulumi.Context,
	name string, args *NodeTemplateArgs, opts ...pulumi.ResourceOption) (*NodeTemplate, error) {
	if args == nil {
		args = &NodeTemplateArgs{}
	}
	var resource NodeTemplate
	err := ctx.RegisterResource("gcp:compute/nodeTemplate:NodeTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNodeTemplate gets an existing NodeTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNodeTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NodeTemplateState, opts ...pulumi.ResourceOption) (*NodeTemplate, error) {
	var resource NodeTemplate
	err := ctx.ReadResource("gcp:compute/nodeTemplate:NodeTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NodeTemplate resources.
type nodeTemplateState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp *string `pulumi:"creationTimestamp"`
	// An optional textual description of the resource.
	Description *string `pulumi:"description"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels map[string]string `pulumi:"nodeAffinityLabels"`
	// Node type to use for nodes group that are created from this template. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	NodeType *string `pulumi:"nodeType"`
	// Flexible properties for the desired node type. Node groups that use this node template will create nodes of a type that
	// matches these properties. Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeTypeFlexibility *NodeTemplateNodeTypeFlexibility `pulumi:"nodeTypeFlexibility"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where nodes using the node template will be created. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The URI of the created resource.
	SelfLink *string `pulumi:"selfLink"`
	// The server binding policy for nodes using this template. Determines where the nodes should restart following a
	// maintenance event.
	ServerBinding *NodeTemplateServerBinding `pulumi:"serverBinding"`
}

type NodeTemplateState struct {
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringPtrInput
	// An optional textual description of the resource.
	Description pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels pulumi.StringMapInput
	// Node type to use for nodes group that are created from this template. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	NodeType pulumi.StringPtrInput
	// Flexible properties for the desired node type. Node groups that use this node template will create nodes of a type that
	// matches these properties. Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where nodes using the node template will be created. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The URI of the created resource.
	SelfLink pulumi.StringPtrInput
	// The server binding policy for nodes using this template. Determines where the nodes should restart following a
	// maintenance event.
	ServerBinding NodeTemplateServerBindingPtrInput
}

func (NodeTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTemplateState)(nil)).Elem()
}

type nodeTemplateArgs struct {
	// An optional textual description of the resource.
	Description *string `pulumi:"description"`
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels map[string]string `pulumi:"nodeAffinityLabels"`
	// Node type to use for nodes group that are created from this template. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	NodeType *string `pulumi:"nodeType"`
	// Flexible properties for the desired node type. Node groups that use this node template will create nodes of a type that
	// matches these properties. Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeTypeFlexibility *NodeTemplateNodeTypeFlexibility `pulumi:"nodeTypeFlexibility"`
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project *string `pulumi:"project"`
	// Region where nodes using the node template will be created. If it is not provided, the provider region is used.
	Region *string `pulumi:"region"`
	// The server binding policy for nodes using this template. Determines where the nodes should restart following a
	// maintenance event.
	ServerBinding *NodeTemplateServerBinding `pulumi:"serverBinding"`
}

// The set of arguments for constructing a NodeTemplate resource.
type NodeTemplateArgs struct {
	// An optional textual description of the resource.
	Description pulumi.StringPtrInput
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Labels to use for node affinity, which will be used in instance scheduling.
	NodeAffinityLabels pulumi.StringMapInput
	// Node type to use for nodes group that are created from this template. Only one of nodeTypeFlexibility and nodeType can
	// be specified.
	NodeType pulumi.StringPtrInput
	// Flexible properties for the desired node type. Node groups that use this node template will create nodes of a type that
	// matches these properties. Only one of nodeTypeFlexibility and nodeType can be specified.
	NodeTypeFlexibility NodeTemplateNodeTypeFlexibilityPtrInput
	// The ID of the project in which the resource belongs.
	// If it is not provided, the provider project is used.
	Project pulumi.StringPtrInput
	// Region where nodes using the node template will be created. If it is not provided, the provider region is used.
	Region pulumi.StringPtrInput
	// The server binding policy for nodes using this template. Determines where the nodes should restart following a
	// maintenance event.
	ServerBinding NodeTemplateServerBindingPtrInput
}

func (NodeTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nodeTemplateArgs)(nil)).Elem()
}
