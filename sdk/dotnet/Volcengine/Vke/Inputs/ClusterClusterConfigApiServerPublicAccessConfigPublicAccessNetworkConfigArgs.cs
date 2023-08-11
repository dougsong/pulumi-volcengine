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

    public sealed class ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The peak bandwidth of the public IP, unit: Mbps.
        /// </summary>
        [Input("bandwidth")]
        public Input<int>? Bandwidth { get; set; }

        /// <summary>
        /// Billing type of public IP, the value is `PostPaidByBandwidth` or `PostPaidByTraffic`.
        /// </summary>
        [Input("billingType")]
        public Input<string>? BillingType { get; set; }

        public ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs()
        {
        }
        public static new ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs Empty => new ClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigArgs();
    }
}
