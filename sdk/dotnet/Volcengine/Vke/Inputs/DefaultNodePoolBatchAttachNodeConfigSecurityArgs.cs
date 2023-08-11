// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vke.Inputs
{

    public sealed class DefaultNodePoolBatchAttachNodeConfigSecurityArgs : global::Pulumi.ResourceArgs
    {
        [Input("logins")]
        private InputList<Inputs.DefaultNodePoolBatchAttachNodeConfigSecurityLoginArgs>? _logins;

        /// <summary>
        /// The Login of Security.
        /// </summary>
        public InputList<Inputs.DefaultNodePoolBatchAttachNodeConfigSecurityLoginArgs> Logins
        {
            get => _logins ?? (_logins = new InputList<Inputs.DefaultNodePoolBatchAttachNodeConfigSecurityLoginArgs>());
            set => _logins = value;
        }

        [Input("securityGroupIds")]
        private InputList<string>? _securityGroupIds;

        /// <summary>
        /// The SecurityGroupIds of Security.
        /// </summary>
        public InputList<string> SecurityGroupIds
        {
            get => _securityGroupIds ?? (_securityGroupIds = new InputList<string>());
            set => _securityGroupIds = value;
        }

        [Input("securityStrategies")]
        private InputList<string>? _securityStrategies;

        /// <summary>
        /// The SecurityStrategies of Security.
        /// </summary>
        public InputList<string> SecurityStrategies
        {
            get => _securityStrategies ?? (_securityStrategies = new InputList<string>());
            set => _securityStrategies = value;
        }

        public DefaultNodePoolBatchAttachNodeConfigSecurityArgs()
        {
        }
        public static new DefaultNodePoolBatchAttachNodeConfigSecurityArgs Empty => new DefaultNodePoolBatchAttachNodeConfigSecurityArgs();
    }
}
