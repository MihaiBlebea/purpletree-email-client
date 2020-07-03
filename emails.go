package ec

// Email types
const (
	WelcomeEmail    string = "welcome"
	ConfirmEmail    string = "confirm-email"
	ConfirmPassword string = "confirm-password"
	ContactEmail    string = "contact"
	ContactPayment  string = "payment"
)

// Subject template constants
const (
	WelcomeSubject         string = "Welcome %s, download link inside"
	ConfirmEmailSubject    string = "One last step, confirm your email %s"
	ConfirmPasswordSubject string = "Reset password, confirm your email"
	ContactEmailSubject    string = "Somebody is looking for you"
	ContactPaymentSubject  string = "Thank you %s, here are your details"
)
