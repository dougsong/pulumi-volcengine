// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Cr
{
    public static class AuthorizationTokens
    {
        /// <summary>
        /// Use this data source to query detailed information of cr authorization tokens
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
        ///     var foo = Volcengine.Cr.AuthorizationTokens.Invoke(new()
        ///     {
        ///         Registry = "tf-1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<AuthorizationTokensResult> InvokeAsync(AuthorizationTokensArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<AuthorizationTokensResult>("volcengine:cr/authorizationTokens:AuthorizationTokens", args ?? new AuthorizationTokensArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to query detailed information of cr authorization tokens
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
        ///     var foo = Volcengine.Cr.AuthorizationTokens.Invoke(new()
        ///     {
        ///         Registry = "tf-1",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<AuthorizationTokensResult> Invoke(AuthorizationTokensInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<AuthorizationTokensResult>("volcengine:cr/authorizationTokens:AuthorizationTokens", args ?? new AuthorizationTokensInvokeArgs(), options.WithDefaults());
    }


    public sealed class AuthorizationTokensArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The cr instance name want to query.
        /// </summary>
        [Input("registry", required: true)]
        public string Registry { get; set; } = null!;

        public AuthorizationTokensArgs()
        {
        }
        public static new AuthorizationTokensArgs Empty => new AuthorizationTokensArgs();
    }

    public sealed class AuthorizationTokensInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// File name where to save data source results.
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The cr instance name want to query.
        /// </summary>
        [Input("registry", required: true)]
        public Input<string> Registry { get; set; } = null!;

        public AuthorizationTokensInvokeArgs()
        {
        }
        public static new AuthorizationTokensInvokeArgs Empty => new AuthorizationTokensInvokeArgs();
    }


    [OutputType]
    public sealed class AuthorizationTokensResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? OutputFile;
        public readonly string Registry;
        /// <summary>
        /// The collection of users.
        /// </summary>
        public readonly ImmutableArray<Outputs.AuthorizationTokensTokenResult> Tokens;
        /// <summary>
        /// The total count of instance query.
        /// </summary>
        public readonly int TotalCount;

        [OutputConstructor]
        private AuthorizationTokensResult(
            string id,

            string? outputFile,

            string registry,

            ImmutableArray<Outputs.AuthorizationTokensTokenResult> tokens,

            int totalCount)
        {
            Id = id;
            OutputFile = outputFile;
            Registry = registry;
            Tokens = tokens;
            TotalCount = totalCount;
        }
    }
}
