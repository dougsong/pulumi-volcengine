// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Mongodb.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceMongoResult
    {
        /// <summary>
        /// The mongos node ID.
        /// </summary>
        public readonly string MongosNodeId;
        /// <summary>
        /// The node spec.
        /// </summary>
        public readonly string NodeSpec;
        /// <summary>
        /// The node status.
        /// </summary>
        public readonly string NodeStatus;
        /// <summary>
        /// The total memory in GB.
        /// </summary>
        public readonly double TotalMemoryGb;
        /// <summary>
        /// The total vCPU.
        /// </summary>
        public readonly double TotalVcpu;
        /// <summary>
        /// The used memory in GB.
        /// </summary>
        public readonly double UsedMemoryGb;
        /// <summary>
        /// The used vCPU.
        /// </summary>
        public readonly double UsedVcpu;
        /// <summary>
        /// The zone ID to query.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private InstancesInstanceMongoResult(
            string mongosNodeId,

            string nodeSpec,

            string nodeStatus,

            double totalMemoryGb,

            double totalVcpu,

            double usedMemoryGb,

            double usedVcpu,

            string zoneId)
        {
            MongosNodeId = mongosNodeId;
            NodeSpec = nodeSpec;
            NodeStatus = nodeStatus;
            TotalMemoryGb = totalMemoryGb;
            TotalVcpu = totalVcpu;
            UsedMemoryGb = usedMemoryGb;
            UsedVcpu = usedVcpu;
            ZoneId = zoneId;
        }
    }
}
