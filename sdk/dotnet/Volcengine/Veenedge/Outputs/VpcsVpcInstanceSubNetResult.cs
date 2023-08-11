// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Veenedge.Outputs
{

    [OutputType]
    public sealed class VpcsVpcInstanceSubNetResult
    {
        /// <summary>
        /// The account id.
        /// </summary>
        public readonly int AccountIdentity;
        /// <summary>
        /// The ip of cidr.
        /// </summary>
        public readonly string CidrIp;
        /// <summary>
        /// The mask of cidr.
        /// </summary>
        public readonly int CidrMask;
        /// <summary>
        /// The creation time.
        /// </summary>
        public readonly int CreateTime;
        /// <summary>
        /// The id of subnet.
        /// </summary>
        public readonly string SubnetIdentity;
        /// <summary>
        /// The update time of VPC.
        /// </summary>
        public readonly int UpdateTime;
        /// <summary>
        /// The id of user.
        /// </summary>
        public readonly int UserIdentity;

        [OutputConstructor]
        private VpcsVpcInstanceSubNetResult(
            int accountIdentity,

            string cidrIp,

            int cidrMask,

            int createTime,

            string subnetIdentity,

            int updateTime,

            int userIdentity)
        {
            AccountIdentity = accountIdentity;
            CidrIp = cidrIp;
            CidrMask = cidrMask;
            CreateTime = createTime;
            SubnetIdentity = subnetIdentity;
            UpdateTime = updateTime;
            UserIdentity = userIdentity;
        }
    }
}
