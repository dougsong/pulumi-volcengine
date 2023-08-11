// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Redis.Inputs
{

    public sealed class BackupInstanceDetailGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of account.
        /// </summary>
        [Input("accountId")]
        public Input<int>? AccountId { get; set; }

        /// <summary>
        /// Arch type of instance(Standard/Cluster).
        /// </summary>
        [Input("archType")]
        public Input<string>? ArchType { get; set; }

        /// <summary>
        /// Charge type of instance(Postpaid/Prepaid).
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Engine version of instance.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// Expired time of instance.
        /// </summary>
        [Input("expiredTime")]
        public Input<string>? ExpiredTime { get; set; }

        /// <summary>
        /// Id of instance to create backup.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The maintainable period (in UTC) of the instance.
        /// </summary>
        [Input("maintenanceTime")]
        public Input<string>? MaintenanceTime { get; set; }

        /// <summary>
        /// Network type of instance.
        /// </summary>
        [Input("networkType")]
        public Input<string>? NetworkType { get; set; }

        /// <summary>
        /// Project name of instance.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// Id of region.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        /// <summary>
        /// Count of replica in which shard.
        /// </summary>
        [Input("replicas")]
        public Input<int>? Replicas { get; set; }

        /// <summary>
        /// Count of cpu cores of instance.
        /// </summary>
        [Input("serverCpu")]
        public Input<int>? ServerCpu { get; set; }

        /// <summary>
        /// Capacity of shard.
        /// </summary>
        [Input("shardCapacity")]
        public Input<int>? ShardCapacity { get; set; }

        /// <summary>
        /// Count of shard.
        /// </summary>
        [Input("shardCount")]
        public Input<int>? ShardCount { get; set; }

        /// <summary>
        /// Total capacity of instance.
        /// </summary>
        [Input("totalCapacity")]
        public Input<int>? TotalCapacity { get; set; }

        /// <summary>
        /// Capacity used of this instance.
        /// </summary>
        [Input("usedCapacity")]
        public Input<int>? UsedCapacity { get; set; }

        [Input("vpcInfos")]
        private InputList<Inputs.BackupInstanceDetailVpcInfoGetArgs>? _vpcInfos;

        /// <summary>
        /// Information of vpc.
        /// </summary>
        public InputList<Inputs.BackupInstanceDetailVpcInfoGetArgs> VpcInfos
        {
            get => _vpcInfos ?? (_vpcInfos = new InputList<Inputs.BackupInstanceDetailVpcInfoGetArgs>());
            set => _vpcInfos = value;
        }

        [Input("zoneIds")]
        private InputList<string>? _zoneIds;

        /// <summary>
        /// List of id of zone.
        /// </summary>
        public InputList<string> ZoneIds
        {
            get => _zoneIds ?? (_zoneIds = new InputList<string>());
            set => _zoneIds = value;
        }

        public BackupInstanceDetailGetArgs()
        {
        }
        public static new BackupInstanceDetailGetArgs Empty => new BackupInstanceDetailGetArgs();
    }
}
