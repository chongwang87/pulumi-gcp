// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Notebooks.Outputs
{

    [OutputType]
    public sealed class InstanceVmImage
    {
        /// <summary>
        /// Use this VM image family to find the image; the newest image in this family will be used.
        /// </summary>
        public readonly string? ImageFamily;
        /// <summary>
        /// Use VM image name to find the image.
        /// </summary>
        public readonly string? ImageName;
        /// <summary>
        /// The name of the Google Cloud project that this VM image belongs to.
        /// Format: projects/{project_id}
        /// </summary>
        public readonly string Project;

        [OutputConstructor]
        private InstanceVmImage(
            string? imageFamily,

            string? imageName,

            string project)
        {
            ImageFamily = imageFamily;
            ImageName = imageName;
            Project = project;
        }
    }
}
