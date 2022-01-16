package entity

type Book struct {
	ID         uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Title      string `gorm:"type:varchar(2355)" json:"title"`
	Decription string `gorm:"type:varchar(255)" json:"decription"`
	UserID     uint64 `gorm:"not null" json:"-"`
	User       User   `gorm:"foreinkey:UserID;constraint:onUpdate;CASCADE,onDelete:CASCADE" json:"user"`
}
