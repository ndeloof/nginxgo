package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `
          ##         .
    ## ## ##        ==
 ## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
 \    \         __/
  \____\_______/

	
Hello from Docker!

`)

	for _, pair := range os.Environ() {
		fmt.Println(pair)
	}

	bucket := "jefaisuntest"
	item := "/tmp/foo"
	file, err := os.Create(item)

	defer file.Close()

	fmt.Println("Load AWS config...")
	sess, err := session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("eu-west-3"),
		},
		SharedConfigState: session.SharedConfigEnable,
	})

	if err != nil {
		fmt.Sprintf("failed to load SDK config: %v", err)
		return
	}

	fmt.Println("Create S3 Manager")
	downloader := s3manager.NewDownloader(sess)

	fmt.Println("Download file from s3")
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String("blabla.txt"),
		})
	if err != nil {
		fmt.Fprintf(w, "Unable to download item %q, %v\n", item, err)
		return
	}

	fmt.Fprintf(w, "Downloaded %s %d bytes", file.Name(), numBytes)

	fmt.Fprintln(w)
	fmt.Fprintln(w)

	bytes, err := ioutil.ReadFile(item)
	if err != nil {
		fmt.Fprintf(w, "Unable to read downloaded item %q, %v\n", item, err)
		return
	}
	w.Write(bytes)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
