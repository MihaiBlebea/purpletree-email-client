package ec

// Email types
const (
	WelcomeEmail    string = "welcome"
	ConfirmEmail    string = "confirm-email"
	ConfirmPassword string = "confirm-password"
	Contact         string = "contact"
	AdminContact    string = "admin-contact"
	Payment         string = "payment"
	AdminPayment    string = "admin-payment"
)

// Subject template constants
const (
	WelcomeSubject         string = "Welcome %s, download link inside"
	ConfirmEmailSubject    string = "One last step, confirm your email %s"
	ConfirmPasswordSubject string = "Reset password, confirm your email"
	ContactSubject         string = "We received your message"
	ContactAdminSubject    string = "Somebody is looking for you"
	PaymentSubject         string = "Thank you %s, here are your details"
	AdminPaymentSubject    string = "Somebody bought a product"
)
