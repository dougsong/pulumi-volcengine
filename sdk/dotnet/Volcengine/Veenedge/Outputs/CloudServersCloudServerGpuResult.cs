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
    public sealed class CloudServersCloudServerGpuResult
    {
        /// <summary>
        /// The list gpu info.
        /// </summary>
        public readonly ImmutableArray<Outputs.CloudServersCloudServerGpuGpusResult> Gpuses;

        [OutputConstructor]
        private CloudServersCloudServerGpuResult(ImmutableArray<Outputs.CloudServersCloudServerGpuGpusResult> gpuses)
        {
            Gpuses = gpuses;
        }
    }
}
