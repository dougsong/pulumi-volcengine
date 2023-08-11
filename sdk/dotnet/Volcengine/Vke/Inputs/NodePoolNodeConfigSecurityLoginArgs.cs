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

    public sealed class NodePoolNodeConfigSecurityLoginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Password of Security, this field must be encoded with base64.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// The SshKeyPairName of Security.
        /// </summary>
        [Input("sshKeyPairName")]
        public Input<string>? SshKeyPairName { get; set; }

        public NodePoolNodeConfigSecurityLoginArgs()
        {
        }
        public static new NodePoolNodeConfigSecurityLoginArgs Empty => new NodePoolNodeConfigSecurityLoginArgs();
    }
}
