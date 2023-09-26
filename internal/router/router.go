package router

import (
	"OpenIDProvider/internal/model"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func InitRouter() {
	//健康检查路由
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		msg := model.NewMessage(http.StatusOK, "服务正常！", nil)
		jsonMsg, err := json.Marshal(msg)
		if err != nil {
			log.Fatal("Json转化失败", err)
		} else {
			io.WriteString(w, string(jsonMsg))
		}
	})

	// http.HandleFunc("/insertUser",controller.InsertUserHandle(w http.ResponseWriter, r *http.Request))
	// http.Post("/insertUser",)
}
