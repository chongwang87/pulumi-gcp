// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Redis
{
    public static class GetInstance
    {
        /// <summary>
        /// Get information about a Google Cloud Redis instance. For more information see
        /// the [official documentation](https://cloud.google.com/memorystore/docs/redis)
        /// and [API](https://cloud.google.com/memorystore/docs/redis/apis).
        /// </summary>
        public static Task<GetInstanceResult> InvokeAsync(GetInstanceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetInstanceResult>("gcp:redis/getInstance:getInstance", args ?? new GetInstanceArgs(), options.WithVersion());
    }


    public sealed class GetInstanceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of a Redis instance.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The region in which the resource belongs. If it
        /// is not provided, the provider region is used.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        public GetInstanceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetInstanceResult
    {
        public readonly string AlternativeLocationId;
        public readonly string AuthorizedNetwork;
        public readonly string ConnectMode;
        public readonly string CreateTime;
        public readonly string CurrentLocationId;
        public readonly string DisplayName;
        /// <summary>
        /// Hostname or IP address of the exposed Redis endpoint used by clients
        /// to connect to the service.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly string LocationId;
        public readonly int MemorySizeGb;
        public readonly string Name;
        /// <summary>
        /// The port number of the exposed Redis endpoint.
        /// </summary>
        public readonly int Port;
        public readonly string? Project;
        public readonly ImmutableDictionary<string, string> RedisConfigs;
        public readonly string RedisVersion;
        public readonly string? Region;
        public readonly string ReservedIpRange;
        public readonly string Tier;

        [OutputConstructor]
        private GetInstanceResult(
            string alternativeLocationId,

            string authorizedNetwork,

            string connectMode,

            string createTime,

            string currentLocationId,

            string displayName,

            string host,

            string id,

            ImmutableDictionary<string, string> labels,

            string locationId,

            int memorySizeGb,

            string name,

            int port,

            string? project,

            ImmutableDictionary<string, string> redisConfigs,

            string redisVersion,

            string? region,

            string reservedIpRange,

            string tier)
        {
            AlternativeLocationId = alternativeLocationId;
            AuthorizedNetwork = authorizedNetwork;
            ConnectMode = connectMode;
            CreateTime = createTime;
            CurrentLocationId = currentLocationId;
            DisplayName = displayName;
            Host = host;
            Id = id;
            Labels = labels;
            LocationId = locationId;
            MemorySizeGb = memorySizeGb;
            Name = name;
            Port = port;
            Project = project;
            RedisConfigs = redisConfigs;
            RedisVersion = redisVersion;
            Region = region;
            ReservedIpRange = reservedIpRange;
            Tier = tier;
        }
    }
}
