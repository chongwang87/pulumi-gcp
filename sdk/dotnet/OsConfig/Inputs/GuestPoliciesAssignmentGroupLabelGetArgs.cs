// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.OsConfig.Inputs
{

    public sealed class GuestPoliciesAssignmentGroupLabelGetArgs : Pulumi.ResourceArgs
    {
        [Input("labels", required: true)]
        private InputMap<string>? _labels;

        /// <summary>
        /// Google Compute Engine instance labels that must be present for an instance to be included in this assignment group.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        public GuestPoliciesAssignmentGroupLabelGetArgs()
        {
        }
    }
}