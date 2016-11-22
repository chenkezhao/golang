package controllers

import (
	"blog/utils/filetool"
	"blog/utils/qiniu"
	"fmt"
	"time"
)

const (
	EDITOR_IMG_DIR = "static/uploads/editor"
)

type ApiController struct {
	BaseController
}

func (this *ApiController) Upload() {
	result := map[string]interface{}{
		"success": false,
	}

	defer func() {
		this.Data["json"] = &result
		this.ServeJSON()
	}()

	_, header, err := this.GetFile("image")
	if err != nil {
		return
	}

	ext := filetool.Ext(header.Filename)
	imgPath := fmt.Sprintf("%s/%d%s", EDITOR_IMG_DIR, time.Now().Unix(), ext)

	filetool.InsureDir(EDITOR_IMG_DIR)
	err = this.SaveToFile("image", imgPath)
	if err != nil {
		return
	}

	imgUrl := "/" + imgPath
	if qiniu.UseQiniu {
		if addr, er := qiniu.UploadFile(imgPath, imgPath); er != nil {
			return
		} else {
			imgUrl = addr
			filetool.Unlink(imgPath)
		}
	}

	result["link"] = imgUrl
	result["success"] = true

}
