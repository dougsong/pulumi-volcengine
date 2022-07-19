// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Inputs
{

    public sealed class NodePoolStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates the status condition of the node pool in the active state.
        /// </summary>
        [Input("conditionsType")]
        public Input<string>? ConditionsType { get; set; }

        /// <summary>
        /// The Phase of Status.
        /// </summary>
        [Input("phase")]
        public Input<string>? Phase { get; set; }

        public NodePoolStatusArgs()
        {
        }
    }
}
