// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Escloud.Outputs
{

    [OutputType]
    public sealed class InstanceInstanceConfigurationNodeSpecsAssign
    {
        /// <summary>
        /// The number of node.
        /// </summary>
        public readonly int Number;
        /// <summary>
        /// The name of compute resource spec, the value is `kibana.x2.small` or `es.x4.medium` or `es.x4.large` or `es.x4.xlarge` or `es.x2.2xlarge` or `es.x4.2xlarge` or `es.x2.3xlarge`.
        /// </summary>
        public readonly string ResourceSpecName;
        /// <summary>
        /// The size of storage. Kibana NodeSpecsAssign should not specify this field.
        /// </summary>
        public readonly int? StorageSize;
        /// <summary>
        /// The name of storage spec. Kibana NodeSpecsAssign should not specify this field.
        /// </summary>
        public readonly string? StorageSpecName;
        /// <summary>
        /// The type of node, the value is `Master` or `Hot` or `Kibana`.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private InstanceInstanceConfigurationNodeSpecsAssign(
            int number,

            string resourceSpecName,

            int? storageSize,

            string? storageSpecName,

            string type)
        {
            Number = number;
            ResourceSpecName = resourceSpecName;
            StorageSize = storageSize;
            StorageSpecName = storageSpecName;
            Type = type;
        }
    }
}
