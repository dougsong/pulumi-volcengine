// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Ecs.Outputs
{

    [OutputType]
    public sealed class InstanceDataVolume
    {
        /// <summary>
        /// The delete with instance flag of volume.
        /// </summary>
        public readonly bool? DeleteWithInstance;
        /// <summary>
        /// The size of volume. The value range of the data volume size is ESSD_PL0: 10~32768, ESSD_FlexPL: 10~32768, PTSSD: 20~8192.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The type of volume, the value is `PTSSD` or `ESSD_PL0` or `ESSD_PL1` or `ESSD_PL2` or `ESSD_FlexPL`.
        /// </summary>
        public readonly string VolumeType;

        [OutputConstructor]
        private InstanceDataVolume(
            bool? deleteWithInstance,

            int size,

            string volumeType)
        {
            DeleteWithInstance = deleteWithInstance;
            Size = size;
            VolumeType = volumeType;
        }
    }
}
