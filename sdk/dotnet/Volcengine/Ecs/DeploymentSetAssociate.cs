// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.PulumiPackage.Volcengine.Ecs
{
    /// <summary>
    /// Provides a resource to manage ecs deployment set associate
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Volcengine = Pulumi.Volcengine;
    /// using Volcengine = Volcengine.PulumiPackage.Volcengine;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var fooZones = Output.Create(Volcengine.Ecs.Zones.InvokeAsync());
    ///         var fooVpc = new Volcengine.Vpc.Vpc("fooVpc", new Volcengine.Vpc.VpcArgs
    ///         {
    ///             VpcName = "acc-test-vpc",
    ///             CidrBlock = "172.16.0.0/16",
    ///         });
    ///         var fooSubnet = new Volcengine.Vpc.Subnet("fooSubnet", new Volcengine.Vpc.SubnetArgs
    ///         {
    ///             SubnetName = "acc-test-subnet",
    ///             CidrBlock = "172.16.0.0/24",
    ///             ZoneId = fooZones.Apply(fooZones =&gt; fooZones.Zones?[0]?.Id),
    ///             VpcId = fooVpc.Id,
    ///         });
    ///         var fooSecurityGroup = new Volcengine.Vpc.SecurityGroup("fooSecurityGroup", new Volcengine.Vpc.SecurityGroupArgs
    ///         {
    ///             SecurityGroupName = "acc-test-security-group",
    ///             VpcId = fooVpc.Id,
    ///         });
    ///         var fooImages = Output.Create(Volcengine.Ecs.Images.InvokeAsync(new Volcengine.Ecs.ImagesArgs
    ///         {
    ///             OsType = "Linux",
    ///             Visibility = "public",
    ///             InstanceTypeId = "ecs.g1.large",
    ///         }));
    ///         var fooInstance = new Volcengine.Ecs.Instance("fooInstance", new Volcengine.Ecs.InstanceArgs
    ///         {
    ///             InstanceName = "acc-test-ecs",
    ///             ImageId = fooImages.Apply(fooImages =&gt; fooImages.Images?[0]?.ImageId),
    ///             InstanceType = "ecs.g1.large",
    ///             Password = "93f0cb0614Aab12",
    ///             InstanceChargeType = "PostPaid",
    ///             SystemVolumeType = "ESSD_PL0",
    ///             SystemVolumeSize = 40,
    ///             SubnetId = fooSubnet.Id,
    ///             SecurityGroupIds = 
    ///             {
    ///                 fooSecurityGroup.Id,
    ///             },
    ///         });
    ///         var fooState = new Volcengine.Ecs.State("fooState", new Volcengine.Ecs.StateArgs
    ///         {
    ///             InstanceId = fooInstance.Id,
    ///             Action = "Stop",
    ///             StoppedMode = "KeepCharging",
    ///         });
    ///         var fooDeploymentSet = new Volcengine.Ecs.DeploymentSet("fooDeploymentSet", new Volcengine.Ecs.DeploymentSetArgs
    ///         {
    ///             DeploymentSetName = "acc-test-ecs-ds",
    ///             Description = "acc-test",
    ///             Granularity = "switch",
    ///             Strategy = "Availability",
    ///         });
    ///         var fooDeploymentSetAssociate = new Volcengine.Ecs.DeploymentSetAssociate("fooDeploymentSetAssociate", new Volcengine.Ecs.DeploymentSetAssociateArgs
    ///         {
    ///             DeploymentSetId = fooDeploymentSet.Id,
    ///             InstanceId = fooInstance.Id,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 fooState,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// ECS deployment set associate can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:ecs/deploymentSetAssociate:DeploymentSetAssociate default dps-ybti5tkpkv2udbfolrft:i-mizl7m1kqccg5smt1bdpijuj
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:ecs/deploymentSetAssociate:DeploymentSetAssociate")]
    public partial class DeploymentSetAssociate : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of ECS DeploymentSet Associate.
        /// </summary>
        [Output("deploymentSetId")]
        public Output<string> DeploymentSetId { get; private set; } = null!;

        /// <summary>
        /// The ID of ECS Instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a DeploymentSetAssociate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeploymentSetAssociate(string name, DeploymentSetAssociateArgs args, CustomResourceOptions? options = null)
            : base("volcengine:ecs/deploymentSetAssociate:DeploymentSetAssociate", name, args ?? new DeploymentSetAssociateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeploymentSetAssociate(string name, Input<string> id, DeploymentSetAssociateState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:ecs/deploymentSetAssociate:DeploymentSetAssociate", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeploymentSetAssociate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeploymentSetAssociate Get(string name, Input<string> id, DeploymentSetAssociateState? state = null, CustomResourceOptions? options = null)
        {
            return new DeploymentSetAssociate(name, id, state, options);
        }
    }

    public sealed class DeploymentSetAssociateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of ECS DeploymentSet Associate.
        /// </summary>
        [Input("deploymentSetId", required: true)]
        public Input<string> DeploymentSetId { get; set; } = null!;

        /// <summary>
        /// The ID of ECS Instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public DeploymentSetAssociateArgs()
        {
        }
    }

    public sealed class DeploymentSetAssociateState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of ECS DeploymentSet Associate.
        /// </summary>
        [Input("deploymentSetId")]
        public Input<string>? DeploymentSetId { get; set; }

        /// <summary>
        /// The ID of ECS Instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public DeploymentSetAssociateState()
        {
        }
    }
}
