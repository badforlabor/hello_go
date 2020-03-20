/**
 * Auth :   liubo
 * Date :   2019/10/18 11:26
 * Comment:
 */

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func postFile(fieldname, filename string, target_url string) error  {
	body_buf := bytes.NewBufferString("")
	body_writer := multipart.NewWriter(body_buf)

	// use the body_writer to write the Part headers to the buffer
	_, err := body_writer.CreateFormFile(fieldname, filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// the file data will be the second part of the body
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	// need to know the boundary to properly close the part myself.
	boundary := body_writer.Boundary()
	//close_string := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

	// use multi-reader to defer the reading of the file data until
	// writing to the socket buffer.
	request_reader := io.MultiReader(body_buf, fh, close_buf)
	fi, err := fh.Stat()
	if err != nil {
		fmt.Printf("Error Stating file: %s", filename)
		return err
	}
	req, err := http.NewRequest("POST", target_url, request_reader)
	if err != nil {
		return err
	}

	// Set headers for multipart, and Content Length
	req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
	req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("post err:", err.Error())
	} else {
		fmt.Println("post done:", resp.ContentLength, resp.Status, resp.StatusCode)
	}
	return err
}

// sample usage
func main() {
	target_url := "https://sentry.io/api/1783431/minidump/?sentry_key=db29c7ac9d954225bcd841df5552d63a"
	filename := "UE4Minidump.dmp"
	postFile("upload_file_minidump", filename, target_url)
	systemPause()
}
func systemPause() {
	fmt.Println("按回车键继续...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	//bufio.NewReader(os.Stdin).ReadByte()
}
