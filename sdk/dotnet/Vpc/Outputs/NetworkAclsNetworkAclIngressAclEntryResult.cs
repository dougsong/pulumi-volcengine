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
    public sealed class NetworkAclsNetworkAclIngressAclEntryResult
    {
        /// <summary>
        /// The description of entry.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The id of entry.
        /// </summary>
        public readonly string NetworkAclEntryId;
        /// <summary>
        /// The name of entry.
        /// </summary>
        public readonly string NetworkAclEntryName;
        /// <summary>
        /// The policy of entry.
        /// </summary>
        public readonly string Policy;
        /// <summary>
        /// The port of entry.
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// The priority of entry.
        /// </summary>
        public readonly int Priority;
        /// <summary>
        /// The protocol of entry.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The SourceCidrIp of entry.
        /// </summary>
        public readonly string SourceCidrIp;

        [OutputConstructor]
        private NetworkAclsNetworkAclIngressAclEntryResult(
            string description,

            string networkAclEntryId,

            string networkAclEntryName,

            string policy,

            string port,

            int priority,

            string protocol,

            string sourceCidrIp)
        {
            Description = description;
            NetworkAclEntryId = networkAclEntryId;
            NetworkAclEntryName = networkAclEntryName;
            Policy = policy;
            Port = port;
            Priority = priority;
            Protocol = protocol;
            SourceCidrIp = sourceCidrIp;
        }
    }
}
