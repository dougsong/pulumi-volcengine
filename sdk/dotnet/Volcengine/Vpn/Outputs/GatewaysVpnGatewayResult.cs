// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vpn.Outputs
{

    [OutputType]
    public sealed class GatewaysVpnGatewayResult
    {
        /// <summary>
        /// The account ID of the VPN gateway.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The bandwidth of the VPN gateway.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// The BillingType of the VPN gateway.
        /// </summary>
        public readonly string BillingType;
        /// <summary>
        /// The business status of the VPN gateway.
        /// </summary>
        public readonly string BusinessStatus;
        /// <summary>
        /// The connection count of the VPN gateway.
        /// </summary>
        public readonly int ConnectionCount;
        /// <summary>
        /// The create time of VPN gateway.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The deleted time of the VPN gateway.
        /// </summary>
        public readonly string DeletedTime;
        /// <summary>
        /// The description of the VPN gateway.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The expired time of the VPN gateway.
        /// </summary>
        public readonly string ExpiredTime;
        /// <summary>
        /// The ID of the VPN gateway.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A IP address of the VPN gateway.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The lock reason of the VPN gateway.
        /// </summary>
        public readonly string LockReason;
        /// <summary>
        /// The route count of the VPN gateway.
        /// </summary>
        public readonly int RouteCount;
        /// <summary>
        /// The status of the VPN gateway.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// A subnet ID of the VPN gateway.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.GatewaysVpnGatewayTagResult> Tags;
        /// <summary>
        /// The update time of VPN gateway.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// A VPC ID of the VPN gateway.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The ID of the VPN gateway.
        /// </summary>
        public readonly string VpnGatewayId;
        /// <summary>
        /// The name of the VPN gateway.
        /// </summary>
        public readonly string VpnGatewayName;

        [OutputConstructor]
        private GatewaysVpnGatewayResult(
            string accountId,

            int bandwidth,

            string billingType,

            string businessStatus,

            int connectionCount,

            string creationTime,

            string deletedTime,

            string description,

            string expiredTime,

            string id,

            string ipAddress,

            string lockReason,

            int routeCount,

            string status,

            string? subnetId,

            ImmutableArray<Outputs.GatewaysVpnGatewayTagResult> tags,

            string updateTime,

            string vpcId,

            string vpnGatewayId,

            string vpnGatewayName)
        {
            AccountId = accountId;
            Bandwidth = bandwidth;
            BillingType = billingType;
            BusinessStatus = businessStatus;
            ConnectionCount = connectionCount;
            CreationTime = creationTime;
            DeletedTime = deletedTime;
            Description = description;
            ExpiredTime = expiredTime;
            Id = id;
            IpAddress = ipAddress;
            LockReason = lockReason;
            RouteCount = routeCount;
            Status = status;
            SubnetId = subnetId;
            Tags = tags;
            UpdateTime = updateTime;
            VpcId = vpcId;
            VpnGatewayId = vpnGatewayId;
            VpnGatewayName = vpnGatewayName;
        }
    }
}
