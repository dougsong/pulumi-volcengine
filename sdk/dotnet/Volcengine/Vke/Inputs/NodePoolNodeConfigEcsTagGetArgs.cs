// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vke.Inputs
{

    public sealed class NodePoolNodeConfigEcsTagGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Key of Tags.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        /// <summary>
        /// The Value of Tags.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public NodePoolNodeConfigEcsTagGetArgs()
        {
        }
        public static new NodePoolNodeConfigEcsTagGetArgs Empty => new NodePoolNodeConfigEcsTagGetArgs();
    }
}
