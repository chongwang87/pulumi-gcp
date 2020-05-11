// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Dataproc.Outputs
{

    [OutputType]
    public sealed class ClusterClusterConfigEncryptionConfig
    {
        /// <summary>
        /// The Cloud KMS key name to use for PD disk encryption for
        /// all instances in the cluster.
        /// </summary>
        public readonly string KmsKeyName;

        [OutputConstructor]
        private ClusterClusterConfigEncryptionConfig(string kmsKeyName)
        {
            KmsKeyName = kmsKeyName;
        }
    }
}