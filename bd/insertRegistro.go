package bd

import (
	"context"
	"time"

	"github.com/coralheyman/twitter/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro se inserta un usuario */
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twitter")
	collection := db.Collection("usuarios")
	u.Password, _ = EncriptarPassword(u.Password)

	result, err := collection.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}
	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil
}
