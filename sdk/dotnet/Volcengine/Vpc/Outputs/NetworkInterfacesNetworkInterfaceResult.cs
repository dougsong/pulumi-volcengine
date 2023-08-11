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
    public sealed class NetworkInterfacesNetworkInterfaceResult
    {
        /// <summary>
        /// The account id of the ENI creator.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The IP address of the EIP to which the ENI associates.
        /// </summary>
        public readonly string AssociatedElasticIpAddress;
        /// <summary>
        /// The allocation id of the EIP to which the ENI associates.
        /// </summary>
        public readonly string AssociatedElasticIpId;
        /// <summary>
        /// The create time of the ENI.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The description of the ENI.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of the device to which the ENI is bound.
        /// </summary>
        public readonly string DeviceId;
        /// <summary>
        /// The id of the ENI.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The mac address of the ENI.
        /// </summary>
        public readonly string MacAddress;
        /// <summary>
        /// The id of the ENI.
        /// </summary>
        public readonly string NetworkInterfaceId;
        /// <summary>
        /// A name of ENI.
        /// </summary>
        public readonly string NetworkInterfaceName;
        /// <summary>
        /// The enable of port security.
        /// </summary>
        public readonly bool PortSecurityEnabled;
        /// <summary>
        /// The primary IP address of the ENI.
        /// </summary>
        public readonly string PrimaryIpAddress;
        /// <summary>
        /// The IP address of secondary private network interface.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfacesNetworkInterfacePrivateIpSetResult> PrivateIpSets;
        /// <summary>
        /// The ProjectName of the ENI.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The list of the security group id to which the secondary ENI belongs.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// Whether the network card has been authorized to be used by other account services.
        /// </summary>
        public readonly bool ServiceManaged;
        /// <summary>
        /// A status of ENI, Optional choice contains `Creating`, `Available`, `Attaching`, `InUse`, `Detaching`, `Deleting`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// An id of the subnet to which the ENI is connected.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.NetworkInterfacesNetworkInterfaceTagResult> Tags;
        /// <summary>
        /// A type of ENI, Optional choice contains `primary`, `secondary`.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The last update time of the ENI.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// An id of the virtual private cloud (VPC) to which the ENI belongs.
        /// </summary>
        public readonly string VpcId;
        /// <summary>
        /// The name of the virtual private cloud (VPC) to which the ENI belongs.
        /// </summary>
        public readonly string VpcName;
        /// <summary>
        /// The zone ID.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private NetworkInterfacesNetworkInterfaceResult(
            string accountId,

            string associatedElasticIpAddress,

            string associatedElasticIpId,

            string createdAt,

            string description,

            string deviceId,

            string id,

            string macAddress,

            string networkInterfaceId,

            string networkInterfaceName,

            bool portSecurityEnabled,

            string primaryIpAddress,

            ImmutableArray<Outputs.NetworkInterfacesNetworkInterfacePrivateIpSetResult> privateIpSets,

            string projectName,

            ImmutableArray<string> securityGroupIds,

            bool serviceManaged,

            string status,

            string subnetId,

            ImmutableArray<Outputs.NetworkInterfacesNetworkInterfaceTagResult> tags,

            string type,

            string updatedAt,

            string vpcId,

            string vpcName,

            string zoneId)
        {
            AccountId = accountId;
            AssociatedElasticIpAddress = associatedElasticIpAddress;
            AssociatedElasticIpId = associatedElasticIpId;
            CreatedAt = createdAt;
            Description = description;
            DeviceId = deviceId;
            Id = id;
            MacAddress = macAddress;
            NetworkInterfaceId = networkInterfaceId;
            NetworkInterfaceName = networkInterfaceName;
            PortSecurityEnabled = portSecurityEnabled;
            PrimaryIpAddress = primaryIpAddress;
            PrivateIpSets = privateIpSets;
            ProjectName = projectName;
            SecurityGroupIds = securityGroupIds;
            ServiceManaged = serviceManaged;
            Status = status;
            SubnetId = subnetId;
            Tags = tags;
            Type = type;
            UpdatedAt = updatedAt;
            VpcId = vpcId;
            VpcName = vpcName;
            ZoneId = zoneId;
        }
    }
}
