package main

import (
	// "bufio"
	"compress/gzip"
	// "context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"log"
	"net/http"
	"os"

	// "strings"
	// "time"
	// "html/template"
	// "github.com/adrianosela/sslmgr"
	// "golang.org/x/crypto/acme/autocert"
	// "github.com/gorilla/handlers"
	// "crypto/tls"

	"github.com/gorilla/mux"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodborg/mongo-driver/bson"
	// "gopkg.in/yaml.v2"
)

///////////////////////////////////////////////////////////////////////////////

func UUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = 0x80
	uuid[4] = 0x40
	boo := hex.EncodeToString(uuid)
	return boo, nil
}

///////////////////////////////////////////////////////////////////////////////

// func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
// 	defer cancel()
// 	defer func() {
// 		if err := client.Disconnect(ctx); err != nil {
// 			panic(err)
// 		}
// 	}()
// }

// func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
// 	return client, ctx, cancel, err
// }

// func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
// 	collection := client.Database(dataBase).Collection(col)
// 	result, err := collection.InsertOne(ctx, doc)
// 	return result, err
// }

// func UpdateOne(client *mongo.Client, ctx context.Context, filter interface{}, dataBase string, col string, update interface{}) (*mongo.UpdateResult, error) {
// 	collection := client.Database(dataBase).Collection(col)
// 	result, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
// 	return result, err
// }

// func Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
// 	collection := client.Database(dataBase).Collection(col)
// 	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
// 	return
// }

// func AlphaT_Insert(db string, coll string, ablob ReviewStruct) {
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	CheckError(err, "AlphaT_Insert_: Connections has failed")
// 	defer Close(client, ctx, cancel)
// 	_, err2 := InsertOne(client, ctx, db, coll, ablob)
// 	CheckError(err2, "AlphaT_Insert_has failed")
// }

// func AlphaT_Insert_Pics(db string, coll string, picinfo PicStruct) {
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	CheckError(err, "AlphaT_Insert_Pics: Connections has failed")
// 	defer Close(client, ctx, cancel)
// 	_, err2 := InsertOne(client, ctx, db, coll, picinfo)
// 	CheckError(err2, "AlphaT_Insert_Picshas failed")
// }

// func AddToQuarantineHandler(w http.ResponseWriter, r *http.Request) {
// 	uuid, _ := UUID()
// 	var name string = r.URL.Query().Get("name")
// 	var email string = r.URL.Query().Get("email")
// 	var message string = r.URL.Query().Get("message")
// 	var sig string
// 	if name != "" {
// 		sig = name
// 	} else if email != "" {
// 		s := strings.Split(email, "@")
// 		sig = s[0]
// 	} else {
// 		sig = ""
// 	}

// 	ct := time.Now()
// 	date := ct.Format("01-01-2021")

// 	var newReview = ReviewStruct{
// 		UUID:       uuid,
// 		Date:       date,
// 		Name:       name,
// 		Email:      email,
// 		Sig:        sig,
// 		Message:    message,
// 		Approved:   "no",
// 		Quarintine: "yes",
// 		Delete:     "no",
// 	}
// 	AlphaT_Insert("maindb", "main", newReview)
// m1 := "<p>A new review was posted</p>"
// m2 := "<a href='http://34.127.50.188/Admin'>AlphaTreeService Admin Page</>"
// m3 := m1 + m2
// m := gomail.NewMessage()
// m.SetHeader("From", "porthose.cjsmo.cjsmo@gmail.com")
// m.SetHeader("To", "porthose.cjsmo.cjsmo@gmail.com", "Alpha.treeservicecdm@gmail.com")
// m.SetHeader("Subject: NEW REVIEW Has Been Posted")
// m.SetBody("text/html", m3)
// d := gomail.NewDialer("smtp.gmail.com", 587, "porthose.cjsmo.cjsmo@gmail.com", "!Porthose1960")
// if err := d.DialAndSend(m); err != nil {
// 	panic(err)
// }
// }

// func AllApprovedReviews() bool {
// 	result := false
// 	filter := bson.M{"approved": "yes", "quarintine": "no", "delete": "no"}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	defer Close(client, ctx, cancel)
// 	CheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("maindb").Collection("main")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	CheckError(err, "AllReviews find has failed")
// 	var allRevs []ReviewStruct
// 	if err = cur.All(context.TODO(), &allRevs); err != nil {
// 		return result
// 	}
// 	if len(allRevs) != 0 {
// 		result = true
// 	}
// 	log.Println("AllReviews Info Complete")
// 	return result

// }

// func BackupReviewHandler(w http.ResponseWriter, r *http.Request) {
// 	filter := bson.M{}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	defer Close(client, ctx, cancel)
// 	CheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("maindb").Collection("main")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	CheckError(err, "AllReviews find has failed")
// 	var allRevs []ReviewStruct
// 	if err = cur.All(context.TODO(), &allRevs); err != nil {
// 		log.Fatal(err)
// 	}
// 	bString, _ := json.Marshal(allRevs)

// 	err = ioutil.WriteFile("/root/backup/backup.json", bString, 0644)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	bStringBytes := []byte(bString)
// 	t := time.Now().Format(time.RFC3339)
// 	new_file_name := "/root/backup/" + t + "_backup.gz"
// 	newFile, _ := os.Create(new_file_name)
// 	ww := gzip.NewWriter(newFile)
// 	ww.Write(bStringBytes)
// 	ww.Close()

// 	fmt.Println("this is new_file_name")
// 	fmt.Println(new_file_name)
// 	// m := gomail.NewMessage()
// 	// m.SetHeader("From", "porthose.cjsmo.cjsmo@gmail.com")
// 	// m.SetHeader("To", "porthose.cjsmo.cjsmo@gmail.com", "Alpha.treeservicecdm@gmail.com")
// 	// m.SetHeader("Subject: AlphaTreeService Reviews Backup")
// 	// m.SetBody("text/html", s)
// 	// m.Attach("/root/backup/" + name_of_file)
// 	// d := gomail.NewDialer("smtp.gmail.com", 587, "porthose.cjsmo.cjsmo@gmail.com", "!Porthose1960")
// 	// if err := d.DialAndSend(m); err != nil {
// 	// 	panic(err)
// 	// }
// }
func CheckError(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
		log.Println(msg)
		log.Println(err)
		panic(err)
	}
}

///////////////////////////////////////////////////////////////////////////////

func ShowIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://alphatreeservice.pages.dev", http.StatusSeeOther)
	// tmppath := "./assets/index.html"
	// tmpl := template.Must(template.ParseFiles(tmppath))
	// tmpl.Execute(w, tmpl)
}

func ShowAdmin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://atsa-dminsvelte.vercel.app", http.StatusSeeOther)
	// showtmppath := "./assets/admin.html"
	// showtmpl := template.Must(template.ParseFiles(showtmppath))
	// showtmpl.Execute(w, showtmpl)
}

func RemoveBackups() {
	err := os.Remove("./backup/backup.gz")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove("./assets/backup.gz")
	if err != nil {
		fmt.Println(err)
	}
}

func WriteJsonFile(alist string) {
	outfile_json := "./assets/backup.json"
	f, _ := os.Create(outfile_json)
	f.Write([]byte(alist))
	f.Close()
}

func WriteGzipFile(alist string) {
	fmt.Println(alist)
	outfile_gzip := os.Getenv("ATSGO_GZIP_PATH")
	f, _ := os.Create(outfile_gzip)
	z, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
	z.Write([]byte(alist))
	z.Close()

}

func ProcessReviewsHandler(w http.ResponseWriter, r *http.Request) {
	// RemoveBackups()

	reviews := r.URL.Query().Get("reviewslist")
	fmt.Printf("%T\n\n", reviews)
	// fmt.Println(reviews)
	WriteJsonFile(reviews)
	WriteGzipFile(reviews)

	// log.Println(reviews)
	// fmt.Println(reviews)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("GZIPISCOMPLETENOW")
	// 	log.Println("AllQuarintineReviews Info Complete")
	// 	tmpl2 := template.Must(template.ParseFiles("./assets/zoom.html"))
	// 	tmpl2.Execute(w, pic2)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("TEST_COMPLETE_IT_WORKS")
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", ShowIndex)
	r.HandleFunc("/admin", ShowAdmin)
	r.HandleFunc("/Backup", ProcessReviewsHandler)
	r.HandleFunc("/Test", TestHandler)
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	// http.ListenAndServe(":80", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
	// 	handlers.AllowedOrigins([]string{"*"}))(r))

	// cert := "fullchain1.pem"
	// key := "privkey1.pem"
	http.ListenAndServeTLS(":80", "cert1.pem", "privkey1.pem", r)

	// http.ListenAndServeTLS(":80", cert, key,
	// 	handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
	// 		handlers.AllowedOrigins([]string{"*"}))(r))

}

// DOO NOT DELETE atsgohttps@atsgo-340504.iam.gserviceaccount.com   DO NO DELETE

// func ProcessQuarantineHandler(w http.ResponseWriter, r *http.Request) {
// 	filter := bson.M{}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	defer Close(client, ctx, cancel)
// 	CheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("maindb").Collection("main")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	CheckError(err, "AllQuarintineReviews find has failed")
// 	var allRevs []ReviewStruct
// 	if err = cur.All(context.TODO(), &allRevs); err != nil {
// 		log.Fatal(err)
// 	}
// 	for _, rev := range allRevs {
// 		filter := bson.M{"uuid": rev.UUID}
// 		update := bson.M{"$set": bson.M{"approved": "yes", "quarintine": "no"}}
// 		client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 		defer Close(client, ctx, cancel)
// 		CheckError(err, "MongoDB connection has failed")
// 		UpdateOne(client, ctx, filter, "maindb", "main", update)
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode("Update complete")
// 	log.Println("AllQuarintineReviews Info Complete")
// }
