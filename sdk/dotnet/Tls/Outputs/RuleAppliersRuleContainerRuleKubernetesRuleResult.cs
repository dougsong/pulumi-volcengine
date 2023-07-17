// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class RuleAppliersRuleContainerRuleKubernetesRuleResult
    {
        /// <summary>
        /// Whether to add Kubernetes Annotation as a log tag to the raw log data.
        /// </summary>
        public readonly ImmutableDictionary<string, object> AnnotationTag;
        /// <summary>
        /// Specify the containers not to be collected through the Pod Label blacklist, and not enable means to collect all containers.
        /// </summary>
        public readonly ImmutableDictionary<string, object> ExcludePodLabelRegex;
        /// <summary>
        /// The Pod Label whitelist is used to specify containers to be collected. When the Pod Label whitelist is not enabled, it means that all containers are collected.
        /// </summary>
        public readonly ImmutableDictionary<string, object> IncludePodLabelRegex;
        /// <summary>
        /// Whether to add Kubernetes Label as a log label to the original log data.
        /// </summary>
        public readonly ImmutableDictionary<string, object> LabelTag;
        /// <summary>
        /// The name of the Kubernetes Namespace to be collected. If no Namespace name is specified, all containers will be collected. Namespace names support regular matching.
        /// </summary>
        public readonly string NamespaceNameRegex;
        /// <summary>
        /// The Pod name is used to specify the container to be collected. When no Pod name is specified, it means to collect all containers.
        /// </summary>
        public readonly string PodNameRegex;
        /// <summary>
        /// Specify the container to be collected by the name of the workload. When no workload name is specified, all containers are collected. The workload name supports regular matching.
        /// </summary>
        public readonly string WorkloadNameRegex;
        /// <summary>
        /// Specify the container to be collected by the type of workload. Only one type can be selected. When no type is specified, it means to collect all types of containers.
        /// </summary>
        public readonly string WorkloadType;

        [OutputConstructor]
        private RuleAppliersRuleContainerRuleKubernetesRuleResult(
            ImmutableDictionary<string, object> annotationTag,

            ImmutableDictionary<string, object> excludePodLabelRegex,

            ImmutableDictionary<string, object> includePodLabelRegex,

            ImmutableDictionary<string, object> labelTag,

            string namespaceNameRegex,

            string podNameRegex,

            string workloadNameRegex,

            string workloadType)
        {
            AnnotationTag = annotationTag;
            ExcludePodLabelRegex = excludePodLabelRegex;
            IncludePodLabelRegex = includePodLabelRegex;
            LabelTag = labelTag;
            NamespaceNameRegex = namespaceNameRegex;
            PodNameRegex = podNameRegex;
            WorkloadNameRegex = workloadNameRegex;
            WorkloadType = workloadType;
        }
    }
}
