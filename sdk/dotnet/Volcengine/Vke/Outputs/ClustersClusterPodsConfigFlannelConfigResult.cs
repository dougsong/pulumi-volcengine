// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vke.Outputs
{

    [OutputType]
    public sealed class ClustersClusterPodsConfigFlannelConfigResult
    {
        /// <summary>
        /// The maximum number of single-node Pod instances for a Flannel container network.
        /// </summary>
        public readonly int MaxPodsPerNode;
        /// <summary>
        /// Pod CIDR for the Flannel container network.
        /// </summary>
        public readonly ImmutableArray<string> PodCidrs;

        [OutputConstructor]
        private ClustersClusterPodsConfigFlannelConfigResult(
            int maxPodsPerNode,

            ImmutableArray<string> podCidrs)
        {
            MaxPodsPerNode = maxPodsPerNode;
            PodCidrs = podCidrs;
        }
    }
}
