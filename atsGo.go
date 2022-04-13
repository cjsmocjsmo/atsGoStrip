package main

import (
	// "bufio"
	"compress/gzip"
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"

	// "html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	// "path/filepath"
	// "sort"

	// "path/filepath"
	// "strconv"
	"strings"
	"time"

	"github.com/adrianosela/sslmgr"

	// "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodborg/mongo-driver/bson"

	// "gopkg.in/gomail.v2"
	"gopkg.in/yaml.v2"
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

func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, ctx, cancel, err
}

func InsertOne(client *mongo.Client, ctx context.Context, dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(ctx, doc)
	return result, err
}

func UpdateOne(client *mongo.Client, ctx context.Context, filter interface{}, dataBase string, col string, update interface{}) (*mongo.UpdateResult, error) {
	collection := client.Database(dataBase).Collection(col)
	result, err := collection.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true))
	return result, err
}

func Query(client *mongo.Client, ctx context.Context, dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	collection := client.Database(dataBase).Collection(col)
	result, err = collection.Find(ctx, query, options.Find().SetProjection(field))
	return
}

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

func AlphaT_Insert(db string, coll string, ablob ReviewStruct) {
	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
	CheckError(err, "AlphaT_Insert_: Connections has failed")
	defer Close(client, ctx, cancel)
	_, err2 := InsertOne(client, ctx, db, coll, ablob)
	CheckError(err2, "AlphaT_Insert_has failed")
}

// func AlphaT_Insert_Pics(db string, coll string, picinfo PicStruct) {
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	CheckError(err, "AlphaT_Insert_Pics: Connections has failed")
// 	defer Close(client, ctx, cancel)
// 	_, err2 := InsertOne(client, ctx, db, coll, picinfo)
// 	CheckError(err2, "AlphaT_Insert_Picshas failed")
// }

func AddToQuarantineHandler(w http.ResponseWriter, r *http.Request) {
	uuid, _ := UUID()
	var name string = r.URL.Query().Get("name")
	var email string = r.URL.Query().Get("email")
	var message string = r.URL.Query().Get("message")
	var sig string
	if name != "" {
		sig = name
	} else if email != "" {
		s := strings.Split(email, "@")
		sig = s[0]
	} else {
		sig = ""
	}

	ct := time.Now()
	date := ct.Format("01-01-2021")

	var newReview = ReviewStruct{
		UUID:       uuid,
		Date:       date,
		Name:       name,
		Email:      email,
		Sig:        sig,
		Message:    message,
		Approved:   "no",
		Quarintine: "yes",
		Delete:     "no",
	}
	AlphaT_Insert("maindb", "main", newReview)
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
}

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

func RemoveBackups() {
	err := os.Remove("/root/backup/backup.gz")
	if err != nil {
		fmt.Println(err)
	}
	err = os.Remove("/root/assets/backup.gz")
	if err != nil {
		fmt.Println(err)
	}
	return
}

func WriteJsonFile(alist string) {
	outfile_json := os.Getenv("ATSGO_JSON_PATH")
	f, _ := os.Create(outfile_json)
	f.Write([]byte(alist))
	f.Close()
	return
}

func WriteGzipFile(alist string) {
	// outfile_json := os.Getenv("ATSGO_JSON_PATH")
	// ofj, _ := os.Open(outfile_json)
	// reader := bufio.NewReader(ofj)
	// content, _ := ioutil.ReadAll(reader)
	fmt.Println(alist)
	// var revs string
	// revs, err := json.Marshal([]byte(alist))
	// if err != nil {
	// 	fmt.Println(err)
	// }
	outfile_gzip := os.Getenv("ATSGO_GZIP_PATH")
	f, _ := os.Create(outfile_gzip)
	z, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
	z.Write([]byte(alist))
	z.Close()
	// 	// ofj.Close()
	// 	// ofgz, _ := os.Open(outfile_gzip)
	// 	// ofgz.Write(content)
	// 	// ofgz.Close()

	return
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

	// revs, err := json.Marshal(reviews)
	// if err != nil {
	// 	log.Println(err)
	// }

	// f, _ := os.Create("/root/backup/backup.gz")
	// z, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
	// z.Write([]byte(reviews))
	// z.Close()

	// a, _ := os.Create("/root/assets/backup.gz")
	// m, _ := gzip.NewWriterLevel(a, gzip.BestCompression)
	// m.Write([]byte(reviews))
	// m.Close()
	// log.Println("finishered")
	// outfile := "/root/backup/backup.json"
	// f, err1 := os.OpenFile(outfile, os.O_APPEND|os.O_WRONLY, 0644)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// 	return
	// }
	// f.Write(revs)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reviews)
	// 	log.Println("AllQuarintineReviews Info Complete")
	// 	tmpl2 := template.Must(template.ParseFiles("./assets/zoom.html"))
	// 	tmpl2.Execute(w, pic2)
}

// func getAllBackupsHandler(w http.ResponseWriter, r *http.Request) {
// 	files, _ := filepath.Glob("/root/backup/*.gz")
// 	sort.Strings(files)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(files)
// 	log.Println("AllQuarintineReviews Info Complete")
// }

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("THIS IS A TEST IT WORKS")
}

type ReviewStruct struct {
	UUID       string `yaml:"UUID"`
	Date       string `yaml:"Date"`
	Name       string `yaml:"Name"`
	Email      string `yaml:"Email"`
	Sig        string `yaml:"Sig"`
	Message    string `yaml:"Message"`
	Approved   string `yaml:"Approved"`
	Quarintine string `yaml:"Quarintine"`
	Delete     string `yaml:"Delete"`
}

func (c *ReviewStruct) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

// type PicStruct struct {
// 	PicID  string `bson:"picid"`
// 	Pic    string `bson:"pic"`
// 	Thumb  string `bson:"thumb"`
// 	Page   string `bson:"page"`
// 	Orient bool   `bson:"orient"`
// }
func RemoveLogFile(logtxtfile string) bool {
	// var logtxtfile string = os.Getenv("AMPGO_SERVER_LOG_PATH")
	var result bool
	_, err := os.Stat(logtxtfile)
	if err == nil {
		log.Printf("logfile %s exists removing", logtxtfile)
		os.Remove(logtxtfile)
		result = true
	} else if os.IsNotExist(err) {
		log.Printf("logfile %s does not exists", logtxtfile)
		result = true
	} else {
		log.Printf("logfile %s stat error: %v", logtxtfile, err)
		result = false
	}
	return result
}

func StartServerLogging() string {
	var logtxtfile string = os.Getenv("ATSGO_SERVER_LOG_PATH")
	result := RemoveLogFile(logtxtfile)
	if result {
		file, err := os.OpenFile(logtxtfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			fmt.Println(err)
		}
		log.SetOutput(file)
	} else {
		fmt.Println("Unable to setup logging")
	}
	return "Logging started"
}

func init() {
	// if AllApprovedReviews() {
	// 	fmt.Println("Db present do nothing")
	// } else {

	data, err := ioutil.ReadFile("./assets/review1.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var rev1 ReviewStruct
	if err := rev1.Parse(data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(rev1)
	AlphaT_Insert("maindb", "main", rev1)
	os.Remove("./assets/review1.yaml")

	data2, err := ioutil.ReadFile("./assets/review2.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var rev2 ReviewStruct
	if err := rev2.Parse(data2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(rev2)
	AlphaT_Insert("maindb", "main", rev2)
	os.Remove("./assets/review2.yaml")

	data3, err := ioutil.ReadFile("./assets/fake1.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var rev3 ReviewStruct
	if err := rev3.Parse(data3); err != nil {
		log.Fatal(err)
	}
	fmt.Println(rev3)
	AlphaT_Insert("maindb", "main", rev3)
	os.Remove("./assets/fake1.yaml")

	data4, err := ioutil.ReadFile("./assets/fake2.yaml")
	if err != nil {
		log.Fatal(err)
	}
	var rev4 ReviewStruct
	if err := rev4.Parse(data4); err != nil {
		log.Fatal(err)
	}
	fmt.Println(rev4)
	AlphaT_Insert("maindb", "main", rev4)
	os.Remove("./assets/fake2.yaml")

	// os.Remove("/root/backup/backup.gz")
	// }
}

func main() {
	// StartServerLogging()
	r := mux.NewRouter()
	r.HandleFunc("/", ShowIndex)
	r.HandleFunc("/admin", ShowAdmin)
	r.HandleFunc("/Backup", ProcessReviewsHandler)
	r.HandleFunc("/test", TestHandler)

	// r.HandleFunc("/atq", AddToQuarantineHandler)
	// r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))

	ss, err := sslmgr.NewSecureServer(r, "atsio.xyz")
	if err != nil {
		log.Fatal(err)
	}
	ss.ListenAndServe()

	// http.ListenAndServe(":80", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 	handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
	// 	handlers.AllowedOrigins([]string{"*"}))(r))

	// http.ListenAndServeTLS(":80", "/root/atsio.crt", "/root/atsio.key",
	// 	handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
	// 		handlers.AllowedOrigins([]string{"*"}))(r))

}

// func AllQuarintineReviewsHandler(w http.ResponseWriter, r *http.Request) {
// 	filter := bson.M{"approved": "no", "quarintine": "yes", "delete": "no"}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	defer Close(client, ctx, cancel)
// 	CheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("maindb").Collection("main")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	CheckError(err, "AllQuarintineReviews find has failed")
// 	var allQRevs []ReviewStruct
// 	if err = cur.All(context.TODO(), &allQRevs); err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Printf("%s this is AllQuarintineReviews-", allQRevs)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allQRevs)
// }

// func ShowGalleryPage1Handler(w http.ResponseWriter, r *http.Request) {
// 	filter := bson.M{"page": "1"}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	defer Close(client, ctx, cancel)
// 	CheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("picdb").Collection("portrait")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	CheckError(err, "AllQuarintineReviews find has failed")
// 	var allPage1 []PicStruct
// 	if err = cur.All(context.TODO(), &allPage1); err != nil {
// 		log.Fatal(err)
// 	}
// 	tmpl2 := template.Must(template.ParseFiles("./assets/gallery.html"))
// 	tmpl2.Execute(w, allPage1)
// }

// func ShowGalleryPage2Handler(w http.ResponseWriter, r *http.Request) {
// 	filter := bson.M{"page": "2"}
// 	opts := options.Find()
// 	opts.SetProjection(bson.M{"_id": 0})
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	defer Close(client, ctx, cancel)
// 	CheckError(err, "MongoDB connection has failed")
// 	coll := client.Database("picdb").Collection("landscape")
// 	cur, err := coll.Find(context.TODO(), filter, opts)
// 	CheckError(err, "AllQuarintineReviews find has failed")
// 	var allPage2 []PicStruct
// 	if err = cur.All(context.TODO(), &allPage2); err != nil {
// 		log.Fatal(err)
// 	}
// 	tmpl2 := template.Must(template.ParseFiles("./assets/gallery.html"))
// 	tmpl2.Execute(w, allPage2)
// }

// func AtsGoFindOnePic(db string, coll string, filtertype string, filterstring string) PicStruct {
// 	filter := bson.M{filtertype: filterstring}
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	defer Close(client, ctx, cancel)
// 	CheckError(err, "AtsGoFindOnePic: MongoDB connection has failed")
// 	collection := client.Database(db).Collection(coll)
// 	var results PicStruct
// 	err = collection.FindOne(context.Background(), filter).Decode(&results)
// 	if err != nil {
// 		log.Println("AtsGoFindOnePic: find one has fucked up")
// 		log.Fatal(err)
// 	}
// 	return results
// }

// func ZoomPic1Handler(w http.ResponseWriter, r *http.Request) {
// 	// portrait
// 	pid := r.URL.Query().Get("picid")
// 	fmt.Println(pid)
// 	pic := AtsGoFindOnePic("picdb", "portrait", "picid", pid)
// 	fmt.Println(pic)
// 	tmpl2 := template.Must(template.ParseFiles("./assets/zoom.html"))
// 	tmpl2.Execute(w, pic)
// }

// func AllApprovedReviewsHandler(w http.ResponseWriter, r *http.Request) {
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
// 		log.Fatal(err)
// 	}
// 	log.Printf("%s this is AllReviews-", allRevs)
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(&allRevs)
// 	log.Println("AllReviews Info Complete")
// }

// func SetReviewToDeleteHandler(w http.ResponseWriter, r *http.Request) {
// 	var delUUID string = r.URL.Query().Get("uuid")
// 	filter := bson.M{"uuid": delUUID}
// 	update := bson.M{"$set": bson.M{"delete": "yes"}}
// 	client, ctx, cancel, err := Connect("mongodb://db:27017/atsgodb")
// 	defer Close(client, ctx, cancel)
// 	CheckError(err, "MongoDB connection has failed")
// 	UpdateOne(client, ctx, filter, "maindb", "main", update)
// }

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
