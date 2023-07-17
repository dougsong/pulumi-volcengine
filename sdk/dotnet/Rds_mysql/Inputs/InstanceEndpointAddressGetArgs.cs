// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql.Inputs
{

    public sealed class InstanceEndpointAddressGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// DNS Visibility.
        /// </summary>
        [Input("dnsVisibility")]
        public Input<bool>? DnsVisibility { get; set; }

        /// <summary>
        /// Connect domain name.
        /// </summary>
        [Input("domain")]
        public Input<string>? Domain { get; set; }

        /// <summary>
        /// The ID of the EIP, only valid for Public addresses.
        /// </summary>
        [Input("eipId")]
        public Input<string>? EipId { get; set; }

        /// <summary>
        /// The IP Address.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Network address type, temporarily Private, Public, PublicService.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// The Port.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// Subnet ID of the RDS instance.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public InstanceEndpointAddressGetArgs()
        {
        }
    }
}
