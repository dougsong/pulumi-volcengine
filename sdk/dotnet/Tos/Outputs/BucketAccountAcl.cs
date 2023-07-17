// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tos.Outputs
{

    [OutputType]
    public sealed class BucketAccountAcl
    {
        /// <summary>
        /// The accountId to control.
        /// </summary>
        public readonly string AccountId;
        /// <summary>
        /// The acl type to control.Valid value is CanonicalUser.
        /// </summary>
        public readonly string? AclType;
        /// <summary>
        /// The permission to control.Valid value is FULL_CONTROL|READ|READ_ACP|WRITE|WRITE_ACP.
        /// </summary>
        public readonly string Permission;

        [OutputConstructor]
        private BucketAccountAcl(
            string accountId,

            string? aclType,

            string permission)
        {
            AccountId = accountId;
            AclType = aclType;
            Permission = permission;
        }
    }
}
