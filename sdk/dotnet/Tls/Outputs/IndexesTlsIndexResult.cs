// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Volcengine.Tls.Outputs
{

    [OutputType]
    public sealed class IndexesTlsIndexResult
    {
        /// <summary>
        /// The create time of the tls index.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The FullText index of the tls topic.
        /// </summary>
        public readonly Outputs.IndexesTlsIndexFullTextResult FullText;
        /// <summary>
        /// The topic id of the tls index.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The KeyValue index of the tls topic.
        /// </summary>
        public readonly ImmutableArray<Outputs.IndexesTlsIndexKeyValueResult> KeyValues;
        /// <summary>
        /// The modify time of the tls index.
        /// </summary>
        public readonly string ModifyTime;
        /// <summary>
        /// The topic id of the tls index.
        /// </summary>
        public readonly string TopicId;

        [OutputConstructor]
        private IndexesTlsIndexResult(
            string createTime,

            Outputs.IndexesTlsIndexFullTextResult fullText,

            string id,

            ImmutableArray<Outputs.IndexesTlsIndexKeyValueResult> keyValues,

            string modifyTime,

            string topicId)
        {
            CreateTime = createTime;
            FullText = fullText;
            Id = id;
            KeyValues = keyValues;
            ModifyTime = modifyTime;
            TopicId = topicId;
        }
    }
}
