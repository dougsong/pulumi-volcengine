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
    public sealed class EndpointsEndpointResult
    {
        /// <summary>
        /// The list of mongodb addresses.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointsEndpointDbAddressResult> DbAddresses;
        /// <summary>
        /// The ID of endpoint.
        /// </summary>
        public readonly string EndpointId;
        /// <summary>
        /// The endpoint information.
        /// </summary>
        public readonly string EndpointStr;
        /// <summary>
        /// The node type corresponding to the endpoint.
        /// </summary>
        public readonly string EndpointType;
        /// <summary>
        /// The network type of endpoint.
        /// </summary>
        public readonly string NetworkType;
        /// <summary>
        /// The object ID corresponding to the endpoint.
        /// </summary>
        public readonly string ObjectId;
        /// <summary>
        /// The subnet ID.
        /// </summary>
        public readonly string SubnetId;
        /// <summary>
        /// The VPC ID.
        /// </summary>
        public readonly string VpcId;

        [OutputConstructor]
        private EndpointsEndpointResult(
            ImmutableArray<Outputs.EndpointsEndpointDbAddressResult> dbAddresses,

            string endpointId,

            string endpointStr,

            string endpointType,

            string networkType,

            string objectId,

            string subnetId,

            string vpcId)
        {
            DbAddresses = dbAddresses;
            EndpointId = endpointId;
            EndpointStr = endpointStr;
            EndpointType = endpointType;
            NetworkType = networkType;
            ObjectId = objectId;
            SubnetId = subnetId;
            VpcId = vpcId;
        }
    }
}
