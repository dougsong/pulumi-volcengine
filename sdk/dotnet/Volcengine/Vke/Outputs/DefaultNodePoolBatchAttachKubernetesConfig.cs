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
    public sealed class DefaultNodePoolBatchAttachKubernetesConfig
    {
        /// <summary>
        /// The Cordon of KubernetesConfig.
        /// </summary>
        public readonly bool? Cordon;
        /// <summary>
        /// The Labels of KubernetesConfig.
        /// </summary>
        public readonly ImmutableArray<Outputs.DefaultNodePoolBatchAttachKubernetesConfigLabel> Labels;
        /// <summary>
        /// The Taints of KubernetesConfig.
        /// </summary>
        public readonly ImmutableArray<Outputs.DefaultNodePoolBatchAttachKubernetesConfigTaint> Taints;

        [OutputConstructor]
        private DefaultNodePoolBatchAttachKubernetesConfig(
            bool? cordon,

            ImmutableArray<Outputs.DefaultNodePoolBatchAttachKubernetesConfigLabel> labels,

            ImmutableArray<Outputs.DefaultNodePoolBatchAttachKubernetesConfigTaint> taints)
        {
            Cordon = cordon;
            Labels = labels;
            Taints = taints;
        }
    }
}
