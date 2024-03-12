package service

import (
	"go_mall/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvatarPath + "user" + bId + "/" //路径拼接
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user" + bId + "/" + userName + ".jpg", nil
	//返回数据库需要存储的路径
}

func UploadProductToLocalStatic(file multipart.File, userId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.ProductPath + "boss" + bId + "/" //路径拼接
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg"
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		return
	}
	return "boss" + bId + "/" + productName + ".jpg", nil
	//返回数据库需要存储的路径
}

// 判断文件路径是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 0755)
	return err == nil
}
