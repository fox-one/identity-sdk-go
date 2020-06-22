package identity

import "time"

// ProfileResponse profile
type ProfileResponse struct {
	UserID      uint64    `json:"user_id"`
	Name        string    `json:"name,omitempty"`
	KycLevel    string    `json:"kyc_level,omitempty"`
	KycStatus   string    `json:"kyc_status,omitempty"`
	KycError    string    `json:"kyc_error,omitempty"`
	IDDigest    string    `json:"id_digest,omitempty"`
	CertifiedAt time.Time `json:"certified_at"`
}
