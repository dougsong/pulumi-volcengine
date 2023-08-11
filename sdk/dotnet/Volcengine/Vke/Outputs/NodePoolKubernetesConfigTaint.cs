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
    public sealed class NodePoolKubernetesConfigTaint
    {
        /// <summary>
        /// The Effect of Taints, the value can be `NoSchedule` or `NoExecute` or `PreferNoSchedule`.
        /// </summary>
        public readonly string? Effect;
        /// <summary>
        /// The Key of Taints.
        /// </summary>
        public readonly string? Key;
        /// <summary>
        /// The Value of Taints.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private NodePoolKubernetesConfigTaint(
            string? effect,

            string? key,

            string? value)
        {
            Effect = effect;
            Key = key;
            Value = value;
        }
    }
}
