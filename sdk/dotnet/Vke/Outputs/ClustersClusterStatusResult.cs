// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Outputs
{

    [OutputType]
    public sealed class ClustersClusterStatusResult
    {
        /// <summary>
        /// The state condition in the current primary state of the cluster, that is, the reason for entering the primary state.
        /// </summary>
        public readonly ImmutableArray<Outputs.ClustersClusterStatusConditionResult> Conditions;
        /// <summary>
        /// The status of cluster. the value contains `Creating`, `Running`, `Updating`, `Deleting`, `Stopped`, `Failed`.
        /// </summary>
        public readonly string Phase;

        [OutputConstructor]
        private ClustersClusterStatusResult(
            ImmutableArray<Outputs.ClustersClusterStatusConditionResult> conditions,

            string phase)
        {
            Conditions = conditions;
            Phase = phase;
        }
    }
}
