//Cryptography: Its the practice of ensuring secure communication in the presence of third parties.
//- It makes use of serveral encryption, decryption & cryptic techniques to ensure that digital data is not expoited by unwanted and harmful entities.
// go provides good support for cryptography and hashing through the package "crypto".

/*The built in Crypto package provides us the implementation for all of these algorithms
-> aes
-> cipher
-> rsa
-> sha
*/

//So imagine you are creating a new registration form and there is a pwd field which you need to encrypt before storing to database.
//so we decided to go with MD5 hashing.

package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// NOTE: MD5 hash is a mathematical algorithm that is used to create a unique code from a text string.
func encodeWithMD5(str string) string {
	var hashCode = md5.Sum([]byte(str))
	return hex.EncodeToString(hashCode[:]) //this return the hexadecimal encoding equivalent of the string passed to the MD5 func.
}

func main() {
	var password string
	fmt.Scanln(&password)
	fmt.Println("Password encrypted to:", encodeWithMD5(password))
}
