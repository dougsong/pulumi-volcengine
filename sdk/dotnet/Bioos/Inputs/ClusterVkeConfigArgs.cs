// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Bioos.Inputs
{

    public sealed class ClusterVkeConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The id of the vke cluster.
        /// </summary>
        [Input("clusterId", required: true)]
        public Input<string> ClusterId { get; set; } = null!;

        /// <summary>
        /// The name of the StorageClass that the vke cluster has installed.
        /// </summary>
        [Input("storageClass", required: true)]
        public Input<string> StorageClass { get; set; } = null!;

        public ClusterVkeConfigArgs()
        {
        }
    }
}
