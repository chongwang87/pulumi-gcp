// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub.Inputs
{

    public sealed class LiteTopicRetentionConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The provisioned storage, in bytes, per partition. If the number of bytes stored
        /// in any of the topic's partitions grows beyond this value, older messages will be
        /// dropped to make room for newer ones, regardless of the value of period.
        /// </summary>
        [Input("perPartitionBytes", required: true)]
        public Input<string> PerPartitionBytes { get; set; } = null!;

        /// <summary>
        /// How long a published message is retained. If unset, messages will be retained as
        /// long as the bytes retained for each partition is below perPartitionBytes.
        /// </summary>
        [Input("period")]
        public Input<string>? Period { get; set; }

        public LiteTopicRetentionConfigGetArgs()
        {
        }
    }
}