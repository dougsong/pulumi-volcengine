// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Escloud.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceInstanceConfigurationMasterNodeStorageSpecResult
    {
        /// <summary>
        /// The description of plugin.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The show name of storage spec.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The max size of storage spec.
        /// </summary>
        public readonly int MaxSize;
        /// <summary>
        /// The min size of storage spec.
        /// </summary>
        public readonly int MinSize;
        /// <summary>
        /// The name of storage spec.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The size of storage spec.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The type of storage spec.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private InstancesInstanceInstanceConfigurationMasterNodeStorageSpecResult(
            string description,

            string displayName,

            int maxSize,

            int minSize,

            string name,

            int size,

            string type)
        {
            Description = description;
            DisplayName = displayName;
            MaxSize = maxSize;
            MinSize = minSize;
            Name = name;
            Size = size;
            Type = type;
        }
    }
}
