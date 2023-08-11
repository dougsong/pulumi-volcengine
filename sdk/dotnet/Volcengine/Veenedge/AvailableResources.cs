// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Veenedge
{
    public static class AvailableResources
    {
        /// <summary>
        /// Use this data source to query detailed information of veenedge available resources
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Volcengine.Veenedge.AvailableResources.Invoke(new()
        ///     {
        ///         BandwithLimit = 20,
        ///         CloudDiskType = "CloudSSD",
        ///         InstanceType = "ve******rge",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<AvailableResourcesResult> InvokeAsync(AvailableResourcesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<AvailableResourcesResult>("volcengine:veenedge/availableResources:AvailableResources", args ?? new AvailableResourcesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of veenedge available resources
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Volcengine = Pulumi.Volcengine;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var @default = Volcengine.Veenedge.AvailableResources.Invoke(new()
        ///     {
        ///         BandwithLimit = 20,
        ///         CloudDiskType = "CloudSSD",
        ///         InstanceType = "ve******rge",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<AvailableResourcesResult> Invoke(AvailableResourcesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<AvailableResourcesResult>("volcengine:veenedge/availableResources:AvailableResources", args ?? new AvailableResourcesInvokeArgs(), options.WithDefaults());
    }


    public sealed class AvailableResourcesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The limit of bandwidth.
        /// </summary>
        [Input("bandwithLimit", required: true)]
        public int BandwithLimit { get; set; }

        /// <summary>
        /// The type of storage. The value can be `CloudHDD` or `CloudSSD`.
        /// </summary>
        [Input("cloudDiskType", required: true)]
        public string CloudDiskType { get; set; } = null!;

        /// <summary>
        /// The type of instance.
        /// </summary>
        [Input("instanceType", required: true)]
        public string InstanceType { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public AvailableResourcesArgs()
        {
        }
        public static new AvailableResourcesArgs Empty => new AvailableResourcesArgs();
    }

    public sealed class AvailableResourcesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The limit of bandwidth.
        /// </summary>
        [Input("bandwithLimit", required: true)]
        public Input<int> BandwithLimit { get; set; } = null!;

        /// <summary>
        /// The type of storage. The value can be `CloudHDD` or `CloudSSD`.
        /// </summary>
        [Input("cloudDiskType", required: true)]
        public Input<string> CloudDiskType { get; set; } = null!;

        /// <summary>
        /// The type of instance.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public AvailableResourcesInvokeArgs()
        {
        }
        public static new AvailableResourcesInvokeArgs Empty => new AvailableResourcesInvokeArgs();
    }


    [OutputType]
    public sealed class AvailableResourcesResult
    {
        public readonly int BandwithLimit;
        public readonly string CloudDiskType;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InstanceType;
        public readonly string? OutputFile;
        /// <summary>
        /// The collection of resource query.
        /// </summary>
        public readonly ImmutableArray<Outputs.AvailableResourcesRegionResult> Regions;
        /// <summary>
        /// The total count of resource query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private AvailableResourcesResult(
            int bandwithLimit,

            string cloudDiskType,

            string id,

            string instanceType,

            string? outputFile,

            ImmutableArray<Outputs.AvailableResourcesRegionResult> regions,

            int totalCount)
        {
            BandwithLimit = bandwithLimit;
            CloudDiskType = cloudDiskType;
            Id = id;
            InstanceType = instanceType;
            OutputFile = outputFile;
            Regions = regions;
            TotalCount = totalCount;
        }
    }
}
