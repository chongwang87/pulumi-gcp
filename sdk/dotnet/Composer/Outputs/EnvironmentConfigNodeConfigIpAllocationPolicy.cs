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
    public sealed class EnvironmentConfigNodeConfigIpAllocationPolicy
    {
        /// <summary>
        /// The IP address range used to allocate IP addresses to pods in the cluster.
        /// Set to blank to have GKE choose a range with the default size.
        /// Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask.
        /// Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks
        /// (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
        /// Specify either `cluster_secondary_range_name` or `cluster_ipv4_cidr_block` but not both.
        /// </summary>
        public readonly string? ClusterIpv4CidrBlock;
        /// <summary>
        /// The name of the cluster's secondary range used to allocate IP addresses to pods.
        /// Specify either `cluster_secondary_range_name` or `cluster_ipv4_cidr_block` but not both.
        /// This field is applicable only when `use_ip_aliases` is true.
        /// </summary>
        public readonly string? ClusterSecondaryRangeName;
        /// <summary>
        /// The IP address range used to allocate IP addresses in this cluster.
        /// Set to blank to have GKE choose a range with the default size.
        /// Set to /netmask (e.g. /14) to have GKE choose a range with a specific netmask.
        /// Set to a CIDR notation (e.g. 10.96.0.0/14) from the RFC-1918 private networks
        /// (e.g. 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16) to pick a specific range to use.
        /// Specify either `services_secondary_range_name` or `services_ipv4_cidr_block` but not both.
        /// </summary>
        public readonly string? ServicesIpv4CidrBlock;
        /// <summary>
        /// The name of the services' secondary range used to allocate IP addresses to the cluster.
        /// Specify either `services_secondary_range_name` or `services_ipv4_cidr_block` but not both.
        /// This field is applicable only when `use_ip_aliases` is true.
        /// </summary>
        public readonly string? ServicesSecondaryRangeName;
        /// <summary>
        /// Whether or not to enable Alias IPs in the GKE cluster. If true, a VPC-native cluster is created.
        /// Defaults to true if the `ip_allocation_policy` block is present in config.
        /// </summary>
        public readonly bool UseIpAliases;

        [OutputConstructor]
        private EnvironmentConfigNodeConfigIpAllocationPolicy(
            string? clusterIpv4CidrBlock,

            string? clusterSecondaryRangeName,

            string? servicesIpv4CidrBlock,

            string? servicesSecondaryRangeName,

            bool useIpAliases)
        {
            ClusterIpv4CidrBlock = clusterIpv4CidrBlock;
            ClusterSecondaryRangeName = clusterSecondaryRangeName;
            ServicesIpv4CidrBlock = servicesIpv4CidrBlock;
            ServicesSecondaryRangeName = servicesSecondaryRangeName;
            UseIpAliases = useIpAliases;
        }
    }
}
