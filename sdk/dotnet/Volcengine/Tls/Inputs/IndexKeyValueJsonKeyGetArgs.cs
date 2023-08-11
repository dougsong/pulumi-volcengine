// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Tls.Inputs
{

    public sealed class IndexKeyValueJsonKeyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key of the subfield key value index.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The type of value. Valid values: `long`, `double`, `text`.
        /// </summary>
        [Input("valueType", required: true)]
        public Input<string> ValueType { get; set; } = null!;

        public IndexKeyValueJsonKeyGetArgs()
        {
        }
        public static new IndexKeyValueJsonKeyGetArgs Empty => new IndexKeyValueJsonKeyGetArgs();
    }
}
