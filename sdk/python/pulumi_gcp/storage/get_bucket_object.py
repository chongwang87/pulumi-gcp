# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetBucketObjectResult:
    """
    A collection of values returned by getBucketObject.
    """
    def __init__(__self__, bucket=None, cache_control=None, content=None, content_disposition=None, content_encoding=None, content_language=None, content_type=None, crc32c=None, detect_md5hash=None, id=None, md5hash=None, metadata=None, name=None, output_name=None, self_link=None, source=None, storage_class=None):
        if bucket and not isinstance(bucket, str):
            raise TypeError("Expected argument 'bucket' to be a str")
        __self__.bucket = bucket
        if cache_control and not isinstance(cache_control, str):
            raise TypeError("Expected argument 'cache_control' to be a str")
        __self__.cache_control = cache_control
        """
        (Computed) [Cache-Control](https://tools.ietf.org/html/rfc7234#section-5.2)
        directive to specify caching behavior of object data. If omitted and object is accessible to all anonymous users, the default will be public, max-age=3600
        """
        if content and not isinstance(content, str):
            raise TypeError("Expected argument 'content' to be a str")
        __self__.content = content
        if content_disposition and not isinstance(content_disposition, str):
            raise TypeError("Expected argument 'content_disposition' to be a str")
        __self__.content_disposition = content_disposition
        """
        (Computed) [Content-Disposition](https://tools.ietf.org/html/rfc6266) of the object data.
        """
        if content_encoding and not isinstance(content_encoding, str):
            raise TypeError("Expected argument 'content_encoding' to be a str")
        __self__.content_encoding = content_encoding
        """
        (Computed) [Content-Encoding](https://tools.ietf.org/html/rfc7231#section-3.1.2.2) of the object data.
        """
        if content_language and not isinstance(content_language, str):
            raise TypeError("Expected argument 'content_language' to be a str")
        __self__.content_language = content_language
        """
        (Computed) [Content-Language](https://tools.ietf.org/html/rfc7231#section-3.1.3.2) of the object data.
        """
        if content_type and not isinstance(content_type, str):
            raise TypeError("Expected argument 'content_type' to be a str")
        __self__.content_type = content_type
        """
        (Computed) [Content-Type](https://tools.ietf.org/html/rfc7231#section-3.1.1.5) of the object data. Defaults to "application/octet-stream" or "text/plain; charset=utf-8".
        """
        if crc32c and not isinstance(crc32c, str):
            raise TypeError("Expected argument 'crc32c' to be a str")
        __self__.crc32c = crc32c
        """
        (Computed) Base 64 CRC32 hash of the uploaded data.
        """
        if detect_md5hash and not isinstance(detect_md5hash, str):
            raise TypeError("Expected argument 'detect_md5hash' to be a str")
        __self__.detect_md5hash = detect_md5hash
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        id is the provider-assigned unique ID for this managed resource.
        """
        if md5hash and not isinstance(md5hash, str):
            raise TypeError("Expected argument 'md5hash' to be a str")
        __self__.md5hash = md5hash
        """
        (Computed) Base 64 MD5 hash of the uploaded data.
        """
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        __self__.metadata = metadata
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if output_name and not isinstance(output_name, str):
            raise TypeError("Expected argument 'output_name' to be a str")
        __self__.output_name = output_name
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        __self__.self_link = self_link
        """
        (Computed) A url reference to this object.
        """
        if source and not isinstance(source, str):
            raise TypeError("Expected argument 'source' to be a str")
        __self__.source = source
        if storage_class and not isinstance(storage_class, str):
            raise TypeError("Expected argument 'storage_class' to be a str")
        __self__.storage_class = storage_class
        """
        (Computed) The [StorageClass](https://cloud.google.com/storage/docs/storage-classes) of the new bucket object.
        Supported values include: `MULTI_REGIONAL`, `REGIONAL`, `NEARLINE`, `COLDLINE`. If not provided, this defaults to the bucket's default
        storage class or to a [standard](https://cloud.google.com/storage/docs/storage-classes#standard) class.
        """
class AwaitableGetBucketObjectResult(GetBucketObjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBucketObjectResult(
            bucket=self.bucket,
            cache_control=self.cache_control,
            content=self.content,
            content_disposition=self.content_disposition,
            content_encoding=self.content_encoding,
            content_language=self.content_language,
            content_type=self.content_type,
            crc32c=self.crc32c,
            detect_md5hash=self.detect_md5hash,
            id=self.id,
            md5hash=self.md5hash,
            metadata=self.metadata,
            name=self.name,
            output_name=self.output_name,
            self_link=self.self_link,
            source=self.source,
            storage_class=self.storage_class)

def get_bucket_object(bucket=None,name=None,opts=None):
    """
    Gets an existing object inside an existing bucket in Google Cloud Storage service (GCS).
    See [the official documentation](https://cloud.google.com/storage/docs/key-terms#objects)
    and
    [API](https://cloud.google.com/storage/docs/json_api/v1/objects).




    > This content is derived from https://github.com/terraform-providers/terraform-provider-google/blob/master/website/docs/d/storage_bucket_object.html.markdown.


    :param str bucket: The name of the containing bucket.
    :param str name: The name of the object.
    """
    __args__ = dict()


    __args__['bucket'] = bucket
    __args__['name'] = name
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('gcp:storage/getBucketObject:getBucketObject', __args__, opts=opts).value

    return AwaitableGetBucketObjectResult(
        bucket=__ret__.get('bucket'),
        cache_control=__ret__.get('cacheControl'),
        content=__ret__.get('content'),
        content_disposition=__ret__.get('contentDisposition'),
        content_encoding=__ret__.get('contentEncoding'),
        content_language=__ret__.get('contentLanguage'),
        content_type=__ret__.get('contentType'),
        crc32c=__ret__.get('crc32c'),
        detect_md5hash=__ret__.get('detectMd5hash'),
        id=__ret__.get('id'),
        md5hash=__ret__.get('md5hash'),
        metadata=__ret__.get('metadata'),
        name=__ret__.get('name'),
        output_name=__ret__.get('outputName'),
        self_link=__ret__.get('selfLink'),
        source=__ret__.get('source'),
        storage_class=__ret__.get('storageClass'))
