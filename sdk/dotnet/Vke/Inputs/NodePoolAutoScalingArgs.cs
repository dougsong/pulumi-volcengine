// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Inputs
{

    public sealed class NodePoolAutoScalingArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DesiredReplicas of AutoScaling, default 0, range in min_replicas to max_replicas.
        /// </summary>
        [Input("desiredReplicas")]
        public Input<int>? DesiredReplicas { get; set; }

        /// <summary>
        /// Is Enabled of AutoScaling.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The MaxReplicas of AutoScaling, default 10, range in 1~2000.
        /// </summary>
        [Input("maxReplicas")]
        public Input<int>? MaxReplicas { get; set; }

        /// <summary>
        /// The MinReplicas of AutoScaling, default 0.
        /// </summary>
        [Input("minReplicas")]
        public Input<int>? MinReplicas { get; set; }

        /// <summary>
        /// The Priority of AutoScaling, default 10, rang in 0~100.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// Multi-subnet scheduling strategy for nodes. The value can be `ZoneBalance` or `Priority`.
        /// </summary>
        [Input("subnetPolicy")]
        public Input<string>? SubnetPolicy { get; set; }

        public NodePoolAutoScalingArgs()
        {
        }
    }
}
