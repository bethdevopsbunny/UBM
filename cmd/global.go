package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func authentication() {

	ApiKey = os.Getenv("UNITY_API_KEY")

	if ApiKey == "" {
		log.Errorf("No api key provided. Please set UNITY_API_KEY environment variable")
		os.Exit(1)
	}
}

type returnValue struct {
	BuildStatus string `json:"BuildStatus"`
	BuildGUID   string `json:"BuildGUID"`
}

var (
	filePath     string
	org          int
	buildNumber  int
	buildTarget  string
	projectID    string
	raw          bool
	ApiKey       string
	requestDelay time.Duration
)

func downloadFile(filepath string, url string) (err error) {

	out, err := os.Create(filepath)
	if err != nil {
		log.Errorf("Failed to create file")
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		log.Errorf("Failed to get data")
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Failed data request: %s", resp.Status)
		return err
	}
	log.Infof("Successfully Requested Artifact")

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Errorf("Failed to write to file")
		return err
	}
	log.Infof("Completed Artifact %s Download", filePath)

	return nil
}

func splitAtUpper(text string) string {
	re := regexp.MustCompile(`[A-Z][^A-Z]*`)
	var returnstring string
	fmt.Printf("Pattern: %v\n", re.String()) // Print Pattern

	submatchall := re.FindAllString(text, -1)
	for _, element := range submatchall {
		returnstring += element
		returnstring += " "
	}

	// if single word
	if submatchall == nil {
		return strings.Title(text)
	}

	return strings.Title(strings.TrimSuffix(returnstring, " "))
}

// A collection of functions to allow you to unzip and edit the contents.
// You cant choose the root directory name so this allowed you to do so for easier deployment.
// But im thinking this might be better left out?

//func cleanDownloadFile(filepath string, url string) (err error) {
//
//	out, err := os.Create(filepath)
//	if err != nil {
//		log.Errorf("Failed to create file")
//		return err
//	}
//	defer out.Close()
//
//	resp, err := http.Get(url)
//	if err != nil {
//		log.Errorf("Failed to get data")
//		return err
//	}
//	defer resp.Body.Close()
//
//	if resp.StatusCode != http.StatusOK {
//		log.Errorf("Failed data request: %s", resp.Status)
//		return err
//	}
//	log.Infof("Successfully Requested Artifact")
//
//	_, err = io.Copy(out, resp.Body)
//	if err != nil {
//		log.Errorf("Failed to write to file")
//		return err
//	}
//	log.Infof("Completed Artifact %s Download", filePath)
//
//	unzipSource("artifact.zip", "tmp")
//	renameFile()
//	zipSource("tmp/artifact", "clean_artifact.zip")
//
//	err = os.RemoveAll("tmp")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	return nil
//}
//
//func unzipSource(source, destination string) error {
//	// 1. Open the zip file
//	reader, err := zip.OpenReader(source)
//	if err != nil {
//		return err
//	}
//	defer reader.Close()
//
//	// 2. Get the absolute destination path
//	destination, err = filepath.Abs(destination)
//	if err != nil {
//		return err
//	}
//
//	// 3. Iterate over zip files inside the archive and unzip each of them
//	for _, f := range reader.File {
//		err := unzipFile(f, destination)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func unzipFile(f *zip.File, destination string) error {
//	// 4. Check if file paths are not vulnerable to Zip Slip
//	filePath := filepath.Join(destination, f.Name)
//	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
//		return fmt.Errorf("invalid file path: %s", filePath)
//	}
//
//	// 5. Create directory tree
//	if f.FileInfo().IsDir() {
//		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
//			return err
//		}
//		return nil
//	}
//
//	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
//		return err
//	}
//
//	// 6. Create a destination file for unzipped content
//	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
//	if err != nil {
//		return err
//	}
//	defer destinationFile.Close()
//
//	// 7. Unzip the content of a file and copy it to the destination file
//	zippedFile, err := f.Open()
//	if err != nil {
//		return err
//	}
//	defer zippedFile.Close()
//
//	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
//		return err
//	}
//	return nil
//}
//
//func zipSource(source, target string) error {
//	// 1. Create a ZIP file and zip.Writer
//	f, err := os.Create(target)
//	if err != nil {
//		return err
//	}
//	defer f.Close()
//
//	writer := zip.NewWriter(f)
//	defer writer.Close()
//
//	// 2. Go through all the files of the source
//	return filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//
//		// 3. Create a local file header
//		header, err := zip.FileInfoHeader(info)
//		if err != nil {
//			return err
//		}
//
//		// set compression
//		header.Method = zip.Deflate
//
//		// 4. Set relative path of a file as the header name
//		header.Name, err = filepath.Rel(filepath.Dir(source), path)
//		if err != nil {
//			return err
//		}
//		if info.IsDir() {
//			header.Name += "/"
//		}
//
//		// 5. Create writer for the file header and save content of the file
//		headerWriter, err := writer.CreateHeader(header)
//		if err != nil {
//			return err
//		}
//
//		if info.IsDir() {
//			return nil
//		}
//
//		f, err := os.Open(path)
//		if err != nil {
//			return err
//		}
//		defer f.Close()
//
//		_, err = io.Copy(headerWriter, f)
//		return err
//	})
//}
//
//func renameFile() {
//	files, err := ioutil.ReadDir("tmp")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	var filez string
//	for _, file := range files {
//		if file.IsDir() {
//			filez = file.Name()
//		}
//	}
//
//	err = os.Rename("tmp/"+filez, "tmp/artifact")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//}
//
//func getFilename() string {
//	files, err := ioutil.ReadDir("tmp")
//	if err != nil {
//		log.Errorf("failed to read tmp dir")
//		os.Exit(1)
//	}
//
//	var filez string
//	for _, file := range files {
//		if file.IsDir() {
//			filez = file.Name()
//		}
//	}
//
//	return filez
//
//}
