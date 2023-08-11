// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Tos.Inputs
{

    public sealed class BucketAccountAclGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The accountId to control.
        /// </summary>
        [Input("accountId", required: true)]
        public Input<string> AccountId { get; set; } = null!;

        /// <summary>
        /// The acl type to control.Valid value is CanonicalUser.
        /// </summary>
        [Input("aclType")]
        public Input<string>? AclType { get; set; }

        /// <summary>
        /// The permission to control.Valid value is FULL_CONTROL|READ|READ_ACP|WRITE|WRITE_ACP.
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        public BucketAccountAclGetArgs()
        {
        }
        public static new BucketAccountAclGetArgs Empty => new BucketAccountAclGetArgs();
    }
}
