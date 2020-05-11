// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Storage.Outputs
{

    [OutputType]
    public sealed class TransferJobTransferSpecAwsS3DataSourceAwsAccessKey
    {
        /// <summary>
        /// AWS Key ID.
        /// </summary>
        public readonly string AccessKeyId;
        /// <summary>
        /// AWS Secret Access Key.
        /// </summary>
        public readonly string SecretAccessKey;

        [OutputConstructor]
        private TransferJobTransferSpecAwsS3DataSourceAwsAccessKey(
            string accessKeyId,

            string secretAccessKey)
        {
            AccessKeyId = accessKeyId;
            SecretAccessKey = secretAccessKey;
        }
    }
}