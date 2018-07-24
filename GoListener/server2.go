package main

import (
    "context"
    "fmt"
    "./kafka-go-master/kafka-go-master"
    "net/http"
    "io/ioutil"
    "encoding/json"
)


func main(){

    r := kafka.NewReader(kafka.ReaderConfig{
        Brokers:   []string{"localhost:9092", "localhost:9093", "localhost:9094"},
        GroupID:   "consumer-group-id",
        Topic:     "my_topic",
        MinBytes:  10e3, // 10KB
        MaxBytes:  10e6, // 10MB
    })

    for {
        m, err := r.ReadMessage(context.Background())
        if err != nil {
            break
        }

		//EXAMPLE CALL TO URL
        getRequest("http://www.interest.jackluo.online/get_interest.php")

        fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

    }

    r.Close()

}


func getRequest(url string){

    resp, err := http.Get(url)
    if err != nil {
        return
    }
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    var data map[string]interface{}
    err2 := json.Unmarshal([]byte(body), &data)
    if err2 != nil {
        panic(err2)
    }
    fmt.Println(data["interest"])

}


//func postRequest(urlF string){
//
//    resp, err := http.PostForm(urlF, url.Values{"key": {"Value"}, "id": {"123"}})
//    if err != nil {
//        return
//    }
//    defer resp.Body.Close()
//    body, err := ioutil.ReadAll(resp.Body)
//
//    var data map[string]interface{}
//    err2 := json.Unmarshal([]byte(body), &data)
//    if err2 != nil {
//        panic(err2)
//    }
//    fmt.Println(data["interest"])
//
//}