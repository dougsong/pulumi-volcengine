// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vke.Outputs
{

    [OutputType]
    public sealed class NodesNodeResult
    {
        /// <summary>
        /// Is Additional Container storage enables.
        /// </summary>
        public readonly bool AdditionalContainerStorageEnabled;
        /// <summary>
        /// The cluster id of node.
        /// </summary>
        public readonly string ClusterId;
        /// <summary>
        /// The Condition of Node.
        /// </summary>
        public readonly ImmutableArray<string> ConditionTypes;
        /// <summary>
        /// The Storage Path.
        /// </summary>
        public readonly string ContainerStoragePath;
        /// <summary>
        /// The Cordon of KubernetesConfig.
        /// </summary>
        public readonly bool Cordon;
        /// <summary>
        /// The Create Client Token.
        /// </summary>
        public readonly string CreateClientToken;
        /// <summary>
        /// The create time of Node.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The ID of Node.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ImageId of NodeConfig.
        /// </summary>
        public readonly string ImageId;
        /// <summary>
        /// The InitializeScript of NodeConfig.
        /// </summary>
        public readonly string InitializeScript;
        /// <summary>
        /// The instance id of node.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// Is virtual node.
        /// </summary>
        public readonly bool IsVirtual;
        /// <summary>
        /// The Label of KubernetesConfig.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodesNodeLabelResult> Labels;
        /// <summary>
        /// The Name of Node.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The node pool id.
        /// </summary>
        public readonly string NodePoolId;
        /// <summary>
        /// The Phase of Node, the value is `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Starting` or `Stopping` or `Stopped`.
        /// </summary>
        public readonly string Phase;
        /// <summary>
        /// The roles of node.
        /// </summary>
        public readonly ImmutableArray<string> Roles;
        /// <summary>
        /// The Taint of KubernetesConfig.
        /// </summary>
        public readonly ImmutableArray<Outputs.NodesNodeTaintResult> Taints;
        /// <summary>
        /// The update time of Node.
        /// </summary>
        public readonly string UpdateTime;
        /// <summary>
        /// The zone id.
        /// </summary>
        public readonly string ZoneId;

        [OutputConstructor]
        private NodesNodeResult(
            bool additionalContainerStorageEnabled,

            string clusterId,

            ImmutableArray<string> conditionTypes,

            string containerStoragePath,

            bool cordon,

            string createClientToken,

            string createTime,

            string id,

            string imageId,

            string initializeScript,

            string instanceId,

            bool isVirtual,

            ImmutableArray<Outputs.NodesNodeLabelResult> labels,

            string name,

            string nodePoolId,

            string phase,

            ImmutableArray<string> roles,

            ImmutableArray<Outputs.NodesNodeTaintResult> taints,

            string updateTime,

            string zoneId)
        {
            AdditionalContainerStorageEnabled = additionalContainerStorageEnabled;
            ClusterId = clusterId;
            ConditionTypes = conditionTypes;
            ContainerStoragePath = containerStoragePath;
            Cordon = cordon;
            CreateClientToken = createClientToken;
            CreateTime = createTime;
            Id = id;
            ImageId = imageId;
            InitializeScript = initializeScript;
            InstanceId = instanceId;
            IsVirtual = isVirtual;
            Labels = labels;
            Name = name;
            NodePoolId = nodePoolId;
            Phase = phase;
            Roles = roles;
            Taints = taints;
            UpdateTime = updateTime;
            ZoneId = zoneId;
        }
    }
}
