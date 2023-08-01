# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['RepositoryArgs', 'Repository']

@pulumi.input_type
class RepositoryArgs:
    def __init__(__self__, *,
                 namespace: pulumi.Input[str],
                 registry: pulumi.Input[str],
                 access_level: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Repository resource.
        :param pulumi.Input[str] namespace: The target namespace name.
        :param pulumi.Input[str] registry: The CrRegistry name.
        :param pulumi.Input[str] access_level: The access level of CrRepository.
        :param pulumi.Input[str] description: The description of CrRepository.
        :param pulumi.Input[str] name: The name of CrRepository.
        """
        pulumi.set(__self__, "namespace", namespace)
        pulumi.set(__self__, "registry", registry)
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Input[str]:
        """
        The target namespace name.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: pulumi.Input[str]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def registry(self) -> pulumi.Input[str]:
        """
        The CrRegistry name.
        """
        return pulumi.get(self, "registry")

    @registry.setter
    def registry(self, value: pulumi.Input[str]):
        pulumi.set(self, "registry", value)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        The access level of CrRepository.
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of CrRepository.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of CrRepository.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RepositoryState:
    def __init__(__self__, *,
                 access_level: Optional[pulumi.Input[str]] = None,
                 create_time: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 registry: Optional[pulumi.Input[str]] = None,
                 update_time: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Repository resources.
        :param pulumi.Input[str] access_level: The access level of CrRepository.
        :param pulumi.Input[str] create_time: The creation time of repository.
        :param pulumi.Input[str] description: The description of CrRepository.
        :param pulumi.Input[str] name: The name of CrRepository.
        :param pulumi.Input[str] namespace: The target namespace name.
        :param pulumi.Input[str] registry: The CrRegistry name.
        :param pulumi.Input[str] update_time: The last update time of repository.
        """
        if access_level is not None:
            pulumi.set(__self__, "access_level", access_level)
        if create_time is not None:
            pulumi.set(__self__, "create_time", create_time)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if registry is not None:
            pulumi.set(__self__, "registry", registry)
        if update_time is not None:
            pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> Optional[pulumi.Input[str]]:
        """
        The access level of CrRepository.
        """
        return pulumi.get(self, "access_level")

    @access_level.setter
    def access_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_level", value)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> Optional[pulumi.Input[str]]:
        """
        The creation time of repository.
        """
        return pulumi.get(self, "create_time")

    @create_time.setter
    def create_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "create_time", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of CrRepository.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of CrRepository.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The target namespace name.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def registry(self) -> Optional[pulumi.Input[str]]:
        """
        The CrRegistry name.
        """
        return pulumi.get(self, "registry")

    @registry.setter
    def registry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "registry", value)

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> Optional[pulumi.Input[str]]:
        """
        The last update time of repository.
        """
        return pulumi.get(self, "update_time")

    @update_time.setter
    def update_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "update_time", value)


class Repository(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 registry: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage cr repository
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cr.Repository("foo",
            access_level="Public",
            description="A test repository created by terraform.",
            namespace="namespace-1",
            registry="tf-2")
        ```

        ## Import

        CR Repository can be imported using the registry:namespace:name, e.g.

        ```sh
         $ pulumi import volcengine:cr/repository:Repository default cr-basic:namespace-1:repo-1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: The access level of CrRepository.
        :param pulumi.Input[str] description: The description of CrRepository.
        :param pulumi.Input[str] name: The name of CrRepository.
        :param pulumi.Input[str] namespace: The target namespace name.
        :param pulumi.Input[str] registry: The CrRegistry name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RepositoryArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage cr repository
        ## Example Usage

        ```python
        import pulumi
        import pulumi_volcengine as volcengine

        foo = volcengine.cr.Repository("foo",
            access_level="Public",
            description="A test repository created by terraform.",
            namespace="namespace-1",
            registry="tf-2")
        ```

        ## Import

        CR Repository can be imported using the registry:namespace:name, e.g.

        ```sh
         $ pulumi import volcengine:cr/repository:Repository default cr-basic:namespace-1:repo-1
        ```

        :param str resource_name: The name of the resource.
        :param RepositoryArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RepositoryArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_level: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 registry: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RepositoryArgs.__new__(RepositoryArgs)

            __props__.__dict__["access_level"] = access_level
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if namespace is None and not opts.urn:
                raise TypeError("Missing required property 'namespace'")
            __props__.__dict__["namespace"] = namespace
            if registry is None and not opts.urn:
                raise TypeError("Missing required property 'registry'")
            __props__.__dict__["registry"] = registry
            __props__.__dict__["create_time"] = None
            __props__.__dict__["update_time"] = None
        super(Repository, __self__).__init__(
            'volcengine:cr/repository:Repository',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_level: Optional[pulumi.Input[str]] = None,
            create_time: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            registry: Optional[pulumi.Input[str]] = None,
            update_time: Optional[pulumi.Input[str]] = None) -> 'Repository':
        """
        Get an existing Repository resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_level: The access level of CrRepository.
        :param pulumi.Input[str] create_time: The creation time of repository.
        :param pulumi.Input[str] description: The description of CrRepository.
        :param pulumi.Input[str] name: The name of CrRepository.
        :param pulumi.Input[str] namespace: The target namespace name.
        :param pulumi.Input[str] registry: The CrRegistry name.
        :param pulumi.Input[str] update_time: The last update time of repository.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RepositoryState.__new__(_RepositoryState)

        __props__.__dict__["access_level"] = access_level
        __props__.__dict__["create_time"] = create_time
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["registry"] = registry
        __props__.__dict__["update_time"] = update_time
        return Repository(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessLevel")
    def access_level(self) -> pulumi.Output[Optional[str]]:
        """
        The access level of CrRepository.
        """
        return pulumi.get(self, "access_level")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation time of repository.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of CrRepository.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of CrRepository.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[str]:
        """
        The target namespace name.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def registry(self) -> pulumi.Output[str]:
        """
        The CrRegistry name.
        """
        return pulumi.get(self, "registry")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The last update time of repository.
        """
        return pulumi.get(self, "update_time")

