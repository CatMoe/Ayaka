package define

import "time"

type ReCaptchaResponse struct {
	Success     bool      `json:"success"`
	ChallengeTS time.Time `json:"challenge_ts"`
	Hostname    string    `json:"hostname,omitempty"`
	ErrorCodes  []string  `json:"error-codes,omitempty"`
}
