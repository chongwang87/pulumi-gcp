// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Gcp.Firestore.Inputs
{

    public sealed class IndexFieldArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates that this field supports operations on arrayValues. Only one of `order` and `arrayConfig` can
        /// be specified.
        /// </summary>
        [Input("arrayConfig")]
        public Input<string>? ArrayConfig { get; set; }

        /// <summary>
        /// Name of the field.
        /// </summary>
        [Input("fieldPath")]
        public Input<string>? FieldPath { get; set; }

        /// <summary>
        /// Indicates that this field supports ordering by the specified order or comparing using =, &lt;, &lt;=, &gt;, &gt;=.
        /// Only one of `order` and `arrayConfig` can be specified.
        /// </summary>
        [Input("order")]
        public Input<string>? Order { get; set; }

        public IndexFieldArgs()
        {
        }
    }
}