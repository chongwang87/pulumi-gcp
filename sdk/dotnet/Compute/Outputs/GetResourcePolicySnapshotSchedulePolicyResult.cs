// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Compute.Outputs
{

    [OutputType]
    public sealed class GetResourcePolicySnapshotSchedulePolicyResult
    {
        public readonly ImmutableArray<Outputs.GetResourcePolicySnapshotSchedulePolicyRetentionPolicyResult> RetentionPolicies;
        public readonly ImmutableArray<Outputs.GetResourcePolicySnapshotSchedulePolicyScheduleResult> Schedules;
        public readonly ImmutableArray<Outputs.GetResourcePolicySnapshotSchedulePolicySnapshotPropertyResult> SnapshotProperties;

        [OutputConstructor]
        private GetResourcePolicySnapshotSchedulePolicyResult(
            ImmutableArray<Outputs.GetResourcePolicySnapshotSchedulePolicyRetentionPolicyResult> retentionPolicies,

            ImmutableArray<Outputs.GetResourcePolicySnapshotSchedulePolicyScheduleResult> schedules,

            ImmutableArray<Outputs.GetResourcePolicySnapshotSchedulePolicySnapshotPropertyResult> snapshotProperties)
        {
            RetentionPolicies = retentionPolicies;
            Schedules = schedules;
            SnapshotProperties = snapshotProperties;
        }
    }
}
