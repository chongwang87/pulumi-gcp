// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class BackendBucketCdnPolicyNegativeCachingPolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("code")]
        public Input<int>? Code { get; set; }

        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public BackendBucketCdnPolicyNegativeCachingPolicyGetArgs()
        {
        }
    }
}
