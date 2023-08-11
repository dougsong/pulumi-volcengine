// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Cr.Outputs
{

    [OutputType]
    public sealed class RepositoriesRepositoryResult
    {
        /// <summary>
        /// The access level of repository.
        /// </summary>
        public readonly string AccessLevel;
        /// <summary>
        /// The creation time of repository.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of repository.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The name of repository.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The namespace of repository.
        /// </summary>
        public readonly string Namespace;
        /// <summary>
        /// The last update time of repository.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private RepositoriesRepositoryResult(
            string accessLevel,

            string createTime,

            string description,

            string name,

            string @namespace,

            string updateTime)
        {
            AccessLevel = accessLevel;
            CreateTime = createTime;
            Description = description;
            Name = name;
            Namespace = @namespace;
            UpdateTime = updateTime;
        }
    }
}
