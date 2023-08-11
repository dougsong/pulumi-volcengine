// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Escloud.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceNodeResult
    {
        /// <summary>
        /// Is cold node.
        /// </summary>
        public readonly bool IsCold;
        /// <summary>
        /// Is hot node.
        /// </summary>
        public readonly bool IsHot;
        /// <summary>
        /// Is kibana node.
        /// </summary>
        public readonly bool IsKibana;
        /// <summary>
        /// Is master node.
        /// </summary>
        public readonly bool IsMaster;
        /// <summary>
        /// Is warm node.
        /// </summary>
        public readonly bool IsWarm;
        /// <summary>
        /// The show name of node.
        /// </summary>
        public readonly string NodeDisplayName;
        /// <summary>
        /// The name of node.
        /// </summary>
        public readonly string NodeName;
        /// <summary>
        /// The node resource spec of master.
        /// </summary>
        public readonly Outputs.InstancesInstanceNodeResourceSpecResult ResourceSpec;
        /// <summary>
        /// The restart times of node.
        /// </summary>
        public readonly int RestartNumber;
        /// <summary>
        /// The start time of node.
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The status of instance.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The node storage spec of master.
        /// </summary>
        public readonly Outputs.InstancesInstanceNodeStorageSpecResult StorageSpec;

        [OutputConstructor]
        private InstancesInstanceNodeResult(
            bool isCold,

            bool isHot,

            bool isKibana,

            bool isMaster,

            bool isWarm,

            string nodeDisplayName,

            string nodeName,

            Outputs.InstancesInstanceNodeResourceSpecResult resourceSpec,

            int restartNumber,

            string startTime,

            string status,

            Outputs.InstancesInstanceNodeStorageSpecResult storageSpec)
        {
            IsCold = isCold;
            IsHot = isHot;
            IsKibana = isKibana;
            IsMaster = isMaster;
            IsWarm = isWarm;
            NodeDisplayName = nodeDisplayName;
            NodeName = nodeName;
            ResourceSpec = resourceSpec;
            RestartNumber = restartNumber;
            StartTime = startTime;
            Status = status;
            StorageSpec = storageSpec;
        }
    }
}
