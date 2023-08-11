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

    public sealed class ClusterClusterConfigGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cluster API Server public network access configuration.
        /// </summary>
        [Input("apiServerPublicAccessConfig")]
        public Input<Inputs.ClusterClusterConfigApiServerPublicAccessConfigGetArgs>? ApiServerPublicAccessConfig { get; set; }

        /// <summary>
        /// Cluster API Server public network access configuration, the value is `true` or `false`.
        /// </summary>
        [Input("apiServerPublicAccessEnabled")]
        public Input<bool>? ApiServerPublicAccessEnabled { get; set; }

        [Input("ipFamily")]
        public Input<string>? IpFamily { get; set; }

        /// <summary>
        /// Node public network access configuration, the value is `true` or `false`.
        /// </summary>
        [Input("resourcePublicAccessDefaultEnabled")]
        public Input<bool>? ResourcePublicAccessDefaultEnabled { get; set; }

        [Input("subnetIds", required: true)]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The subnet ID for the cluster control plane to communicate within the private network.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        public ClusterClusterConfigGetArgs()
        {
        }
        public static new ClusterClusterConfigGetArgs Empty => new ClusterClusterConfigGetArgs();
    }
}
