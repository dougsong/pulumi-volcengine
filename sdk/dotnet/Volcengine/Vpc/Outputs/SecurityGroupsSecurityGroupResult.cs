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
    public sealed class SecurityGroupsSecurityGroupResult
    {
        /// <summary>
        /// The creation time of SecurityGroup.
        /// </summary>
        public readonly string CreationTime;
        /// <summary>
        /// The description of SecurityGroup.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of SecurityGroup.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ProjectName of SecurityGroup.
        /// </summary>
        public readonly string ProjectName;
        /// <summary>
        /// The ID of SecurityGroup.
        /// </summary>
        public readonly string SecurityGroupId;
        /// <summary>
        /// The Name of SecurityGroup.
        /// </summary>
        public readonly string SecurityGroupName;
        /// <summary>
        /// The Status of SecurityGroup.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Tags.
        /// </summary>
        public readonly ImmutableArray<Outputs.SecurityGroupsSecurityGroupTagResult> Tags;
        /// <summary>
        /// A Name Regex of SecurityGroup.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The ID of vpc where security group is located.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private SecurityGroupsSecurityGroupResult(
            string creationTime,

            string description,

            string id,

            string projectName,

            string securityGroupId,

            string securityGroupName,

            string status,

            ImmutableArray<Outputs.SecurityGroupsSecurityGroupTagResult> tags,

            string type,

            string vpcId)
        {
            CreationTime = creationTime;
            Description = description;
            Id = id;
            ProjectName = projectName;
            SecurityGroupId = securityGroupId;
            SecurityGroupName = securityGroupName;
            Status = status;
            Tags = tags;
            Type = type;
            VpcId = vpcId;
        }
    }
}
