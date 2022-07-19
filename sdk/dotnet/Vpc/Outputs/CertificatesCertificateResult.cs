// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Vpc.Outputs
{

    [OutputType]
    public sealed class CertificatesCertificateResult
    {
        /// <summary>
        /// The ID of the Certificate.
        /// </summary>
        public readonly string CertificateId;
        /// <summary>
        /// The name of the Certificate.
        /// </summary>
        public readonly string CertificateName;
        /// <summary>
        /// The create time of the Certificate.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The description of the Certificate.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The domain name of the Certificate.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// The expire time of the Certificate.
        /// </summary>
        public readonly string ExpiredAt;
        /// <summary>
        /// The ID of the Certificate.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The ID list of the Listener.
        /// </summary>
        public readonly ImmutableArray<string> Listeners;

        [OutputConstructor]
        private CertificatesCertificateResult(
            string certificateId,

            string certificateName,

            string createTime,

            string description,

            string domainName,

            string expiredAt,

            string id,

            ImmutableArray<string> listeners)
        {
            CertificateId = certificateId;
            CertificateName = certificateName;
            CreateTime = createTime;
            Description = description;
            DomainName = domainName;
            ExpiredAt = expiredAt;
            Id = id;
            Listeners = listeners;
        }
    }
}
