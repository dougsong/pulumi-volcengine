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
    public sealed class ConnectionsVpnConnectionResult
    {
        /// <summary>
        /// The account ID of the VPN connection.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The IPsec attach status.
        /// </summary>
        public readonly string AttachStatus;
        /// <summary>
        /// The IPsec attach type.
        /// </summary>
        public readonly string AttachType;
        /// <summary>
        /// The business status of IPsec connection, valid when the attach type is 'TransitRouter'.
        /// </summary>
        public readonly string BusinessStatus;
        /// <summary>
        /// The connect status of the VPN connection.
        /// </summary>
        public readonly string ConnectStatus;
        /// <summary>
        /// The create time of VPN connection.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// An ID of customer gateway.
        /// </summary>
        public readonly string CustomerGatewayId;
        /// <summary>
        /// The delete time of resource, valid when the attach type is 'TransitRouter'.
        /// </summary>
        public readonly string DeletedTime;
        /// <summary>
        /// The description of the VPN connection.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The dpd action of the VPN connection.
        /// </summary>
        public readonly string DpdAction;
        /// <summary>
        /// The ID of the VPN connection.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The auth alg of the ike config of the VPN connection.
        /// </summary>
        public readonly string IkeConfigAuthAlg;
        /// <summary>
        /// The dk group of the ike config of the VPN connection.
        /// </summary>
        public readonly string IkeConfigDhGroup;
        /// <summary>
        /// The enc alg of the ike config of the VPN connection.
        /// </summary>
        public readonly string IkeConfigEncAlg;
        /// <summary>
        /// The lifetime of the ike config of the VPN connection.
        /// </summary>
        public readonly int IkeConfigLifetime;
        /// <summary>
        /// The local_id of the ike config of the VPN connection.
        /// </summary>
        public readonly string IkeConfigLocalId;
        /// <summary>
        /// The mode of the ike config of the VPN connection.
        /// </summary>
        public readonly string IkeConfigMode;
        /// <summary>
        /// The psk of the ike config of the VPN connection.
        /// </summary>
        public readonly string IkeConfigPsk;
        /// <summary>
        /// The remote id of the ike config of the VPN connection.
        /// </summary>
        public readonly string IkeConfigRemoteId;
        /// <summary>
        /// The version of the ike config of the VPN connection.
        /// </summary>
        public readonly string IkeConfigVersion;
        /// <summary>
        /// The ip address of transit router, valid when the attach type is 'TransitRouter'.
        /// </summary>
        public readonly string IpAddress;
        /// <summary>
        /// The auth alg of the ipsec config of the VPN connection.
        /// </summary>
        public readonly string IpsecConfigAuthAlg;
        /// <summary>
        /// The dh group of the ipsec config of the VPN connection.
        /// </summary>
        public readonly string IpsecConfigDhGroup;
        /// <summary>
        /// The enc alg of the ipsec config of the VPN connection.
        /// </summary>
        public readonly string IpsecConfigEncAlg;
        /// <summary>
        /// The lifetime of the ike config of the VPN connection.
        /// </summary>
        public readonly int IpsecConfigLifetime;
        /// <summary>
        /// The local subnet of the VPN connection.
        /// </summary>
        public readonly ImmutableArray<string> LocalSubnets;
        /// <summary>
        /// Whether to enable the connection log.
        /// </summary>
        public readonly bool LogEnabled;
        /// <summary>
        /// The nat traversal of the VPN connection.
        /// </summary>
        public readonly bool NatTraversal;
        /// <summary>
        /// Whether to initiate negotiation mode immediately.
        /// </summary>
        public readonly bool NegotiateInstantly;
        /// <summary>
        /// The overdue time of resource, valid when the attach type is 'TransitRouter'.
        /// </summary>
        public readonly string OverdueTime;
        /// <summary>
        /// The remote subnet of the VPN connection.
        /// </summary>
        public readonly ImmutableArray<string> RemoteSubnets;
        /// <summary>
        /// The status of the VPN connection.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The id of transit router, valid when the attach type is 'TransitRouter'.
        /// </summary>
        public readonly string TransitRouterId;
        /// <summary>
        /// The update time of VPN connection.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The ID of the VPN connection.
        /// </summary>
        public readonly string VpnConnectionId;
        /// <summary>
        /// The name of the VPN connection.
        /// </summary>
        public readonly string VpnConnectionName;
        /// <summary>
        /// An ID of VPN gateway.
        /// </summary>
        public readonly string VpnGatewayId;
        /// <summary>
        /// The zone id of transit router, valid when the attach type is 'TransitRouter'.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private ConnectionsVpnConnectionResult(
            string accountId,

            string attachStatus,

            string attachType,

            string businessStatus,

            string connectStatus,

            string creationTime,

            string customerGatewayId,

            string deletedTime,

            string description,

            string dpdAction,

            string id,

            string ikeConfigAuthAlg,

            string ikeConfigDhGroup,

            string ikeConfigEncAlg,

            int ikeConfigLifetime,

            string ikeConfigLocalId,

            string ikeConfigMode,

            string ikeConfigPsk,

            string ikeConfigRemoteId,

            string ikeConfigVersion,

            string ipAddress,

            string ipsecConfigAuthAlg,

            string ipsecConfigDhGroup,

            string ipsecConfigEncAlg,

            int ipsecConfigLifetime,

            ImmutableArray<string> localSubnets,

            bool logEnabled,

            bool natTraversal,

            bool negotiateInstantly,

            string overdueTime,

            ImmutableArray<string> remoteSubnets,

            string status,

            string transitRouterId,

            string updateTime,

            string vpnConnectionId,

            string vpnConnectionName,

            string vpnGatewayId,

            string zoneId)
        {
            AccountId = accountId;
            AttachStatus = attachStatus;
            AttachType = attachType;
            BusinessStatus = businessStatus;
            ConnectStatus = connectStatus;
            CreationTime = creationTime;
            CustomerGatewayId = customerGatewayId;
            DeletedTime = deletedTime;
            Description = description;
            DpdAction = dpdAction;
            Id = id;
            IkeConfigAuthAlg = ikeConfigAuthAlg;
            IkeConfigDhGroup = ikeConfigDhGroup;
            IkeConfigEncAlg = ikeConfigEncAlg;
            IkeConfigLifetime = ikeConfigLifetime;
            IkeConfigLocalId = ikeConfigLocalId;
            IkeConfigMode = ikeConfigMode;
            IkeConfigPsk = ikeConfigPsk;
            IkeConfigRemoteId = ikeConfigRemoteId;
            IkeConfigVersion = ikeConfigVersion;
            IpAddress = ipAddress;
            IpsecConfigAuthAlg = ipsecConfigAuthAlg;
            IpsecConfigDhGroup = ipsecConfigDhGroup;
            IpsecConfigEncAlg = ipsecConfigEncAlg;
            IpsecConfigLifetime = ipsecConfigLifetime;
            LocalSubnets = localSubnets;
            LogEnabled = logEnabled;
            NatTraversal = natTraversal;
            NegotiateInstantly = negotiateInstantly;
            OverdueTime = overdueTime;
            RemoteSubnets = remoteSubnets;
            Status = status;
            TransitRouterId = transitRouterId;
            UpdateTime = updateTime;
            VpnConnectionId = vpnConnectionId;
            VpnConnectionName = vpnConnectionName;
            VpnGatewayId = vpnGatewayId;
            ZoneId = zoneId;
        }
    }
}
