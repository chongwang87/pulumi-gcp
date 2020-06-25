// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute
{
    /// <summary>
    /// Represents a NodeGroup resource to manage a group of sole-tenant nodes.
    /// 
    /// To get more information about NodeGroup, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/nodeGroups)
    /// * How-to Guides
    ///     * [Sole-Tenant Nodes](https://cloud.google.com/compute/docs/nodes/)
    /// 
    /// &gt; **Warning:** Due to limitations of the API, this provider cannot update the
    /// number of nodes in a node group and changes to node group size either
    /// through provider config or through external changes will cause
    /// the provider to delete and recreate the node group.
    /// 
    /// ## Example Usage
    /// ### Node Group Basic
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var soletenant_tmpl = new Gcp.Compute.NodeTemplate("soletenant-tmpl", new Gcp.Compute.NodeTemplateArgs
    ///         {
    ///             Region = "us-central1",
    ///             NodeType = "n1-node-96-624",
    ///         });
    ///         var nodes = new Gcp.Compute.NodeGroup("nodes", new Gcp.Compute.NodeGroupArgs
    ///         {
    ///             Zone = "us-central1-a",
    ///             Description = "example google_compute_node_group for the Google Provider",
    ///             Size = 1,
    ///             NodeTemplate = soletenant_tmpl.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class NodeGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// -
        /// If you use sole-tenant nodes for your workloads, you can use the node
        /// group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
        /// </summary>
        [Output("autoscalingPolicy")]
        public Output<Outputs.NodeGroupAutoscalingPolicy> AutoscalingPolicy { get; private set; } = null!;

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Output("creationTimestamp")]
        public Output<string> CreationTimestamp { get; private set; } = null!;

        /// <summary>
        /// An optional textual description of the resource.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The URL of the node template to which this node group belongs.
        /// </summary>
        [Output("nodeTemplate")]
        public Output<string> NodeTemplate { get; private set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Output("selfLink")]
        public Output<string> SelfLink { get; private set; } = null!;

        /// <summary>
        /// The total number of nodes in the node group.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// Zone where this node group is located
        /// </summary>
        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a NodeGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NodeGroup(string name, NodeGroupArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/nodeGroup:NodeGroup", name, args ?? new NodeGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NodeGroup(string name, Input<string> id, NodeGroupState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/nodeGroup:NodeGroup", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NodeGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NodeGroup Get(string name, Input<string> id, NodeGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new NodeGroup(name, id, state, options);
        }
    }

    public sealed class NodeGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// If you use sole-tenant nodes for your workloads, you can use the node
        /// group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
        /// </summary>
        [Input("autoscalingPolicy")]
        public Input<Inputs.NodeGroupAutoscalingPolicyArgs>? AutoscalingPolicy { get; set; }

        /// <summary>
        /// An optional textual description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The URL of the node template to which this node group belongs.
        /// </summary>
        [Input("nodeTemplate", required: true)]
        public Input<string> NodeTemplate { get; set; } = null!;

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The total number of nodes in the node group.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        /// <summary>
        /// Zone where this node group is located
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public NodeGroupArgs()
        {
        }
    }

    public sealed class NodeGroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// -
        /// If you use sole-tenant nodes for your workloads, you can use the node
        /// group autoscaler to automatically manage the sizes of your node groups.  Structure is documented below.
        /// </summary>
        [Input("autoscalingPolicy")]
        public Input<Inputs.NodeGroupAutoscalingPolicyGetArgs>? AutoscalingPolicy { get; set; }

        /// <summary>
        /// Creation timestamp in RFC3339 text format.
        /// </summary>
        [Input("creationTimestamp")]
        public Input<string>? CreationTimestamp { get; set; }

        /// <summary>
        /// An optional textual description of the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The URL of the node template to which this node group belongs.
        /// </summary>
        [Input("nodeTemplate")]
        public Input<string>? NodeTemplate { get; set; }

        /// <summary>
        /// The ID of the project in which the resource belongs.
        /// If it is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The URI of the created resource.
        /// </summary>
        [Input("selfLink")]
        public Input<string>? SelfLink { get; set; }

        /// <summary>
        /// The total number of nodes in the node group.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Zone where this node group is located
        /// </summary>
        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public NodeGroupState()
        {
        }
    }
}
