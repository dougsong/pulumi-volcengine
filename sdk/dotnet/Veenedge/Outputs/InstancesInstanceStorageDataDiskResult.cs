// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceStorageDataDiskResult
    {
        /// <summary>
        /// The capacity of storage.
        /// </summary>
        public readonly string Capacity;
        /// <summary>
        /// The type of storage.
        /// </summary>
        public readonly string StorageType;

        [OutputConstructor]
        private InstancesInstanceStorageDataDiskResult(
            string capacity,

            string storageType)
        {
            Capacity = capacity;
            StorageType = storageType;
        }
    }
}
