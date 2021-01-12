// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Container.Outputs
{

    [OutputType]
    public sealed class ClusterClusterAutoscalingAutoProvisioningDefaults
    {
        /// <summary>
        /// Minimum CPU platform to be used by this instance.
        /// The instance may be scheduled on the specified or newer CPU platform. Applicable
        /// values are the friendly names of CPU platforms, such as `Intel Haswell`. See the
        /// [official documentation](https://cloud.google.com/compute/docs/instances/specify-min-cpu-platform)
        /// for more information.
        /// </summary>
        public readonly string? MinCpuPlatform;
        /// <summary>
        /// The set of Google API scopes to be made available
        /// on all of the node VMs under the "default" service account.
        /// Use the "https://www.googleapis.com/auth/cloud-platform" scope to grant access to all APIs. It is recommended that you set `service_account` to a non-default service account and grant IAM roles to that service account for only the resources that it needs.
        /// </summary>
        public readonly ImmutableArray<string> OauthScopes;
        /// <summary>
        /// The service account to be used by the Node VMs.
        /// If not specified, the "default" service account is used.
        /// </summary>
        public readonly string? ServiceAccount;

        [OutputConstructor]
        private ClusterClusterAutoscalingAutoProvisioningDefaults(
            string? minCpuPlatform,

            ImmutableArray<string> oauthScopes,

            string? serviceAccount)
        {
            MinCpuPlatform = minCpuPlatform;
            OauthScopes = oauthScopes;
            ServiceAccount = serviceAccount;
        }
    }
}
