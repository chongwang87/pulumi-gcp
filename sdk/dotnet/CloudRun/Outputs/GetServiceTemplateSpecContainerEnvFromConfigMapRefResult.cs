// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.CloudRun.Outputs
{

    [OutputType]
    public sealed class GetServiceTemplateSpecContainerEnvFromConfigMapRefResult
    {
        public readonly ImmutableArray<Outputs.GetServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReferenceResult> LocalObjectReferences;
        public readonly bool Optional;

        [OutputConstructor]
        private GetServiceTemplateSpecContainerEnvFromConfigMapRefResult(
            ImmutableArray<Outputs.GetServiceTemplateSpecContainerEnvFromConfigMapRefLocalObjectReferenceResult> localObjectReferences,

            bool optional)
        {
            LocalObjectReferences = localObjectReferences;
            Optional = optional;
        }
    }
}