// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Veenedge.Outputs
{

    [OutputType]
    public sealed class CloudServerStorageConfig
    {
        /// <summary>
        /// The disk list info of data.
        /// </summary>
        public readonly ImmutableArray<Outputs.CloudServerStorageConfigDataDiskList> DataDiskLists;
        /// <summary>
        /// The disk info of system.
        /// </summary>
        public readonly Outputs.CloudServerStorageConfigSystemDisk SystemDisk;

        [OutputConstructor]
        private CloudServerStorageConfig(
            ImmutableArray<Outputs.CloudServerStorageConfigDataDiskList> dataDiskLists,

            Outputs.CloudServerStorageConfigSystemDisk systemDisk)
        {
            DataDiskLists = dataDiskLists;
            SystemDisk = systemDisk;
        }
    }
}
