// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Inputs
{

    public sealed class ClusterClusterConfigPreemptibleWorkerConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disk Config
        /// </summary>
        [Input("diskConfig")]
        public Input<Inputs.ClusterClusterConfigPreemptibleWorkerConfigDiskConfigGetArgs>? DiskConfig { get; set; }

        [Input("instanceNames")]
        private InputList<string>? _instanceNames;
        public InputList<string> InstanceNames
        {
            get => _instanceNames ?? (_instanceNames = new InputList<string>());
            set => _instanceNames = value;
        }

        /// <summary>
        /// Specifies the number of preemptible nodes to create.
        /// Defaults to 0.
        /// </summary>
        [Input("numInstances")]
        public Input<int>? NumInstances { get; set; }

        public ClusterClusterConfigPreemptibleWorkerConfigGetArgs()
        {
        }
    }
}
