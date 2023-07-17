// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Redis.Inputs
{

    public sealed class AllowListAssociatedInstanceGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Name of instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// Id of virtual private cloud.
        /// </summary>
        [Input("vpc")]
        public Input<string>? Vpc { get; set; }

        public AllowListAssociatedInstanceGetArgs()
        {
        }
    }
}
