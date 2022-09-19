package config

type AUTHKey struct {
	Token            string `mapstructure:"token" json:"token" yaml:"token"`                                        // Token Key  			like: x-token
	ExpiresAt        int64  `mapstructure:"expires-at" json:"expires-at" yaml:"expires-at"`                         // Token expires Key    like: expiresAt
	RefreshToken     string `mapstructure:"refresh-token" json:"refresh-token" yaml:"refresh-token"`                // Token Key  			like: new-token
	RefreshExpiresAt string `mapstructure:"refresh-expires-at" json:"refresh-expires-at" yaml:"refresh-expires-at"` // Token Key  			like: new-expires-at
	User             string `mapstructure:"user" json:"user" yaml:"user"`                                           // User Key   			like: clims
	UserId           string `mapstructure:"user-id" json:"user-id" yaml:"user-id"`                                  // UserId Key 			like: x-user-id
}
