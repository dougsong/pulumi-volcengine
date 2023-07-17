// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Veenedge
{
    public static class InstanceTypes
    {
        /// <summary>
        /// Use this data source to query detailed information of veenedge instance types
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Volcengine.Veenedge.InstanceTypes.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<InstanceTypesResult> InvokeAsync(InstanceTypesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<InstanceTypesResult>("volcengine:veenedge/instanceTypes:InstanceTypes", args ?? new InstanceTypesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of veenedge instance types
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var @default = Output.Create(Volcengine.Veenedge.InstanceTypes.InvokeAsync());
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<InstanceTypesResult> Invoke(InstanceTypesInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<InstanceTypesResult>("volcengine:veenedge/instanceTypes:InstanceTypes", args ?? new InstanceTypesInvokeArgs(), options.WithDefaults());
    }


    public sealed class InstanceTypesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public InstanceTypesArgs()
        {
        }
    }

    public sealed class InstanceTypesInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public InstanceTypesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class InstanceTypesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The collection of instance types query.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceTypesInstanceTypeConfigResult> InstanceTypeConfigs;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of instance types query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private InstanceTypesResult(
            string id,

            ImmutableArray<Outputs.InstanceTypesInstanceTypeConfigResult> instanceTypeConfigs,

            string? outputFile,

            int totalCount)
        {
            Id = id;
            InstanceTypeConfigs = instanceTypeConfigs;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
