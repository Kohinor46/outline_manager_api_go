package outline_manager_api

import (
	"encoding/json"
	"crypto/tls"
	"log"
	"net/http"
	"bytes"
	"mime/multipart"
	"os"
	"strings"
	"fmt"
	"io"
)

type Keys map[string][]Key

type Key struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	AccessUrl string `json:"accessUrl"`
}

func Get_all_keys(API_URL string) (Keys, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", os.ExpandEnv(API_URL+"/access-keys/"), nil)
	if err != nil {
		log.Fatal(err)
		return Keys{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return Keys{}, err
	}
	defer resp.Body.Close()
	var k Keys
	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&k)
		if err != nil {
			log.Fatal("Ошибка декодирования всех ключей:\n", err)
			return Keys{}, err
		}
		return k, nil
	}
	return Keys{}, err
}

func Rename_key(API_URL string, key_id string, name string) error {
	form := new(bytes.Buffer)
	writer := multipart.NewWriter(form)
	formField, err := writer.CreateFormField("name")
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = formField.Write([]byte(name))

	writer.Close()

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("PUT", API_URL+ "/access-keys/"+key_id+"/name", form)
	if err != nil {
		log.Fatal(err)
		return err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func Delete_key(API_URL string, key_id string) error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("DELETE", os.ExpandEnv(API_URL+"/access-keys/"+key_id), nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	return nil
}

func Create_key(API_URL string) (Key, error){
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("POST", API_URL + "/access-keys", nil)
	if err != nil {
		log.Fatal(err)
		return Key{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return Key{}, err
	}
	defer resp.Body.Close()
	var k Key
	if resp.StatusCode == http.StatusOK {
		err = json.NewDecoder(resp.Body).Decode(&k)
		if err != nil {
			log.Fatal("Ошибка декодирования всех ключей:\n", err)
			return Key{}, err
		}
		return k, nil
	}
	return k, nil
}

func Set_data_limit_for_server(API_URL string, limit string)error{
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	var data = strings.NewReader("{\"limit\": {\"bytes\":"+ limit+"}}")
	req, err := http.NewRequest("PUT", API_URL + "/experimental/access-key-data-limit", data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("%s\n", bodyText)
	return nil
}

func Set_data_limit_for_key(API_URL string,key_id string, limit string)error{
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	var data = strings.NewReader("{\"limit\": {\"bytes\":"+ limit+"}}")
	req, err := http.NewRequest("PUT", API_URL+ "/access-keys/"+key_id+"/data-limit", data)
	if err != nil {
		log.Fatal(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("%s\n", bodyText)
	return nil
}

func Remove_data_limit_for_server (API_URL string)error{
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("DELETE", API_URL + "/experimental/access-key-data-limit", nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("%s\n", bodyText)
	return nil
}

func Remove_data_limit_for_key(API_URL string, key_id string) error{
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("DELETE", API_URL + "/access-keys/"+key_id+"/data-limit", nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("%s\n", bodyText)
	return nil
}
