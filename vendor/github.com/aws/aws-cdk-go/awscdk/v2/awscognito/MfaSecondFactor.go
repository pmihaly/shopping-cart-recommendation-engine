package awscognito


// The different ways in which a user pool can obtain their MFA token for sign in.
//
// Example:
//   cognito.NewUserPool(this, jsii.String("myuserpool"), &UserPoolProps{
//   	// ...
//   	Mfa: cognito.Mfa_REQUIRED,
//   	MfaSecondFactor: &MfaSecondFactor{
//   		Sms: jsii.Boolean(true),
//   		Otp: jsii.Boolean(true),
//   	},
//   })
//
// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa.html
//
type MfaSecondFactor struct {
	// The MFA token is a time-based one time password that is generated by a hardware or software token.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa-totp.html
	//
	// Default: false.
	//
	Otp *bool `field:"required" json:"otp" yaml:"otp"`
	// The MFA token is sent to the user via SMS to their verified phone numbers.
	// See: https://docs.aws.amazon.com/cognito/latest/developerguide/user-pool-settings-mfa-sms-text-message.html
	//
	// Default: true.
	//
	Sms *bool `field:"required" json:"sms" yaml:"sms"`
}
