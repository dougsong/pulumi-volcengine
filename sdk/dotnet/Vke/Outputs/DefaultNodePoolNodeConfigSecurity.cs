// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vke.Outputs
{

    [OutputType]
    public sealed class DefaultNodePoolNodeConfigSecurity
    {
        /// <summary>
        /// The Login of Security.
        /// </summary>
        public readonly Outputs.DefaultNodePoolNodeConfigSecurityLogin Login;
        /// <summary>
        /// The SecurityGroupIds of Security.
        /// </summary>
        public readonly ImmutableArray<string> SecurityGroupIds;
        /// <summary>
        /// The SecurityStrategies of Security.
        /// </summary>
        public readonly ImmutableArray<string> SecurityStrategies;

        [OutputConstructor]
        private DefaultNodePoolNodeConfigSecurity(
            Outputs.DefaultNodePoolNodeConfigSecurityLogin login,

            ImmutableArray<string> securityGroupIds,

            ImmutableArray<string> securityStrategies)
        {
            Login = login;
            SecurityGroupIds = securityGroupIds;
            SecurityStrategies = securityStrategies;
        }
    }
}
