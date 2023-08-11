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

    public sealed class ClusterPodsConfigFlannelConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum number of single-node Pod instances for a Flannel container network, the value can be `16` or `32` or `64` or `128` or `256`.
        /// </summary>
        [Input("maxPodsPerNode")]
        public Input<int>? MaxPodsPerNode { get; set; }

        [Input("podCidrs")]
        private InputList<string>? _podCidrs;

        /// <summary>
        /// Pod CIDR for the Flannel container network.
        /// </summary>
        public InputList<string> PodCidrs
        {
            get => _podCidrs ?? (_podCidrs = new InputList<string>());
            set => _podCidrs = value;
        }

        public ClusterPodsConfigFlannelConfigArgs()
        {
        }
        public static new ClusterPodsConfigFlannelConfigArgs Empty => new ClusterPodsConfigFlannelConfigArgs();
    }
}
