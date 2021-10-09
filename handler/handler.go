package handler

import (
	"Rest_api/encrypt"
	"Rest_api/model"
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"regexp"
	"time"
)

var (
	getUserCred = regexp.MustCompile(`^\/users\/(\w+)$`)
	getPostCred = regexp.MustCompile(`^\/posts\/(\w+)$`)
	getFinal    = regexp.MustCompile(`\/posts/users\/(\w+)$`)
)

func ConnectDB() (*mongo.Collection, *mongo.Collection) {

	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	UserSchema := client.Database("abhinav").Collection("users")
	PostSchema := client.Database("abhinav").Collection("posts")

	return UserSchema, PostSchema
}

var UserSchema, PostSchema = ConnectDB()

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	key := "653569012340000004"
	hashed_password := encrypt.Encrypt(key, user.Password)
	user.Password = hashed_password

	result, err := UserSchema.InsertOne(context.TODO(), &user)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(result)

}

func GetUserByIDEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User
	Path := getUserCred.FindStringSubmatch(r.URL.Path)
	id := Path[1]

	filter := bson.M{"_id": id}
	err := UserSchema.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		w.WriteHeader(404)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func CreatePostEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post model.Post
	post.TimeStamp = time.Now()

	_ = json.NewDecoder(r.Body).Decode(&post)
	result, err := PostSchema.InsertOne(context.TODO(), &post)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func GetPostByIDEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post model.Post

	Path := getPostCred.FindStringSubmatch(r.URL.Path)
	id := Path[1]

	filter := bson.M{"_id": id}
	err := PostSchema.FindOne(context.TODO(), filter).Decode(&post)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return

	}

	json.NewEncoder(w).Encode(post)
}

func GetUsersPostByIdEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts []model.Post

	Path := getFinal.FindStringSubmatch(r.URL.Path)

	id := Path[1]

	cur, err := PostSchema.Find(context.TODO(), bson.M{})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	for cur.Next(context.TODO()) {
		var single_post model.Post

		err := cur.Decode(&single_post)
		if err != nil {
			log.Fatal(err)
		}
		if (single_post.UserID) == id {
			posts = append(posts, single_post)
		}
	}

	if err := cur.Err(); err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(posts)
}
