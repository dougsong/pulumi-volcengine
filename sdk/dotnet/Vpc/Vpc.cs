// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc
{
    /// <summary>
    /// Provides a resource to manage vpc
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Pulumi.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Volcengine.Vpc.Vpc("foo", new Volcengine.Vpc.VpcArgs
    ///         {
    ///             CidrBlock = "172.16.0.0/16",
    ///             DnsServers = 
    ///             {
    ///                 "8.8.8.8",
    ///                 "114.114.114.114",
    ///             },
    ///             VpcName = "tf-test-2",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VPC can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:Vpc/vpc:Vpc default vpc-mizl7m1kqccg5smt1bdpijuj
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:Vpc/vpc:Vpc")]
    public partial class Vpc : Pulumi.CustomResource
    {
        /// <summary>
        /// The account ID of VPC.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The associate cen list of VPC.
        /// </summary>
        [Output("associateCens")]
        public Output<ImmutableArray<Outputs.VpcAssociateCen>> AssociateCens { get; private set; } = null!;

        /// <summary>
        /// The auxiliary cidr block list of VPC.
        /// </summary>
        [Output("auxiliaryCidrBlocks")]
        public Output<ImmutableArray<string>> AuxiliaryCidrBlocks { get; private set; } = null!;

        /// <summary>
        /// A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
        /// </summary>
        [Output("cidrBlock")]
        public Output<string> CidrBlock { get; private set; } = null!;

        /// <summary>
        /// Creation time of VPC.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The description of the VPC.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
        /// </summary>
        [Output("dnsServers")]
        public Output<ImmutableArray<string>> DnsServers { get; private set; } = null!;

        /// <summary>
        /// The nat gateway ID list of VPC.
        /// </summary>
        [Output("natGatewayIds")]
        public Output<ImmutableArray<string>> NatGatewayIds { get; private set; } = null!;

        /// <summary>
        /// The route table ID list of VPC.
        /// </summary>
        [Output("routeTableIds")]
        public Output<ImmutableArray<string>> RouteTableIds { get; private set; } = null!;

        /// <summary>
        /// The security group ID list of VPC.
        /// </summary>
        [Output("securityGroupIds")]
        public Output<ImmutableArray<string>> SecurityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Status of VPC.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The subnet ID list of VPC.
        /// </summary>
        [Output("subnetIds")]
        public Output<ImmutableArray<string>> SubnetIds { get; private set; } = null!;

        /// <summary>
        /// The update time of VPC.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The ID of VPC.
        /// </summary>
        [Output("vpcId")]
        public Output<string> VpcId { get; private set; } = null!;

        /// <summary>
        /// The name of the VPC.
        /// </summary>
        [Output("vpcName")]
        public Output<string> VpcName { get; private set; } = null!;


        /// <summary>
        /// Create a Vpc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vpc(string name, VpcArgs args, CustomResourceOptions? options = null)
            : base("volcengine:Vpc/vpc:Vpc", name, args ?? new VpcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vpc(string name, Input<string> id, VpcState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:Vpc/vpc:Vpc", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Vpc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vpc Get(string name, Input<string> id, VpcState? state = null, CustomResourceOptions? options = null)
        {
            return new Vpc(name, id, state, options);
        }
    }

    public sealed class VpcArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
        /// </summary>
        [Input("cidrBlock", required: true)]
        public Input<string> CidrBlock { get; set; } = null!;

        /// <summary>
        /// The description of the VPC.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// The name of the VPC.
        /// </summary>
        [Input("vpcName")]
        public Input<string>? VpcName { get; set; }

        public VpcArgs()
        {
        }
    }

    public sealed class VpcState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account ID of VPC.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        [Input("associateCens")]
        private InputList<Inputs.VpcAssociateCenGetArgs>? _associateCens;

        /// <summary>
        /// The associate cen list of VPC.
        /// </summary>
        public InputList<Inputs.VpcAssociateCenGetArgs> AssociateCens
        {
            get => _associateCens ?? (_associateCens = new InputList<Inputs.VpcAssociateCenGetArgs>());
            set => _associateCens = value;
        }

        [Input("auxiliaryCidrBlocks")]
        private InputList<string>? _auxiliaryCidrBlocks;

        /// <summary>
        /// The auxiliary cidr block list of VPC.
        /// </summary>
        public InputList<string> AuxiliaryCidrBlocks
        {
            get => _auxiliaryCidrBlocks ?? (_auxiliaryCidrBlocks = new InputList<string>());
            set => _auxiliaryCidrBlocks = value;
        }

        /// <summary>
        /// A network address block which should be a subnet of the three internal network segments (10.0.0.0/16, 172.16.0.0/12 and 192.168.0.0/16).
        /// </summary>
        [Input("cidrBlock")]
        public Input<string>? CidrBlock { get; set; }

        /// <summary>
        /// Creation time of VPC.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// The description of the VPC.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// The DNS server list of the VPC. And you can specify 0 to 5 servers to this list.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        [Input("natGatewayIds")]
        private InputList<string>? _natGatewayIds;

        /// <summary>
        /// The nat gateway ID list of VPC.
        /// </summary>
        public InputList<string> NatGatewayIds
        {
            get => _natGatewayIds ?? (_natGatewayIds = new InputList<string>());
            set => _natGatewayIds = value;
        }

        [Input("routeTableIds")]
        private InputList<string>? _routeTableIds;

        /// <summary>
        /// The route table ID list of VPC.
        /// </summary>
        public InputList<string> RouteTableIds
        {
            get => _routeTableIds ?? (_routeTableIds = new InputList<string>());
            set => _routeTableIds = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The security group ID list of VPC.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        /// <summary>
        /// Status of VPC.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("subnetIds")]
        private InputList<string>? _subnetIds;

        /// <summary>
        /// The subnet ID list of VPC.
        /// </summary>
        public InputList<string> SubnetIds
        {
            get => _subnetIds ?? (_subnetIds = new InputList<string>());
            set => _subnetIds = value;
        }

        /// <summary>
        /// The update time of VPC.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// The ID of VPC.
        /// </summary>
        [Input("vpcId")]
        public Input<string>? VpcId { get; set; }

        /// <summary>
        /// The name of the VPC.
        /// </summary>
        [Input("vpcName")]
        public Input<string>? VpcName { get; set; }

        public VpcState()
        {
        }
    }
}
