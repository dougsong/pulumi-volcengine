# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['TopicArgs', 'Topic']

@pulumi.input_type
class TopicArgs:
    def __init__(__self__, *,
                 project_id: pulumi.Input[str],
                 shard_count: pulumi.Input[int],
                 topic_name: pulumi.Input[str],
                 ttl: pulumi.Input[int],
                 auto_split: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_tracking: Optional[pulumi.Input[bool]] = None,
                 max_split_shard: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TopicTagArgs']]]] = None,
                 time_format: Optional[pulumi.Input[str]] = None,
                 time_key: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Topic resource.
        :param pulumi.Input[str] project_id: The project id of the tls topic.
        :param pulumi.Input[int] shard_count: The count of shards in the tls topic. Valid value range: 1-10.
        :param pulumi.Input[str] topic_name: The name of the tls topic.
        :param pulumi.Input[int] ttl: The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
        :param pulumi.Input[bool] auto_split: Whether to enable automatic partition splitting function of the tls topic.
        :param pulumi.Input[str] description: The description of the tls project.
        :param pulumi.Input[bool] enable_tracking: Whether to enable WebTracking function of the tls topic.
        :param pulumi.Input[int] max_split_shard: The max count of shards in the tls topic.
        :param pulumi.Input[Sequence[pulumi.Input['TopicTagArgs']]] tags: Tags.
        :param pulumi.Input[str] time_format: The format of the time field.
        :param pulumi.Input[str] time_key: The name of the time field.
        """
        pulumi.set(__self__, "project_id", project_id)
        pulumi.set(__self__, "shard_count", shard_count)
        pulumi.set(__self__, "topic_name", topic_name)
        pulumi.set(__self__, "ttl", ttl)
        if auto_split is not None:
            pulumi.set(__self__, "auto_split", auto_split)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_tracking is not None:
            pulumi.set(__self__, "enable_tracking", enable_tracking)
        if max_split_shard is not None:
            pulumi.set(__self__, "max_split_shard", max_split_shard)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if time_format is not None:
            pulumi.set(__self__, "time_format", time_format)
        if time_key is not None:
            pulumi.set(__self__, "time_key", time_key)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Input[str]:
        """
        The project id of the tls topic.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> pulumi.Input[int]:
        """
        The count of shards in the tls topic. Valid value range: 1-10.
        """
        return pulumi.get(self, "shard_count")

    @shard_count.setter
    def shard_count(self, value: pulumi.Input[int]):
        pulumi.set(self, "shard_count", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Input[str]:
        """
        The name of the tls topic.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Input[int]:
        """
        The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: pulumi.Input[int]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter(name="autoSplit")
    def auto_split(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable automatic partition splitting function of the tls topic.
        """
        return pulumi.get(self, "auto_split")

    @auto_split.setter
    def auto_split(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_split", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the tls project.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableTracking")
    def enable_tracking(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable WebTracking function of the tls topic.
        """
        return pulumi.get(self, "enable_tracking")

    @enable_tracking.setter
    def enable_tracking(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_tracking", value)

    @property
    @pulumi.getter(name="maxSplitShard")
    def max_split_shard(self) -> Optional[pulumi.Input[int]]:
        """
        The max count of shards in the tls topic.
        """
        return pulumi.get(self, "max_split_shard")

    @max_split_shard.setter
    def max_split_shard(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_split_shard", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TopicTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TopicTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="timeFormat")
    def time_format(self) -> Optional[pulumi.Input[str]]:
        """
        The format of the time field.
        """
        return pulumi.get(self, "time_format")

    @time_format.setter
    def time_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_format", value)

    @property
    @pulumi.getter(name="timeKey")
    def time_key(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the time field.
        """
        return pulumi.get(self, "time_key")

    @time_key.setter
    def time_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_key", value)


@pulumi.input_type
class _TopicState:
    def __init__(__self__, *,
                 auto_split: Optional[pulumi.Input[bool]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_tracking: Optional[pulumi.Input[bool]] = None,
                 max_split_shard: Optional[pulumi.Input[int]] = None,
                 modify_time: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 shard_count: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['TopicTagArgs']]]] = None,
                 time_format: Optional[pulumi.Input[str]] = None,
                 time_key: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering Topic resources.
        :param pulumi.Input[bool] auto_split: Whether to enable automatic partition splitting function of the tls topic.
        :param pulumi.Input[str] create_time: The create time of the tls topic.
        :param pulumi.Input[str] description: The description of the tls project.
        :param pulumi.Input[bool] enable_tracking: Whether to enable WebTracking function of the tls topic.
        :param pulumi.Input[int] max_split_shard: The max count of shards in the tls topic.
        :param pulumi.Input[str] modify_time: The modify time of the tls topic.
        :param pulumi.Input[str] project_id: The project id of the tls topic.
        :param pulumi.Input[int] shard_count: The count of shards in the tls topic. Valid value range: 1-10.
        :param pulumi.Input[Sequence[pulumi.Input['TopicTagArgs']]] tags: Tags.
        :param pulumi.Input[str] time_format: The format of the time field.
        :param pulumi.Input[str] time_key: The name of the time field.
        :param pulumi.Input[str] topic_name: The name of the tls topic.
        :param pulumi.Input[int] ttl: The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
        """
        if auto_split is not None:
            pulumi.set(__self__, "auto_split", auto_split)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if enable_tracking is not None:
            pulumi.set(__self__, "enable_tracking", enable_tracking)
        if max_split_shard is not None:
            pulumi.set(__self__, "max_split_shard", max_split_shard)
        if modify_time is not None:
            pulumi.set(__self__, "modify_time", modify_time)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if shard_count is not None:
            pulumi.set(__self__, "shard_count", shard_count)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if time_format is not None:
            pulumi.set(__self__, "time_format", time_format)
        if time_key is not None:
            pulumi.set(__self__, "time_key", time_key)
        if topic_name is not None:
            pulumi.set(__self__, "topic_name", topic_name)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter(name="autoSplit")
    def auto_split(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable automatic partition splitting function of the tls topic.
        """
        return pulumi.get(self, "auto_split")

    @auto_split.setter
    def auto_split(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_split", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The create time of the tls topic.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the tls project.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="enableTracking")
    def enable_tracking(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether to enable WebTracking function of the tls topic.
        """
        return pulumi.get(self, "enable_tracking")

    @enable_tracking.setter
    def enable_tracking(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_tracking", value)

    @property
    @pulumi.getter(name="maxSplitShard")
    def max_split_shard(self) -> Optional[pulumi.Input[int]]:
        """
        The max count of shards in the tls topic.
        """
        return pulumi.get(self, "max_split_shard")

    @max_split_shard.setter
    def max_split_shard(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_split_shard", value)

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> Optional[pulumi.Input[str]]:
        """
        The modify time of the tls topic.
        """
        return pulumi.get(self, "modify_time")

    @modify_time.setter
    def modify_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "modify_time", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The project id of the tls topic.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> Optional[pulumi.Input[int]]:
        """
        The count of shards in the tls topic. Valid value range: 1-10.
        """
        return pulumi.get(self, "shard_count")

    @shard_count.setter
    def shard_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "shard_count", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['TopicTagArgs']]]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['TopicTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="timeFormat")
    def time_format(self) -> Optional[pulumi.Input[str]]:
        """
        The format of the time field.
        """
        return pulumi.get(self, "time_format")

    @time_format.setter
    def time_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_format", value)

    @property
    @pulumi.getter(name="timeKey")
    def time_key(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the time field.
        """
        return pulumi.get(self, "time_key")

    @time_key.setter
    def time_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "time_key", value)

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the tls topic.
        """
        return pulumi.get(self, "topic_name")

    @topic_name.setter
    def topic_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "topic_name", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)


class Topic(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_split: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_tracking: Optional[pulumi.Input[bool]] = None,
                 max_split_shard: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 shard_count: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TopicTagArgs']]]]] = None,
                 time_format: Optional[pulumi.Input[str]] = None,
                 time_key: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Provides a resource to manage tls topic
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.tls.Topic("foo",
            auto_split=True,
            description="test",
            enable_tracking=True,
            max_split_shard=10,
            project_id="e020c978-4f05-40e1-9167-0113d3ef****",
            shard_count=2,
            tags=[volcengine.tls.TopicTagArgs(
                key="k1",
                value="v1",
            )],
            time_format="%Y-%m-%dT%H:%M:%S,%f",
            time_key="request_time",
            topic_name="tf-test-topic",
            ttl=10)
        ```

        ## Import

        Tls Topic can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:tls/topic:Topic default edf051ed-3c46-49ba-9339-bea628fe****
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_split: Whether to enable automatic partition splitting function of the tls topic.
        :param pulumi.Input[str] description: The description of the tls project.
        :param pulumi.Input[bool] enable_tracking: Whether to enable WebTracking function of the tls topic.
        :param pulumi.Input[int] max_split_shard: The max count of shards in the tls topic.
        :param pulumi.Input[str] project_id: The project id of the tls topic.
        :param pulumi.Input[int] shard_count: The count of shards in the tls topic. Valid value range: 1-10.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TopicTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] time_format: The format of the time field.
        :param pulumi.Input[str] time_key: The name of the time field.
        :param pulumi.Input[str] topic_name: The name of the tls topic.
        :param pulumi.Input[int] ttl: The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TopicArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage tls topic
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.tls.Topic("foo",
            auto_split=True,
            description="test",
            enable_tracking=True,
            max_split_shard=10,
            project_id="e020c978-4f05-40e1-9167-0113d3ef****",
            shard_count=2,
            tags=[volcengine.tls.TopicTagArgs(
                key="k1",
                value="v1",
            )],
            time_format="%Y-%m-%dT%H:%M:%S,%f",
            time_key="request_time",
            topic_name="tf-test-topic",
            ttl=10)
        ```

        ## Import

        Tls Topic can be imported using the id, e.g.

        ```sh
         $ pulumi import volcengine:tls/topic:Topic default edf051ed-3c46-49ba-9339-bea628fe****
        ```

        :param str resource_name: The name of the resource.
        :param TopicArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TopicArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_split: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 enable_tracking: Optional[pulumi.Input[bool]] = None,
                 max_split_shard: Optional[pulumi.Input[int]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 shard_count: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TopicTagArgs']]]]] = None,
                 time_format: Optional[pulumi.Input[str]] = None,
                 time_key: Optional[pulumi.Input[str]] = None,
                 topic_name: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TopicArgs.__new__(TopicArgs)

            __props__.__dict__["auto_split"] = auto_split
            __props__.__dict__["description"] = description
            __props__.__dict__["enable_tracking"] = enable_tracking
            __props__.__dict__["max_split_shard"] = max_split_shard
            if project_id is None and not opts.urn:
                raise TypeError("Missing required property 'project_id'")
            __props__.__dict__["project_id"] = project_id
            if shard_count is None and not opts.urn:
                raise TypeError("Missing required property 'shard_count'")
            __props__.__dict__["shard_count"] = shard_count
            __props__.__dict__["tags"] = tags
            __props__.__dict__["time_format"] = time_format
            __props__.__dict__["time_key"] = time_key
            if topic_name is None and not opts.urn:
                raise TypeError("Missing required property 'topic_name'")
            __props__.__dict__["topic_name"] = topic_name
            if ttl is None and not opts.urn:
                raise TypeError("Missing required property 'ttl'")
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["create_time"] = None
            __props__.__dict__["modify_time"] = None
        super(Topic, __self__).__init__(
            'volcengine:tls/topic:Topic',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_split: Optional[pulumi.Input[bool]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            enable_tracking: Optional[pulumi.Input[bool]] = None,
            max_split_shard: Optional[pulumi.Input[int]] = None,
            modify_time: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            shard_count: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TopicTagArgs']]]]] = None,
            time_format: Optional[pulumi.Input[str]] = None,
            time_key: Optional[pulumi.Input[str]] = None,
            topic_name: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[int]] = None) -> 'Topic':
        """
        Get an existing Topic resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_split: Whether to enable automatic partition splitting function of the tls topic.
        :param pulumi.Input[str] create_time: The create time of the tls topic.
        :param pulumi.Input[str] description: The description of the tls project.
        :param pulumi.Input[bool] enable_tracking: Whether to enable WebTracking function of the tls topic.
        :param pulumi.Input[int] max_split_shard: The max count of shards in the tls topic.
        :param pulumi.Input[str] modify_time: The modify time of the tls topic.
        :param pulumi.Input[str] project_id: The project id of the tls topic.
        :param pulumi.Input[int] shard_count: The count of shards in the tls topic. Valid value range: 1-10.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['TopicTagArgs']]]] tags: Tags.
        :param pulumi.Input[str] time_format: The format of the time field.
        :param pulumi.Input[str] time_key: The name of the time field.
        :param pulumi.Input[str] topic_name: The name of the tls topic.
        :param pulumi.Input[int] ttl: The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TopicState.__new__(_TopicState)

        __props__.__dict__["auto_split"] = auto_split
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["enable_tracking"] = enable_tracking
        __props__.__dict__["max_split_shard"] = max_split_shard
        __props__.__dict__["modify_time"] = modify_time
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["shard_count"] = shard_count
        __props__.__dict__["tags"] = tags
        __props__.__dict__["time_format"] = time_format
        __props__.__dict__["time_key"] = time_key
        __props__.__dict__["topic_name"] = topic_name
        __props__.__dict__["ttl"] = ttl
        return Topic(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoSplit")
    def auto_split(self) -> pulumi.Output[bool]:
        """
        Whether to enable automatic partition splitting function of the tls topic.
        """
        return pulumi.get(self, "auto_split")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The create time of the tls topic.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the tls project.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableTracking")
    def enable_tracking(self) -> pulumi.Output[bool]:
        """
        Whether to enable WebTracking function of the tls topic.
        """
        return pulumi.get(self, "enable_tracking")

    @property
    @pulumi.getter(name="maxSplitShard")
    def max_split_shard(self) -> pulumi.Output[int]:
        """
        The max count of shards in the tls topic.
        """
        return pulumi.get(self, "max_split_shard")

    @property
    @pulumi.getter(name="modifyTime")
    def modify_time(self) -> pulumi.Output[str]:
        """
        The modify time of the tls topic.
        """
        return pulumi.get(self, "modify_time")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[str]:
        """
        The project id of the tls topic.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="shardCount")
    def shard_count(self) -> pulumi.Output[int]:
        """
        The count of shards in the tls topic. Valid value range: 1-10.
        """
        return pulumi.get(self, "shard_count")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.TopicTag']]]:
        """
        Tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="timeFormat")
    def time_format(self) -> pulumi.Output[str]:
        """
        The format of the time field.
        """
        return pulumi.get(self, "time_format")

    @property
    @pulumi.getter(name="timeKey")
    def time_key(self) -> pulumi.Output[str]:
        """
        The name of the time field.
        """
        return pulumi.get(self, "time_key")

    @property
    @pulumi.getter(name="topicName")
    def topic_name(self) -> pulumi.Output[str]:
        """
        The name of the tls topic.
        """
        return pulumi.get(self, "topic_name")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[int]:
        """
        The data storage time of the tls topic. Unit: Day. Valid value range: 1-3650.
        """
        return pulumi.get(self, "ttl")

