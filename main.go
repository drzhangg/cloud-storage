package main

import (
	"cloud-storage/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload",handler.UploadHandler)	//处理文件上传接口
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)		//文件上传成功接口
	http.HandleFunc("/file/meta",handler.GetFileMetaHandler)
	err:=http.ListenAndServe(":8080",nil)
	if err != nil {
		fmt.Printf("failed to start server,err:%s",err.Error())
	}

}
