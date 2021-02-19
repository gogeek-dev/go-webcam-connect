package controller

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
	mysqldb "webcam/connection"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var Token string

func Loginview(c *gin.Context) {

	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {

	name := c.Request.PostFormValue("name")

	cpwds := c.Request.PostFormValue("password")

	Logauth(name, cpwds, c)

}

func Register(c *gin.Context) {

	c.HTML(200, "register.html", nil)
}

func Regsave(c *gin.Context) {

	db := mysqldb.SetupDB()

	key := hex.EncodeToString(make([]byte, 32))

	name := c.Request.PostFormValue("name")

	email := c.Request.PostFormValue("email")

	mobilenumber := c.Request.PostFormValue("mobilenumber")

	pwd := c.Request.PostFormValue("password")

	encryptpwd := encrypt(pwd, key)

	created_date, _ := time.Parse("2006-01-02 15:04:05 ", time.Now().Format("2006-01-02 15:04:05"))

	var emailid int

	_ = db.QueryRow("select count(emailid) from `user` where emailid=?", email).Scan(&emailid)

	log.Println("emailid count is", emailid)

	if emailid < 1 {

		insForm, err := db.Prepare("INSERT INTO `user`(name,emailid,mobilenumber,password,created_date) VALUES(?,?,?,?,?)")

		if err != nil {
			panic(err.Error())
		}

		insForm.Exec(name, email, mobilenumber, encryptpwd, created_date)

		Logauth(email, pwd, c)

	} else {

		c.HTML(200, "register.html", gin.H{"error1": "*Already have a mail id give another one"})
	}

}

func encrypt(stringToEncrypt string, keyString string) (encryptedString string) {

	key, _ := hex.DecodeString(keyString)

	plaintext := []byte(stringToEncrypt)

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonce := make([]byte, aesGCM.NonceSize())

	log.Println("nounce size is", nonce)

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return fmt.Sprintf("%x", ciphertext)
}

func decrypt(encryptedString string, keyString string) (decryptedString string) {

	key, _ := hex.DecodeString(keyString)

	enc, _ := hex.DecodeString(encryptedString)

	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	aesGCM, err := cipher.NewGCM(block)

	if err != nil {
		panic(err.Error())
	}

	nonceSize := aesGCM.NonceSize()

	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	log.Println("nonce,cipher", nonce, ciphertext)
	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)

	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}

func Logauth(Emailid string, Cpassword string, c *gin.Context) {

	db := mysqldb.SetupDB()

	key := hex.EncodeToString(make([]byte, 32))

	var id int
	var password, username string

	err1 := db.QueryRow("select id,name,password from `user` where emailid=?", Emailid).Scan(&id, &username, &password)

	if err1 != nil {
		c.HTML(200, "login.html", gin.H{"error": "Please singup,you are not register"})
	}

	decrypted := decrypt(password, key)

	fmt.Printf("decrypted : %s\n", decrypted)

	if Cpassword == decrypted {

		token, err := CreateToken(id, username)

		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, err.Error())
			return
		}

		Token = token
		log.Println("Userdata.Tokens", Token)

		c.Redirect(301, "/index")
	} else {
		c.HTML(200, "login.html", gin.H{"error": "*Password is incorrect", "email": Emailid})
	}
}

func Logout(c *gin.Context) {

	Token = " "

	time.Sleep(1 * time.Second)

	c.Writer.Header().Set("Pragma", "no-cache")

	c.Redirect(301, "/")
}

func CreateToken(userId int, userName string) (string, error) {

	var err error
	//Creating Access Token
	atClaims := jwt.MapClaims{}

	atClaims["authorized"] = true

	atClaims["user_id"] = userId

	atClaims["user_name"] = userName

	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("JWT_Key")))
	log.Println("Created Token", token)

	if err != nil {
		return "", err
	}

	return token, nil
}
