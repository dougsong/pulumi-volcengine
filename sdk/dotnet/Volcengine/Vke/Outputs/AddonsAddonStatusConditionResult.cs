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
    public sealed class AddonsAddonStatusConditionResult
    {
        /// <summary>
        /// The state condition in the current main state of the addon, that is, the reason for entering the main state, there can be multiple reasons, the value contains `Progressing`, `Ok`, `Degraded`,`Unknown`, `ClusterNotRunning`, `CrashLoopBackOff`, `SchedulingFailed`, `NameConflict`, `ResourceCleanupFailed`, `ClusterVersionUpgrading`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private AddonsAddonStatusConditionResult(string type)
        {
            Type = type;
        }
    }
}
