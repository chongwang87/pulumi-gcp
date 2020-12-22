// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Inputs
{

    public sealed class RegionBackendServiceConsistentHashHttpCookieArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the cookie.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Path to set for the cookie.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("ttl")]
        public Input<Inputs.RegionBackendServiceConsistentHashHttpCookieTtlArgs>? Ttl { get; set; }

        public RegionBackendServiceConsistentHashHttpCookieArgs()
        {
        }
    }
}
