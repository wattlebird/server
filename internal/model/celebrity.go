package model

type Celebrity struct {
	Category       CelebrityCat
	Name           string
	Image          string
	Infobox        string
	Summary        string
	ID             uint32
	Redirect       uint32
	CollectCount   uint32
	CommentCount   uint32
	FieldBirthYear uint16
	Producer       bool
	Type           uint8
	Artist         bool
	Seiyu          bool
	Writer         bool
	Illustrator    bool
	Actor          bool
	FieldBloodType uint8
	FieldGender    uint8
	FieldBirthMon  uint8
	Locked         bool
	FieldBirthDay  uint8
}
