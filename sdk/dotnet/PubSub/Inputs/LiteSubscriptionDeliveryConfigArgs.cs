// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.PubSub.Inputs
{

    public sealed class LiteSubscriptionDeliveryConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When this subscription should send messages to subscribers relative to messages persistence in storage.
        /// Possible values are `DELIVER_IMMEDIATELY`, `DELIVER_AFTER_STORED`, and `DELIVERY_REQUIREMENT_UNSPECIFIED`.
        /// </summary>
        [Input("deliveryRequirement", required: true)]
        public Input<string> DeliveryRequirement { get; set; } = null!;

        public LiteSubscriptionDeliveryConfigArgs()
        {
        }
    }
}
