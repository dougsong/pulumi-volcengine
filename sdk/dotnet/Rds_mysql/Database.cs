// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Rds_mysql
{
    /// <summary>
    /// Provides a resource to manage rds mysql database
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
    ///         var @default = new Volcengine.Rds_mysql.Database("default", new Volcengine.Rds_mysql.DatabaseArgs
    ///         {
    ///             CharacterSetName = "utf8",
    ///             DbName = "xxx",
    ///             InstanceId = "mysql-xxx",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Database can be imported using the instanceId:dbName, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import volcengine:rds_mysql/database:Database default mysql-42b38c769c4b:dbname
    /// ```
    /// </summary>
    [VolcengineResourceType("volcengine:rds_mysql/database:Database")]
    public partial class Database : Pulumi.CustomResource
    {
        /// <summary>
        /// Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
        /// </summary>
        [Output("characterSetName")]
        public Output<string?> CharacterSetName { get; private set; } = null!;

        /// <summary>
        /// Name database.
        /// illustrate:
        /// Unique name.
        /// The length is 2~64 characters.
        /// Start with a letter and end with a letter or number.
        /// Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
        /// Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
        /// </summary>
        [Output("dbName")]
        public Output<string> DbName { get; private set; } = null!;

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Output("instanceId")]
        public Output<string> InstanceId { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("volcengine:rds_mysql/database:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
            : base("volcengine:rds_mysql/database:Database", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, DatabaseState? state = null, CustomResourceOptions? options = null)
        {
            return new Database(name, id, state, options);
        }
    }

    public sealed class DatabaseArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
        /// </summary>
        [Input("characterSetName")]
        public Input<string>? CharacterSetName { get; set; }

        /// <summary>
        /// Name database.
        /// illustrate:
        /// Unique name.
        /// The length is 2~64 characters.
        /// Start with a letter and end with a letter or number.
        /// Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
        /// Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
        /// </summary>
        [Input("dbName", required: true)]
        public Input<string> DbName { get; set; } = null!;

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Input("instanceId", required: true)]
        public Input<string> InstanceId { get; set; } = null!;

        public DatabaseArgs()
        {
        }
    }

    public sealed class DatabaseState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Database character set. Currently supported character sets include: utf8, utf8mb4, latin1, ascii.
        /// </summary>
        [Input("characterSetName")]
        public Input<string>? CharacterSetName { get; set; }

        /// <summary>
        /// Name database.
        /// illustrate:
        /// Unique name.
        /// The length is 2~64 characters.
        /// Start with a letter and end with a letter or number.
        /// Consists of lowercase letters, numbers, and underscores (_) or dashes (-).
        /// Database names are disabled [keywords](https://www.volcengine.com/docs/6313/66162).
        /// </summary>
        [Input("dbName")]
        public Input<string>? DbName { get; set; }

        /// <summary>
        /// The ID of the RDS instance.
        /// </summary>
        [Input("instanceId")]
        public Input<string>? InstanceId { get; set; }

        public DatabaseState()
        {
        }
    }
}
