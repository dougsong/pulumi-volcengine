// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_v2.Inputs
{

    public sealed class RdsInstanceV2ConnectionInfoReadOnlyNodeWeightGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the node.
        /// </summary>
        [Input("nodeId")]
        public Input<string>? NodeId { get; set; }

        /// <summary>
        /// Node type, the value is "Primary", "Secondary", "ReadOnly".
        /// </summary>
        [Input("nodeType")]
        public Input<string>? NodeType { get; set; }

        /// <summary>
        /// The weight of the node.
        /// </summary>
        [Input("weight")]
        public Input<int>? Weight { get; set; }

        public RdsInstanceV2ConnectionInfoReadOnlyNodeWeightGetArgs()
        {
        }
    }
}
