// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vpc.Outputs
{

    [OutputType]
    public sealed class VpcsVpcResult
    {
        /// <summary>
        /// The account ID of VPC.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The associate cen list of VPC.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpcsVpcAssociateCenResult> AssociateCens;
        /// <summary>
        /// The auxiliary cidr block list of VPC.
        /// </summary>
        public readonly ImmutableArray<string> AuxiliaryCidrBlocks;
        /// <summary>
        /// The cidr block of VPC.
        /// </summary>
        public readonly string CidrBlock;
        /// <summary>
        /// The create time of VPC.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The description of VPC.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The dns server list of VPC.
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// The IPv6 CIDR block of the VPC.
        /// </summary>
        public readonly string Ipv6CidrBlock;
        /// <summary>
        /// The nat gateway ID list of VPC.
        /// </summary>
        public readonly ImmutableArray<string> NatGatewayIds;
        /// <summary>
        /// The ProjectName of the VPC.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The route table ID list of VPC.
        /// </summary>
        public readonly ImmutableArray<string> RouteTableIds;
        /// <summary>
        /// The security group ID list of VPC.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The status of VPC.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The subnet ID list of VPC.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpcsVpcTagResult> Tags;
        /// <summary>
        /// The update time of VPC.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The ID of VPC.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The vpc name to query.
        /// </summary>
        public readonly string VpcName;

        [OutputConstructor]
        private VpcsVpcResult(
            string accountId,

            ImmutableArray<Outputs.VpcsVpcAssociateCenResult> associateCens,

            ImmutableArray<string> auxiliaryCidrBlocks,

            string cidrBlock,

            string creationTime,

            string description,

            ImmutableArray<string> dnsServers,

            string ipv6CidrBlock,

            ImmutableArray<string> natGatewayIds,

            string projectName,

            ImmutableArray<string> routeTableIds,

            ImmutableArray<string> securityGroupIds,

            string status,

            ImmutableArray<string> subnetIds,

            ImmutableArray<Outputs.VpcsVpcTagResult> tags,

            string updateTime,

            string vpcId,

            string vpcName)
        {
            AccountId = accountId;
            AssociateCens = associateCens;
            AuxiliaryCidrBlocks = auxiliaryCidrBlocks;
            CidrBlock = cidrBlock;
            CreationTime = creationTime;
            Description = description;
            DnsServers = dnsServers;
            Ipv6CidrBlock = ipv6CidrBlock;
            NatGatewayIds = natGatewayIds;
            ProjectName = projectName;
            RouteTableIds = routeTableIds;
            SecurityGroupIds = securityGroupIds;
            Status = status;
            SubnetIds = subnetIds;
            Tags = tags;
            UpdateTime = updateTime;
            VpcId = vpcId;
            VpcName = vpcName;
        }
    }
}
