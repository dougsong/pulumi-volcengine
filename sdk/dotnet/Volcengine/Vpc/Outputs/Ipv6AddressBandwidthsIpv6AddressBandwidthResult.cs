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
    public sealed class Ipv6AddressBandwidthsIpv6AddressBandwidthResult
    {
        /// <summary>
        /// The ID of the Ipv6AddressBandwidth.
        /// </summary>
        public readonly string AllocationId;
        /// <summary>
        /// Peek bandwidth of the Ipv6 address.
        /// </summary>
        public readonly int Bandwidth;
        /// <summary>
        /// BillingType of the Ipv6 bandwidth.
        /// </summary>
        public readonly string BillingType;
        /// <summary>
        /// The BusinessStatus of the Ipv6AddressBandwidth.
        /// </summary>
        public readonly string BusinessStatus;
        /// <summary>
        /// Creation time of the Ipv6AddressBandwidth.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// Delete time of the Ipv6AddressBandwidth.
        /// </summary>
        public readonly string DeleteTime;
        /// <summary>
        /// The ID of the Ipv6AddressBandwidth.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID of the associated instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The type of the associated instance.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The IPv6 address.
        /// </summary>
        public readonly string Ipv6Address;
        /// <summary>
        /// ISP of the ipv6 address.
        /// </summary>
        public readonly string Isp;
        /// <summary>
        /// The BusinessStatus of the Ipv6AddressBandwidth.
        /// </summary>
        public readonly string LockReason;
        /// <summary>
        /// The network type of the ipv6 address.
        /// </summary>
        public readonly string NetworkType;
        /// <summary>
        /// Overdue time of the Ipv6AddressBandwidth.
        /// </summary>
        public readonly string OverdueTime;
        /// <summary>
        /// The status of the Ipv6AddressBandwidth.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Update time of the Ipv6AddressBandwidth.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private Ipv6AddressBandwidthsIpv6AddressBandwidthResult(
            string allocationId,

            int bandwidth,

            string billingType,

            string businessStatus,

            string creationTime,

            string deleteTime,

            string id,

            string instanceId,

            string instanceType,

            string ipv6Address,

            string isp,

            string lockReason,

            string networkType,

            string overdueTime,

            string status,

            string updateTime)
        {
            AllocationId = allocationId;
            Bandwidth = bandwidth;
            BillingType = billingType;
            BusinessStatus = businessStatus;
            CreationTime = creationTime;
            DeleteTime = deleteTime;
            Id = id;
            InstanceId = instanceId;
            InstanceType = instanceType;
            Ipv6Address = ipv6Address;
            Isp = isp;
            LockReason = lockReason;
            NetworkType = networkType;
            OverdueTime = overdueTime;
            Status = status;
            UpdateTime = updateTime;
        }
    }
}
