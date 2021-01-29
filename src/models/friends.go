package models

type Friends struct {
	UsersID uint
	FriendID uint
}
func(f *Friends)TableName()string{
	return "users_friends"
}
