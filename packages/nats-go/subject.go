package natsGo


type ListenerSubject struct {
	CreateUser string
	UpdateUser string
	PasswordResetRequest string
	BookingReminder string
}

var Subject = NewListenerSubject()

func NewListenerSubject() *ListenerSubject {
	return &ListenerSubject{
		CreateUser: "CreateUser",
		UpdateUser: "UpdateUser",
		PasswordResetRequest: "PasswordResetRequest",
		BookingReminder: "BookingReminder",
	}
}


