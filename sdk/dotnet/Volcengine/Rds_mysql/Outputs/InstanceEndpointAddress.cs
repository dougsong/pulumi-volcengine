// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Rds_mysql.Outputs
{

    [OutputType]
    public sealed class InstanceEndpointAddress
    {
        /// <summary>
        /// DNS Visibility.
        /// </summary>
        public readonly bool? DnsVisibility;
        /// <summary>
        /// Connect domain name.
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// The ID of the EIP, only valid for Public addresses.
        /// </summary>
        public readonly string? EipId;
        /// <summary>
        /// The IP Address.
        /// </summary>
        public readonly string? IpAddress;
        /// <summary>
        /// Network address type, temporarily Private, Public, PublicService.
        /// </summary>
        public readonly string? NetworkType;
        /// <summary>
        /// The Port.
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// Subnet ID of the RDS instance.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private InstanceEndpointAddress(
            bool? dnsVisibility,

            string? domain,

            string? eipId,

            string? ipAddress,

            string? networkType,

            string? port,

            string? subnetId)
        {
            DnsVisibility = dnsVisibility;
            Domain = domain;
            EipId = eipId;
            IpAddress = ipAddress;
            NetworkType = networkType;
            Port = port;
            SubnetId = subnetId;
        }
    }
}
