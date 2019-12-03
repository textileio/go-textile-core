package jwt

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	crypto "github.com/libp2p/go-libp2p-core/crypto"
	peer "github.com/libp2p/go-libp2p-core/peer"
	protocol "github.com/libp2p/go-libp2p-core/protocol"
)

var ErrClaimsInvalid = fmt.Errorf("claims invalid")
var ErrNoToken = fmt.Errorf("no token found")
var ErrExpired = fmt.Errorf("token expired")
var ErrInvalid = fmt.Errorf("token invalid")

type Claims struct {
	Scope Scope `json:"scopes"`
	jwt.StandardClaims
}

type Scope string

const (
	Access  Scope = "access"
	Refresh Scope = "refresh"
)

type Session struct {
	ID      string
	Access  string
	Exp     time.Time
	Refresh string
	Rexp    time.Time
	Subject string
	Type    string
}

func NewSession(sk crypto.PrivKey, pid peer.ID, proto protocol.ID, duration time.Duration) (*Session, error) {
	issuer, err := peer.IDFromPrivateKey(sk)
	if err != nil {
		return nil, err
	}
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	// build access token
	now := time.Now()
	exp := now.Add(duration)
	claims := &Claims{
		Scope: Access,
		StandardClaims: jwt.StandardClaims{
			Audience:  string(proto),
			ExpiresAt: exp.Unix(),
			Id:        id.String(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer.Pretty(),
			Subject:   pid.Pretty(),
		},
	}
	access, err := jwt.NewWithClaims(SigningMethodEd25519i, claims).SignedString(sk)
	if err != nil {
		return nil, err
	}

	// build refresh token
	rexp := now.Add(duration * 2)
	rclaims := &Claims{
		Scope: Refresh,
		StandardClaims: jwt.StandardClaims{
			Audience:  string(proto),
			ExpiresAt: rexp.Unix(),
			Id:        "r" + id.String(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    issuer.Pretty(),
			Subject:   pid.Pretty(),
		},
	}
	refresh, err := jwt.NewWithClaims(SigningMethodEd25519i, rclaims).SignedString(sk)
	if err != nil {
		return nil, err
	}

	// build session
	return &Session{
		ID:      issuer.Pretty(),
		Access:  access,
		Exp:     exp,
		Refresh: refresh,
		Rexp:    rexp,
		Subject: pid.Pretty(),
		Type:    "JWT",
	}, nil
}

func ParseClaims(claims jwt.Claims) (*Claims, error) {
	mapClaims, ok := claims.(jwt.MapClaims)
	if !ok {
		return nil, ErrClaimsInvalid
	}
	claimsb, err := json.Marshal(mapClaims)
	if err != nil {
		return nil, ErrClaimsInvalid
	}
	var tclaims *Claims
	if err := json.Unmarshal(claimsb, &tclaims); err != nil {
		return nil, ErrClaimsInvalid
	}
	return tclaims, nil
}

func Validate(tokenString string, keyfunc jwt.Keyfunc, refreshing bool, audience string, subject *string) (*Claims, error) {
	token, pErr := jwt.Parse(tokenString, keyfunc)
	if token == nil {
		return nil, ErrNoToken
	}

	claims, err := ParseClaims(token.Claims)
	if err != nil {
		return nil, ErrInvalid
	}

	if pErr != nil {
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return nil, ErrExpired
		}
		return nil, ErrInvalid
	}

	switch claims.Scope {
	case Access:
		if refreshing {
			return nil, ErrInvalid
		}
	case Refresh:
		if !refreshing {
			return nil, ErrInvalid
		}
	default:
		return nil, ErrInvalid
	}

	// verify owner
	if subject != nil && *subject != claims.Subject {
		return nil, ErrInvalid
	}

	// verify protocol
	if !claims.VerifyAudience(audience, true) {
		return nil, ErrInvalid
	}
	return claims, nil
}
