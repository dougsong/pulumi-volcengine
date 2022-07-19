// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Outputs
{

    [OutputType]
    public sealed class ClustersClusterClusterConfigApiServerPublicAccessConfigResult
    {
        /// <summary>
        /// IPv4 public network access whitelist. A null value means all network segments (0.0.0.0/0) are allowed to pass.
        /// </summary>
        public readonly ImmutableArray<string> AccessSourceIpsv4s;
        /// <summary>
        /// Public network access network configuration.
        /// </summary>
        public readonly Outputs.ClustersClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigResult PublicAccessNetworkConfig;

        [OutputConstructor]
        private ClustersClusterClusterConfigApiServerPublicAccessConfigResult(
            ImmutableArray<string> accessSourceIpsv4s,

            Outputs.ClustersClusterClusterConfigApiServerPublicAccessConfigPublicAccessNetworkConfigResult publicAccessNetworkConfig)
        {
            AccessSourceIpsv4s = accessSourceIpsv4s;
            PublicAccessNetworkConfig = publicAccessNetworkConfig;
        }
    }
}
