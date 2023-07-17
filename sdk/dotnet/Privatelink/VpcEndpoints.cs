// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Privatelink
{
    public static class VpcEndpoints
    {
        /// <summary>
        /// Use this data source to query detailed information of privatelink vpc endpoints
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
        ///         var @default = Output.Create(Volcengine.Privatelink.VpcEndpoints.InvokeAsync(new Volcengine.Privatelink.VpcEndpointsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "ep-3rel74u229dz45zsk2i6l****",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<VpcEndpointsResult> InvokeAsync(VpcEndpointsArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<VpcEndpointsResult>("volcengine:privatelink/vpcEndpoints:VpcEndpoints", args ?? new VpcEndpointsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of privatelink vpc endpoints
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
        ///         var @default = Output.Create(Volcengine.Privatelink.VpcEndpoints.InvokeAsync(new Volcengine.Privatelink.VpcEndpointsArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "ep-3rel74u229dz45zsk2i6l****",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<VpcEndpointsResult> Invoke(VpcEndpointsInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<VpcEndpointsResult>("volcengine:privatelink/vpcEndpoints:VpcEndpoints", args ?? new VpcEndpointsInvokeArgs(), options.WithDefaults());
    }


    public sealed class VpcEndpointsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of vpc endpoint.
        /// </summary>
        [Input("endpointName")]
        public string? EndpointName { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// The IDs of vpc endpoint.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of vpc endpoint.
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

        /// <summary>
        /// The status of vpc endpoint. Valid values: `Creating`, `Pending`, `Available`, `Deleting`, `Inactive`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        /// <summary>
        /// The vpc id of vpc endpoint.
        /// </summary>
        [Input("vpcId")]
        public string? VpcId { get; set; }

        public VpcEndpointsArgs()
        {
        }
    }

    public sealed class VpcEndpointsInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of vpc endpoint.
        /// </summary>
        [Input("endpointName")]
        public Input<string>? EndpointName { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// The IDs of vpc endpoint.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of vpc endpoint.
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

        /// <summary>
        /// The status of vpc endpoint. Valid values: `Creating`, `Pending`, `Available`, `Deleting`, `Inactive`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The vpc id of vpc endpoint.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        public VpcEndpointsInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class VpcEndpointsResult
    {
        /// <summary>
        /// The name of vpc endpoint.
        /// </summary>
        public readonly string? EndpointName;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The name of vpc endpoint service.
        /// </summary>
        public readonly string? ServiceName;
        /// <summary>
        /// The status of vpc endpoint.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Returns the total amount of the data list.
        /// </summary>
        public readonly int TotalCount;
        /// <summary>
        /// The collection of query.
        /// </summary>
        public readonly ImmutableArray<Outputs.VpcEndpointsVpcEndpointResult> VpcEndpoints;
        /// <summary>
        /// The vpc id of vpc endpoint.
        /// </summary>
        public readonly string? VpcId;

        [OutputConstructor]
        private VpcEndpointsResult(
            string? endpointName,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            string? serviceName,

            string? status,

            int totalCount,

            ImmutableArray<Outputs.VpcEndpointsVpcEndpointResult> vpcEndpoints,

            string? vpcId)
        {
            EndpointName = endpointName;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ServiceName = serviceName;
            Status = status;
            TotalCount = totalCount;
            VpcEndpoints = vpcEndpoints;
            VpcId = vpcId;
        }
    }
}
