// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Veenedge.Inputs
{

    public sealed class CloudServerNetworkConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The peak of bandwidth.
        /// </summary>
        [Input("bandwidthPeak", required: true)]
        public Input<string> BandwidthPeak { get; set; } = null!;

        /// <summary>
        /// The name of custom external interface.
        /// </summary>
        [Input("customExternalInterfaceName")]
        public Input<string>? CustomExternalInterfaceName { get; set; }

        /// <summary>
        /// The name of custom internal interface.
        /// </summary>
        [Input("customInternalInterfaceName")]
        public Input<string>? CustomInternalInterfaceName { get; set; }

        /// <summary>
        /// Whether enable ipv6.
        /// </summary>
        [Input("enableIpv6")]
        public Input<bool>? EnableIpv6 { get; set; }

        /// <summary>
        /// The internal peak of bandwidth.
        /// </summary>
        [Input("internalBandwidthPeak")]
        public Input<string>? InternalBandwidthPeak { get; set; }

        public CloudServerNetworkConfigGetArgs()
        {
        }
        public static new CloudServerNetworkConfigGetArgs Empty => new CloudServerNetworkConfigGetArgs();
    }
}
