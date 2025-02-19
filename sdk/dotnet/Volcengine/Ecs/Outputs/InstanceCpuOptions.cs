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
    public sealed class InstanceCpuOptions
    {
        /// <summary>
        /// The per core of threads.
        /// </summary>
        public readonly int ThreadsPerCore;

        [OutputConstructor]
        private InstanceCpuOptions(int threadsPerCore)
        {
            ThreadsPerCore = threadsPerCore;
        }
    }
}
