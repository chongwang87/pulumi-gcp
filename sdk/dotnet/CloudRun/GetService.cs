// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun
{
    public static class GetService
    {
        /// <summary>
        /// Get information about a Google Cloud Run. For more information see
        /// the [official documentation](https://cloud.google.com/run/docs/)
        /// and [API](https://cloud.google.com/run/docs/apis).
        /// </summary>
        public static Task<GetServiceResult> InvokeAsync(GetServiceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceResult>("gcp:cloudrun/getService:getService", args ?? new GetServiceArgs(), options.WithVersion());
    }


    public sealed class GetServiceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The location of the cloud run instance. eg us-central1
        /// </summary>
        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        /// <summary>
        /// The name of the Cloud Run Service.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        public GetServiceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceResult
    {
        public readonly bool AutogenerateRevisionName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Location;
        public readonly ImmutableArray<Outputs.GetServiceMetadataResult> Metadatas;
        public readonly string Name;
        public readonly string? Project;
        public readonly ImmutableArray<Outputs.GetServiceStatusResult> Statuses;
        public readonly ImmutableArray<Outputs.GetServiceTemplateResult> Templates;
        public readonly ImmutableArray<Outputs.GetServiceTrafficResult> Traffics;

        [OutputConstructor]
        private GetServiceResult(
            bool autogenerateRevisionName,

            string id,

            string location,

            ImmutableArray<Outputs.GetServiceMetadataResult> metadatas,

            string name,

            string? project,

            ImmutableArray<Outputs.GetServiceStatusResult> statuses,

            ImmutableArray<Outputs.GetServiceTemplateResult> templates,

            ImmutableArray<Outputs.GetServiceTrafficResult> traffics)
        {
            AutogenerateRevisionName = autogenerateRevisionName;
            Id = id;
            Location = location;
            Metadatas = metadatas;
            Name = name;
            Project = project;
            Statuses = statuses;
            Templates = templates;
            Traffics = traffics;
        }
    }
}
