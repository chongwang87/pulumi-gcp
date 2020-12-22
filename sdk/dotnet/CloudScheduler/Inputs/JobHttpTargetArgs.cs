// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudScheduler.Inputs
{

    public sealed class JobHttpTargetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// HTTP request body.
        /// A request body is allowed only if the HTTP method is POST, PUT, or PATCH.
        /// It is an error to set body on a job with an incompatible HttpMethod.
        /// A base64-encoded string.
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        [Input("headers")]
        private InputMap<string>? _headers;

        /// <summary>
        /// This map contains the header field names and values.
        /// Repeated headers are not supported, but a header value can contain commas.
        /// </summary>
        public InputMap<string> Headers
        {
            get => _headers ?? (_headers = new InputMap<string>());
            set => _headers = value;
        }

        /// <summary>
        /// Which HTTP method to use for the request.
        /// </summary>
        [Input("httpMethod")]
        public Input<string>? HttpMethod { get; set; }

        /// <summary>
        /// Contains information needed for generating an OAuth token.
        /// This type of authorization should be used when sending requests to a GCP endpoint.
        /// Structure is documented below.
        /// </summary>
        [Input("oauthToken")]
        public Input<Inputs.JobHttpTargetOauthTokenArgs>? OauthToken { get; set; }

        /// <summary>
        /// Contains information needed for generating an OpenID Connect token.
        /// This type of authorization should be used when sending requests to third party endpoints or Cloud Run.
        /// Structure is documented below.
        /// </summary>
        [Input("oidcToken")]
        public Input<Inputs.JobHttpTargetOidcTokenArgs>? OidcToken { get; set; }

        /// <summary>
        /// The full URI path that the request will be sent to.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        public JobHttpTargetArgs()
        {
        }
    }
}
