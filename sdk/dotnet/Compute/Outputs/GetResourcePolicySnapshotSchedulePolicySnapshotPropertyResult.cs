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
    public sealed class GetResourcePolicySnapshotSchedulePolicySnapshotPropertyResult
    {
        public readonly bool GuestFlush;
        public readonly ImmutableDictionary<string, string> Labels;
        public readonly ImmutableArray<string> StorageLocations;

        [OutputConstructor]
        private GetResourcePolicySnapshotSchedulePolicySnapshotPropertyResult(
            bool guestFlush,

            ImmutableDictionary<string, string> labels,

            ImmutableArray<string> storageLocations)
        {
            GuestFlush = guestFlush;
            Labels = labels;
            StorageLocations = storageLocations;
        }
    }
}
