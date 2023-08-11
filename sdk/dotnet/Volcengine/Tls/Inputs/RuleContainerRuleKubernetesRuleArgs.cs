// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Tls.Inputs
{

    public sealed class RuleContainerRuleKubernetesRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotationTag")]
        private InputMap<string>? _annotationTag;

        /// <summary>
        /// Whether to add Kubernetes Annotation as a log tag to the raw log data.
        /// </summary>
        public InputMap<string> AnnotationTag
        {
            get => _annotationTag ?? (_annotationTag = new InputMap<string>());
            set => _annotationTag = value;
        }

        [Input("excludePodLabelRegex")]
        private InputMap<string>? _excludePodLabelRegex;

        /// <summary>
        /// Specify the containers not to be collected through the Pod Label blacklist, and not enable means to collect all containers.
        /// </summary>
        public InputMap<string> ExcludePodLabelRegex
        {
            get => _excludePodLabelRegex ?? (_excludePodLabelRegex = new InputMap<string>());
            set => _excludePodLabelRegex = value;
        }

        [Input("includePodLabelRegex")]
        private InputMap<string>? _includePodLabelRegex;

        /// <summary>
        /// The Pod Label whitelist is used to specify containers to be collected. When the Pod Label whitelist is not enabled, it means that all containers are collected.
        /// </summary>
        public InputMap<string> IncludePodLabelRegex
        {
            get => _includePodLabelRegex ?? (_includePodLabelRegex = new InputMap<string>());
            set => _includePodLabelRegex = value;
        }

        [Input("labelTag")]
        private InputMap<string>? _labelTag;

        /// <summary>
        /// Whether to add Kubernetes Label as a log label to the original log data.
        /// </summary>
        public InputMap<string> LabelTag
        {
            get => _labelTag ?? (_labelTag = new InputMap<string>());
            set => _labelTag = value;
        }

        /// <summary>
        /// The name of the Kubernetes Namespace to be collected. If no Namespace name is specified, all containers will be collected. Namespace names support regular matching.
        /// </summary>
        [Input("namespaceNameRegex")]
        public Input<string>? NamespaceNameRegex { get; set; }

        /// <summary>
        /// The Pod name is used to specify the container to be collected. When no Pod name is specified, it means to collect all containers.
        /// </summary>
        [Input("podNameRegex")]
        public Input<string>? PodNameRegex { get; set; }

        /// <summary>
        /// Specify the container to be collected by the name of the workload. When no workload name is specified, all containers are collected. The workload name supports regular matching.
        /// </summary>
        [Input("workloadNameRegex")]
        public Input<string>? WorkloadNameRegex { get; set; }

        /// <summary>
        /// Specify the container to be collected by the type of workload. Only one type can be selected. When no type is specified, it means to collect all types of containers.
        /// </summary>
        [Input("workloadType")]
        public Input<string>? WorkloadType { get; set; }

        public RuleContainerRuleKubernetesRuleArgs()
        {
        }
        public static new RuleContainerRuleKubernetesRuleArgs Empty => new RuleContainerRuleKubernetesRuleArgs();
    }
}
