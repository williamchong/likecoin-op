package model

type LikerIdInfo struct {
	DisplayName       string `json:"displayName"`
	IsEmailEnabled    bool   `json:"isEmailEnabled"`
	Locale            string `json:"locale"`
	LikeWallet        string `json:"likeWallet"`
	Description       string `json:"description"`
	Email             string `json:"email"`
	IsEmailVerified   bool   `json:"isEmailVerified"`
	NormalizedEmail   string `json:"normalizedEmail"`
	IsEmailDuplicated bool   `json:"isEmailDuplicated"`
	LastVerifyTs      int64  `json:"lastVerifyTs"`
	VerificationUUID  string `json:"verificationUUID"`
	Timestamp         int64  `json:"timestamp"`
	BonusCooldown     int64  `json:"bonusCooldown"`
	User              string `json:"user"`
	Avatar            string `json:"avatar"`
}
