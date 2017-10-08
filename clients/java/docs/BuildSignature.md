
# BuildSignature

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**publicKey** | **String** | Public key of the builder which can be used to verify that related Findings are valid and unchanged.  If &#x60;key_type&#x60; is empty this defaults to PEM encoded public keys.  This field may be empty if &#x60;key_id&#x60; references an external key.  For Cloud Container Builder based signatures this is a PEM encoded public key. To verify the Cloud Container Builder signature, place the contents of this field into a file (public.pem). The signature field is base64-decoded into its binary representation in signature.bin, and the provenance bytes from BuildDetails are base64-decoded into a binary representation in signed.bin. OpenSSL can then verify the signature: &#x60;openssl sha256 -verify public.pem -signature signature.bin signed.bin&#x60; |  [optional]
**signature** | **String** | Signature of the related BuildProvenance, encoded in a base64 string. |  [optional]
**keyId** | **String** | An ID for the key used to sign.  This could be either an ID for the key stored in &#x60;public_key&#x60; (e.g., the ID or fingerprint for a PGP key, or the CN for a cert), or a reference to an external key (e.g., a reference to a key in Cloud KMS). |  [optional]
**keyType** | [**KeyTypeEnum**](#KeyTypeEnum) | The type of the key, either stored in &#x60;public_key&#x60; or referenced in &#x60;key_id&#x60; |  [optional]


<a name="KeyTypeEnum"></a>
## Enum: KeyTypeEnum
Name | Value
---- | -----
UNSET | &quot;UNSET&quot;
PGP_ASCII_ARMORED | &quot;PGP_ASCII_ARMORED&quot;
PKIX_PEM | &quot;PKIX_PEM&quot;



