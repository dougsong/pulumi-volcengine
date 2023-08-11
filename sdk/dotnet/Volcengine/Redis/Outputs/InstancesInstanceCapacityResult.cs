// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Redis.Outputs
{

    [OutputType]
    public sealed class InstancesInstanceCapacityResult
    {
        /// <summary>
        /// The total memory capacity of the redis instance. Unit: MiB.
        /// </summary>
        public readonly int Total;
        /// <summary>
        /// The used memory capacity of the redis instance. Unit: MiB.
        /// </summary>
        public readonly int Used;

        [OutputConstructor]
        private InstancesInstanceCapacityResult(
            int total,

            int used)
        {
            Total = total;
            Used = used;
        }
    }
}
