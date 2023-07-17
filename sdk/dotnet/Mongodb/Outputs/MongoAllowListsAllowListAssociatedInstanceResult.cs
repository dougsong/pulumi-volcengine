// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Mongodb.Outputs
{

    [OutputType]
    public sealed class MongoAllowListsAllowListAssociatedInstanceResult
    {
        /// <summary>
        /// The instance ID to query.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The instance name that bound to the allow list.
        /// </summary>
        public readonly string InstanceName;
        /// <summary>
        /// The VPC ID.
        /// </summary>
        public readonly string Vpc;

        [OutputConstructor]
        private MongoAllowListsAllowListAssociatedInstanceResult(
            string instanceId,

            string instanceName,

            string vpc)
        {
            InstanceId = instanceId;
            InstanceName = instanceName;
            Vpc = vpc;
        }
    }
}
