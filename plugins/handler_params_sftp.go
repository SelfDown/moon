package plugins

import (
	"fmt"
	common "github.com/SelfDown/collect/src/collect/common"
	config "github.com/SelfDown/collect/src/collect/config"
	templateService "github.com/SelfDown/collect/src/collect/service_imp"
	utils "github.com/SelfDown/collect/src/collect/utils"
	"github.com/demdxx/gocast"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"math"
	"strconv"
)

type Sftp struct {
	templateService.BaseHandler
}

type SftpFile struct {
	templateService.TsFile
	File *sftp.File
}

func (s *SftpFile) ReadAt(b []byte, off int64) (int, error) {
	return s.File.ReadAt(b, off)
}
func (s *SftpFile) Size() int64 {
	stat, _ := s.File.Stat()
	return stat.Size()

}

func getTargetPath(field string, params map[string]interface{}) (string, error) {
	tpl, error := config.CastTemplate(field)
	if error != nil {
		return "", error
	}
	targetPath := utils.RenderTplData(tpl, params).(string)
	if targetPath == "" || targetPath == "~" {
		targetPath = "./"
	}
	return targetPath, error
}
func (si *Sftp) HandlerData(template *config.Template, handlerParam *config.HandlerParam, ts *templateService.TemplateService) *common.Result {
	client := ts.GetThirdData(SshName).(*ssh.Client)
	//field := handlerParam.Field
	params := template.GetParams()
	result := make(map[string]interface{})
	sftpClient, err := sftp.NewClient(client)
	if err != nil {
		return common.NotOk(err.Error())
	}
	defer sftpClient.Close()
	fields := handlerParam.Fields
	// 依次编译模板
	for _, subField := range fields {
		saveField := subField.SaveField
		if subField.Type == "pwd" {
			cwd, err := sftpClient.Getwd()
			if err != nil {
				return common.NotOk(err.Error())
			}
			targetPath, err := getTargetPath(subField.Field, params)
			if targetPath == "." || targetPath == "./" || targetPath == "" || targetPath == "~" {
				result[saveField] = cwd
			} else {
				result[saveField] = targetPath
			}
		} else if subField.Type == "dir" {
			targetPath, err := getTargetPath(subField.Field, params)
			if err != nil {
				return common.NotOk(err.Error())
			}
			fList, err := sftpClient.ReadDir(targetPath)
			if err != nil {
				return common.NotOk(err.Error())
			}
			infoList := make([]map[string]interface{}, 0)
			for _, file := range fList {
				item := make(map[string]interface{})
				item["mode"] = file.Mode()
				item["name"] = file.Name()
				item["is_dir"] = file.IsDir()
				item["modify_time"] = utils.DateFormat(file.ModTime(), "")
				item["size"] = file.Size()
				item["size_info"] = getSize(file.Size())
				item["sys"] = file.Sys()
				infoList = append(infoList, item)
			}
			result[saveField] = infoList
		} else if subField.Type == "remove" {
			targetPath, err := getTargetPath(subField.Field, params)
			if err != nil {
				return common.NotOk(err.Error())
			}
			err = sftpClient.Remove(targetPath)
			if err != nil {
				return common.NotOk(err.Error())
			}
		} else if subField.Type == "upload" {
			file := ts.File

			if file == nil {
				return common.NotOk("上传文件不能为空")
			}
			targetDir, err := getTargetPath(subField.Field, params)
			if err != nil {
				return common.NotOk(err.Error())
			}

			sftpClient.MkdirAll(targetDir)
			targetPath := targetDir + "/" + ts.FileHeader.Filename
			remoteFile, err := sftpClient.Create(targetPath)
			defer remoteFile.Close()
			var offset int64 = 0
			var bufsize int64 = 1024 * 1024
			buf := make([]byte, bufsize)
			for {
				n, err := file.ReadAt(buf, offset)
				if err != nil && err != io.EOF {
					log.Panicln("read file error", err)
					break
				}
				if n == 0 {
					break
				}
				_, err = remoteFile.Write(buf[:n])
				if err != nil {
					log.Panicln("write file error", err)
					break
				}
				offset += int64(n)
			}

		} else if subField.Type == "download" {
			targetPath, err := getTargetPath(subField.Field, params)
			if err != nil {
				return common.NotOk(err.Error())
			}
			srcFile, err := sftpClient.Open(targetPath)
			if err != nil {
				common.NotOk(err.Error())
			}
			var f templateService.TsFile
			f = &SftpFile{File: srcFile}
			ts.IsFileResponse = true
			ts.SetTsFile(f)
			ts.ResponseFileName = utils.FileName(srcFile.Name())
			tsFile := ts.GetTsFile()
			size := tsFile.Size()
			c := ts.GetContext()
			c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", ts.ResponseFileName)) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
			c.Writer.Header().Add("Content-Type", "application/octet-stream")
			c.Writer.Header().Add("Content-Length", strconv.FormatInt(size, 10))
			var offset int64 = 0
			var bufsize int64 = 1024 * 1024
			buf := make([]byte, bufsize)
			for {
				n, err := tsFile.ReadAt(buf, offset)
				if err != nil && err != io.EOF {
					log.Panicln("read file error", err)
					break
				}
				if n == 0 {
					break
				}
				_, err = c.Writer.Write(buf[:n])
				if err != nil {
					log.Panicln("write file error", err)
					break
				}
				offset += int64(n)
			}
			c.Writer.Flush()

		}

	}

	r := common.Ok(result, "处理参数成功")
	return r
}
func getSize(size int64) string {

	a := gocast.ToFloat(size)
	c := gocast.ToFloat(1024)
	p := math.Pow10(2)
	if a < c {
		return gocast.ToString(math.Round(a*p)/p) + " B"
	} else if a < (c * c) {
		return gocast.ToString(math.Round(a/c*p)/p) + " KB"
	} else if a < (c * c * c) {
		return gocast.ToString(math.Round(a/(c*c)*p)/p) + " MB"
	} else if a < (c * c * c * c) {
		return gocast.ToString(math.Round(a/(c*c*c)*p)/p) + " GB"
	} else if a < (c * c * c * c * c) {
		return gocast.ToString(math.Round(a/(c*c*c*c)*p)/p) + " TB"
	} else {
		return gocast.ToString(math.Round(a/(c*c*c*c*c)*p)/p) + " PB"
	}
}
