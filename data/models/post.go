package models

type Post struct {
	BaseModel
	Title    string `gorm:"type:varchar(256);not null"`
	Content  string `gorm:"type:text;not null"`
	User     User   `gorm:"foreignKey:AuthorId;constraint:OnUpdate:RESTRICT,OnDelete:RESTRICT"`
	AuthorId uint   `gorm:"not null;index"`
}
