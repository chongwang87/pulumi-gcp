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
    public sealed class GetEnvironmentConfigPrivateEnvironmentConfigResult
    {
        public readonly string CloudSqlIpv4CidrBlock;
        public readonly bool EnablePrivateEndpoint;
        public readonly string MasterIpv4CidrBlock;
        public readonly string WebServerIpv4CidrBlock;

        [OutputConstructor]
        private GetEnvironmentConfigPrivateEnvironmentConfigResult(
            string cloudSqlIpv4CidrBlock,

            bool enablePrivateEndpoint,

            string masterIpv4CidrBlock,

            string webServerIpv4CidrBlock)
        {
            CloudSqlIpv4CidrBlock = cloudSqlIpv4CidrBlock;
            EnablePrivateEndpoint = enablePrivateEndpoint;
            MasterIpv4CidrBlock = masterIpv4CidrBlock;
            WebServerIpv4CidrBlock = webServerIpv4CidrBlock;
        }
    }
}
