// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Eip
{
    /// <summary>
    /// Provides a resource to manage eip associate
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Volcengine = Pulumi.Volcengine;
    /// using Volcengine = Volcengine.Pulumi.Volcengine;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var fooZones = Volcengine.Ecs.Zones.Invoke();
    /// 
    ///     var fooImages = Volcengine.Ecs.Images.Invoke(new()
    ///     {
    ///         OsType = "Linux",
    ///         Visibility = "public",
    ///         InstanceTypeId = "ecs.g1.large",
    ///     });
    /// 
    ///     var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new()
    ///     {
    ///         VpcName = "acc-test-vpc",
    ///         CidrBlock = "172.16.0.0/16",
    ///     });
    /// 
    ///     var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new()
    ///     {
    ///         SubnetName = "acc-test-subnet",
    ///         CidrBlock = "172.16.0.0/24",
    ///         ZoneId = fooZones.Apply(zonesResult =&gt; zonesResult.Zones[0]?.Id),
    ///         VpcId = fooVpc.Id,
    ///     });
    /// 
    ///     var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new()
    ///     {
    ///         VpcId = fooVpc.Id,
    ///         SecurityGroupName = "acc-test-security-group",
    ///     });
    /// 
    ///     var fooInstance = new Volcengine.Ecs.Instance("fooInstance", new()
    ///     {
    ///         ImageId = fooImages.Apply(imagesResult =&gt; imagesResult.Images[0]?.ImageId),
    ///         InstanceType = "ecs.g1.large",
    ///         InstanceName = "acc-test-ecs-name",
    ///         Password = "93f0cb0614Aab12",
    ///         InstanceChargeType = "PostPaid",
    ///         SystemVolumeType = "ESSD_PL0",
    ///         SystemVolumeSize = 40,
    ///         SubnetId = fooSubnet.Id,
    ///         SecurityGroupIds = new[]
    ///         {
    ///             fooSecurityGroup.Id,
    ///         },
    ///     });
    /// 
    ///     var fooAddress = new Volcengine.Eip.Address("fooAddress", new()
    ///     {
    ///         BillingType = "PostPaidByTraffic",
    ///     });
    /// 
    ///     var fooAssociate = new Volcengine.Eip.Associate("fooAssociate", new()
    ///     {
    ///         AllocationId = fooAddress.Id,
    ///         InstanceId = fooInstance.Id,
    ///         InstanceType = "EcsInstance",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Eip associate can be imported using the eip allocation_id:instance_id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:eip/associate:Associate default eip-274oj9a8rs9a87fap8sf9515b:i-cm9t9ug9lggu79yr5tcw
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:eip/associate:Associate")]
    public partial class Associate : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The allocation id of the EIP.
        /// </summary>
        [Output("allocationId")]
        public Output<string> AllocationId { get; private set; } = null!;

        /// <summary>
        /// The instance id which be associated to the EIP.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// The type of the associated instance,the value is `Nat` or `NetworkInterface` or `ClbInstance` or `EcsInstance` or `HaVip`.
        /// </summary>
        [Output("instanceType")]
        public Output<string> InstanceType { get; private set; } = null!;

        /// <summary>
        /// The private IP address of the instance will be associated to the EIP.
        /// </summary>
        [Output("privateIpAddress")]
        public Output<string> PrivateIpAddress { get; private set; } = null!;


        /// <summary>
        /// Create a Associate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Associate(string name, AssociateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:eip/associate:Associate", name, args ?? new AssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Associate(string name, Input<string> id, AssociateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:eip/associate:Associate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Associate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Associate Get(string name, Input<string> id, AssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new Associate(name, id, state, options);
        }
    }

    public sealed class AssociateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allocation id of the EIP.
        /// </summary>
        [Input("allocationId", required: true)]
        public Input<string> AllocationId { get; set; } = null!;

        /// <summary>
        /// The instance id which be associated to the EIP.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        /// <summary>
        /// The type of the associated instance,the value is `Nat` or `NetworkInterface` or `ClbInstance` or `EcsInstance` or `HaVip`.
        /// </summary>
        [Input("instanceType", required: true)]
        public Input<string> InstanceType { get; set; } = null!;

        /// <summary>
        /// The private IP address of the instance will be associated to the EIP.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        public AssociateArgs()
        {
        }
        public static new AssociateArgs Empty => new AssociateArgs();
    }

    public sealed class AssociateState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The allocation id of the EIP.
        /// </summary>
        [Input("allocationId")]
        public Input<string>? AllocationId { get; set; }

        /// <summary>
        /// The instance id which be associated to the EIP.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// The type of the associated instance,the value is `Nat` or `NetworkInterface` or `ClbInstance` or `EcsInstance` or `HaVip`.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The private IP address of the instance will be associated to the EIP.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<string>? PrivateIpAddress { get; set; }

        public AssociateState()
        {
        }
        public static new AssociateState Empty => new AssociateState();
    }
}
