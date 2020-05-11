// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Inputs
{

    public sealed class ClusterWorkloadIdentityConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Currently, the only supported identity namespace is the project's default.
        /// </summary>
        [Input("identityNamespace", required: true)]
        public Input<string> IdentityNamespace { get; set; } = null!;

        public ClusterWorkloadIdentityConfigArgs()
        {
        }
    }
}