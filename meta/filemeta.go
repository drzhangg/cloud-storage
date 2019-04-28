package meta

import "sort"

//FileMeta：文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string //文件存放路径
	UploadAt string //时间戳
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//UpdateFileMeta:新增/更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//GetFileMeta：通过sha1值获取文件的元信息对象
func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

//GetLastFileMeta：获取批量的文件元信息列表
func GetLastFileMeta(count int) []FileMeta {
	fMetaArray := make([]FileMeta,len(fileMetas))
	for _,v := range fileMetas{
		fMetaArray = append(fMetaArray,v)
	}

	sort.Sort(ByUploadTime(fMetaArray))
	return fMetaArray[0:count]
}

//RemoveFileMeta：删除元信息
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas,fileSha1)
}