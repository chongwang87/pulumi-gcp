# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables
from . import outputs

__all__ = [
    'AttestorAttestationAuthorityNote',
    'AttestorAttestationAuthorityNotePublicKey',
    'AttestorAttestationAuthorityNotePublicKeyPkixPublicKey',
    'AttestorIamBindingCondition',
    'AttestorIamMemberCondition',
    'PolicyAdmissionWhitelistPattern',
    'PolicyClusterAdmissionRule',
    'PolicyDefaultAdmissionRule',
]

@pulumi.output_type
class AttestorAttestationAuthorityNote(dict):
    def __init__(__self__, *,
                 note_reference: str,
                 delegation_service_account_email: Optional[str] = None,
                 public_keys: Optional[List['outputs.AttestorAttestationAuthorityNotePublicKey']] = None):
        """
        :param str note_reference: The resource name of a ATTESTATION_AUTHORITY Note, created by the
               user. If the Note is in a different project from the Attestor, it
               should be specified in the format `projects/*/notes/*` (or the legacy
               `providers/*/notes/*`). This field may not be updated.
               An attestation by this attestor is stored as a Container Analysis
               ATTESTATION_AUTHORITY Occurrence that names a container image
               and that links to this Note.
        :param str delegation_service_account_email: -
               This field will contain the service account email address that
               this Attestor will use as the principal when querying Container
               Analysis. Attestor administrators must grant this service account
               the IAM role needed to read attestations from the noteReference in
               Container Analysis (containeranalysis.notes.occurrences.viewer).
               This email address is fixed for the lifetime of the Attestor, but
               callers should not make any other assumptions about the service
               account email; future versions may use an email based on a
               different naming pattern.
        :param List['AttestorAttestationAuthorityNotePublicKeyArgs'] public_keys: Public keys that verify attestations signed by this attestor. This
               field may be updated.
               If this field is non-empty, one of the specified public keys must
               verify that an attestation was signed by this attestor for the
               image specified in the admission request.
               If this field is empty, this attestor always returns that no valid
               attestations exist.
               Structure is documented below.
        """
        pulumi.set(__self__, "note_reference", note_reference)
        if delegation_service_account_email is not None:
            pulumi.set(__self__, "delegation_service_account_email", delegation_service_account_email)
        if public_keys is not None:
            pulumi.set(__self__, "public_keys", public_keys)

    @property
    @pulumi.getter(name="noteReference")
    def note_reference(self) -> str:
        """
        The resource name of a ATTESTATION_AUTHORITY Note, created by the
        user. If the Note is in a different project from the Attestor, it
        should be specified in the format `projects/*/notes/*` (or the legacy
        `providers/*/notes/*`). This field may not be updated.
        An attestation by this attestor is stored as a Container Analysis
        ATTESTATION_AUTHORITY Occurrence that names a container image
        and that links to this Note.
        """
        return pulumi.get(self, "note_reference")

    @property
    @pulumi.getter(name="delegationServiceAccountEmail")
    def delegation_service_account_email(self) -> Optional[str]:
        """
        -
        This field will contain the service account email address that
        this Attestor will use as the principal when querying Container
        Analysis. Attestor administrators must grant this service account
        the IAM role needed to read attestations from the noteReference in
        Container Analysis (containeranalysis.notes.occurrences.viewer).
        This email address is fixed for the lifetime of the Attestor, but
        callers should not make any other assumptions about the service
        account email; future versions may use an email based on a
        different naming pattern.
        """
        return pulumi.get(self, "delegation_service_account_email")

    @property
    @pulumi.getter(name="publicKeys")
    def public_keys(self) -> Optional[List['outputs.AttestorAttestationAuthorityNotePublicKey']]:
        """
        Public keys that verify attestations signed by this attestor. This
        field may be updated.
        If this field is non-empty, one of the specified public keys must
        verify that an attestation was signed by this attestor for the
        image specified in the admission request.
        If this field is empty, this attestor always returns that no valid
        attestations exist.
        Structure is documented below.
        """
        return pulumi.get(self, "public_keys")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AttestorAttestationAuthorityNotePublicKey(dict):
    def __init__(__self__, *,
                 ascii_armored_pgp_public_key: Optional[str] = None,
                 comment: Optional[str] = None,
                 id: Optional[str] = None,
                 pkix_public_key: Optional['outputs.AttestorAttestationAuthorityNotePublicKeyPkixPublicKey'] = None):
        """
        :param str ascii_armored_pgp_public_key: ASCII-armored representation of a PGP public key, as the
               entire output by the command
               `gpg --export --armor foo@example.com` (either LF or CRLF
               line endings). When using this field, id should be left
               blank. The BinAuthz API handlers will calculate the ID
               and fill it in automatically. BinAuthz computes this ID
               as the OpenPGP RFC4880 V4 fingerprint, represented as
               upper-case hex. If id is provided by the caller, it will
               be overwritten by the API-calculated ID.
        :param str comment: A descriptive comment. This field may be updated.
        :param str id: The ID of this public key. Signatures verified by BinAuthz
               must include the ID of the public key that can be used to
               verify them, and that ID must match the contents of this
               field exactly. Additional restrictions on this field can
               be imposed based on which public key type is encapsulated.
               See the documentation on publicKey cases below for details.
        :param 'AttestorAttestationAuthorityNotePublicKeyPkixPublicKeyArgs' pkix_public_key: A raw PKIX SubjectPublicKeyInfo format public key.
               NOTE: id may be explicitly provided by the caller when using this
               type of public key, but it MUST be a valid RFC3986 URI. If id is left
               blank, a default one will be computed based on the digest of the DER
               encoding of the public key.
               Structure is documented below.
        """
        if ascii_armored_pgp_public_key is not None:
            pulumi.set(__self__, "ascii_armored_pgp_public_key", ascii_armored_pgp_public_key)
        if comment is not None:
            pulumi.set(__self__, "comment", comment)
        if id is not None:
            pulumi.set(__self__, "id", id)
        if pkix_public_key is not None:
            pulumi.set(__self__, "pkix_public_key", pkix_public_key)

    @property
    @pulumi.getter(name="asciiArmoredPgpPublicKey")
    def ascii_armored_pgp_public_key(self) -> Optional[str]:
        """
        ASCII-armored representation of a PGP public key, as the
        entire output by the command
        `gpg --export --armor foo@example.com` (either LF or CRLF
        line endings). When using this field, id should be left
        blank. The BinAuthz API handlers will calculate the ID
        and fill it in automatically. BinAuthz computes this ID
        as the OpenPGP RFC4880 V4 fingerprint, represented as
        upper-case hex. If id is provided by the caller, it will
        be overwritten by the API-calculated ID.
        """
        return pulumi.get(self, "ascii_armored_pgp_public_key")

    @property
    @pulumi.getter
    def comment(self) -> Optional[str]:
        """
        A descriptive comment. This field may be updated.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def id(self) -> Optional[str]:
        """
        The ID of this public key. Signatures verified by BinAuthz
        must include the ID of the public key that can be used to
        verify them, and that ID must match the contents of this
        field exactly. Additional restrictions on this field can
        be imposed based on which public key type is encapsulated.
        See the documentation on publicKey cases below for details.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="pkixPublicKey")
    def pkix_public_key(self) -> Optional['outputs.AttestorAttestationAuthorityNotePublicKeyPkixPublicKey']:
        """
        A raw PKIX SubjectPublicKeyInfo format public key.
        NOTE: id may be explicitly provided by the caller when using this
        type of public key, but it MUST be a valid RFC3986 URI. If id is left
        blank, a default one will be computed based on the digest of the DER
        encoding of the public key.
        Structure is documented below.
        """
        return pulumi.get(self, "pkix_public_key")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AttestorAttestationAuthorityNotePublicKeyPkixPublicKey(dict):
    def __init__(__self__, *,
                 public_key_pem: Optional[str] = None,
                 signature_algorithm: Optional[str] = None):
        """
        :param str public_key_pem: A PEM-encoded public key, as described in
               `https://tools.ietf.org/html/rfc7468#section-13`
        :param str signature_algorithm: The signature algorithm used to verify a message against
               a signature using this key. These signature algorithm must
               match the structure and any object identifiers encoded in
               publicKeyPem (i.e. this algorithm must match that of the
               public key).
        """
        if public_key_pem is not None:
            pulumi.set(__self__, "public_key_pem", public_key_pem)
        if signature_algorithm is not None:
            pulumi.set(__self__, "signature_algorithm", signature_algorithm)

    @property
    @pulumi.getter(name="publicKeyPem")
    def public_key_pem(self) -> Optional[str]:
        """
        A PEM-encoded public key, as described in
        `https://tools.ietf.org/html/rfc7468#section-13`
        """
        return pulumi.get(self, "public_key_pem")

    @property
    @pulumi.getter(name="signatureAlgorithm")
    def signature_algorithm(self) -> Optional[str]:
        """
        The signature algorithm used to verify a message against
        a signature using this key. These signature algorithm must
        match the structure and any object identifiers encoded in
        publicKeyPem (i.e. this algorithm must match that of the
        public key).
        """
        return pulumi.get(self, "signature_algorithm")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AttestorIamBindingCondition(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class AttestorIamMemberCondition(dict):
    def __init__(__self__, *,
                 expression: str,
                 title: str,
                 description: Optional[str] = None):
        pulumi.set(__self__, "expression", expression)
        pulumi.set(__self__, "title", title)
        if description is not None:
            pulumi.set(__self__, "description", description)

    @property
    @pulumi.getter
    def expression(self) -> str:
        return pulumi.get(self, "expression")

    @property
    @pulumi.getter
    def title(self) -> str:
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyAdmissionWhitelistPattern(dict):
    def __init__(__self__, *,
                 name_pattern: str):
        """
        :param str name_pattern: An image name pattern to whitelist, in the form
               `registry/path/to/image`. This supports a trailing * as a
               wildcard, but this is allowed only in text after the registry/
               part.
        """
        pulumi.set(__self__, "name_pattern", name_pattern)

    @property
    @pulumi.getter(name="namePattern")
    def name_pattern(self) -> str:
        """
        An image name pattern to whitelist, in the form
        `registry/path/to/image`. This supports a trailing * as a
        wildcard, but this is allowed only in text after the registry/
        part.
        """
        return pulumi.get(self, "name_pattern")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyClusterAdmissionRule(dict):
    def __init__(__self__, *,
                 cluster: str,
                 enforcement_mode: str,
                 evaluation_mode: str,
                 require_attestations_bies: Optional[List[str]] = None):
        """
        :param str cluster: The identifier for this object. Format specified above.
        :param str enforcement_mode: The action when a pod creation is denied by the admission rule.
               Possible values are `ENFORCED_BLOCK_AND_AUDIT_LOG` and `DRYRUN_AUDIT_LOG_ONLY`.
        :param str evaluation_mode: How this admission rule will be evaluated.
               Possible values are `ALWAYS_ALLOW`, `REQUIRE_ATTESTATION`, and `ALWAYS_DENY`.
        :param List[str] require_attestations_bies: The resource names of the attestors that must attest to a
               container image. If the attestor is in a different project from the
               policy, it should be specified in the format `projects/*/attestors/*`.
               Each attestor must exist before a policy can reference it. To add an
               attestor to a policy the principal issuing the policy change
               request must be able to read the attestor resource.
               Note: this field must be non-empty when the evaluation_mode field
               specifies REQUIRE_ATTESTATION, otherwise it must be empty.
        """
        pulumi.set(__self__, "cluster", cluster)
        pulumi.set(__self__, "enforcement_mode", enforcement_mode)
        pulumi.set(__self__, "evaluation_mode", evaluation_mode)
        if require_attestations_bies is not None:
            pulumi.set(__self__, "require_attestations_bies", require_attestations_bies)

    @property
    @pulumi.getter
    def cluster(self) -> str:
        """
        The identifier for this object. Format specified above.
        """
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter(name="enforcementMode")
    def enforcement_mode(self) -> str:
        """
        The action when a pod creation is denied by the admission rule.
        Possible values are `ENFORCED_BLOCK_AND_AUDIT_LOG` and `DRYRUN_AUDIT_LOG_ONLY`.
        """
        return pulumi.get(self, "enforcement_mode")

    @property
    @pulumi.getter(name="evaluationMode")
    def evaluation_mode(self) -> str:
        """
        How this admission rule will be evaluated.
        Possible values are `ALWAYS_ALLOW`, `REQUIRE_ATTESTATION`, and `ALWAYS_DENY`.
        """
        return pulumi.get(self, "evaluation_mode")

    @property
    @pulumi.getter(name="requireAttestationsBies")
    def require_attestations_bies(self) -> Optional[List[str]]:
        """
        The resource names of the attestors that must attest to a
        container image. If the attestor is in a different project from the
        policy, it should be specified in the format `projects/*/attestors/*`.
        Each attestor must exist before a policy can reference it. To add an
        attestor to a policy the principal issuing the policy change
        request must be able to read the attestor resource.
        Note: this field must be non-empty when the evaluation_mode field
        specifies REQUIRE_ATTESTATION, otherwise it must be empty.
        """
        return pulumi.get(self, "require_attestations_bies")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class PolicyDefaultAdmissionRule(dict):
    def __init__(__self__, *,
                 enforcement_mode: str,
                 evaluation_mode: str,
                 require_attestations_bies: Optional[List[str]] = None):
        """
        :param str enforcement_mode: The action when a pod creation is denied by the admission rule.
               Possible values are `ENFORCED_BLOCK_AND_AUDIT_LOG` and `DRYRUN_AUDIT_LOG_ONLY`.
        :param str evaluation_mode: How this admission rule will be evaluated.
               Possible values are `ALWAYS_ALLOW`, `REQUIRE_ATTESTATION`, and `ALWAYS_DENY`.
        :param List[str] require_attestations_bies: The resource names of the attestors that must attest to a
               container image. If the attestor is in a different project from the
               policy, it should be specified in the format `projects/*/attestors/*`.
               Each attestor must exist before a policy can reference it. To add an
               attestor to a policy the principal issuing the policy change
               request must be able to read the attestor resource.
               Note: this field must be non-empty when the evaluation_mode field
               specifies REQUIRE_ATTESTATION, otherwise it must be empty.
        """
        pulumi.set(__self__, "enforcement_mode", enforcement_mode)
        pulumi.set(__self__, "evaluation_mode", evaluation_mode)
        if require_attestations_bies is not None:
            pulumi.set(__self__, "require_attestations_bies", require_attestations_bies)

    @property
    @pulumi.getter(name="enforcementMode")
    def enforcement_mode(self) -> str:
        """
        The action when a pod creation is denied by the admission rule.
        Possible values are `ENFORCED_BLOCK_AND_AUDIT_LOG` and `DRYRUN_AUDIT_LOG_ONLY`.
        """
        return pulumi.get(self, "enforcement_mode")

    @property
    @pulumi.getter(name="evaluationMode")
    def evaluation_mode(self) -> str:
        """
        How this admission rule will be evaluated.
        Possible values are `ALWAYS_ALLOW`, `REQUIRE_ATTESTATION`, and `ALWAYS_DENY`.
        """
        return pulumi.get(self, "evaluation_mode")

    @property
    @pulumi.getter(name="requireAttestationsBies")
    def require_attestations_bies(self) -> Optional[List[str]]:
        """
        The resource names of the attestors that must attest to a
        container image. If the attestor is in a different project from the
        policy, it should be specified in the format `projects/*/attestors/*`.
        Each attestor must exist before a policy can reference it. To add an
        attestor to a policy the principal issuing the policy change
        request must be able to read the attestor resource.
        Note: this field must be non-empty when the evaluation_mode field
        specifies REQUIRE_ATTESTATION, otherwise it must be empty.
        """
        return pulumi.get(self, "require_attestations_bies")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


