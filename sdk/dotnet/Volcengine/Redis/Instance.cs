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
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Volcengine = Volcengine.Pulumi.Volcengine;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var foo = new Volcengine.Redis.Instance("foo", new()
    ///     {
    ///         ApplyImmediately = true,
    ///         BackupActive = true,
    ///         BackupHour = 4,
    ///         BackupPeriods = new[]
    ///         {
    ///             1,
    ///             2,
    ///             3,
    ///         },
    ///         ChargeType = "PostPaid",
    ///         CreateBackup = false,
    ///         DeletionProtection = "disabled",
    ///         EngineVersion = "5.0",
    ///         InstanceName = "tf-test",
    ///         NodeNumber = 2,
    ///         ParamValues = new[]
    ///         {
    ///             new Volcengine.Redis.Inputs.InstanceParamValueArgs
    ///             {
    ///                 Name = "active-defrag-cycle-min",
    ///                 Value = "5",
    ///             },
    ///             new Volcengine.Redis.Inputs.InstanceParamValueArgs
    ///             {
    ///                 Name = "active-defrag-cycle-max",
    ///                 Value = "28",
    ///             },
    ///         },
    ///         Password = "1qaz!QAZ12",
    ///         Port = 6381,
    ///         ProjectName = "default",
    ///         ShardCapacity = 1024,
    ///         ShardNumber = 2,
    ///         ShardedCluster = 1,
    ///         SubnetId = "subnet-13g7c3lot0lc03n6nu4wj****",
    ///         Tags = new[]
    ///         {
    ///             new Volcengine.Redis.Inputs.InstanceTagArgs
    ///             {
    ///                 Key = "k1",
    ///                 Value = "v1",
    ///             },
    ///             new Volcengine.Redis.Inputs.InstanceTagArgs
    ///             {
    ///                 Key = "k3",
    ///                 Value = "v3",
    ///             },
    ///         },
    ///         VpcAuthMode = "close",
    ///         ZoneIds = new[]
    ///         {
    ///             "cn-beijing-a",
    ///             "cn-beijing-b",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// redis instance can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:redis/instance:Instance default redis-n769ewmjjqyqh5dv
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:redis/instance:Instance")]
    public partial class Instance : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether to apply the instance configuration change operation immediately. The value of this field is false, means that the change operation will be applied within maintenance time.
        /// </summary>
        [Output("applyImmediately")]
        public Output<bool?> ApplyImmediately { get; private set; } = null!;

        /// <summary>
        /// Whether to enable automatic renewal. This field is valid only when `ChargeType` is `PrePaid`, the default value is false. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("autoRenew")]
        public Output<bool?> AutoRenew { get; private set; } = null!;

        /// <summary>
        /// Whether enable auto backup for redis instance. This field is valid and required when updating the backup plan of primary and secondary instance.
        /// </summary>
        [Output("backupActive")]
        public Output<bool> BackupActive { get; private set; } = null!;

        /// <summary>
        /// The time period to start performing the backup. The valid value range is any integer between 0 and 23, where 0 means that the system will perform the backup in the period of 00:00~01:00, 1 means that the backup will be performed in the period of 01:00~02:00, and so on. 
        /// This field is valid and required when updating the backup plan of primary and secondary instance.
        /// </summary>
        [Output("backupHour")]
        public Output<int> BackupHour { get; private set; } = null!;

        /// <summary>
        /// The backup period. The valid value can be any integer between 1 and 7. Among them, 1 means backup every Monday, 2 means backup every Tuesday, and so on. 
        /// This field is valid and required when updating the backup plan of primary and secondary instance.
        /// </summary>
        [Output("backupPeriods")]
        public Output<ImmutableArray<int>> BackupPeriods { get; private set; } = null!;

        /// <summary>
        /// The charge type of redis instance. Valid value: `PostPaid`, `PrePaid`.
        /// </summary>
        [Output("chargeType")]
        public Output<string?> ChargeType { get; private set; } = null!;

        /// <summary>
        /// Whether to create a final backup when modify the instance configuration or destroy the redis instance.
        /// </summary>
        [Output("createBackup")]
        public Output<bool?> CreateBackup { get; private set; } = null!;

        /// <summary>
        /// Whether enable deletion protection for redis instance. Valid values: `enabled`, `disabled`(default).
        /// </summary>
        [Output("deletionProtection")]
        public Output<string?> DeletionProtection { get; private set; } = null!;

        /// <summary>
        /// The engine version of redis instance. Valid value: `4.0`, `5.0`, `6.0`.
        /// </summary>
        [Output("engineVersion")]
        public Output<string> EngineVersion { get; private set; } = null!;

        /// <summary>
        /// The name of the redis instance.
        /// </summary>
        [Output("instanceName")]
        public Output<string?> InstanceName { get; private set; } = null!;

        /// <summary>
        /// The number of nodes in each shard, the valid value range is `1-6`. When the value is 1, it means creating a single node instance, and this field can not be modified. When the value is greater than 1, it means creating a primary and secondary instance, and this field can be modified.
        /// </summary>
        [Output("nodeNumber")]
        public Output<int> NodeNumber { get; private set; } = null!;

        /// <summary>
        /// The configuration item information to be modified. This field can only be added or modified. Deleting this field is invalid.
        /// </summary>
        [Output("paramValues")]
        public Output<ImmutableArray<Outputs.InstanceParamValue>> ParamValues { get; private set; } = null!;

        /// <summary>
        /// The account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The port of custom define private network address. The valid value range is `1024-65535`. The default value is `6379`.
        /// </summary>
        [Output("port")]
        public Output<int?> Port { get; private set; } = null!;

        /// <summary>
        /// The project name to which the redis instance belongs, if this parameter is empty, the new redis instance will not be added to any project.
        /// </summary>
        [Output("projectName")]
        public Output<string> ProjectName { get; private set; } = null!;

        /// <summary>
        /// The purchase months of redis instance, the unit is month. the valid value range is as fallows: `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36`. This field is valid and required when `ChargeType` is `Prepaid`. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Output("purchaseMonths")]
        public Output<int?> PurchaseMonths { get; private set; } = null!;

        /// <summary>
        /// The memory capacity of each shard, unit is MiB. The valid value range is as fallows: When the value of `ShardedCluster` is 0: 256, 1024, 2048, 4096, 8192, 16384, 32768, 65536. When the value of `ShardedCluster` is 1: 1024, 2048, 4096, 8192, 16384. When the value of `node_number` is 1, the value of this field can not be 256.
        /// </summary>
        [Output("shardCapacity")]
        public Output<int> ShardCapacity { get; private set; } = null!;

        /// <summary>
        /// The number of shards in redis instance, the valid value range is `2-256`. This field is valid and required when the value of `ShardedCluster` is 1.
        /// </summary>
        [Output("shardNumber")]
        public Output<int> ShardNumber { get; private set; } = null!;

        /// <summary>
        /// Whether enable sharded cluster for the current redis instance. Valid values: 0, 1. 0 means disable, 1 means enable.
        /// </summary>
        [Output("shardedCluster")]
        public Output<int> ShardedCluster { get; private set; } = null!;

        /// <summary>
        /// The subnet id of the redis instance. The specified subnet id must belong to the zone ids.
        /// </summary>
        [Output("subnetId")]
        public Output<string> SubnetId { get; private set; } = null!;

        /// <summary>
        /// Tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.InstanceTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// Whether to enable password-free access when connecting to an instance through a private network. Valid values: `open`, `close`. Works only on modified scenes.
        /// </summary>
        [Output("vpcAuthMode")]
        public Output<string> VpcAuthMode { get; private set; } = null!;

        /// <summary>
        /// The list of zone IDs of instance. When creating a single node instance, only one zone id can be specified.
        /// </summary>
        [Output("zoneIds")]
        public Output<ImmutableArray<string>> ZoneIds { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs args, CustomResourceOptions? options = null)
            : base("volcengine:redis/instance:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:redis/instance:Instance", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/volcengine",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, InstanceState? state = null, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, state, options);
        }
    }

    public sealed class InstanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to apply the instance configuration change operation immediately. The value of this field is false, means that the change operation will be applied within maintenance time.
        /// </summary>
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        /// <summary>
        /// Whether to enable automatic renewal. This field is valid only when `ChargeType` is `PrePaid`, the default value is false. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Whether enable auto backup for redis instance. This field is valid and required when updating the backup plan of primary and secondary instance.
        /// </summary>
        [Input("backupActive")]
        public Input<bool>? BackupActive { get; set; }

        /// <summary>
        /// The time period to start performing the backup. The valid value range is any integer between 0 and 23, where 0 means that the system will perform the backup in the period of 00:00~01:00, 1 means that the backup will be performed in the period of 01:00~02:00, and so on. 
        /// This field is valid and required when updating the backup plan of primary and secondary instance.
        /// </summary>
        [Input("backupHour")]
        public Input<int>? BackupHour { get; set; }

        [Input("backupPeriods")]
        private InputList<int>? _backupPeriods;

        /// <summary>
        /// The backup period. The valid value can be any integer between 1 and 7. Among them, 1 means backup every Monday, 2 means backup every Tuesday, and so on. 
        /// This field is valid and required when updating the backup plan of primary and secondary instance.
        /// </summary>
        public InputList<int> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<int>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// The charge type of redis instance. Valid value: `PostPaid`, `PrePaid`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Whether to create a final backup when modify the instance configuration or destroy the redis instance.
        /// </summary>
        [Input("createBackup")]
        public Input<bool>? CreateBackup { get; set; }

        /// <summary>
        /// Whether enable deletion protection for redis instance. Valid values: `enabled`, `disabled`(default).
        /// </summary>
        [Input("deletionProtection")]
        public Input<string>? DeletionProtection { get; set; }

        /// <summary>
        /// The engine version of redis instance. Valid value: `4.0`, `5.0`, `6.0`.
        /// </summary>
        [Input("engineVersion", required: true)]
        public Input<string> EngineVersion { get; set; } = null!;

        /// <summary>
        /// The name of the redis instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The number of nodes in each shard, the valid value range is `1-6`. When the value is 1, it means creating a single node instance, and this field can not be modified. When the value is greater than 1, it means creating a primary and secondary instance, and this field can be modified.
        /// </summary>
        [Input("nodeNumber", required: true)]
        public Input<int> NodeNumber { get; set; } = null!;

        [Input("paramValues")]
        private InputList<Inputs.InstanceParamValueArgs>? _paramValues;

        /// <summary>
        /// The configuration item information to be modified. This field can only be added or modified. Deleting this field is invalid.
        /// </summary>
        public InputList<Inputs.InstanceParamValueArgs> ParamValues
        {
            get => _paramValues ?? (_paramValues = new InputList<Inputs.InstanceParamValueArgs>());
            set => _paramValues = value;
        }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The port of custom define private network address. The valid value range is `1024-65535`. The default value is `6379`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The project name to which the redis instance belongs, if this parameter is empty, the new redis instance will not be added to any project.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The purchase months of redis instance, the unit is month. the valid value range is as fallows: `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36`. This field is valid and required when `ChargeType` is `Prepaid`. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("purchaseMonths")]
        public Input<int>? PurchaseMonths { get; set; }

        /// <summary>
        /// The memory capacity of each shard, unit is MiB. The valid value range is as fallows: When the value of `ShardedCluster` is 0: 256, 1024, 2048, 4096, 8192, 16384, 32768, 65536. When the value of `ShardedCluster` is 1: 1024, 2048, 4096, 8192, 16384. When the value of `node_number` is 1, the value of this field can not be 256.
        /// </summary>
        [Input("shardCapacity", required: true)]
        public Input<int> ShardCapacity { get; set; } = null!;

        /// <summary>
        /// The number of shards in redis instance, the valid value range is `2-256`. This field is valid and required when the value of `ShardedCluster` is 1.
        /// </summary>
        [Input("shardNumber")]
        public Input<int>? ShardNumber { get; set; }

        /// <summary>
        /// Whether enable sharded cluster for the current redis instance. Valid values: 0, 1. 0 means disable, 1 means enable.
        /// </summary>
        [Input("shardedCluster", required: true)]
        public Input<int> ShardedCluster { get; set; } = null!;

        /// <summary>
        /// The subnet id of the redis instance. The specified subnet id must belong to the zone ids.
        /// </summary>
        [Input("subnetId", required: true)]
        public Input<string> SubnetId { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.InstanceTagArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.InstanceTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.InstanceTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether to enable password-free access when connecting to an instance through a private network. Valid values: `open`, `close`. Works only on modified scenes.
        /// </summary>
        [Input("vpcAuthMode")]
        public Input<string>? VpcAuthMode { get; set; }

        [Input("zoneIds", required: true)]
        private InputList<string>? _zoneIds;

        /// <summary>
        /// The list of zone IDs of instance. When creating a single node instance, only one zone id can be specified.
        /// </summary>
        public InputList<string> ZoneIds
        {
            get => _zoneIds ?? (_zoneIds = new InputList<string>());
            set => _zoneIds = value;
        }

        public InstanceArgs()
        {
        }
        public static new InstanceArgs Empty => new InstanceArgs();
    }

    public sealed class InstanceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether to apply the instance configuration change operation immediately. The value of this field is false, means that the change operation will be applied within maintenance time.
        /// </summary>
        [Input("applyImmediately")]
        public Input<bool>? ApplyImmediately { get; set; }

        /// <summary>
        /// Whether to enable automatic renewal. This field is valid only when `ChargeType` is `PrePaid`, the default value is false. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Whether enable auto backup for redis instance. This field is valid and required when updating the backup plan of primary and secondary instance.
        /// </summary>
        [Input("backupActive")]
        public Input<bool>? BackupActive { get; set; }

        /// <summary>
        /// The time period to start performing the backup. The valid value range is any integer between 0 and 23, where 0 means that the system will perform the backup in the period of 00:00~01:00, 1 means that the backup will be performed in the period of 01:00~02:00, and so on. 
        /// This field is valid and required when updating the backup plan of primary and secondary instance.
        /// </summary>
        [Input("backupHour")]
        public Input<int>? BackupHour { get; set; }

        [Input("backupPeriods")]
        private InputList<int>? _backupPeriods;

        /// <summary>
        /// The backup period. The valid value can be any integer between 1 and 7. Among them, 1 means backup every Monday, 2 means backup every Tuesday, and so on. 
        /// This field is valid and required when updating the backup plan of primary and secondary instance.
        /// </summary>
        public InputList<int> BackupPeriods
        {
            get => _backupPeriods ?? (_backupPeriods = new InputList<int>());
            set => _backupPeriods = value;
        }

        /// <summary>
        /// The charge type of redis instance. Valid value: `PostPaid`, `PrePaid`.
        /// </summary>
        [Input("chargeType")]
        public Input<string>? ChargeType { get; set; }

        /// <summary>
        /// Whether to create a final backup when modify the instance configuration or destroy the redis instance.
        /// </summary>
        [Input("createBackup")]
        public Input<bool>? CreateBackup { get; set; }

        /// <summary>
        /// Whether enable deletion protection for redis instance. Valid values: `enabled`, `disabled`(default).
        /// </summary>
        [Input("deletionProtection")]
        public Input<string>? DeletionProtection { get; set; }

        /// <summary>
        /// The engine version of redis instance. Valid value: `4.0`, `5.0`, `6.0`.
        /// </summary>
        [Input("engineVersion")]
        public Input<string>? EngineVersion { get; set; }

        /// <summary>
        /// The name of the redis instance.
        /// </summary>
        [Input("instanceName")]
        public Input<string>? InstanceName { get; set; }

        /// <summary>
        /// The number of nodes in each shard, the valid value range is `1-6`. When the value is 1, it means creating a single node instance, and this field can not be modified. When the value is greater than 1, it means creating a primary and secondary instance, and this field can be modified.
        /// </summary>
        [Input("nodeNumber")]
        public Input<int>? NodeNumber { get; set; }

        [Input("paramValues")]
        private InputList<Inputs.InstanceParamValueGetArgs>? _paramValues;

        /// <summary>
        /// The configuration item information to be modified. This field can only be added or modified. Deleting this field is invalid.
        /// </summary>
        public InputList<Inputs.InstanceParamValueGetArgs> ParamValues
        {
            get => _paramValues ?? (_paramValues = new InputList<Inputs.InstanceParamValueGetArgs>());
            set => _paramValues = value;
        }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The account password. When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The port of custom define private network address. The valid value range is `1024-65535`. The default value is `6379`.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// The project name to which the redis instance belongs, if this parameter is empty, the new redis instance will not be added to any project.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        /// <summary>
        /// The purchase months of redis instance, the unit is month. the valid value range is as fallows: `1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36`. This field is valid and required when `ChargeType` is `Prepaid`. 
        /// When importing resources, this attribute will not be imported. If this attribute is set, please use lifecycle and ignore_changes ignore changes in fields.
        /// </summary>
        [Input("purchaseMonths")]
        public Input<int>? PurchaseMonths { get; set; }

        /// <summary>
        /// The memory capacity of each shard, unit is MiB. The valid value range is as fallows: When the value of `ShardedCluster` is 0: 256, 1024, 2048, 4096, 8192, 16384, 32768, 65536. When the value of `ShardedCluster` is 1: 1024, 2048, 4096, 8192, 16384. When the value of `node_number` is 1, the value of this field can not be 256.
        /// </summary>
        [Input("shardCapacity")]
        public Input<int>? ShardCapacity { get; set; }

        /// <summary>
        /// The number of shards in redis instance, the valid value range is `2-256`. This field is valid and required when the value of `ShardedCluster` is 1.
        /// </summary>
        [Input("shardNumber")]
        public Input<int>? ShardNumber { get; set; }

        /// <summary>
        /// Whether enable sharded cluster for the current redis instance. Valid values: 0, 1. 0 means disable, 1 means enable.
        /// </summary>
        [Input("shardedCluster")]
        public Input<int>? ShardedCluster { get; set; }

        /// <summary>
        /// The subnet id of the redis instance. The specified subnet id must belong to the zone ids.
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        [Input("tags")]
        private InputList<Inputs.InstanceTagGetArgs>? _tags;

        /// <summary>
        /// Tags.
        /// </summary>
        public InputList<Inputs.InstanceTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.InstanceTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// Whether to enable password-free access when connecting to an instance through a private network. Valid values: `open`, `close`. Works only on modified scenes.
        /// </summary>
        [Input("vpcAuthMode")]
        public Input<string>? VpcAuthMode { get; set; }

        [Input("zoneIds")]
        private InputList<string>? _zoneIds;

        /// <summary>
        /// The list of zone IDs of instance. When creating a single node instance, only one zone id can be specified.
        /// </summary>
        public InputList<string> ZoneIds
        {
            get => _zoneIds ?? (_zoneIds = new InputList<string>());
            set => _zoneIds = value;
        }

        public InstanceState()
        {
        }
        public static new InstanceState Empty => new InstanceState();
    }
}
