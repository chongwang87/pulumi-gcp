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
    public sealed class ClusterMasterAuth
    {
        public readonly string? ClientCertificate;
        /// <summary>
        /// Whether client certificate authorization is enabled for this cluster.  For example:
        /// </summary>
        public readonly Outputs.ClusterMasterAuthClientCertificateConfig? ClientCertificateConfig;
        public readonly string? ClientKey;
        public readonly string? ClusterCaCertificate;
        /// <summary>
        /// The password to use for HTTP basic authentication when accessing
        /// the Kubernetes master endpoint. This has been deprecated as of GKE 1.19.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// The username to use for HTTP basic authentication when accessing
        /// the Kubernetes master endpoint. If not present basic auth will be disabled. This has been deprecated as of GKE 1.19.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private ClusterMasterAuth(
            string? clientCertificate,

            Outputs.ClusterMasterAuthClientCertificateConfig? clientCertificateConfig,

            string? clientKey,

            string? clusterCaCertificate,

            string? password,

            string? username)
        {
            ClientCertificate = clientCertificate;
            ClientCertificateConfig = clientCertificateConfig;
            ClientKey = clientKey;
            ClusterCaCertificate = clusterCaCertificate;
            Password = password;
            Username = username;
        }
    }
}
