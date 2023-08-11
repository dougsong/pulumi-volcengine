// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Autoscaling.Outputs
{

    [OutputType]
    public sealed class ScalingConfigurationVolume
    {
        /// <summary>
        /// The delete with instance flag of volume. Valid values: true, false. Default value: true.
        /// </summary>
        public readonly bool? DeleteWithInstance;
        /// <summary>
        /// The size of volume. System disk value range: 10 - 500. The value range of the data disk: 10 - 8192.
        /// </summary>
        public readonly int Size;
        /// <summary>
        /// The type of volume.
        /// </summary>
        public readonly string VolumeType;

        [OutputConstructor]
        private ScalingConfigurationVolume(
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
