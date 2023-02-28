/*
 * grafeas.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas
// JustificationJustificationType : Provides the type of justification.   - JUSTIFICATION_TYPE_UNSPECIFIED: JUSTIFICATION_TYPE_UNSPECIFIED.  - COMPONENT_NOT_PRESENT: The vulnerable component is not present in the product.  - VULNERABLE_CODE_NOT_PRESENT: The vulnerable code is not present. Typically this case occurs when source code is configured or built in a way that excludes the vulnerable code.  - VULNERABLE_CODE_NOT_IN_EXECUTE_PATH: The vulnerable code can not be executed. Typically this case occurs when the product includes the vulnerable code but does not call or use the vulnerable code.  - VULNERABLE_CODE_CANNOT_BE_CONTROLLED_BY_ADVERSARY: The vulnerable code cannot be controlled by an attacker to exploit the vulnerability.  - INLINE_MITIGATIONS_ALREADY_EXIST: The product includes built-in protections or features that prevent exploitation of the vulnerability. These built-in protections cannot be subverted by the attacker and cannot be configured or disabled by the user. These mitigations completely prevent exploitation based on known attack vectors.
type JustificationJustificationType string

// List of JustificationJustificationType
const (
	JUSTIFICATION_TYPE_UNSPECIFIED_JustificationJustificationType JustificationJustificationType = "JUSTIFICATION_TYPE_UNSPECIFIED"
	COMPONENT_NOT_PRESENT_JustificationJustificationType JustificationJustificationType = "COMPONENT_NOT_PRESENT"
	VULNERABLE_CODE_NOT_PRESENT_JustificationJustificationType JustificationJustificationType = "VULNERABLE_CODE_NOT_PRESENT"
	VULNERABLE_CODE_NOT_IN_EXECUTE_PATH_JustificationJustificationType JustificationJustificationType = "VULNERABLE_CODE_NOT_IN_EXECUTE_PATH"
	VULNERABLE_CODE_CANNOT_BE_CONTROLLED_BY_ADVERSARY_JustificationJustificationType JustificationJustificationType = "VULNERABLE_CODE_CANNOT_BE_CONTROLLED_BY_ADVERSARY"
	INLINE_MITIGATIONS_ALREADY_EXIST_JustificationJustificationType JustificationJustificationType = "INLINE_MITIGATIONS_ALREADY_EXIST"
)