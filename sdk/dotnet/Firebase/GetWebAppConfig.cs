// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Firebase
{
    public static class GetWebAppConfig
    {
        /// <summary>
        /// A Google Cloud Firebase web application configuration
        /// 
        /// To get more information about WebApp, see:
        /// 
        /// * [API documentation](https://firebase.google.com/docs/projects/api/reference/rest/v1beta1/projects.webApps)
        /// * How-to Guides
        ///     * [Official Documentation](https://firebase.google.com/)
        /// </summary>
        public static Task<GetWebAppConfigResult> InvokeAsync(GetWebAppConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWebAppConfigResult>("gcp:firebase/getWebAppConfig:getWebAppConfig", args ?? new GetWebAppConfigArgs(), options.WithVersion());
    }


    public sealed class GetWebAppConfigArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// the id of the firebase web app
        /// </summary>
        [Input("webAppId", required: true)]
        public string WebAppId { get; set; } = null!;

        public GetWebAppConfigArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWebAppConfigResult
    {
        public readonly string ApiKey;
        public readonly string AuthDomain;
        public readonly string DatabaseUrl;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string LocationId;
        public readonly string MeasurementId;
        public readonly string MessagingSenderId;
        public readonly string? Project;
        public readonly string StorageBucket;
        public readonly string WebAppId;

        [OutputConstructor]
        private GetWebAppConfigResult(
            string apiKey,

            string authDomain,

            string databaseUrl,

            string id,

            string locationId,

            string measurementId,

            string messagingSenderId,

            string? project,

            string storageBucket,

            string webAppId)
        {
            ApiKey = apiKey;
            AuthDomain = authDomain;
            DatabaseUrl = databaseUrl;
            Id = id;
            LocationId = locationId;
            MeasurementId = measurementId;
            MessagingSenderId = messagingSenderId;
            Project = project;
            StorageBucket = storageBucket;
            WebAppId = webAppId;
        }
    }
}