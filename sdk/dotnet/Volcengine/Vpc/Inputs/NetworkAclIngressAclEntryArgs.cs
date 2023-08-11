// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Vpc.Inputs
{

    public sealed class NetworkAclIngressAclEntryArgs : global::Pulumi.ResourceArgs
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
        /// The policy of entry, default is `accept`. The value can be `accept` or `drop`.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// The port of entry. Default is `-1/-1`. When Protocol is `all`, `icmp` or `gre`, the port range is `-1/-1`, which means no port restriction. When the Protocol is `tcp` or `udp`, the port range is `1~65535`, and the format is `1/200`, `80/80`, which means port 1 to port 200, port 80.
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The protocol of entry, default is `all`. The value can be `icmp` or `gre` or `tcp` or `udp` or `all`.
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
        public static new NetworkAclIngressAclEntryArgs Empty => new NetworkAclIngressAclEntryArgs();
    }
}
