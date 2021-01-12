// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub.Inputs
{

    public sealed class LiteTopicPartitionConfigGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The capacity configuration.
        /// Structure is documented below.
        /// </summary>
        [Input("capacity")]
        public Input<Inputs.LiteTopicPartitionConfigCapacityGetArgs>? Capacity { get; set; }

        /// <summary>
        /// The number of partitions in the topic. Must be at least 1.
        /// </summary>
        [Input("count", required: true)]
        public Input<int> Count { get; set; } = null!;

        public LiteTopicPartitionConfigGetArgs()
        {
        }
    }
}
