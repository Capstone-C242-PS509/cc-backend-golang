package entity

import (
	"backend-bangkit/pkg/errs"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid();column:user_id" json:"user_id"`
	Username  string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string    `gorm:"type:text;not null" json:"-"`
	Role      string    `gorm:"type:varchar(50);default:'user'" json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var jwtSecret = "SangatRahasia"

func (u *User) HashPassword() errs.MessageErr {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errs.NewInternalServerError("Failed to hash password")
	}
	u.Password = string(hashedPassword)

	return nil
}

func (u *User) ConvertStructToMap(user interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	// Get reflection of the struct
	value := reflect.ValueOf(user)

	// If it's a pointer, we need to dereference it
	if value.Kind() == reflect.Ptr {
		value = value.Elem()
	}

	typ := value.Type()

	// Iterate over all fields of the struct
	for i := 0; i < value.NumField(); i++ {
		fieldName := typ.Field(i).Name           // Get the field name
		fieldValue := value.Field(i).Interface() // Get the field value

		// Skip Password field
		if fieldName == "Password" {
			continue
		}

		// Insert the field into the map
		result[fieldName] = fieldValue
	}

	return result
}

func (u *User) ComparePassword(password string) errs.MessageErr {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	if err != nil {
		fmt.Println("Password doesn't match:", err)
		return errs.NewBadRequest("Password is not valid!")
	}

	return nil
}

func (u *User) CreateToken() (string, errs.MessageErr) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":    u.ID,
			"email": u.Email,
			"role":  u.Role,
		})

	signedString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Println("Error:", err.Error())
		return "", errs.NewInternalServerError("Failed to sign jwt token")
	}

	return signedString, nil
}

func (u *User) ParseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.NewUnauthenticated("Token method is not valid")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		fmt.Print("Hallo ", err)
		return nil, errs.NewUnauthenticated("Token is not valid")
	}

	return token, nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	if isBearer := strings.HasPrefix(bearerToken, "Bearer"); !isBearer {
		return errs.NewUnauthenticated("Token type should be Bearer")
	}

	splitToken := strings.Fields(bearerToken)
	if len(splitToken) != 2 {
		return errs.NewUnauthenticated("Token is not valid")
	}

	tokenString := splitToken[1]
	token, err := u.ParseToken(tokenString)
	if err != nil {
		fmt.Println(err)
		return err
	}

	var mapClaims jwt.MapClaims

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {

		return errs.NewUnauthenticated("Token is not valid")
	}
	mapClaims = claims

	return u.bindTokenToUserEntity(mapClaims)
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	id, okForId := claim["id"].(string)
	email, okForEmail := claim["email"].(string)
	role, okForRole := claim["role"].(string)

	if !okForId {
		return errs.NewUnauthenticated("Token doesn't contains id")
	}

	if !okForEmail {
		return errs.NewUnauthenticated("Token doesn't contains email")
	}

	if !okForRole {
		return errs.NewUnauthenticated("Token doesn't contains role")
	}
	uuid, err := uuid.Parse(id)

	if err != nil {
		return errs.NewBadRequest("ID invalid")
	}
	u.ID = uuid
	u.Email = email
	u.Role = role

	return nil
}
