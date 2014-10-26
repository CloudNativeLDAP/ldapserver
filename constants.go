package ldapserver

// filter substrings expressions
const (
	FilterSubstringsInitial = 0
	FilterSubstringsAny     = 1
	FilterSubstringsFinal   = 2
)

// Filter expressions
const (
	FilterAnd             = 0
	FilterOr              = 1
	FilterNot             = 2
	FilterEqualityMatch   = 3
	FilterSubstrings      = 4
	FilterGreaterOrEqual  = 5
	FilterLessOrEqual     = 6
	FilterPresent         = 7
	FilterApproxMatch     = 8
	FilterExtensibleMatch = 9
)

// LDAP Application Codes
const (
	ApplicationBindRequest           = 0
	ApplicationBindResponse          = 1
	ApplicationUnbindRequest         = 2
	ApplicationSearchRequest         = 3
	ApplicationSearchResultEntry     = 4
	ApplicationSearchResultDone      = 5
	ApplicationModifyRequest         = 6
	ApplicationModifyResponse        = 7
	ApplicationAddRequest            = 8
	ApplicationAddResponse           = 9
	ApplicationDelRequest            = 10
	ApplicationDelResponse           = 11
	ApplicationModifyDNRequest       = 12
	ApplicationModifyDNResponse      = 13
	ApplicationCompareRequest        = 14
	ApplicationCompareResponse       = 15
	ApplicationAbandonRequest        = 16
	ApplicationSearchResultReference = 19
	ApplicationExtendedRequest       = 23
	ApplicationExtendedResponse      = 24
)

// LDAP Result Codes
const (
	LDAPResultSuccess                      = 0
	LDAPResultOperationsError              = 1
	LDAPResultProtocolError                = 2
	LDAPResultTimeLimitExceeded            = 3
	LDAPResultSizeLimitExceeded            = 4
	LDAPResultCompareFalse                 = 5
	LDAPResultCompareTrue                  = 6
	LDAPResultAuthMethodNotSupported       = 7
	LDAPResultStrongAuthRequired           = 8
	LDAPResultReferral                     = 10
	LDAPResultAdminLimitExceeded           = 11
	LDAPResultUnavailableCriticalExtension = 12
	LDAPResultConfidentialityRequired      = 13
	LDAPResultSaslBindInProgress           = 14
	LDAPResultNoSuchAttribute              = 16
	LDAPResultUndefinedAttributeType       = 17
	LDAPResultInappropriateMatching        = 18
	LDAPResultConstraintViolation          = 19
	LDAPResultAttributeOrValueExists       = 20
	LDAPResultInvalidAttributeSyntax       = 21
	LDAPResultNoSuchObject                 = 32
	LDAPResultAliasProblem                 = 33
	LDAPResultInvalidDNSyntax              = 34
	LDAPResultAliasDereferencingProblem    = 36
	LDAPResultInappropriateAuthentication  = 48
	LDAPResultInvalidCredentials           = 49
	LDAPResultInsufficientAccessRights     = 50
	LDAPResultBusy                         = 51
	LDAPResultUnavailable                  = 52
	LDAPResultUnwillingToPerform           = 53
	LDAPResultLoopDetect                   = 54
	LDAPResultNamingViolation              = 64
	LDAPResultObjectClassViolation         = 65
	LDAPResultNotAllowedOnNonLeaf          = 66
	LDAPResultNotAllowedOnRDN              = 67
	LDAPResultEntryAlreadyExists           = 68
	LDAPResultObjectClassModsProhibited    = 69
	LDAPResultAffectsMultipleDSAs          = 71
	LDAPResultOther                        = 80

	ErrorNetwork         = 200
	ErrorFilterCompile   = 201
	ErrorFilterDecompile = 202
	ErrorDebugging       = 203
)

// LDAPOID is a notational convenience to indicate that the
// permitted value of this string is a (UTF-8 encoded) dotted-decimal
// representation of an OBJECT IDENTIFIER.  Although an LDAPOID is
type LDAPOID string

// Extended operation responseName and requestName
const (
	NoticeOfDisconnection LDAPOID = "1.3.6.1.4.1.1466.2003"
)

// ###################################################
// ################# LDAP TYPEs
// ###################################################

// LDAPDN is defined to be the representation of a Distinguished Name
// (DN) after encoding according to the specification
type LDAPDN string

// respectively, and have the same single octet UTF-8 encoding.  Other
// Unicode characters have a multiple octet UTF-8 encoding.
type LDAPString string

//
//        Attribute ::= PartialAttribute(WITH COMPONENTS {
//             ...,
//             vals (SIZE(1..MAX))})
type Attribute PartialAttribute

func (p *Attribute) GetDescription() AttributeDescription {
	return p.type_
}
func (p *Attribute) GetValues() []AttributeValue {
	return p.vals
}

//
//        PartialAttribute ::= SEQUENCE {
//             type       AttributeDescription,
//             vals       SET OF value AttributeValue }
type PartialAttribute struct {
	type_ AttributeDescription
	vals  []AttributeValue
}

func (p *PartialAttribute) GetDescription() AttributeDescription {
	return p.type_
}
func (p *PartialAttribute) GetValues() []AttributeValue {
	return p.vals
}

//
//        AttributeDescription ::= LDAPString
//                                -- Constrained to <attributedescription>
//                                -- [RFC4512]
type AttributeDescription LDAPString

//        AttributeValue ::= OCTET STRING
type AttributeValue OCTETSTRING

type OCTETSTRING string

//
//        AttributeList ::= SEQUENCE OF attribute Attribute
type AttributeList []Attribute
