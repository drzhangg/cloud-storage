package main

import (
	"cloud-storage/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)        //处理文件上传接口
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler) //文件上传成功接口
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)       //文件查询信息接口
	http.HandleFunc("/file/download", handler.DownloadHandler)     //文件下载接口
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler) //更新元信息接口(重命名)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)     //删除文件接口
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		fmt.Printf("failed to start server,err:%s", err.Error())
	}

}
