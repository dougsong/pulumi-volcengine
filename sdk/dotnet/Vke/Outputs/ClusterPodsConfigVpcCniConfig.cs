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
    public sealed class ClusterPodsConfigVpcCniConfig
    {
        /// <summary>
        /// A list of Pod subnet IDs for the VPC-CNI container network.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// The private network where the cluster control plane network resides.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private ClusterPodsConfigVpcCniConfig(
            ImmutableArray<string> subnetIds,

            string? vpcId)
        {
            SubnetIds = subnetIds;
            VpcId = vpcId;
        }
    }
}
