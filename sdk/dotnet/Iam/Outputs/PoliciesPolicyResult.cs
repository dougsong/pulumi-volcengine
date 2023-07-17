// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Iam.Outputs
{

    [OutputType]
    public sealed class PoliciesPolicyResult
    {
        /// <summary>
        /// The create time of the Policy.
        /// </summary>
        public readonly string CreateDate;
        /// <summary>
        /// The description of the Policy.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The ID of the Policy.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The document of the Policy.
        /// </summary>
        public readonly string PolicyDocument;
        /// <summary>
        /// The name of the Policy.
        /// </summary>
        public readonly string PolicyName;
        /// <summary>
        /// The resource name of the Policy.
        /// </summary>
        public readonly string PolicyTrn;
        /// <summary>
        /// The type of the Policy.
        /// </summary>
        public readonly string PolicyType;
        /// <summary>
        /// The role attach time of the Policy.The data show only query with role_name.
        /// </summary>
        public readonly string RoleAttachDate;
        /// <summary>
        /// The name of the IAM role.
        /// </summary>
        public readonly string RoleName;
        /// <summary>
        /// The update time of the Policy.
        /// </summary>
        public readonly string UpdateDate;
        /// <summary>
        /// The user attach time of the Policy.The data show only query with user_name.
        /// </summary>
        public readonly string UserAttachDate;
        /// <summary>
        /// The name of the IAM user.
        /// </summary>
        public readonly string UserName;

        [OutputConstructor]
        private PoliciesPolicyResult(
            string createDate,

            string description,

            string id,

            string policyDocument,

            string policyName,

            string policyTrn,

            string policyType,

            string roleAttachDate,

            string roleName,

            string updateDate,

            string userAttachDate,

            string userName)
        {
            CreateDate = createDate;
            Description = description;
            Id = id;
            PolicyDocument = policyDocument;
            PolicyName = policyName;
            PolicyTrn = policyTrn;
            PolicyType = policyType;
            RoleAttachDate = roleAttachDate;
            RoleName = roleName;
            UpdateDate = updateDate;
            UserAttachDate = userAttachDate;
            UserName = userName;
        }
    }
}
