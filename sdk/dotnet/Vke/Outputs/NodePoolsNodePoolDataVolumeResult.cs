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
    public sealed class NodePoolsNodePoolDataVolumeResult
    {
        /// <summary>
        /// The target mount directory of the disk.
        /// </summary>
        public readonly string MountPoint;
        /// <summary>
        /// The Size of SystemVolume.
        /// </summary>
        public readonly string Size;
        /// <summary>
        /// The Type of Tags.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private NodePoolsNodePoolDataVolumeResult(
            string mountPoint,

            string size,

            string type)
        {
            MountPoint = mountPoint;
            Size = size;
            Type = type;
        }
    }
}
