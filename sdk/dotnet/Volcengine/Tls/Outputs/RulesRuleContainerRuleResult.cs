// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class RulesRuleContainerRuleResult
    {
        /// <summary>
        /// The name of the container to be collected.
        /// </summary>
        public readonly string ContainerNameRegex;
        /// <summary>
        /// Whether to add environment variables as log tags to raw log data.
        /// </summary>
        public readonly ImmutableDictionary<string, object> EnvTag;
        /// <summary>
        /// The container environment variable blacklist is used to specify the range of containers not to be collected.
        /// </summary>
        public readonly ImmutableDictionary<string, object> ExcludeContainerEnvRegex;
        /// <summary>
        /// The container Label blacklist is used to specify the range of containers not to be collected.
        /// </summary>
        public readonly ImmutableDictionary<string, object> ExcludeContainerLabelRegex;
        /// <summary>
        /// The container environment variable whitelist specifies the container to be collected through the container environment variable. If the whitelist is not enabled, it means that all containers are specified to be collected.
        /// </summary>
        public readonly ImmutableDictionary<string, object> IncludeContainerEnvRegex;
        /// <summary>
        /// The container label whitelist specifies the containers to be collected through the container label. If the whitelist is not enabled, all containers are specified to be collected.
        /// </summary>
        public readonly ImmutableDictionary<string, object> IncludeContainerLabelRegex;
        /// <summary>
        /// Collection rules for Kubernetes containers.
        /// </summary>
        public readonly ImmutableArray<Outputs.RulesRuleContainerRuleKubernetesRuleResult> KubernetesRules;
        /// <summary>
        /// The collection mode.
        /// </summary>
        public readonly string Stream;

        [OutputConstructor]
        private RulesRuleContainerRuleResult(
            string containerNameRegex,

            ImmutableDictionary<string, object> envTag,

            ImmutableDictionary<string, object> excludeContainerEnvRegex,

            ImmutableDictionary<string, object> excludeContainerLabelRegex,

            ImmutableDictionary<string, object> includeContainerEnvRegex,

            ImmutableDictionary<string, object> includeContainerLabelRegex,

            ImmutableArray<Outputs.RulesRuleContainerRuleKubernetesRuleResult> kubernetesRules,

            string stream)
        {
            ContainerNameRegex = containerNameRegex;
            EnvTag = envTag;
            ExcludeContainerEnvRegex = excludeContainerEnvRegex;
            ExcludeContainerLabelRegex = excludeContainerLabelRegex;
            IncludeContainerEnvRegex = includeContainerEnvRegex;
            IncludeContainerLabelRegex = includeContainerLabelRegex;
            KubernetesRules = kubernetesRules;
            Stream = stream;
        }
    }
}
