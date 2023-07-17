// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc.Inputs
{

    public sealed class NetworkAclIngressAclEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of entry.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("networkAclEntryId")]
        public Input<string>? NetworkAclEntryId { get; set; }

        /// <summary>
        /// The name of entry.
        /// </summary>
        [Input("networkAclEntryName")]
        public Input<string>? NetworkAclEntryName { get; set; }

        /// <summary>
        /// The policy of entry.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The port of entry.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The protocol of entry.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The SourceCidrIp of entry.
        /// </summary>
        [Input("sourceCidrIp")]
        public Input<string>? SourceCidrIp { get; set; }

        public NetworkAclIngressAclEntryArgs()
        {
        }
    }
}
