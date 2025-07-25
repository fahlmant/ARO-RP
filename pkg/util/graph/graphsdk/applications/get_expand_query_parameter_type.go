// Code generated by Microsoft Kiota - DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package applications

// Provides operations to manage the collection of application entities.
type GetExpandQueryParameterType int

const (
	ASTERISK_GETEXPANDQUERYPARAMETERTYPE GetExpandQueryParameterType = iota
	APPMANAGEMENTPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	CREATEDONBEHALFOF_GETEXPANDQUERYPARAMETERTYPE
	EXTENSIONPROPERTIES_GETEXPANDQUERYPARAMETERTYPE
	FEDERATEDIDENTITYCREDENTIALS_GETEXPANDQUERYPARAMETERTYPE
	HOMEREALMDISCOVERYPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	OWNERS_GETEXPANDQUERYPARAMETERTYPE
	TOKENISSUANCEPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	TOKENLIFETIMEPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	SYNCHRONIZATION_GETEXPANDQUERYPARAMETERTYPE
)

func (i GetExpandQueryParameterType) String() string {
	return []string{"*", "appManagementPolicies", "createdOnBehalfOf", "extensionProperties", "federatedIdentityCredentials", "homeRealmDiscoveryPolicies", "owners", "tokenIssuancePolicies", "tokenLifetimePolicies", "synchronization"}[i]
}
func ParseGetExpandQueryParameterType(v string) (any, error) {
	result := ASTERISK_GETEXPANDQUERYPARAMETERTYPE
	switch v {
	case "*":
		result = ASTERISK_GETEXPANDQUERYPARAMETERTYPE
	case "appManagementPolicies":
		result = APPMANAGEMENTPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	case "createdOnBehalfOf":
		result = CREATEDONBEHALFOF_GETEXPANDQUERYPARAMETERTYPE
	case "extensionProperties":
		result = EXTENSIONPROPERTIES_GETEXPANDQUERYPARAMETERTYPE
	case "federatedIdentityCredentials":
		result = FEDERATEDIDENTITYCREDENTIALS_GETEXPANDQUERYPARAMETERTYPE
	case "homeRealmDiscoveryPolicies":
		result = HOMEREALMDISCOVERYPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	case "owners":
		result = OWNERS_GETEXPANDQUERYPARAMETERTYPE
	case "tokenIssuancePolicies":
		result = TOKENISSUANCEPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	case "tokenLifetimePolicies":
		result = TOKENLIFETIMEPOLICIES_GETEXPANDQUERYPARAMETERTYPE
	case "synchronization":
		result = SYNCHRONIZATION_GETEXPANDQUERYPARAMETERTYPE
	default:
		return nil, nil
	}
	return &result, nil
}
func SerializeGetExpandQueryParameterType(values []GetExpandQueryParameterType) []string {
	result := make([]string, len(values))
	for i, v := range values {
		result[i] = v.String()
	}
	return result
}
func (i GetExpandQueryParameterType) isMultiValue() bool {
	return false
}
