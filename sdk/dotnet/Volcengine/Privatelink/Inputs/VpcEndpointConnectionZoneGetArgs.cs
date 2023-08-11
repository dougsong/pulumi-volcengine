// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Privatelink.Inputs
{

    public sealed class VpcEndpointConnectionZoneGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the network interface.
        /// </summary>
        [Input("networkInterfaceId")]
        public Input<string>? NetworkInterfaceId { get; set; }

        /// <summary>
        /// The ip address of the network interface.
        /// </summary>
        [Input("networkInterfaceIp")]
        public Input<string>? NetworkInterfaceIp { get; set; }

        /// <summary>
        /// The id of the resource.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The id of the subnet.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        /// <summary>
        /// The domain of the zone.
        /// </summary>
        [Input("zoneDomain")]
        public Input<string>? ZoneDomain { get; set; }

        /// <summary>
        /// The id of the zone.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        /// <summary>
        /// The status of the zone.
        /// </summary>
        [Input("zoneStatus")]
        public Input<string>? ZoneStatus { get; set; }

        public VpcEndpointConnectionZoneGetArgs()
        {
        }
        public static new VpcEndpointConnectionZoneGetArgs Empty => new VpcEndpointConnectionZoneGetArgs();
    }
}
