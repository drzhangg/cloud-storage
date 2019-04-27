package handler

import (
	"cloud-storage/meta"
	"cloud-storage/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

//UploadHandler:处理文件上传
func UploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		//返回上传页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		//接收文件流及存储到本地目录
		file, head, err := r.FormFile("file") // 从页面读取接收的文件
		if err != nil {
			fmt.Printf("Failed to get data,err:%s\n", err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName:head.Filename,
			Location:"/tmp/" +head.Filename,
			UploadAt:time.Now().Format("2006-01-02 15:04:05"),
		}

		//本地创建存储文件路径
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			fmt.Printf("Failed to create file,err:%s\n",err.Error())
			return
		}
		defer newFile.Close()

		//将内存中文件copy到新的文件buf区
		fileMeta.FileSize,err = io.Copy(newFile,file)
		if err != nil {
			fmt.Printf("Failed to save data into faile,err:%s\n",err.Error())
			return
		}
		newFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		meta.UpdateFileMeta(fileMeta)

		//保存返回正确信息
		http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	}
}

//UploadSucHandler:上传已完成
func UploadSucHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w,"Upload finished!")
}
