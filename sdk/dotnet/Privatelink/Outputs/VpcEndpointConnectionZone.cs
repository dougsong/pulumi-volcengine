// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Privatelink.Outputs
{

    [OutputType]
    public sealed class VpcEndpointConnectionZone
    {
        /// <summary>
        /// The id of the network interface.
        /// </summary>
        public readonly string? NetworkInterfaceId;
        /// <summary>
        /// The ip address of the network interface.
        /// </summary>
        public readonly string? NetworkInterfaceIp;
        /// <summary>
        /// The id of the resource.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// The id of the subnet.
        /// </summary>
        public readonly string? SubnetId;
        /// <summary>
        /// The domain of the zone.
        /// </summary>
        public readonly string? ZoneDomain;
        /// <summary>
        /// The id of the zone.
        /// </summary>
        public readonly string? ZoneId;
        /// <summary>
        /// The status of the zone.
        /// </summary>
        public readonly string? ZoneStatus;

        [OutputConstructor]
        private VpcEndpointConnectionZone(
            string? networkInterfaceId,

            string? networkInterfaceIp,

            string? resourceId,

            string? subnetId,

            string? zoneDomain,

            string? zoneId,

            string? zoneStatus)
        {
            NetworkInterfaceId = networkInterfaceId;
            NetworkInterfaceIp = networkInterfaceIp;
            ResourceId = resourceId;
            SubnetId = subnetId;
            ZoneDomain = zoneDomain;
            ZoneId = zoneId;
            ZoneStatus = zoneStatus;
        }
    }
}
