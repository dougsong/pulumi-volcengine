// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Redis
{
    public static class Regions
    {
        /// <summary>
        /// Use this data source to query detailed information of redis regions
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
        ///     var @default = Volcengine.Redis.Regions.Invoke(new()
        ///     {
        ///         RegionId = "cn-north-3",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<RegionsResult> InvokeAsync(RegionsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<RegionsResult>("volcengine:redis/regions:Regions", args ?? new RegionsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of redis regions
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
        ///     var @default = Volcengine.Redis.Regions.Invoke(new()
        ///     {
        ///         RegionId = "cn-north-3",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<RegionsResult> Invoke(RegionsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<RegionsResult>("volcengine:redis/regions:Regions", args ?? new RegionsInvokeArgs(), options.WithDefaults());
    }


    public sealed class RegionsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// Target region info.
        /// </summary>
        [Input("regionId")]
        public string? RegionId { get; set; }

        public RegionsArgs()
        {
        }
        public static new RegionsArgs Empty => new RegionsArgs();
    }

    public sealed class RegionsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// Target region info.
        /// </summary>
        [Input("regionId")]
        public Input<string>? RegionId { get; set; }

        public RegionsInvokeArgs()
        {
        }
        public static new RegionsInvokeArgs Empty => new RegionsInvokeArgs();
    }


    [OutputType]
    public sealed class RegionsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        /// <summary>
        /// The id of the region.
        /// </summary>
        public readonly string? RegionId;
        /// <summary>
        /// The collection of region query.
        /// </summary>
        public readonly ImmutableArray<Outputs.RegionsRegionResult> Regions;
        /// <summary>
        /// The total count of region query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private RegionsResult(
            string id,

            string? outputFile,

            string? regionId,

            ImmutableArray<Outputs.RegionsRegionResult> regions,

            int totalCount)
        {
            Id = id;
            OutputFile = outputFile;
            RegionId = regionId;
            Regions = regions;
            TotalCount = totalCount;
        }
    }
}
