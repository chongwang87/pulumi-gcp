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
    /// Enables the Google Compute Engine
    /// [Shared VPC](https://cloud.google.com/compute/docs/shared-vpc)
    /// feature for a project, assigning it as a Shared VPC service project associated
    /// with a given host project.
    /// 
    /// For more information, see,
    /// [the Project API documentation](https://cloud.google.com/compute/docs/reference/latest/projects),
    /// where the Shared VPC feature is referred to by its former name "XPN".
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var service1 = new Gcp.Compute.SharedVPCServiceProject("service1", new Gcp.Compute.SharedVPCServiceProjectArgs
    ///         {
    ///             HostProject = "host-project-id",
    ///             ServiceProject = "service-project-id-1",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// For a complete Shared VPC example with both host and service projects, see
    /// [`gcp.compute.SharedVPCHostProject`](https://www.terraform.io/docs/providers/google/r/compute_shared_vpc_host_project.html).
    /// </summary>
    public partial class SharedVPCServiceProject : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of a host project to associate.
        /// </summary>
        [Output("hostProject")]
        public Output<string> HostProject { get; private set; } = null!;

        /// <summary>
        /// The ID of the project that will serve as a Shared VPC service project.
        /// </summary>
        [Output("serviceProject")]
        public Output<string> ServiceProject { get; private set; } = null!;


        /// <summary>
        /// Create a SharedVPCServiceProject resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SharedVPCServiceProject(string name, SharedVPCServiceProjectArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/sharedVPCServiceProject:SharedVPCServiceProject", name, args ?? new SharedVPCServiceProjectArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SharedVPCServiceProject(string name, Input<string> id, SharedVPCServiceProjectState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/sharedVPCServiceProject:SharedVPCServiceProject", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SharedVPCServiceProject resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SharedVPCServiceProject Get(string name, Input<string> id, SharedVPCServiceProjectState? state = null, CustomResourceOptions? options = null)
        {
            return new SharedVPCServiceProject(name, id, state, options);
        }
    }

    public sealed class SharedVPCServiceProjectArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of a host project to associate.
        /// </summary>
        [Input("hostProject", required: true)]
        public Input<string> HostProject { get; set; } = null!;

        /// <summary>
        /// The ID of the project that will serve as a Shared VPC service project.
        /// </summary>
        [Input("serviceProject", required: true)]
        public Input<string> ServiceProject { get; set; } = null!;

        public SharedVPCServiceProjectArgs()
        {
        }
    }

    public sealed class SharedVPCServiceProjectState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of a host project to associate.
        /// </summary>
        [Input("hostProject")]
        public Input<string>? HostProject { get; set; }

        /// <summary>
        /// The ID of the project that will serve as a Shared VPC service project.
        /// </summary>
        [Input("serviceProject")]
        public Input<string>? ServiceProject { get; set; }

        public SharedVPCServiceProjectState()
        {
        }
    }
}
