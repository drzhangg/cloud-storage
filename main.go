package main

import (
	"cloud-storage/handler"
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	//文件上传相关接口
	http.HandleFunc("/file/upload", handler.UploadFile)        //处理文件上传接口
	http.HandleFunc("/file/upload/suc", handler.UploadSucHandler) //文件上传成功接口
	http.HandleFunc("/file/meta", handler.GetFileMetaHandler)
	http.HandleFunc("/file/query", handler.FileQueryHandler)       //文件查询信息接口
	http.HandleFunc("/file/download", handler.DownloadHandler)     //文件下载接口
	http.HandleFunc("/file/update", handler.FileMetaUpdateHandler) //更新元信息接口(重命名)
	http.HandleFunc("/file/delete", handler.FileDeleteHandler)     //删除文件接口

	//用户相关接口
	http.HandleFunc("/user/signup", handler.SignupHandler)  //用户注册接口
	http.HandleFunc("/user/signin", handler.SiginInHandler) //用户登录接口
	http.HandleFunc("/user/info", handler.HTTPInterceptor(handler.UserInfoHandler))  //用户查询接口

	fmt.Println("上传服务正在启动，监听端口:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("failed to start server,err:%s", err.Error())
	}

}
