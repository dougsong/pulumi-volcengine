// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Privatelink.Inputs
{

    public sealed class VpcEndpointServiceResourceGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        /// <summary>
        /// The type of resource.
        /// </summary>
        [Input("resourceType", required: true)]
        public Input<string> ResourceType { get; set; } = null!;

        public VpcEndpointServiceResourceGetArgs()
        {
        }
        public static new VpcEndpointServiceResourceGetArgs Empty => new VpcEndpointServiceResourceGetArgs();
    }
}
