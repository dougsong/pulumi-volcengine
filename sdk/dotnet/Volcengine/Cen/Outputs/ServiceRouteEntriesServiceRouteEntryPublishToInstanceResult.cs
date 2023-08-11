// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Cen.Outputs
{

    [OutputType]
    public sealed class ServiceRouteEntriesServiceRouteEntryPublishToInstanceResult
    {
        /// <summary>
        /// Cloud service access routes need to publish the network instance ID.
        /// </summary>
        public readonly string InstanceId;
        /// <summary>
        /// The region where the cloud service access route needs to be published.
        /// </summary>
        public readonly string InstanceRegionId;
        /// <summary>
        /// The network instance type that needs to be published for cloud service access routes. The values are as follows: `VPC`, `DCGW`.
        /// </summary>
        public readonly string InstanceType;

        [OutputConstructor]
        private ServiceRouteEntriesServiceRouteEntryPublishToInstanceResult(
            string instanceId,

            string instanceRegionId,

            string instanceType)
        {
            InstanceId = instanceId;
            InstanceRegionId = instanceRegionId;
            InstanceType = instanceType;
        }
    }
}
