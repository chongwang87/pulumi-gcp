// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Iam
{
    public static class GetWorkloadIdentityPool
    {
        public static Task<GetWorkloadIdentityPoolResult> InvokeAsync(GetWorkloadIdentityPoolArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetWorkloadIdentityPoolResult>("gcp:iam/getWorkloadIdentityPool:getWorkloadIdentityPool", args ?? new GetWorkloadIdentityPoolArgs(), options.WithVersion());
    }


    public sealed class GetWorkloadIdentityPoolArgs : Pulumi.InvokeArgs
    {
        [Input("description")]
        public string? Description { get; set; }

        [Input("disabled")]
        public bool? Disabled { get; set; }

        [Input("displayName")]
        public string? DisplayName { get; set; }

        /// <summary>
        /// The project in which the resource belongs. If it
        /// is not provided, the provider project is used.
        /// </summary>
        [Input("project")]
        public string? Project { get; set; }

        /// <summary>
        /// The id of the pool which is the
        /// final component of the resource name.
        /// </summary>
        [Input("workloadIdentityPoolId", required: true)]
        public string WorkloadIdentityPoolId { get; set; } = null!;

        public GetWorkloadIdentityPoolArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetWorkloadIdentityPoolResult
    {
        public readonly string? Description;
        public readonly bool? Disabled;
        public readonly string? DisplayName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string? Project;
        public readonly string State;
        public readonly string WorkloadIdentityPoolId;

        [OutputConstructor]
        private GetWorkloadIdentityPoolResult(
            string? description,

            bool? disabled,

            string? displayName,

            string id,

            string name,

            string? project,

            string state,

            string workloadIdentityPoolId)
        {
            Description = description;
            Disabled = disabled;
            DisplayName = displayName;
            Id = id;
            Name = name;
            Project = project;
            State = state;
            WorkloadIdentityPoolId = workloadIdentityPoolId;
        }
    }
}