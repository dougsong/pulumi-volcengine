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
    /// Provides a resource to manage redis backup
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
    ///     var @default = new Volcengine.Redis.Backup("default", new()
    ///     {
    ///         InstanceId = "redis-cnlfvrv4qye6u4lpa",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Redis Backup can be imported using the instanceId:backupId, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:redis/backup:Backup default redis-cn02aqusft7ws****:b-cn02xmmrp751i9cdzcphjmk4****
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:redis/backup:Backup")]
    public partial class Backup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The id of backup point.
        /// </summary>
        [Output("backupPointId")]
        public Output<string> BackupPointId { get; private set; } = null!;

        /// <summary>
        /// Backup strategy.
        /// </summary>
        [Output("backupStrategy")]
        public Output<string> BackupStrategy { get; private set; } = null!;

        /// <summary>
        /// Backup type.
        /// </summary>
        [Output("backupType")]
        public Output<string> BackupType { get; private set; } = null!;

        /// <summary>
        /// End time of backup.
        /// </summary>
        [Output("endTime")]
        public Output<string> EndTime { get; private set; } = null!;

        /// <summary>
        /// Information of instance.
        /// </summary>
        [Output("instanceDetails")]
        public Output<ImmutableArray<Outputs.BackupInstanceDetail>> InstanceDetails { get; private set; } = null!;

        /// <summary>
        /// Id of instance to create backup.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;

        /// <summary>
        /// Size in MiB.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// Start time of backup.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// Status of backup (Creating/Available/Unavailable/Deleting).
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Backup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Backup(string name, BackupArgs args, CustomResourceOptions? options = null)
            : base("volcengine:redis/backup:Backup", name, args ?? new BackupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Backup(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:redis/backup:Backup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Backup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Backup Get(string name, Input<string> id, BackupState? state = null, CustomResourceOptions? options = null)
        {
            return new Backup(name, id, state, options);
        }
    }

    public sealed class BackupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of instance to create backup.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public BackupArgs()
        {
        }
        public static new BackupArgs Empty => new BackupArgs();
    }

    public sealed class BackupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of backup point.
        /// </summary>
        [Input("backupPointId")]
        public Input<string>? BackupPointId { get; set; }

        /// <summary>
        /// Backup strategy.
        /// </summary>
        [Input("backupStrategy")]
        public Input<string>? BackupStrategy { get; set; }

        /// <summary>
        /// Backup type.
        /// </summary>
        [Input("backupType")]
        public Input<string>? BackupType { get; set; }

        /// <summary>
        /// End time of backup.
        /// </summary>
        [Input("endTime")]
        public Input<string>? EndTime { get; set; }

        [Input("instanceDetails")]
        private InputList<Inputs.BackupInstanceDetailGetArgs>? _instanceDetails;

        /// <summary>
        /// Information of instance.
        /// </summary>
        public InputList<Inputs.BackupInstanceDetailGetArgs> InstanceDetails
        {
            get => _instanceDetails ?? (_instanceDetails = new InputList<Inputs.BackupInstanceDetailGetArgs>());
            set => _instanceDetails = value;
        }

        /// <summary>
        /// Id of instance to create backup.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        /// <summary>
        /// Size in MiB.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// Start time of backup.
        /// </summary>
        [Input("startTime")]
        public Input<string>? StartTime { get; set; }

        /// <summary>
        /// Status of backup (Creating/Available/Unavailable/Deleting).
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public BackupState()
        {
        }
        public static new BackupState Empty => new BackupState();
    }
}
