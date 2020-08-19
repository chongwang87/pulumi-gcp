# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'SubscriptionDeadLetterPolicyArgs',
    'SubscriptionExpirationPolicyArgs',
    'SubscriptionIAMBindingConditionArgs',
    'SubscriptionIAMMemberConditionArgs',
    'SubscriptionPushConfigArgs',
    'SubscriptionPushConfigOidcTokenArgs',
    'TopicIAMBindingConditionArgs',
    'TopicIAMMemberConditionArgs',
    'TopicMessageStoragePolicyArgs',
]

@pulumi.input_type
class SubscriptionDeadLetterPolicyArgs:
    def __init__(__self__, *,
                 dead_letter_topic: Optional[pulumi.Input[str]] = None,
                 max_delivery_attempts: Optional[pulumi.Input[float]] = None):
        """
        :param pulumi.Input[str] dead_letter_topic: The name of the topic to which dead letter messages should be published.
               Format is `projects/{project}/topics/{topic}`.
               The Cloud Pub/Sub service account associated with the enclosing subscription's
               parent project (i.e.,
               service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
               permission to Publish() to this topic.
               The operation will fail if the topic does not exist.
               Users should ensure that there is a subscription attached to this topic
               since messages published to a topic with no subscriptions are lost.
        :param pulumi.Input[float] max_delivery_attempts: The maximum number of delivery attempts for any message. The value must be
               between 5 and 100.
               The number of delivery attempts is defined as 1 + (the sum of number of
               NACKs and number of times the acknowledgement deadline has been exceeded for the message).
               A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
               client libraries may automatically extend ack_deadlines.
               This field will be honored on a best effort basis.
               If this parameter is 0, a default value of 5 is used.
        """
        if dead_letter_topic is not None:
            pulumi.set(__self__, "dead_letter_topic", dead_letter_topic)
        if max_delivery_attempts is not None:
            pulumi.set(__self__, "max_delivery_attempts", max_delivery_attempts)

    @property
    @pulumi.getter(name="deadLetterTopic")
    def dead_letter_topic(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the topic to which dead letter messages should be published.
        Format is `projects/{project}/topics/{topic}`.
        The Cloud Pub/Sub service account associated with the enclosing subscription's
        parent project (i.e.,
        service-{project_number}@gcp-sa-pubsub.iam.gserviceaccount.com) must have
        permission to Publish() to this topic.
        The operation will fail if the topic does not exist.
        Users should ensure that there is a subscription attached to this topic
        since messages published to a topic with no subscriptions are lost.
        """
        return pulumi.get(self, "dead_letter_topic")

    @dead_letter_topic.setter
    def dead_letter_topic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dead_letter_topic", value)

    @property
    @pulumi.getter(name="maxDeliveryAttempts")
    def max_delivery_attempts(self) -> Optional[pulumi.Input[float]]:
        """
        The maximum number of delivery attempts for any message. The value must be
        between 5 and 100.
        The number of delivery attempts is defined as 1 + (the sum of number of
        NACKs and number of times the acknowledgement deadline has been exceeded for the message).
        A NACK is any call to ModifyAckDeadline with a 0 deadline. Note that
        client libraries may automatically extend ack_deadlines.
        This field will be honored on a best effort basis.
        If this parameter is 0, a default value of 5 is used.
        """
        return pulumi.get(self, "max_delivery_attempts")

    @max_delivery_attempts.setter
    def max_delivery_attempts(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "max_delivery_attempts", value)


@pulumi.input_type
class SubscriptionExpirationPolicyArgs:
    def __init__(__self__, *,
                 ttl: pulumi.Input[str]):
        """
        :param pulumi.Input[str] ttl: Specifies the "time-to-live" duration for an associated resource. The
               resource expires if it is not active for a period of ttl.
               If ttl is not set, the associated resource never expires.
               A duration in seconds with up to nine fractional digits, terminated by 's'.
               Example - "3.5s".
        """
        pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Input[str]:
        """
        Specifies the "time-to-live" duration for an associated resource. The
        resource expires if it is not active for a period of ttl.
        If ttl is not set, the associated resource never expires.
        A duration in seconds with up to nine fractional digits, terminated by 's'.
        Example - "3.5s".
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: pulumi.Input[str]):
        pulumi.set(self, "ttl", value)


@pulumi.input_type
class SubscriptionIAMBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class SubscriptionIAMMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class SubscriptionPushConfigArgs:
    def __init__(__self__, *,
                 push_endpoint: pulumi.Input[str],
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 oidc_token: Optional[pulumi.Input['SubscriptionPushConfigOidcTokenArgs']] = None):
        """
        :param pulumi.Input[str] push_endpoint: A URL locating the endpoint to which messages should be pushed.
               For example, a Webhook endpoint might use
               "https://example.com/push".
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] attributes: Endpoint configuration attributes.
               Every endpoint has a set of API supported attributes that can
               be used to control different aspects of the message delivery.
               The currently supported attribute is x-goog-version, which you
               can use to change the format of the pushed message. This
               attribute indicates the version of the data expected by
               the endpoint. This controls the shape of the pushed message
               (i.e., its fields and metadata). The endpoint version is
               based on the version of the Pub/Sub API.
               If not present during the subscriptions.create call,
               it will default to the version of the API used to make
               such call. If not present during a subscriptions.modifyPushConfig
               call, its value will not be changed. subscriptions.get
               calls will always return a valid version, even if the
               subscription was created without this attribute.
               The possible values for this attribute are:
               - v1beta1: uses the push format defined in the v1beta1 Pub/Sub API.
               - v1 or v1beta2: uses the push format defined in the v1 Pub/Sub API.
        :param pulumi.Input['SubscriptionPushConfigOidcTokenArgs'] oidc_token: If specified, Pub/Sub will generate and attach an OIDC JWT token as
               an Authorization header in the HTTP request for every pushed message.
               Structure is documented below.
        """
        pulumi.set(__self__, "push_endpoint", push_endpoint)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if oidc_token is not None:
            pulumi.set(__self__, "oidc_token", oidc_token)

    @property
    @pulumi.getter(name="pushEndpoint")
    def push_endpoint(self) -> pulumi.Input[str]:
        """
        A URL locating the endpoint to which messages should be pushed.
        For example, a Webhook endpoint might use
        "https://example.com/push".
        """
        return pulumi.get(self, "push_endpoint")

    @push_endpoint.setter
    def push_endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "push_endpoint", value)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Endpoint configuration attributes.
        Every endpoint has a set of API supported attributes that can
        be used to control different aspects of the message delivery.
        The currently supported attribute is x-goog-version, which you
        can use to change the format of the pushed message. This
        attribute indicates the version of the data expected by
        the endpoint. This controls the shape of the pushed message
        (i.e., its fields and metadata). The endpoint version is
        based on the version of the Pub/Sub API.
        If not present during the subscriptions.create call,
        it will default to the version of the API used to make
        such call. If not present during a subscriptions.modifyPushConfig
        call, its value will not be changed. subscriptions.get
        calls will always return a valid version, even if the
        subscription was created without this attribute.
        The possible values for this attribute are:
        - v1beta1: uses the push format defined in the v1beta1 Pub/Sub API.
        - v1 or v1beta2: uses the push format defined in the v1 Pub/Sub API.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter(name="oidcToken")
    def oidc_token(self) -> Optional[pulumi.Input['SubscriptionPushConfigOidcTokenArgs']]:
        """
        If specified, Pub/Sub will generate and attach an OIDC JWT token as
        an Authorization header in the HTTP request for every pushed message.
        Structure is documented below.
        """
        return pulumi.get(self, "oidc_token")

    @oidc_token.setter
    def oidc_token(self, value: Optional[pulumi.Input['SubscriptionPushConfigOidcTokenArgs']]):
        pulumi.set(self, "oidc_token", value)


@pulumi.input_type
class SubscriptionPushConfigOidcTokenArgs:
    def __init__(__self__, *,
                 service_account_email: pulumi.Input[str],
                 audience: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] service_account_email: Service account email to be used for generating the OIDC token.
               The caller (for subscriptions.create, subscriptions.patch, and
               subscriptions.modifyPushConfig RPCs) must have the
               iam.serviceAccounts.actAs permission for the service account.
        :param pulumi.Input[str] audience: Audience to be used when generating OIDC token. The audience claim
               identifies the recipients that the JWT is intended for. The audience
               value is a single case-sensitive string. Having multiple values (array)
               for the audience field is not supported. More info about the OIDC JWT
               token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
               Note: if not specified, the Push endpoint URL will be used.
        """
        pulumi.set(__self__, "service_account_email", service_account_email)
        if audience is not None:
            pulumi.set(__self__, "audience", audience)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Input[str]:
        """
        Service account email to be used for generating the OIDC token.
        The caller (for subscriptions.create, subscriptions.patch, and
        subscriptions.modifyPushConfig RPCs) must have the
        iam.serviceAccounts.actAs permission for the service account.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter
    def audience(self) -> Optional[pulumi.Input[str]]:
        """
        Audience to be used when generating OIDC token. The audience claim
        identifies the recipients that the JWT is intended for. The audience
        value is a single case-sensitive string. Having multiple values (array)
        for the audience field is not supported. More info about the OIDC JWT
        token audience here: https://tools.ietf.org/html/rfc7519#section-4.1.3
        Note: if not specified, the Push endpoint URL will be used.
        """
        return pulumi.get(self, "audience")

    @audience.setter
    def audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "audience", value)


@pulumi.input_type
class TopicIAMBindingConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class TopicIAMMemberConditionArgs:
    def __init__(__self__, *,
                 expression: pulumi.Input[str],
                 title: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression")

    @expression.setter
    def expression(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)


@pulumi.input_type
class TopicMessageStoragePolicyArgs:
    def __init__(__self__, *,
                 allowed_persistence_regions: pulumi.Input[List[pulumi.Input[str]]]):
        """
        :param pulumi.Input[List[pulumi.Input[str]]] allowed_persistence_regions: A list of IDs of GCP regions where messages that are published to
               the topic may be persisted in storage. Messages published by
               publishers running in non-allowed GCP regions (or running outside
               of GCP altogether) will be routed for storage in one of the
               allowed regions. An empty list means that no regions are allowed,
               and is not a valid configuration.
        """
        pulumi.set(__self__, "allowed_persistence_regions", allowed_persistence_regions)

    @property
    @pulumi.getter(name="allowedPersistenceRegions")
    def allowed_persistence_regions(self) -> pulumi.Input[List[pulumi.Input[str]]]:
        """
        A list of IDs of GCP regions where messages that are published to
        the topic may be persisted in storage. Messages published by
        publishers running in non-allowed GCP regions (or running outside
        of GCP altogether) will be routed for storage in one of the
        allowed regions. An empty list means that no regions are allowed,
        and is not a valid configuration.
        """
        return pulumi.get(self, "allowed_persistence_regions")

    @allowed_persistence_regions.setter
    def allowed_persistence_regions(self, value: pulumi.Input[List[pulumi.Input[str]]]):
        pulumi.set(self, "allowed_persistence_regions", value)


