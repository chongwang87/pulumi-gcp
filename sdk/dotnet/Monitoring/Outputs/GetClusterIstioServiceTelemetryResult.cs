// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Monitoring.Outputs
{

    [OutputType]
    public sealed class GetClusterIstioServiceTelemetryResult
    {
        public readonly string ResourceName;

        [OutputConstructor]
        private GetClusterIstioServiceTelemetryResult(string resourceName)
        {
            ResourceName = resourceName;
        }
    }
}
