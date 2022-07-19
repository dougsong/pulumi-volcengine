// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc.Outputs
{

    [OutputType]
    public sealed class SnatEntriesSnatEntryResult
    {
        /// <summary>
        /// The public ip address used by the SNAT entry.
        /// </summary>
        public readonly string EipAddress;
        /// <summary>
        /// An id of the public ip address used by the SNAT entry.
        /// </summary>
        public readonly string EipId;
        /// <summary>
        /// The id of the SNAT entry.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// An id of the nat gateway to which the entry belongs.
        /// </summary>
        public readonly string NatGatewayId;
        /// <summary>
        /// The id of the SNAT entry.
        /// </summary>
        public readonly string SnatEntryId;
        /// <summary>
        /// A name of SNAT entry.
        /// </summary>
        public readonly string SnatEntryName;
        /// <summary>
        /// The status of the SNAT entry.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// An id of the subnet that is required to access the Internet.
        /// </summary>
        public readonly string SubnetId;

        [OutputConstructor]
        private SnatEntriesSnatEntryResult(
            string eipAddress,

            string eipId,

            string id,

            string natGatewayId,

            string snatEntryId,

            string snatEntryName,

            string status,

            string subnetId)
        {
            EipAddress = eipAddress;
            EipId = eipId;
            Id = id;
            NatGatewayId = natGatewayId;
            SnatEntryId = snatEntryId;
            SnatEntryName = snatEntryName;
            Status = status;
            SubnetId = subnetId;
        }
    }
}
