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
    /// Persistent disks can be attached to a compute instance using the `attached_disk`
    /// section within the compute instance configuration.
    /// However there may be situations where managing the attached disks via the compute
    /// instance config isn't preferable or possible, such as attaching dynamic
    /// numbers of disks using the `count` variable.
    /// 
    /// 
    /// To get more information about attaching disks, see:
    /// 
    /// * [API documentation](https://cloud.google.com/compute/docs/reference/rest/v1/instances/attachDisk)
    /// * How-to Guides
    ///     * [Adding a persistent disk](https://cloud.google.com/compute/docs/disks/add-persistent-disk)
    /// 
    /// **Note:** When using `gcp.compute.AttachedDisk` you **must** use `lifecycle.ignore_changes = ["attached_disk"]` on the `gcp.compute.Instance` resource that has the disks attached. Otherwise the two resources will fight for control of the attached disk block.
    /// 
    /// 
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/r/compute_attached_disk.html.markdown.
    /// </summary>
    public partial class AttachedDisk : Pulumi.CustomResource
    {
        [Output("deviceName")]
        public Output<string> DeviceName { get; private set; } = null!;

        [Output("disk")]
        public Output<string> Disk { get; private set; } = null!;

        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("zone")]
        public Output<string> Zone { get; private set; } = null!;


        /// <summary>
        /// Create a AttachedDisk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AttachedDisk(string name, AttachedDiskArgs args, CustomResourceOptions? options = null)
            : base("gcp:compute/attachedDisk:AttachedDisk", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private AttachedDisk(string name, Input<string> id, AttachedDiskState? state = null, CustomResourceOptions? options = null)
            : base("gcp:compute/attachedDisk:AttachedDisk", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AttachedDisk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AttachedDisk Get(string name, Input<string> id, AttachedDiskState? state = null, CustomResourceOptions? options = null)
        {
            return new AttachedDisk(name, id, state, options);
        }
    }

    public sealed class AttachedDiskArgs : Pulumi.ResourceArgs
    {
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        [Input("disk", required: true)]
        public Input<string> Disk { get; set; } = null!;

        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AttachedDiskArgs()
        {
        }
    }

    public sealed class AttachedDiskState : Pulumi.ResourceArgs
    {
        [Input("deviceName")]
        public Input<string>? DeviceName { get; set; }

        [Input("disk")]
        public Input<string>? Disk { get; set; }

        [Input("instance")]
        public Input<string>? Instance { get; set; }

        [Input("mode")]
        public Input<string>? Mode { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("zone")]
        public Input<string>? Zone { get; set; }

        public AttachedDiskState()
        {
        }
    }
}
