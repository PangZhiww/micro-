package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/micro/go-micro/client"
	//web "path/to/service/proto/web"
	web "day1/micro/rpc/srv/proto/srv"
)

func WebCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("request map:", request)

	// call the backend service
	webClient := web.NewSrvService("go.micro.srv.srv", client.DefaultClient)
	rsp, err := webClient.Call(context.TODO(), &web.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	fmt.Println("rep:", rsp)

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
