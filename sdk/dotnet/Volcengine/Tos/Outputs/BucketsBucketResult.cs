// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Volcengine.Pulumi.Volcengine.Tos.Outputs
{

    [OutputType]
    public sealed class BucketsBucketResult
    {
        /// <summary>
        /// The create date of the TOS bucket.
        /// </summary>
        public readonly string CreationDate;
        /// <summary>
        /// The extranet endpoint of the TOS bucket.
        /// </summary>
        public readonly string ExtranetEndpoint;
        /// <summary>
        /// The intranet endpoint the TOS bucket.
        /// </summary>
        public readonly string IntranetEndpoint;
        /// <summary>
        /// (**Deprecated**) The Field is Deprecated. The truncated the TOS bucket.
        /// </summary>
        public readonly bool IsTruncated;
        /// <summary>
        /// The location of the TOS bucket.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// (**Deprecated**) The Field is Deprecated. The marker the TOS bucket.
        /// </summary>
        public readonly string Marker;
        /// <summary>
        /// (**Deprecated**) The Field is Deprecated. The max keys the TOS bucket.
        /// </summary>
        public readonly int MaxKeys;
        /// <summary>
        /// The name the TOS bucket.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// (**Deprecated**) The Field is Deprecated. The prefix the TOS bucket.
        /// </summary>
        public readonly string Prefix;

        [OutputConstructor]
        private BucketsBucketResult(
            string creationDate,

            string extranetEndpoint,

            string intranetEndpoint,

            bool isTruncated,

            string location,

            string marker,

            int maxKeys,

            string name,

            string prefix)
        {
            CreationDate = creationDate;
            ExtranetEndpoint = extranetEndpoint;
            IntranetEndpoint = intranetEndpoint;
            IsTruncated = isTruncated;
            Location = location;
            Marker = marker;
            MaxKeys = maxKeys;
            Name = name;
            Prefix = prefix;
        }
    }
}
