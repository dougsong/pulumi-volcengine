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

    public sealed class NodeKubernetesConfigLabelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Key of Labels.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// The Value of Labels.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public NodeKubernetesConfigLabelArgs()
        {
        }
        public static new NodeKubernetesConfigLabelArgs Empty => new NodeKubernetesConfigLabelArgs();
    }
}
