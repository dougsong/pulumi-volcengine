// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpn
{
    public static class CustomerGateways
    {
        /// <summary>
        /// Use this data source to query detailed information of customer gateways
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
        ///         var foo = Output.Create(Volcengine.Vpn.CustomerGateways.InvokeAsync(new Volcengine.Vpn.CustomerGatewaysArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "cgw-2d68c4zglycjk58ozfe96norh",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<CustomerGatewaysResult> InvokeAsync(CustomerGatewaysArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<CustomerGatewaysResult>("volcengine:vpn/customerGateways:CustomerGateways", args ?? new CustomerGatewaysArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of customer gateways
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
        ///         var foo = Output.Create(Volcengine.Vpn.CustomerGateways.InvokeAsync(new Volcengine.Vpn.CustomerGatewaysArgs
        ///         {
        ///             Ids = 
        ///             {
        ///                 "cgw-2d68c4zglycjk58ozfe96norh",
        ///             },
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<CustomerGatewaysResult> Invoke(CustomerGatewaysInvokeArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<CustomerGatewaysResult>("volcengine:vpn/customerGateways:CustomerGateways", args ?? new CustomerGatewaysInvokeArgs(), options.WithDefaults());
    }


    public sealed class CustomerGatewaysArgs : Pulumi.InvokeArgs
    {
        [Input("customerGatewayNames")]
        private List<string>? _customerGatewayNames;

        /// <summary>
        /// A list of customer gateway names.
        /// </summary>
        public List<string> CustomerGatewayNames
        {
            get => _customerGatewayNames ?? (_customerGatewayNames = new List<string>());
            set => _customerGatewayNames = value;
        }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of customer gateway ids.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A IP address of the customer gateway.
        /// </summary>
        [Input("ipAddress")]
        public string? IpAddress { get; set; }

        /// <summary>
        /// A Name Regex of customer gateway.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        public CustomerGatewaysArgs()
        {
        }
    }

    public sealed class CustomerGatewaysInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("customerGatewayNames")]
        private InputList<string>? _customerGatewayNames;

        /// <summary>
        /// A list of customer gateway names.
        /// </summary>
        public InputList<string> CustomerGatewayNames
        {
            get => _customerGatewayNames ?? (_customerGatewayNames = new InputList<string>());
            set => _customerGatewayNames = value;
        }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of customer gateway ids.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A IP address of the customer gateway.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// A Name Regex of customer gateway.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        public CustomerGatewaysInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class CustomerGatewaysResult
    {
        public readonly ImmutableArray<string> CustomerGatewayNames;
        /// <summary>
        /// The collection of customer gateway query.
        /// </summary>
        public readonly ImmutableArray<Outputs.CustomerGatewaysCustomerGatewayResult> CustomerGateways;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// The IP address of the customer gateway.
        /// </summary>
        public readonly string? IpAddress;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The total count of customer gateway query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private CustomerGatewaysResult(
            ImmutableArray<string> customerGatewayNames,

            ImmutableArray<Outputs.CustomerGatewaysCustomerGatewayResult> customerGateways,

            string id,

            ImmutableArray<string> ids,

            string? ipAddress,

            string? nameRegex,

            string? outputFile,

            int totalCount)
        {
            CustomerGatewayNames = customerGatewayNames;
            CustomerGateways = customerGateways;
            Id = id;
            Ids = ids;
            IpAddress = ipAddress;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            TotalCount = totalCount;
        }
    }
}
