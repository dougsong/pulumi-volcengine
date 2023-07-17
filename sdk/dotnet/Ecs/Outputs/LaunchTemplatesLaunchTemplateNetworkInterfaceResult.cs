// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Ecs.Outputs
{

    [OutputType]
    public sealed class LaunchTemplatesLaunchTemplateNetworkInterfaceResult
    {
        /// <summary>
        /// The security group ID associated with the NIC.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The private network subnet ID of the instance, when creating the instance, supports binding the secondary NIC at the same time.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private LaunchTemplatesLaunchTemplateNetworkInterfaceResult(
            ImmutableArray<string> securityGroupIds,

            string subnetId)
        {
            SecurityGroupIds = securityGroupIds;
            SubnetId = subnetId;
        }
    }
}
