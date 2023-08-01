// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Vpn
{
    /// <summary>
    /// Provides a resource to manage vpn connection
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Volcengine.PulumiPackage.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var foo = new Volcengine.Vpn.Connection("foo", new Volcengine.Vpn.ConnectionArgs
    ///         {
    ///             CustomerGatewayId = "cgw-12ayj1s157gn417q7y29bixqy",
    ///             Description = "tf-test",
    ///             DpdAction = "none",
    ///             IkeConfigAuthAlg = "md5",
    ///             IkeConfigDhGroup = "group2",
    ///             IkeConfigEncAlg = "aes",
    ///             IkeConfigLifetime = 9000,
    ///             IkeConfigLocalId = "tf_test",
    ///             IkeConfigMode = "main",
    ///             IkeConfigPsk = "tftest@!3",
    ///             IkeConfigRemoteId = "tf_test",
    ///             IkeConfigVersion = "ikev1",
    ///             IpsecConfigAuthAlg = "sha256",
    ///             IpsecConfigDhGroup = "group2",
    ///             IpsecConfigEncAlg = "aes",
    ///             IpsecConfigLifetime = 9000,
    ///             LocalSubnets = 
    ///             {
    ///                 "192.168.0.0/22",
    ///             },
    ///             NatTraversal = true,
    ///             ProjectName = "default",
    ///             RemoteSubnets = 
    ///             {
    ///                 "192.161.0.0/20",
    ///             },
    ///             VpnConnectionName = "tf-test",
    ///             VpnGatewayId = "vgw-2feq19gnyc9hc59gp68914u6o",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// VpnConnection can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:vpn/connection:Connection default vgc-3tex2x1cwd4c6c0v****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:vpn/connection:Connection")]
    public partial class Connection : Pulumi.CustomResource
    {
        /// <summary>
        /// The account ID of the VPN connection.
        /// </summary>
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// The IPsec attach status.
        /// </summary>
        [Output("attachStatus")]
        public Output<string> AttachStatus { get; private set; } = null!;

        /// <summary>
        /// The attach type of the VPN connection, the value can be `VpnGateway` or `TransitRouter`.
        /// </summary>
        [Output("attachType")]
        public Output<string?> AttachType { get; private set; } = null!;

        /// <summary>
        /// The business status of IPsec connection, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Output("businessStatus")]
        public Output<string> BusinessStatus { get; private set; } = null!;

        /// <summary>
        /// The connect status of the VPN connection.
        /// </summary>
        [Output("connectStatus")]
        public Output<string> ConnectStatus { get; private set; } = null!;

        /// <summary>
        /// The create time of VPN connection.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the customer gateway.
        /// </summary>
        [Output("customerGatewayId")]
        public Output<string> CustomerGatewayId { get; private set; } = null!;

        /// <summary>
        /// The delete time of resource, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Output("deletedTime")]
        public Output<string> DeletedTime { get; private set; } = null!;

        /// <summary>
        /// The description of the VPN connection.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// The dpd action of the VPN connection.
        /// </summary>
        [Output("dpdAction")]
        public Output<string?> DpdAction { get; private set; } = null!;

        /// <summary>
        /// The auth alg of the ike config of the VPN connection. Valid value are `sha1`, `md5`, `sha256`, `sha384`, `sha512`, `sm3`. The default value is `sha1`.
        /// </summary>
        [Output("ikeConfigAuthAlg")]
        public Output<string?> IkeConfigAuthAlg { get; private set; } = null!;

        /// <summary>
        /// The dk group of the ike config of the VPN connection. Valid value are `group1`, `group2`, `group5`, `group14`. The default value is `group2`.
        /// </summary>
        [Output("ikeConfigDhGroup")]
        public Output<string?> IkeConfigDhGroup { get; private set; } = null!;

        /// <summary>
        /// The enc alg of the ike config of the VPN connection. Valid value are `aes`, `aes192`, `aes256`, `des`, `3des`, `sm4`. The default value is `aes`.
        /// </summary>
        [Output("ikeConfigEncAlg")]
        public Output<string?> IkeConfigEncAlg { get; private set; } = null!;

        /// <summary>
        /// The lifetime of the ike config of the VPN connection. Value: 900~86400.
        /// </summary>
        [Output("ikeConfigLifetime")]
        public Output<int?> IkeConfigLifetime { get; private set; } = null!;

        /// <summary>
        /// The local_id of the ike config of the VPN connection.
        /// </summary>
        [Output("ikeConfigLocalId")]
        public Output<string> IkeConfigLocalId { get; private set; } = null!;

        /// <summary>
        /// The mode of the ike config of the VPN connection. Valid values are `main`, `aggressive`, and default value is `main`.
        /// </summary>
        [Output("ikeConfigMode")]
        public Output<string?> IkeConfigMode { get; private set; } = null!;

        /// <summary>
        /// The psk of the ike config of the VPN connection. The length does not exceed 100 characters, and only uppercase and lowercase letters, special symbols and numbers are allowed.
        /// </summary>
        [Output("ikeConfigPsk")]
        public Output<string> IkeConfigPsk { get; private set; } = null!;

        /// <summary>
        /// The remote id of the ike config of the VPN connection.
        /// </summary>
        [Output("ikeConfigRemoteId")]
        public Output<string> IkeConfigRemoteId { get; private set; } = null!;

        /// <summary>
        /// The version of the ike config of the VPN connection. The value can be `ikev1` or `ikev2`. The default value is `ikev1`.
        /// </summary>
        [Output("ikeConfigVersion")]
        public Output<string?> IkeConfigVersion { get; private set; } = null!;

        /// <summary>
        /// The ip address of transit router, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The auth alg of the ipsec config of the VPN connection. Valid value are `sha1`, `md5`, `sha256`, `sha384`, `sha512`, `sm3`. The default value is `sha1`.
        /// </summary>
        [Output("ipsecConfigAuthAlg")]
        public Output<string?> IpsecConfigAuthAlg { get; private set; } = null!;

        /// <summary>
        /// The dh group of the ipsec config of the VPN connection. Valid value are `group1`, `group2`, `group5`, `group14` and `disable`. The default value is `group2`.
        /// </summary>
        [Output("ipsecConfigDhGroup")]
        public Output<string?> IpsecConfigDhGroup { get; private set; } = null!;

        /// <summary>
        /// The enc alg of the ipsec config of the VPN connection. Valid value are `aes`, `aes192`, `aes256`, `des`, `3des`, `sm4`. The default value is `aes`.
        /// </summary>
        [Output("ipsecConfigEncAlg")]
        public Output<string?> IpsecConfigEncAlg { get; private set; } = null!;

        /// <summary>
        /// The ipsec config of the ike config of the VPN connection. Value: 900~86400.
        /// </summary>
        [Output("ipsecConfigLifetime")]
        public Output<int?> IpsecConfigLifetime { get; private set; } = null!;

        /// <summary>
        /// The local subnet of the VPN connection. Up to 5 network segments are supported.
        /// </summary>
        [Output("localSubnets")]
        public Output<ImmutableArray<string>> LocalSubnets { get; private set; } = null!;

        /// <summary>
        /// Whether to enable connection logging. After enabling Connection Day, you can view and download IPsec connection logs, and use the log information to troubleshoot IPsec connection problems yourself.
        /// </summary>
        [Output("logEnabled")]
        public Output<bool?> LogEnabled { get; private set; } = null!;

        /// <summary>
        /// The nat traversal of the VPN connection.
        /// </summary>
        [Output("natTraversal")]
        public Output<bool?> NatTraversal { get; private set; } = null!;

        /// <summary>
        /// Whether to initiate negotiation mode immediately.
        /// </summary>
        [Output("negotiateInstantly")]
        public Output<bool?> NegotiateInstantly { get; private set; } = null!;

        /// <summary>
        /// The overdue time of resource, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Output("overdueTime")]
        public Output<string> OverdueTime { get; private set; } = null!;

        /// <summary>
        /// The project name of the VPN connection.
        /// </summary>
        [Output("projectName")]
        public Output<string?> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The remote subnet of the VPN connection. Up to 5 network segments are supported.
        /// </summary>
        [Output("remoteSubnets")]
        public Output<ImmutableArray<string>> RemoteSubnets { get; private set; } = null!;

        /// <summary>
        /// The status of the VPN connection.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The id of transit router, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Output("transitRouterId")]
        public Output<string> TransitRouterId { get; private set; } = null!;

        /// <summary>
        /// The update time of VPN connection.
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;

        /// <summary>
        /// The ID of the VPN connection.
        /// </summary>
        [Output("vpnConnectionId")]
        public Output<string> VpnConnectionId { get; private set; } = null!;

        /// <summary>
        /// The name of the VPN connection.
        /// </summary>
        [Output("vpnConnectionName")]
        public Output<string> VpnConnectionName { get; private set; } = null!;

        /// <summary>
        /// The ID of the vpn gateway. If the `AttachType` is not passed or the passed value is `VpnGateway`, this parameter must be filled. If the value of `AttachType` is `TransitRouter`, this parameter does not need to be filled.
        /// </summary>
        [Output("vpnGatewayId")]
        public Output<string?> VpnGatewayId { get; private set; } = null!;

        /// <summary>
        /// The zone id of transit router, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Output("zoneId")]
        public Output<string> ZoneId { get; private set; } = null!;


        /// <summary>
        /// Create a Connection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Connection(string name, ConnectionArgs args, CustomResourceOptions? options = null)
            : base("volcengine:vpn/connection:Connection", name, args ?? new ConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Connection(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:vpn/connection:Connection", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Connection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Connection Get(string name, Input<string> id, ConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new Connection(name, id, state, options);
        }
    }

    public sealed class ConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The attach type of the VPN connection, the value can be `VpnGateway` or `TransitRouter`.
        /// </summary>
        [Input("attachType")]
        public Input<string>? AttachType { get; set; }

        /// <summary>
        /// The ID of the customer gateway.
        /// </summary>
        [Input("customerGatewayId", required: true)]
        public Input<string> CustomerGatewayId { get; set; } = null!;

        /// <summary>
        /// The description of the VPN connection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The dpd action of the VPN connection.
        /// </summary>
        [Input("dpdAction")]
        public Input<string>? DpdAction { get; set; }

        /// <summary>
        /// The auth alg of the ike config of the VPN connection. Valid value are `sha1`, `md5`, `sha256`, `sha384`, `sha512`, `sm3`. The default value is `sha1`.
        /// </summary>
        [Input("ikeConfigAuthAlg")]
        public Input<string>? IkeConfigAuthAlg { get; set; }

        /// <summary>
        /// The dk group of the ike config of the VPN connection. Valid value are `group1`, `group2`, `group5`, `group14`. The default value is `group2`.
        /// </summary>
        [Input("ikeConfigDhGroup")]
        public Input<string>? IkeConfigDhGroup { get; set; }

        /// <summary>
        /// The enc alg of the ike config of the VPN connection. Valid value are `aes`, `aes192`, `aes256`, `des`, `3des`, `sm4`. The default value is `aes`.
        /// </summary>
        [Input("ikeConfigEncAlg")]
        public Input<string>? IkeConfigEncAlg { get; set; }

        /// <summary>
        /// The lifetime of the ike config of the VPN connection. Value: 900~86400.
        /// </summary>
        [Input("ikeConfigLifetime")]
        public Input<int>? IkeConfigLifetime { get; set; }

        /// <summary>
        /// The local_id of the ike config of the VPN connection.
        /// </summary>
        [Input("ikeConfigLocalId")]
        public Input<string>? IkeConfigLocalId { get; set; }

        /// <summary>
        /// The mode of the ike config of the VPN connection. Valid values are `main`, `aggressive`, and default value is `main`.
        /// </summary>
        [Input("ikeConfigMode")]
        public Input<string>? IkeConfigMode { get; set; }

        /// <summary>
        /// The psk of the ike config of the VPN connection. The length does not exceed 100 characters, and only uppercase and lowercase letters, special symbols and numbers are allowed.
        /// </summary>
        [Input("ikeConfigPsk", required: true)]
        public Input<string> IkeConfigPsk { get; set; } = null!;

        /// <summary>
        /// The remote id of the ike config of the VPN connection.
        /// </summary>
        [Input("ikeConfigRemoteId")]
        public Input<string>? IkeConfigRemoteId { get; set; }

        /// <summary>
        /// The version of the ike config of the VPN connection. The value can be `ikev1` or `ikev2`. The default value is `ikev1`.
        /// </summary>
        [Input("ikeConfigVersion")]
        public Input<string>? IkeConfigVersion { get; set; }

        /// <summary>
        /// The auth alg of the ipsec config of the VPN connection. Valid value are `sha1`, `md5`, `sha256`, `sha384`, `sha512`, `sm3`. The default value is `sha1`.
        /// </summary>
        [Input("ipsecConfigAuthAlg")]
        public Input<string>? IpsecConfigAuthAlg { get; set; }

        /// <summary>
        /// The dh group of the ipsec config of the VPN connection. Valid value are `group1`, `group2`, `group5`, `group14` and `disable`. The default value is `group2`.
        /// </summary>
        [Input("ipsecConfigDhGroup")]
        public Input<string>? IpsecConfigDhGroup { get; set; }

        /// <summary>
        /// The enc alg of the ipsec config of the VPN connection. Valid value are `aes`, `aes192`, `aes256`, `des`, `3des`, `sm4`. The default value is `aes`.
        /// </summary>
        [Input("ipsecConfigEncAlg")]
        public Input<string>? IpsecConfigEncAlg { get; set; }

        /// <summary>
        /// The ipsec config of the ike config of the VPN connection. Value: 900~86400.
        /// </summary>
        [Input("ipsecConfigLifetime")]
        public Input<int>? IpsecConfigLifetime { get; set; }

        [Input("localSubnets", required: true)]
        private InputList<string>? _localSubnets;

        /// <summary>
        /// The local subnet of the VPN connection. Up to 5 network segments are supported.
        /// </summary>
        public InputList<string> LocalSubnets
        {
            get => _localSubnets ?? (_localSubnets = new InputList<string>());
            set => _localSubnets = value;
        }

        /// <summary>
        /// Whether to enable connection logging. After enabling Connection Day, you can view and download IPsec connection logs, and use the log information to troubleshoot IPsec connection problems yourself.
        /// </summary>
        [Input("logEnabled")]
        public Input<bool>? LogEnabled { get; set; }

        /// <summary>
        /// The nat traversal of the VPN connection.
        /// </summary>
        [Input("natTraversal")]
        public Input<bool>? NatTraversal { get; set; }

        /// <summary>
        /// Whether to initiate negotiation mode immediately.
        /// </summary>
        [Input("negotiateInstantly")]
        public Input<bool>? NegotiateInstantly { get; set; }

        /// <summary>
        /// The project name of the VPN connection.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("remoteSubnets", required: true)]
        private InputList<string>? _remoteSubnets;

        /// <summary>
        /// The remote subnet of the VPN connection. Up to 5 network segments are supported.
        /// </summary>
        public InputList<string> RemoteSubnets
        {
            get => _remoteSubnets ?? (_remoteSubnets = new InputList<string>());
            set => _remoteSubnets = value;
        }

        /// <summary>
        /// The name of the VPN connection.
        /// </summary>
        [Input("vpnConnectionName")]
        public Input<string>? VpnConnectionName { get; set; }

        /// <summary>
        /// The ID of the vpn gateway. If the `AttachType` is not passed or the passed value is `VpnGateway`, this parameter must be filled. If the value of `AttachType` is `TransitRouter`, this parameter does not need to be filled.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        public ConnectionArgs()
        {
        }
    }

    public sealed class ConnectionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The account ID of the VPN connection.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The IPsec attach status.
        /// </summary>
        [Input("attachStatus")]
        public Input<string>? AttachStatus { get; set; }

        /// <summary>
        /// The attach type of the VPN connection, the value can be `VpnGateway` or `TransitRouter`.
        /// </summary>
        [Input("attachType")]
        public Input<string>? AttachType { get; set; }

        /// <summary>
        /// The business status of IPsec connection, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Input("businessStatus")]
        public Input<string>? BusinessStatus { get; set; }

        /// <summary>
        /// The connect status of the VPN connection.
        /// </summary>
        [Input("connectStatus")]
        public Input<string>? ConnectStatus { get; set; }

        /// <summary>
        /// The create time of VPN connection.
        /// </summary>
        [Input("creationTime")]
        public Input<string>? CreationTime { get; set; }

        /// <summary>
        /// The ID of the customer gateway.
        /// </summary>
        [Input("customerGatewayId")]
        public Input<string>? CustomerGatewayId { get; set; }

        /// <summary>
        /// The delete time of resource, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Input("deletedTime")]
        public Input<string>? DeletedTime { get; set; }

        /// <summary>
        /// The description of the VPN connection.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The dpd action of the VPN connection.
        /// </summary>
        [Input("dpdAction")]
        public Input<string>? DpdAction { get; set; }

        /// <summary>
        /// The auth alg of the ike config of the VPN connection. Valid value are `sha1`, `md5`, `sha256`, `sha384`, `sha512`, `sm3`. The default value is `sha1`.
        /// </summary>
        [Input("ikeConfigAuthAlg")]
        public Input<string>? IkeConfigAuthAlg { get; set; }

        /// <summary>
        /// The dk group of the ike config of the VPN connection. Valid value are `group1`, `group2`, `group5`, `group14`. The default value is `group2`.
        /// </summary>
        [Input("ikeConfigDhGroup")]
        public Input<string>? IkeConfigDhGroup { get; set; }

        /// <summary>
        /// The enc alg of the ike config of the VPN connection. Valid value are `aes`, `aes192`, `aes256`, `des`, `3des`, `sm4`. The default value is `aes`.
        /// </summary>
        [Input("ikeConfigEncAlg")]
        public Input<string>? IkeConfigEncAlg { get; set; }

        /// <summary>
        /// The lifetime of the ike config of the VPN connection. Value: 900~86400.
        /// </summary>
        [Input("ikeConfigLifetime")]
        public Input<int>? IkeConfigLifetime { get; set; }

        /// <summary>
        /// The local_id of the ike config of the VPN connection.
        /// </summary>
        [Input("ikeConfigLocalId")]
        public Input<string>? IkeConfigLocalId { get; set; }

        /// <summary>
        /// The mode of the ike config of the VPN connection. Valid values are `main`, `aggressive`, and default value is `main`.
        /// </summary>
        [Input("ikeConfigMode")]
        public Input<string>? IkeConfigMode { get; set; }

        /// <summary>
        /// The psk of the ike config of the VPN connection. The length does not exceed 100 characters, and only uppercase and lowercase letters, special symbols and numbers are allowed.
        /// </summary>
        [Input("ikeConfigPsk")]
        public Input<string>? IkeConfigPsk { get; set; }

        /// <summary>
        /// The remote id of the ike config of the VPN connection.
        /// </summary>
        [Input("ikeConfigRemoteId")]
        public Input<string>? IkeConfigRemoteId { get; set; }

        /// <summary>
        /// The version of the ike config of the VPN connection. The value can be `ikev1` or `ikev2`. The default value is `ikev1`.
        /// </summary>
        [Input("ikeConfigVersion")]
        public Input<string>? IkeConfigVersion { get; set; }

        /// <summary>
        /// The ip address of transit router, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// The auth alg of the ipsec config of the VPN connection. Valid value are `sha1`, `md5`, `sha256`, `sha384`, `sha512`, `sm3`. The default value is `sha1`.
        /// </summary>
        [Input("ipsecConfigAuthAlg")]
        public Input<string>? IpsecConfigAuthAlg { get; set; }

        /// <summary>
        /// The dh group of the ipsec config of the VPN connection. Valid value are `group1`, `group2`, `group5`, `group14` and `disable`. The default value is `group2`.
        /// </summary>
        [Input("ipsecConfigDhGroup")]
        public Input<string>? IpsecConfigDhGroup { get; set; }

        /// <summary>
        /// The enc alg of the ipsec config of the VPN connection. Valid value are `aes`, `aes192`, `aes256`, `des`, `3des`, `sm4`. The default value is `aes`.
        /// </summary>
        [Input("ipsecConfigEncAlg")]
        public Input<string>? IpsecConfigEncAlg { get; set; }

        /// <summary>
        /// The ipsec config of the ike config of the VPN connection. Value: 900~86400.
        /// </summary>
        [Input("ipsecConfigLifetime")]
        public Input<int>? IpsecConfigLifetime { get; set; }

        [Input("localSubnets")]
        private InputList<string>? _localSubnets;

        /// <summary>
        /// The local subnet of the VPN connection. Up to 5 network segments are supported.
        /// </summary>
        public InputList<string> LocalSubnets
        {
            get => _localSubnets ?? (_localSubnets = new InputList<string>());
            set => _localSubnets = value;
        }

        /// <summary>
        /// Whether to enable connection logging. After enabling Connection Day, you can view and download IPsec connection logs, and use the log information to troubleshoot IPsec connection problems yourself.
        /// </summary>
        [Input("logEnabled")]
        public Input<bool>? LogEnabled { get; set; }

        /// <summary>
        /// The nat traversal of the VPN connection.
        /// </summary>
        [Input("natTraversal")]
        public Input<bool>? NatTraversal { get; set; }

        /// <summary>
        /// Whether to initiate negotiation mode immediately.
        /// </summary>
        [Input("negotiateInstantly")]
        public Input<bool>? NegotiateInstantly { get; set; }

        /// <summary>
        /// The overdue time of resource, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Input("overdueTime")]
        public Input<string>? OverdueTime { get; set; }

        /// <summary>
        /// The project name of the VPN connection.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        [Input("remoteSubnets")]
        private InputList<string>? _remoteSubnets;

        /// <summary>
        /// The remote subnet of the VPN connection. Up to 5 network segments are supported.
        /// </summary>
        public InputList<string> RemoteSubnets
        {
            get => _remoteSubnets ?? (_remoteSubnets = new InputList<string>());
            set => _remoteSubnets = value;
        }

        /// <summary>
        /// The status of the VPN connection.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The id of transit router, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Input("transitRouterId")]
        public Input<string>? TransitRouterId { get; set; }

        /// <summary>
        /// The update time of VPN connection.
        /// </summary>
        [Input("updateTime")]
        public Input<string>? UpdateTime { get; set; }

        /// <summary>
        /// The ID of the VPN connection.
        /// </summary>
        [Input("vpnConnectionId")]
        public Input<string>? VpnConnectionId { get; set; }

        /// <summary>
        /// The name of the VPN connection.
        /// </summary>
        [Input("vpnConnectionName")]
        public Input<string>? VpnConnectionName { get; set; }

        /// <summary>
        /// The ID of the vpn gateway. If the `AttachType` is not passed or the passed value is `VpnGateway`, this parameter must be filled. If the value of `AttachType` is `TransitRouter`, this parameter does not need to be filled.
        /// </summary>
        [Input("vpnGatewayId")]
        public Input<string>? VpnGatewayId { get; set; }

        /// <summary>
        /// The zone id of transit router, valid when the attach type is 'TransitRouter'.
        /// </summary>
        [Input("zoneId")]
        public Input<string>? ZoneId { get; set; }

        public ConnectionState()
        {
        }
    }
}
