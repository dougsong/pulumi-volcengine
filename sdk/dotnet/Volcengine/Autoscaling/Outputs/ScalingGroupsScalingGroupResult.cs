// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Autoscaling.Outputs
{

    [OutputType]
    public sealed class ScalingGroupsScalingGroupResult
    {
        /// <summary>
        /// The scaling configuration id which used by the scaling group.
        /// </summary>
        public readonly string ActiveScalingConfigurationId;
        /// <summary>
        /// The create time of the scaling group.
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// The list of db instance ids.
        /// </summary>
        public readonly ImmutableArray<string> DbInstanceIds;
        /// <summary>
        /// The default cooldown interval of the scaling group.
        /// </summary>
        public readonly int DefaultCooldown;
        /// <summary>
        /// The desire instance number of the scaling group.
        /// </summary>
        public readonly int DesireInstanceNumber;
        /// <summary>
        /// The id of the scaling group.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The instance terminate policy of the scaling group.
        /// </summary>
        public readonly string InstanceTerminatePolicy;
        /// <summary>
        /// The ID of the launch template bound to the scaling group.
        /// </summary>
        public readonly string LaunchTemplateId;
        /// <summary>
        /// The version of the launch template bound to the scaling group.
        /// </summary>
        public readonly string LaunchTemplateVersion;
        /// <summary>
        /// The lifecycle state of the scaling group.
        /// </summary>
        public readonly string LifecycleState;
        /// <summary>
        /// The max instance number of the scaling group.
        /// </summary>
        public readonly int MaxInstanceNumber;
        /// <summary>
        /// The min instance number of the scaling group.
        /// </summary>
        public readonly int MinInstanceNumber;
        /// <summary>
        /// The multi az policy of the scaling group. Valid values: PRIORITY, BALANCE.
        /// </summary>
        public readonly string MultiAzPolicy;
        /// <summary>
        /// The ProjectName of scaling group.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The id of the scaling group.
        /// </summary>
        public readonly string ScalingGroupId;
        /// <summary>
        /// The name of the scaling group.
        /// </summary>
        public readonly string ScalingGroupName;
        /// <summary>
        /// The list of server group attributes.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalingGroupsScalingGroupServerGroupAttributeResult> ServerGroupAttributes;
        /// <summary>
        /// The list of the subnet id to which the ENI is connected.
        /// </summary>
        public readonly ImmutableArray<string> SubnetIds;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.ScalingGroupsScalingGroupTagResult> Tags;
        /// <summary>
        /// The total instance count of the scaling group.
        /// </summary>
        public readonly int TotalInstanceCount;
        /// <summary>
        /// The create time of the scaling group.
        /// </summary>
        public readonly string UpdatedAt;
        /// <summary>
        /// The VPC id of the scaling group.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private ScalingGroupsScalingGroupResult(
            string activeScalingConfigurationId,

            string createdAt,

            ImmutableArray<string> dbInstanceIds,

            int defaultCooldown,

            int desireInstanceNumber,

            string id,

            string instanceTerminatePolicy,

            string launchTemplateId,

            string launchTemplateVersion,

            string lifecycleState,

            int maxInstanceNumber,

            int minInstanceNumber,

            string multiAzPolicy,

            string projectName,

            string scalingGroupId,

            string scalingGroupName,

            ImmutableArray<Outputs.ScalingGroupsScalingGroupServerGroupAttributeResult> serverGroupAttributes,

            ImmutableArray<string> subnetIds,

            ImmutableArray<Outputs.ScalingGroupsScalingGroupTagResult> tags,

            int totalInstanceCount,

            string updatedAt,

            string vpcId)
        {
            ActiveScalingConfigurationId = activeScalingConfigurationId;
            CreatedAt = createdAt;
            DbInstanceIds = dbInstanceIds;
            DefaultCooldown = defaultCooldown;
            DesireInstanceNumber = desireInstanceNumber;
            Id = id;
            InstanceTerminatePolicy = instanceTerminatePolicy;
            LaunchTemplateId = launchTemplateId;
            LaunchTemplateVersion = launchTemplateVersion;
            LifecycleState = lifecycleState;
            MaxInstanceNumber = maxInstanceNumber;
            MinInstanceNumber = minInstanceNumber;
            MultiAzPolicy = multiAzPolicy;
            ProjectName = projectName;
            ScalingGroupId = scalingGroupId;
            ScalingGroupName = scalingGroupName;
            ServerGroupAttributes = serverGroupAttributes;
            SubnetIds = subnetIds;
            Tags = tags;
            TotalInstanceCount = totalInstanceCount;
            UpdatedAt = updatedAt;
            VpcId = vpcId;
        }
    }
}
