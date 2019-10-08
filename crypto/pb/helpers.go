package textile_crypto_pb

import (
	"github.com/gogo/protobuf/proto"
	ic "github.com/libp2p/go-libp2p-core/crypto"
)

// PubKeyUnmarshaller is a func that creates a PubKey from a given slice of bytes
type PubKeyUnmarshaller func(data []byte) (ic.PubKey, error)

// PrivKeyUnmarshaller is a func that creates a PrivKey from a given slice of bytes
type PrivKeyUnmarshaller func(data []byte) (ic.PrivKey, error)

// PubKeyUnmarshallers is a map of unmarshallers by key type
var PubKeyUnmarshallers = map[KeyType]PubKeyUnmarshaller{
	KeyType_RSA:       ic.UnmarshalRsaPublicKey,
	KeyType_Ed25519:   ic.UnmarshalEd25519PublicKey,
	KeyType_Secp256k1: ic.UnmarshalSecp256k1PublicKey,
	KeyType_ECDSA:     ic.UnmarshalECDSAPublicKey,
}

// PrivKeyUnmarshallers is a map of unmarshallers by key type
var PrivKeyUnmarshallers = map[KeyType]PrivKeyUnmarshaller{
	KeyType_RSA:       ic.UnmarshalRsaPrivateKey,
	KeyType_Ed25519:   ic.UnmarshalEd25519PrivateKey,
	KeyType_Secp256k1: ic.UnmarshalSecp256k1PrivateKey,
	KeyType_ECDSA:     ic.UnmarshalECDSAPrivateKey,
}

// UnmarshalPublicKey converts a protobuf serialized public key into its
// representative object
func UnmarshalPublicKey(data []byte) (ic.PubKey, error) {
	pmes := new(PublicKey)
	err := proto.Unmarshal(data, pmes)
	if err != nil {
		return nil, err
	}

	return PublicKeyFromProto(pmes)
}

// PublicKeyFromProto converts an unserialized protobuf PublicKey message
// into its representative object.
func PublicKeyFromProto(pmes *PublicKey) (ic.PubKey, error) {
	um, ok := PubKeyUnmarshallers[pmes.GetType()]
	if !ok {
		return nil, ic.ErrBadKeyType
	}

	return um(pmes.GetData())
}

// MarshalPublicKey converts a public key object into a protobuf serialized
// public key
func MarshalPublicKey(k ic.PubKey) ([]byte, error) {
	pbmes, err := PublicKeyToProto(k)
	if err != nil {
		return nil, err
	}

	return proto.Marshal(pbmes)
}

// PublicKeyToProto converts a public key object into an unserialized
// protobuf PublicKey message.
func PublicKeyToProto(k ic.PubKey) (*PublicKey, error) {
	pbmes := new(PublicKey)
	pbmes.Type = KeyType(k.Type())
	data, err := k.Raw()
	if err != nil {
		return nil, err
	}
	pbmes.Data = data
	return pbmes, nil
}

// UnmarshalPrivateKey converts a protobuf serialized private key into its
// representative object
func UnmarshalPrivateKey(data []byte) (ic.PrivKey, error) {
	pmes := new(PrivateKey)
	err := proto.Unmarshal(data, pmes)
	if err != nil {
		return nil, err
	}

	um, ok := PrivKeyUnmarshallers[pmes.GetType()]
	if !ok {
		return nil, ic.ErrBadKeyType
	}

	return um(pmes.GetData())
}

// MarshalPrivateKey converts a key object into its protobuf serialized form.
func MarshalPrivateKey(k ic.PrivKey) ([]byte, error) {
	pbmes := new(PrivateKey)
	pbmes.Type = KeyType(k.Type())
	data, err := k.Raw()
	if err != nil {
		return nil, err
	}

	pbmes.Data = data
	return proto.Marshal(pbmes)
}
