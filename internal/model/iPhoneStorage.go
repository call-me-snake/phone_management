package model

//IPhoneStorage - реализует методы Crude для структуры PhoneOwner
type IPhoneStorage interface {
	GetPhone(owner string) (*PhoneOwner, error)
	CreateOwner(data PhoneOwner) error
	UpdatePhone(data PhoneOwner) (isUpdated bool, err error)
	DeleteOwner(owner string) error
}
