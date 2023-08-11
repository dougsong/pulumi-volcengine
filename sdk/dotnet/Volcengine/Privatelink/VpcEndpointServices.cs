// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Privatelink
{
    public static class VpcEndpointServices
    {
        /// <summary>
        /// Use this data source to query detailed information of privatelink vpc endpoint services
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
        ///     var @default = Volcengine.Privatelink.VpcEndpointServices.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "epsvc-3rel73uf2ewao5zsk2j2l58ro",
        ///             "epsvc-2d72mxjgq02yo58ozfe5tndeh",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<VpcEndpointServicesResult> InvokeAsync(VpcEndpointServicesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<VpcEndpointServicesResult>("volcengine:privatelink/vpcEndpointServices:VpcEndpointServices", args ?? new VpcEndpointServicesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of privatelink vpc endpoint services
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
        ///     var @default = Volcengine.Privatelink.VpcEndpointServices.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "epsvc-3rel73uf2ewao5zsk2j2l58ro",
        ///             "epsvc-2d72mxjgq02yo58ozfe5tndeh",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<VpcEndpointServicesResult> Invoke(VpcEndpointServicesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<VpcEndpointServicesResult>("volcengine:privatelink/vpcEndpointServices:VpcEndpointServices", args ?? new VpcEndpointServicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class VpcEndpointServicesArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// The IDs of vpc endpoint service.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of vpc endpoint service.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The name of vpc endpoint service.
        /// </summary>
        [Input("serviceName")]
        public string? ServiceName { get; set; }

        public VpcEndpointServicesArgs()
        {
        }
        public static new VpcEndpointServicesArgs Empty => new VpcEndpointServicesArgs();
    }

    public sealed class VpcEndpointServicesInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// The IDs of vpc endpoint service.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of vpc endpoint service.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The name of vpc endpoint service.
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public VpcEndpointServicesInvokeArgs()
        {
        }
        public static new VpcEndpointServicesInvokeArgs Empty => new VpcEndpointServicesInvokeArgs();
    }


    [OutputType]
    public sealed class VpcEndpointServicesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The name of service.
        /// </summary>
        public readonly string? ServiceName;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpcEndpointServicesServiceResult> Services;
        /// <summary>
        /// Returns the total amount of the data list.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private VpcEndpointServicesResult(
            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            string? serviceName,

            ImmutableArray<Outputs.VpcEndpointServicesServiceResult> services,

            int totalCount)
        {
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ServiceName = serviceName;
            Services = services;
            TotalCount = totalCount;
        }
    }
}
