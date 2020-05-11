// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Sql.Inputs
{

    public sealed class DatabaseInstanceSettingsDatabaseFlagGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name for this whitelist entry.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// A CIDR notation IPv4 or IPv6 address that is allowed to
        /// access this instance. Must be set even if other two attributes are not for
        /// the whitelist to become active.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public DatabaseInstanceSettingsDatabaseFlagGetArgs()
        {
        }
    }
}