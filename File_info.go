package lib1

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	//"strconv"
	"strings"
	//"syscall"
)

//文件专业操作模块

//获取文件详细信息
//获取文件MD5
//获取文件夹详细列表

//内置变量
//全局文件信息接收变量
var File_info_list = make(map[string]string)




func ListFile(myfolder string) (map[string]string,error) {

	files, err := ioutil.ReadDir(myfolder)

	out_put:= make(map[string]string)

	if err!=nil {
		return out_put,err
	}

	for _, file := range files {
		if file.IsDir() {
			//ListFile(myfolder + "/" + file.Name())
			out_put[myfolder + "/" + file.Name()]=myfolder
		} else {
			//fmt.Println(myfolder + "/" + file.Name())
			out_put[myfolder + "/" + file.Name()]=myfolder
		}
	}

	return out_put,err
}

func Visit(path string, f os.FileInfo, err error) error {
	//fmt.Printf("Visited: %s\n", path)

	File_info_list[path]=path

	return nil
}

/*


//文件内核结构

type Stat_t struct {
    Dev       uint64
    Ino       uint64
    Nlink     uint64
    Mode      uint32
    Uid       uint32
    Gid       uint32
    X__pad0   int32
    Rdev      uint64
    Size      int64
    Blksize   int64
    Blocks    int64
    Atim      Timespec
    Mtim      Timespec
    Ctim      Timespec
    X__unused [3]int64

}


基础文件信息

type FileInfo interface {
    Name() string       // base name of the file
    Size() int64        // length in bytes for regular files; system-dependent for others
    Mode() FileMode     // file mode bits
    ModTime() time.Time // modification time
    IsDir() bool        // abbreviation for Mode().IsDir()
    Sys() interface{}   // underlying data source (can return nil)

}


 */

 //参考信息

 //https://studygolang.com/topics/287
 //https://blog.csdn.net/phachon/article/details/78196874

func Create_File_info(path string) (map[string]string,error){

	var file_info = make(map[string]string)

	//fileInfo, err := os.Stat(path)
	
	/*
	
	//多平台参考
	if runtime.GOOS == "windows" {
    fileinfo, _ := os.Stat(path)
    stat := fileinfo.Sys().(*syscall.Win32FileAttributeData)
    aTimeSince = time.Since(time.Unix(0, stat.LastAccessTime.Nanoseconds()))
    cTimeSince = time.Since(time.Unix(0, stat.CreationTime.Nanoseconds()))
    mTimeSince = time.Since(time.Unix(0, stat.LastWriteTime.Nanoseconds()))
} else {
    fileinfo, _ := os.Stat(path)
    aTime = fileinfo.Sys().(*syscall.Stat_t).Atim
    cTime = fileinfo.Sys().(*syscall.Stat_t).Ctim
    mTime = fileinfo.Sys().(*syscall.Stat_t).Mtim
    aTimeSince = time.Since(time.Unix(aTime.Sec, aTime.Nsec))
    cTimeSince = time.Since(time.Unix(cTime.Sec, cTime.Nsec))
    mTimeSince = time.Since(time.Unix(mTime.Sec, mTime.Nsec))
}

	if(os.IsNotExist(err)) {
		fmt.Println("file not exist!")
		return file_info,err
	}

	if err == nil {
		//Atim := reflect.ValueOf(fileInfo.Sys()).Elem().FieldByName("Atim").Field(0).Int()
		//println("文件的访问时间：\n", Atim, )

		//修改时间
		//modTime := fileInfo.ModTime()
		file_info["mod_time"]=fileInfo.ModTime().Format("2006-01-02 15:04:05")
		file_info["name"]=fileInfo.Name()
		file_info["size"]=strconv.FormatInt(fileInfo.Size(),10)

		//文件内核信息
		fileSys := fileInfo.Sys().(*syscall.Win32FileAttributeData)

		//创建时间
		nanoseconds := fileSys.CreationTime.Nanoseconds() // 返回的是纳秒
		file_info["create_core_time"]=strconv.FormatInt(nanoseconds,10)
		file_info["create_time"]=Convert_appoint_number(File_info_list["create_core_time"][0:10])

		//写入时间
		fileSys = fileInfo.Sys().(*syscall.Win32FileAttributeData)
		nanoseconds = fileSys.LastWriteTime.Nanoseconds() // 返回的是纳秒
		file_info["write_core_time"]=strconv.FormatInt(nanoseconds,10)
		file_info["write_time"]=Convert_appoint_number(File_info_list["write_core_time"][0:10])

		//访问时间
		fileSys = fileInfo.Sys().(*syscall.Win32FileAttributeData)
		nanoseconds = fileSys.LastAccessTime.Nanoseconds() // 返回的是纳秒
		file_info["read_core_time"]=strconv.FormatInt(nanoseconds,10)
		file_info["write_time"]=Convert_appoint_number(File_info_list["read_core_time"][0:10])

		//文件访问权限
		file_info["mode"]=fmt.Sprintf("%d",fileInfo.Mode())
	}
	*/

	//return  file_info,err
	
		return  file_info,nil

}

func File_MD5(path string) string {

	f, err := os.Open(path)

	if err != nil {

		fmt.Println("Open", err)
		return ""
	}

	defer f.Close()
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f);
		err != nil {
		fmt.Println("Copy", err)
		return ""
	}

	//tmp:=md5hash.Sum(nil)
	//fmt.Printf("%x\n", md5hash.Sum(nil))

	return strings.ToUpper(fmt.Sprintf("%x", md5hash.Sum(nil)))

}
