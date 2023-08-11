// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vke.Inputs
{

    public sealed class NodePoolsStatusArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Indicates the status condition of the node pool in the active state. The value can be `Progressing` or `Ok` or `VersionPartlyUpgraded` or `StockOut` or `LimitedByQuota` or `Balance` or `Degraded` or `ClusterVersionUpgrading` or `Cluster` or `ResourceCleanupFailed` or `Unknown` or `ClusterNotRunning` or `SetByProvider`.
        /// </summary>
        [Input("conditionsType")]
        public string? ConditionsType { get; set; }

        /// <summary>
        /// The Phase of Status. The value can be `Creating` or `Running` or `Updating` or `Deleting` or `Failed` or `Scaling`.
        /// </summary>
        [Input("phase")]
        public string? Phase { get; set; }

        public NodePoolsStatusArgs()
        {
        }
        public static new NodePoolsStatusArgs Empty => new NodePoolsStatusArgs();
    }
}
