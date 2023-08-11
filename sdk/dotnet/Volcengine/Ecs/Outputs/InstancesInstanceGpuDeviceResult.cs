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
    public sealed class InstancesInstanceGpuDeviceResult
    {
        /// <summary>
        /// The Count of GPU device.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// The Encrypted Memory Size of GPU device.
        /// </summary>
        public readonly int EncryptedMemorySize;
        /// <summary>
        /// The memory size of ECS instance.
        /// </summary>
        public readonly int MemorySize;
        /// <summary>
        /// The Product Name of GPU device.
        /// </summary>
        public readonly string ProductName;

        [OutputConstructor]
        private InstancesInstanceGpuDeviceResult(
            int count,

            int encryptedMemorySize,

            int memorySize,

            string productName)
        {
            Count = count;
            EncryptedMemorySize = encryptedMemorySize;
            MemorySize = memorySize;
            ProductName = productName;
        }
    }
}
