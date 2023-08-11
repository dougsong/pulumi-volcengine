// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Clb
{
    public static class Acls
    {
        /// <summary>
        /// Use this data source to query detailed information of acls
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
        ///     var @default = Volcengine.Clb.Acls.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "acl-3ti8n0rurx4bwbh9jzdy",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<AclsResult> InvokeAsync(AclsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<AclsResult>("volcengine:clb/acls:Acls", args ?? new AclsArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of acls
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
        ///     var @default = Volcengine.Clb.Acls.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             "acl-3ti8n0rurx4bwbh9jzdy",
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<AclsResult> Invoke(AclsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<AclsResult>("volcengine:clb/acls:Acls", args ?? new AclsInvokeArgs(), options.WithDefaults());
    }


    public sealed class AclsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of acl.
        /// </summary>
        [Input("aclName")]
        public string? AclName { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Acl IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of Acl.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The ProjectName of Acl.
        /// </summary>
        [Input("projectName")]
        public string? ProjectName { get; set; }

        public AclsArgs()
        {
        }
        public static new AclsArgs Empty => new AclsArgs();
    }

    public sealed class AclsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of acl.
        /// </summary>
        [Input("aclName")]
        public Input<string>? AclName { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Acl IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A Name Regex of Acl.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The ProjectName of Acl.
        /// </summary>
        [Input("projectName")]
        public Input<string>? ProjectName { get; set; }

        public AclsInvokeArgs()
        {
        }
        public static new AclsInvokeArgs Empty => new AclsInvokeArgs();
    }


    [OutputType]
    public sealed class AclsResult
    {
        /// <summary>
        /// The Name of Acl.
        /// </summary>
        public readonly string? AclName;
        /// <summary>
        /// The collection of Acl query.
        /// </summary>
        public readonly ImmutableArray<Outputs.AclsAclResult> Acls;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly string? OutputFile;
        /// <summary>
        /// The ProjectName of Acl.
        /// </summary>
        public readonly string? ProjectName;
        /// <summary>
        /// The total count of Acl query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private AclsResult(
            string? aclName,

            ImmutableArray<Outputs.AclsAclResult> acls,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            string? outputFile,

            string? projectName,

            int totalCount)
        {
            AclName = aclName;
            Acls = acls;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            OutputFile = outputFile;
            ProjectName = projectName;
            TotalCount = totalCount;
        }
    }
}
