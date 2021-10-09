package encrypt

import (
	"crypto/aes"
	"encoding/base64" 
	"crypto/cipher" 
)

var iv = []byte{35, 46, 57, 22, 85, 35, 24, 84, 87, 35, 88, 98, 66, 32, 14, 05}                                                                                                                             
                                                                                                                                                                                                            
func encodeMain(b []byte) string {                                                                                                                                                                        
    return base64.StdEncoding.EncodeToString(b)                                                                                                                                                             
}                                                                                                                                                                                                           
                                                                                                                                                                                                            
func decodeMain(s string) []byte {                                                                                                                                                                        
    data, err := base64.StdEncoding.DecodeString(s)                                                                                                                                                         
    if err != nil { panic(err) }                                                                                                                                                                            
    return data                                                                                                                                                                                             
}                                                                                                                                                                                                           
                                                                                                                                                                                                            
func Encrypt(key, text string) string {                                                                                                                                                                     
                                                                                                                                                                                        
    block, err := aes.NewCipher([]byte(key))                                                                                                                                                                
    if err != nil { panic(err) }                                                                                                                                                                            
    plaintext := []byte(text)                                                                                                                                                                               
    cfb := cipher.NewCFBEncrypter(block, iv)                                                                                                                                                                
    ciphertext := make([]byte, len(plaintext))                                                                                                                                                              
    cfb.XORKeyStream(ciphertext, plaintext)                                                                                                                                                                 
    return encodeMain(ciphertext)                                                                                                                                                                         
}                                                                                                                                                                                                           
                                                                                                                                                                                                            
func Decrypt(key, text string) string {                                                                                                                                                                     
                                                                                                                                                                                           
    block, err := aes.NewCipher([]byte(key))                                                                                                                                                                
    if err != nil { panic(err) }                                                                                                                                                                            
    ciphertext := decodeMain(text)                                                                                                                                                                        
    cfb := cipher.NewCFBEncrypter(block, iv)                                                                                                                                                                
    plaintext := make([]byte, len(ciphertext))                                                                                                                                                              
    cfb.XORKeyStream(plaintext, ciphertext)                                                                                                                                                                 
    return string(plaintext)                                                                                                                                                                                
} 