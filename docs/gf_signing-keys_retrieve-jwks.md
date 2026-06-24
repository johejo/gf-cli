## gf signing-keys retrieve-jwks

Gets JSON web key set j w k s with all the keys that can be used to verify tokens public keys

### Synopsis

Gets JSON web key set j w k s with all the keys that can be used to verify tokens public keys

Required permissions None

Response schema (RetrieveJWKSOK.Payload):
  keys                                                        array<object>
  keys[].Algorithm                                            string                Key algorithm, parsed from `alg` header.
  keys[].CertificateThumbprintSHA1                            array<number>         X.509 certificate thumbprint (SHA-1), parsed from `x5t` header.
  keys[].CertificateThumbprintSHA256                          array<number>         X.509 certificate thumbprint (SHA-256), parsed from `x5t#S256` header.
  keys[].Certificates                                         array<object>         X.509 certificate chain, parsed from `x5c` header.
  keys[].Certificates[].AuthorityKeyId                        array<number>
  keys[].Certificates[].BasicConstraintsValid                 boolean               BasicConstraintsValid indicates whether IsCA, MaxPathLen,
                                                                                    and MaxPathLenZero are valid.
  keys[].Certificates[].CRLDistributionPoints                 array<string>
  keys[].Certificates[].DNSNames                              array<string>         Subject Alternate Name values. (Note that these values may not be valid
                                                                                    if invalid values were contained within a parsed certificate. For
                                                                                    example, an element of DNSNames may not be a valid DNS domain name.)
  keys[].Certificates[].EmailAddresses                        array<string>
  keys[].Certificates[].ExcludedDNSDomains                    array<string>
  keys[].Certificates[].ExcludedEmailAddresses                array<string>
  keys[].Certificates[].ExcludedIPRanges                      array<object>
  keys[].Certificates[].ExcludedIPRanges[].IP                 string
  keys[].Certificates[].ExcludedIPRanges[].Mask               array<number>
  keys[].Certificates[].ExcludedURIDomains                    array<string>
  keys[].Certificates[].ExtKeyUsage                           array<number>
  keys[].Certificates[].Extensions                            array<object>         Extensions contains raw X.509 extensions. When parsing certificates,
                                                                                    this can be used to extract non-critical extensions that are not
                                                                                    parsed by this package. When marshaling certificates, the Extensions
                                                                                    field is ignored, see ExtraExtensions.
  keys[].Certificates[].Extensions[].Critical                 boolean
  keys[].Certificates[].Extensions[].Id                       array<number>
  keys[].Certificates[].Extensions[].Value                    array<number>
  keys[].Certificates[].ExtraExtensions                       array<object>         ExtraExtensions contains extensions to be copied, raw, into any
                                                                                    marshaled certificates. Values override any extensions that would
                                                                                    otherwise be produced based on the other fields. The ExtraExtensions
                                                                                    field is not populated when parsing certificates, see Extensions.
  keys[].Certificates[].ExtraExtensions[].Critical            boolean
  keys[].Certificates[].ExtraExtensions[].Id                  array<number>
  keys[].Certificates[].ExtraExtensions[].Value               array<number>
  keys[].Certificates[].IPAddresses                           array<string>
  keys[].Certificates[].InhibitAnyPolicy                      number                InhibitAnyPolicy and InhibitAnyPolicyZero indicate the presence and value
                                                                                    of the inhibitAnyPolicy extension.
                                                                                    The value of InhibitAnyPolicy indicates the number of additional
                                                                                    certificates in the path after this certificate that may use the
                                                                                    anyPolicy policy OID to indicate a match with any other policy.
                                                                                    When parsing a certificate, a positive non-zero InhibitAnyPolicy means
                                                                                    that the field was specified, -1 means it was unset, and
                                                                                    InhibitAnyPolicyZero being true mean that the field was explicitly set to
                                                                                    zero. The case of InhibitAnyPolicy==0 with InhibitAnyPolicyZero==false
                                                                                    should be treated equivalent to -1 (unset).
  keys[].Certificates[].InhibitAnyPolicyZero                  boolean               InhibitAnyPolicyZero indicates that InhibitAnyPolicy==0 should be
                                                                                    interpreted as an actual maximum path length of zero. Otherwise, that
                                                                                    combination is interpreted as InhibitAnyPolicy not being set.
  keys[].Certificates[].InhibitPolicyMapping                  number                InhibitPolicyMapping and InhibitPolicyMappingZero indicate the presence
                                                                                    and value of the inhibitPolicyMapping field of the policyConstraints
                                                                                    extension.
                                                                                    The value of InhibitPolicyMapping indicates the number of additional
                                                                                    certificates in the path after this certificate that may use policy
                                                                                    mapping.
                                                                                    When parsing a certificate, a positive non-zero InhibitPolicyMapping
                                                                                    means that the field was specified, -1 means it was unset, and
                                                                                    InhibitPolicyMappingZero being true mean that the field was explicitly
                                                                                    set to zero. The case of InhibitPolicyMapping==0 with
                                                                                    InhibitPolicyMappingZero==false should be treated equivalent to -1
                                                                                    (unset).
  keys[].Certificates[].InhibitPolicyMappingZero              boolean               InhibitPolicyMappingZero indicates that InhibitPolicyMapping==0 should be
                                                                                    interpreted as an actual maximum path length of zero. Otherwise, that
                                                                                    combination is interpreted as InhibitAnyPolicy not being set.
  keys[].Certificates[].IsCA                                  boolean
  keys[].Certificates[].Issuer                                object
  keys[].Certificates[].Issuer.Country                        array<string>
  keys[].Certificates[].Issuer.ExtraNames                     array<object>         ExtraNames contains attributes to be copied, raw, into any marshaled
                                                                                    distinguished names. Values override any attributes with the same OID.
                                                                                    The ExtraNames field is not populated when parsing, see Names.
  keys[].Certificates[].Issuer.ExtraNames[].Type              array<number>
  keys[].Certificates[].Issuer.ExtraNames[].Value             object
  keys[].Certificates[].Issuer.Locality                       array<string>
  keys[].Certificates[].Issuer.Names                          array<object>         Names contains all parsed attributes. When parsing distinguished names,
                                                                                    this can be used to extract non-standard attributes that are not parsed
                                                                                    by this package. When marshaling to RDNSequences, the Names field is
                                                                                    ignored, see ExtraNames.
  keys[].Certificates[].Issuer.Names[].Type                   array<number>
  keys[].Certificates[].Issuer.Names[].Value                  object
  keys[].Certificates[].Issuer.SerialNumber                   string
  keys[].Certificates[].Issuer.StreetAddress                  array<string>
  keys[].Certificates[].IssuingCertificateURL                 array<string>
  keys[].Certificates[].KeyUsage                              number
  keys[].Certificates[].MaxPathLen                            number                MaxPathLen and MaxPathLenZero indicate the presence and
                                                                                    value of the BasicConstraints' "pathLenConstraint".
                                                                                    When parsing a certificate, a positive non-zero MaxPathLen
                                                                                    means that the field was specified, -1 means it was unset,
                                                                                    and MaxPathLenZero being true mean that the field was
                                                                                    explicitly set to zero. The case of MaxPathLen==0 with MaxPathLenZero==false
                                                                                    should be treated equivalent to -1 (unset).
                                                                                    When generating a certificate, an unset pathLenConstraint
                                                                                    can be requested with either MaxPathLen == -1 or using the
                                                                                    zero value for both MaxPathLen and MaxPathLenZero.
  keys[].Certificates[].MaxPathLenZero                        boolean               MaxPathLenZero indicates that BasicConstraintsValid==true
                                                                                    and MaxPathLen==0 should be interpreted as an actual
                                                                                    maximum path length of zero. Otherwise, that combination is
                                                                                    interpreted as MaxPathLen not being set.
  keys[].Certificates[].NotBefore                             string
  keys[].Certificates[].OCSPServer                            array<string>         RFC 5280, 4.2.2.1 (Authority Information Access)
  keys[].Certificates[].PermittedDNSDomains                   array<string>
  keys[].Certificates[].PermittedDNSDomainsCritical           boolean               Name constraints
  keys[].Certificates[].PermittedEmailAddresses               array<string>
  keys[].Certificates[].PermittedIPRanges                     array<object>
  keys[].Certificates[].PermittedIPRanges[].IP                string
  keys[].Certificates[].PermittedIPRanges[].Mask              array<number>
  keys[].Certificates[].PermittedURIDomains                   array<string>
  keys[].Certificates[].Policies                              array<string>         Policies contains all policy identifiers included in the certificate.
                                                                                    See CreateCertificate for context about how this field and the PolicyIdentifiers field
                                                                                    interact.
                                                                                    In Go 1.22, encoding/gob cannot handle and ignores this field.
  keys[].Certificates[].PolicyIdentifiers                     array<array<number>>  PolicyIdentifiers contains asn1.ObjectIdentifiers, the components
                                                                                    of which are limited to int32. If a certificate contains a policy which
                                                                                    cannot be represented by asn1.ObjectIdentifier, it will not be included in
                                                                                    PolicyIdentifiers, but will be present in Policies, which contains all parsed
                                                                                    policy OIDs.
                                                                                    See CreateCertificate for context about how this field and the Policies field
                                                                                    interact.
  keys[].Certificates[].PolicyMappings                        array<object>         PolicyMappings contains a list of policy mappings included in the certificate.
  keys[].Certificates[].PolicyMappings[].IssuerDomainPolicy   string                IssuerDomainPolicy contains a policy OID the issuing certificate considers
                                                                                    equivalent to SubjectDomainPolicy in the subject certificate.
  keys[].Certificates[].PolicyMappings[].SubjectDomainPolicy  string                SubjectDomainPolicy contains a OID the issuing certificate considers
                                                                                    equivalent to IssuerDomainPolicy in the subject certificate.
  keys[].Certificates[].PublicKey                             object
  keys[].Certificates[].PublicKeyAlgorithm                    number
  keys[].Certificates[].Raw                                   array<number>
  keys[].Certificates[].RawIssuer                             array<number>
  keys[].Certificates[].RawSubject                            array<number>
  keys[].Certificates[].RawSubjectPublicKeyInfo               array<number>
  keys[].Certificates[].RawTBSCertificate                     array<number>
  keys[].Certificates[].RequireExplicitPolicy                 number                RequireExplicitPolicy and RequireExplicitPolicyZero indicate the presence
                                                                                    and value of the requireExplicitPolicy field of the policyConstraints
                                                                                    extension.
                                                                                    The value of RequireExplicitPolicy indicates the number of additional
                                                                                    certificates in the path after this certificate before an explicit policy
                                                                                    is required for the rest of the path. When an explicit policy is required,
                                                                                    each subsequent certificate in the path must contain a required policy OID,
                                                                                    or a policy OID which has been declared as equivalent through the policy
                                                                                    mapping extension.
                                                                                    When parsing a certificate, a positive non-zero RequireExplicitPolicy
                                                                                    means that the field was specified, -1 means it was unset, and
                                                                                    RequireExplicitPolicyZero being true mean that the field was explicitly
                                                                                    set to zero. The case of RequireExplicitPolicy==0 with
                                                                                    RequireExplicitPolicyZero==false should be treated equivalent to -1
                                                                                    (unset).
  keys[].Certificates[].RequireExplicitPolicyZero             boolean               RequireExplicitPolicyZero indicates that RequireExplicitPolicy==0 should be
                                                                                    interpreted as an actual maximum path length of zero. Otherwise, that
                                                                                    combination is interpreted as InhibitAnyPolicy not being set.
  keys[].Certificates[].SerialNumber                          string
  keys[].Certificates[].Signature                             array<number>
  keys[].Certificates[].SignatureAlgorithm                    number
  keys[].Certificates[].Subject                               object
  keys[].Certificates[].Subject.Country                       array<string>
  keys[].Certificates[].Subject.ExtraNames                    array<object>         ExtraNames contains attributes to be copied, raw, into any marshaled
                                                                                    distinguished names. Values override any attributes with the same OID.
                                                                                    The ExtraNames field is not populated when parsing, see Names.
  keys[].Certificates[].Subject.ExtraNames[].Type             array<number>
  keys[].Certificates[].Subject.ExtraNames[].Value            object
  keys[].Certificates[].Subject.Locality                      array<string>
  keys[].Certificates[].Subject.Names                         array<object>         Names contains all parsed attributes. When parsing distinguished names,
                                                                                    this can be used to extract non-standard attributes that are not parsed
                                                                                    by this package. When marshaling to RDNSequences, the Names field is
                                                                                    ignored, see ExtraNames.
  keys[].Certificates[].Subject.Names[].Type                  array<number>
  keys[].Certificates[].Subject.Names[].Value                 object
  keys[].Certificates[].Subject.SerialNumber                  string
  keys[].Certificates[].Subject.StreetAddress                 array<string>
  keys[].Certificates[].SubjectKeyId                          array<number>
  keys[].Certificates[].URIs                                  array<string>
  keys[].Certificates[].UnhandledCriticalExtensions           array<array<number>>  UnhandledCriticalExtensions contains a list of extension IDs that
                                                                                    were not (fully) processed when parsing. Verify will fail if this
                                                                                    slice is non-empty, unless verification is delegated to an OS
                                                                                    library which understands all the critical extensions.
                                                                                    Users can access these extensions using Extensions and can remove
                                                                                    elements from this slice if they believe that they have been
                                                                                    handled.
  keys[].Certificates[].UnknownExtKeyUsage                    array<array<number>>
  keys[].Certificates[].Version                               number
  keys[].CertificatesURL                                      string
  keys[].Key                                                  object                Key is the Go in-memory representation of this key. It must have one
                                                                                    of these types:
                                                                                    ed25519.PublicKey
                                                                                    ed25519.PrivateKey
                                                                                    ecdsa.PublicKey
                                                                                    ecdsa.PrivateKey
                                                                                    rsa.PublicKey
                                                                                    rsa.PrivateKey
                                                                                    []byte (a symmetric key)
                                                                                    When marshaling this JSONWebKey into JSON, the "kty" header parameter
                                                                                    will be automatically set based on the type of this field.
  keys[].KeyID                                                string                Key identifier, parsed from `kid` header.
  keys[].Use                                                  string                Key use, parsed from `use` header.

```
gf signing-keys retrieve-jwks [flags]
```

### Options

```
      --describe-response-jsonschema   Print the JSON Schema of the response payload and exit without calling the API
  -h, --help                           help for retrieve-jwks
      --raw                            Print the raw HTTP response body instead of the decoded payload
```

### Options inherited from parent commands

```
      --api-key string               API Key to authenticate to grafana server (env: GF_API_KEY)
      --base-path string             Base path for server: useful when using server behind reverse proxy (env: GF_BASE_PATH) (default "/api")
      --basic-user-password string   Basic authentication password (env: GF_BASIC_AUTH_PASSWORD)
      --basic-user-username string   Basic authentication username (env: GF_BASIC_AUTH_USERNAME)
      --debug                        Enable debug logging (env: GF_DEBUG)
      --host string                  Grafana server host (env: GF_HOST) (default "localhost:3000")
      --org-id int                   Organization ID (env: GF_ORG_ID)
      --timeout duration             Timeout for the HTTP request to the Grafana server; 0 disables it (env: GF_TIMEOUT) (default 30s)
```

### SEE ALSO

* [gf signing-keys](gf_signing-keys.md)	 - Signing keys API

