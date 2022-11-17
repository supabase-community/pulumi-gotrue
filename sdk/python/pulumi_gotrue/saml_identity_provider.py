# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['SamlIdentityProviderArgs', 'SamlIdentityProvider']

@pulumi.input_type
class SamlIdentityProviderArgs:
    def __init__(__self__, *,
                 attribute_mapping: Optional[pulumi.Input[str]] = None,
                 domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metadata_url: Optional[pulumi.Input[str]] = None,
                 metadata_xml: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SamlIdentityProvider resource.
        """
        if attribute_mapping is not None:
            pulumi.set(__self__, "attribute_mapping", attribute_mapping)
        if domains is not None:
            pulumi.set(__self__, "domains", domains)
        if metadata_url is not None:
            pulumi.set(__self__, "metadata_url", metadata_url)
        if metadata_xml is not None:
            pulumi.set(__self__, "metadata_xml", metadata_xml)

    @property
    @pulumi.getter(name="attributeMapping")
    def attribute_mapping(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "attribute_mapping")

    @attribute_mapping.setter
    def attribute_mapping(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attribute_mapping", value)

    @property
    @pulumi.getter
    def domains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "domains")

    @domains.setter
    def domains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "domains", value)

    @property
    @pulumi.getter(name="metadataUrl")
    def metadata_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "metadata_url")

    @metadata_url.setter
    def metadata_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata_url", value)

    @property
    @pulumi.getter(name="metadataXml")
    def metadata_xml(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "metadata_xml")

    @metadata_xml.setter
    def metadata_xml(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata_xml", value)


@pulumi.input_type
class _SamlIdentityProviderState:
    def __init__(__self__, *,
                 attribute_mapping: Optional[pulumi.Input[str]] = None,
                 created_at: Optional[pulumi.Input[str]] = None,
                 domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metadata_url: Optional[pulumi.Input[str]] = None,
                 metadata_xml: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SamlIdentityProvider resources.
        """
        if attribute_mapping is not None:
            pulumi.set(__self__, "attribute_mapping", attribute_mapping)
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if domains is not None:
            pulumi.set(__self__, "domains", domains)
        if metadata_url is not None:
            pulumi.set(__self__, "metadata_url", metadata_url)
        if metadata_xml is not None:
            pulumi.set(__self__, "metadata_xml", metadata_xml)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter(name="attributeMapping")
    def attribute_mapping(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "attribute_mapping")

    @attribute_mapping.setter
    def attribute_mapping(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attribute_mapping", value)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def domains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "domains")

    @domains.setter
    def domains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "domains", value)

    @property
    @pulumi.getter(name="metadataUrl")
    def metadata_url(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "metadata_url")

    @metadata_url.setter
    def metadata_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata_url", value)

    @property
    @pulumi.getter(name="metadataXml")
    def metadata_xml(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "metadata_xml")

    @metadata_xml.setter
    def metadata_xml(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata_xml", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class SamlIdentityProvider(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_mapping: Optional[pulumi.Input[str]] = None,
                 domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metadata_url: Optional[pulumi.Input[str]] = None,
                 metadata_xml: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SamlIdentityProvider resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SamlIdentityProviderArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SamlIdentityProvider resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SamlIdentityProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SamlIdentityProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_mapping: Optional[pulumi.Input[str]] = None,
                 domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metadata_url: Optional[pulumi.Input[str]] = None,
                 metadata_xml: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SamlIdentityProviderArgs.__new__(SamlIdentityProviderArgs)

            __props__.__dict__["attribute_mapping"] = attribute_mapping
            __props__.__dict__["domains"] = domains
            __props__.__dict__["metadata_url"] = metadata_url
            __props__.__dict__["metadata_xml"] = metadata_xml
            __props__.__dict__["created_at"] = None
            __props__.__dict__["updated_at"] = None
        super(SamlIdentityProvider, __self__).__init__(
            'gotrue:index/samlIdentityProvider:SamlIdentityProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attribute_mapping: Optional[pulumi.Input[str]] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            metadata_url: Optional[pulumi.Input[str]] = None,
            metadata_xml: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'SamlIdentityProvider':
        """
        Get an existing SamlIdentityProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SamlIdentityProviderState.__new__(_SamlIdentityProviderState)

        __props__.__dict__["attribute_mapping"] = attribute_mapping
        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["domains"] = domains
        __props__.__dict__["metadata_url"] = metadata_url
        __props__.__dict__["metadata_xml"] = metadata_xml
        __props__.__dict__["updated_at"] = updated_at
        return SamlIdentityProvider(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attributeMapping")
    def attribute_mapping(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "attribute_mapping")

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def domains(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "domains")

    @property
    @pulumi.getter(name="metadataUrl")
    def metadata_url(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "metadata_url")

    @property
    @pulumi.getter(name="metadataXml")
    def metadata_xml(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "metadata_xml")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        return pulumi.get(self, "updated_at")

