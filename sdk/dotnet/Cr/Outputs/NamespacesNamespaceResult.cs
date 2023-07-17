// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Cr.Outputs
{

    [OutputType]
    public sealed class NamespacesNamespaceResult
    {
        /// <summary>
        /// The time when namespace created.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The name of OCI repository.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private NamespacesNamespaceResult(
            string createTime,

            string name)
        {
            CreateTime = createTime;
            Name = name;
        }
    }
}
