// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Composer.Outputs
{

    [OutputType]
    public sealed class GetEnvironmentConfigWebServerConfigResult
    {
        public readonly string MachineType;

        [OutputConstructor]
        private GetEnvironmentConfigWebServerConfigResult(string machineType)
        {
            MachineType = machineType;
        }
    }
}
