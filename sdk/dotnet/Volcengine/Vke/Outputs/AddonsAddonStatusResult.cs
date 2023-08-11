// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vke.Outputs
{

    [OutputType]
    public sealed class AddonsAddonStatusResult
    {
        /// <summary>
        /// The state condition in the current primary state of the cluster, that is, the reason for entering the primary state.
        /// </summary>
        public readonly ImmutableArray<Outputs.AddonsAddonStatusConditionResult> Conditions;
        /// <summary>
        /// The status of addon. the value contains `Creating`, `Running`, `Updating`, `Deleting`, `Failed`.
        /// </summary>
        public readonly string Phase;

        [OutputConstructor]
        private AddonsAddonStatusResult(
            ImmutableArray<Outputs.AddonsAddonStatusConditionResult> conditions,

            string phase)
        {
            Conditions = conditions;
            Phase = phase;
        }
    }
}
