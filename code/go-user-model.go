type User struct {
	gorm.Model
// se defineste cheia primara a tabelei ca tip uuid
	ID                      uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid"`
// se pot specifica proprietatii de unicitate sau de nonnullable
	Email                   string     `json:"email" gorm:"unique;not null"`
	FirstName               string     `json:"firstName" gorm:"not null;"`
	LastName                string     `json:"lastName" gorm:"not null;"`
	Password                string     `json:"password" gorm:"not null;"`
	Salt                    string     `json:"salt" gorm:"not null;"`
	S3ID                    string     `json:"s3Id" gorm:"unique"`
// set pot specifica si valori implicite pentru coloane
	S3MaxNumberOfMediaFiles int32      `json:"s3MaxNumberOfMediaFiles" gorm:"not null;default:1000"`
	Verified                bool       `json:"verified" gorm:"not null;default:false;"`
	LoginCanCheck2FA        bool       `json:"passLogin" gorm:"not null;default:false;"`
	O2FARequestedAt         *time.Time `json:"o2FARequestedAt" gorm:"default:null;"`
	CreatedAt               *time.Time `json:"createdAt"`
	UpdatedAt               *time.Time `json:"updatedAt"`
	DeletedAt               *time.Time `json:"deletedAt" gorm:"default:null"`
// in model se pot specifica relatiile dintre tabelul(modelul) curent si alte tabele
	MarketDeck       []MarketDeck       `json:"marketDeck" gorm:"foreignKey:UserID;references:ID;constraint: OnUpdate:CASCADE,OnDelete:CASCADE"`
	MarketDeckReview []MarketDeckReview `json:"marketDeckReview" gorm:"foreignKey:UserID;references:ID;constraint: OnUpdate:CASCADE,OnDelete:CASCADE"`
}