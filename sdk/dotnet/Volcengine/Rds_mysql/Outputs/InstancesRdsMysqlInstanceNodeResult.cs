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
    public sealed class InstancesRdsMysqlInstanceNodeResult
    {
        /// <summary>
        /// Node creation local time.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The id of the RDS instance.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Memory size in GB.
        /// </summary>
        public readonly int Memory;
        /// <summary>
        /// Node ID.
        /// </summary>
        public readonly string NodeId;
        /// <summary>
        /// General instance type, different from Custom instance type.
        /// </summary>
        public readonly string NodeSpec;
        /// <summary>
        /// Node state, value: aligned with instance state.
        /// </summary>
        public readonly string NodeStatus;
        /// <summary>
        /// Node type. Value: Primary: Primary node.
        /// Secondary: Standby node.
        /// ReadOnly: Read-only node.
        /// </summary>
        public readonly string NodeType;
        /// <summary>
        /// The region of the RDS instance.
        /// </summary>
        public readonly string RegionId;
        /// <summary>
        /// The update time of the RDS instance.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// CPU size.
        /// </summary>
        public readonly int VCpu;
        /// <summary>
        /// The available zone of the RDS instance.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private InstancesRdsMysqlInstanceNodeResult(
            string createTime,

            string instanceId,

            int memory,

            string nodeId,

            string nodeSpec,

            string nodeStatus,

            string nodeType,

            string regionId,

            string updateTime,

            int vCpu,

            string zoneId)
        {
            CreateTime = createTime;
            InstanceId = instanceId;
            Memory = memory;
            NodeId = nodeId;
            NodeSpec = nodeSpec;
            NodeStatus = nodeStatus;
            NodeType = nodeType;
            RegionId = regionId;
            UpdateTime = updateTime;
            VCpu = vCpu;
            ZoneId = zoneId;
        }
    }
}
