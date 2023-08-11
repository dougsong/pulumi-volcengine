// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Cen.Inputs
{

    public sealed class ServiceRouteEntryPublishToInstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cloud service access routes need to publish the network instance ID.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The region where the cloud service access route needs to be published.
        /// </summary>
        [Input("instanceRegionId")]
        public Input<string>? InstanceRegionId { get; set; }

        /// <summary>
        /// The network instance type that needs to be published for cloud service access routes. The values are as follows: `VPC`, `DCGW`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        public ServiceRouteEntryPublishToInstanceArgs()
        {
        }
        public static new ServiceRouteEntryPublishToInstanceArgs Empty => new ServiceRouteEntryPublishToInstanceArgs();
    }
}
