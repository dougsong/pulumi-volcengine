// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cr
{
    public static class Repositories
    {
        /// <summary>
        /// Use this data source to query detailed information of cr repositories
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
        ///         var foo = Output.Create(Volcengine.Cr.Repositories.InvokeAsync(new Volcengine.Cr.RepositoriesArgs
        ///         {
        ///             Names = 
        ///             {
        ///                 "repo*",
        ///             },
        ///             Registry = "tf-1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<RepositoriesResult> InvokeAsync(RepositoriesArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<RepositoriesResult>("volcengine:cr/repositories:Repositories", args ?? new RepositoriesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cr repositories
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
        ///         var foo = Output.Create(Volcengine.Cr.Repositories.InvokeAsync(new Volcengine.Cr.RepositoriesArgs
        ///         {
        ///             Names = 
        ///             {
        ///                 "repo*",
        ///             },
        ///             Registry = "tf-1",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<RepositoriesResult> Invoke(RepositoriesInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<RepositoriesResult>("volcengine:cr/repositories:Repositories", args ?? new RepositoriesInvokeArgs(), options.WithDefaults());
    }


    public sealed class RepositoriesArgs : Pulumi.InvokeArgs
    {
        [Input("accessLevels")]
        private List<string>? _accessLevels;

        /// <summary>
        /// The list of instance access level.
        /// </summary>
        public List<string> AccessLevels
        {
            get => _accessLevels ?? (_accessLevels = new List<string>());
            set => _accessLevels = value;
        }

        [Input("names")]
        private List<string>? _names;

        /// <summary>
        /// The list of instance names.
        /// </summary>
        public List<string> Names
        {
            get => _names ?? (_names = new List<string>());
            set => _names = value;
        }

        [Input("namespaces")]
        private List<string>? _namespaces;

        /// <summary>
        /// The list of instance namespace.
        /// </summary>
        public List<string> Namespaces
        {
            get => _namespaces ?? (_namespaces = new List<string>());
            set => _namespaces = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The CR instance name.
        /// </summary>
        [Input("registry", required: true)]
        public string Registry { get; set; } = null!;

        public RepositoriesArgs()
        {
        }
    }

    public sealed class RepositoriesInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("accessLevels")]
        private InputList<string>? _accessLevels;

        /// <summary>
        /// The list of instance access level.
        /// </summary>
        public InputList<string> AccessLevels
        {
            get => _accessLevels ?? (_accessLevels = new InputList<string>());
            set => _accessLevels = value;
        }

        [Input("names")]
        private InputList<string>? _names;

        /// <summary>
        /// The list of instance names.
        /// </summary>
        public InputList<string> Names
        {
            get => _names ?? (_names = new InputList<string>());
            set => _names = value;
        }

        [Input("namespaces")]
        private InputList<string>? _namespaces;

        /// <summary>
        /// The list of instance namespace.
        /// </summary>
        public InputList<string> Namespaces
        {
            get => _namespaces ?? (_namespaces = new InputList<string>());
            set => _namespaces = value;
        }

        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The CR instance name.
        /// </summary>
        [Input("registry", required: true)]
        public Input<string> Registry { get; set; } = null!;

        public RepositoriesInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class RepositoriesResult
    {
        public readonly ImmutableArray<string> AccessLevels;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Names;
        public readonly ImmutableArray<string> Namespaces;
        public readonly string? OutputFile;
        public readonly string Registry;
        /// <summary>
        /// The collection of repository query.
        /// </summary>
        public readonly ImmutableArray<Outputs.RepositoriesRepositoryResult> Repositories;
        /// <summary>
        /// The total count of instance query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private RepositoriesResult(
            ImmutableArray<string> accessLevels,

            string id,

            ImmutableArray<string> names,

            ImmutableArray<string> namespaces,

            string? outputFile,

            string registry,

            ImmutableArray<Outputs.RepositoriesRepositoryResult> repositories,

            int totalCount)
        {
            AccessLevels = accessLevels;
            Id = id;
            Names = names;
            Namespaces = namespaces;
            OutputFile = outputFile;
            Registry = registry;
            Repositories = repositories;
            TotalCount = totalCount;
        }
    }
}
