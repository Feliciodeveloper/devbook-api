package models

type Likes struct {
	UsersID uint
	PostsID uint
}

func(p *Likes)TableName()string{
	return "posts_users"
}
