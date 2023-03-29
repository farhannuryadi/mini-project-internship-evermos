package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

)

const DefaultPathAssetImage = "./public/files/"

func HandleMultipleFile(ctx *fiber.Ctx) error {
	form, errForm := ctx.MultipartForm()
	if errForm != nil {
		log.Println("Error read multiple file")
	}

	files := form.File["photos"]
	var fileNames []string

	for i, file := range files {
		var fileName string
		if file != nil {
			fileName = fmt.Sprintf("%d-%s", i, file.Filename)

			errSave := ctx.SaveFile(file, fmt.Sprintf("./public/files/%s", fileName))
			if errSave != nil {
				log.Println("Fail to store file")
			}
		} else {
			log.Println("nothing file to upload")
		}

		if fileName != "" {
			fileNames = append(fileNames, fileName)
		}
	}

	ctx.Locals("fileNames", fileNames)

	return ctx.Next()
}

func HandleSingleFile(ctx *fiber.Ctx) error {
	file, errFile := ctx.FormFile("photo")
	if errFile != nil {
		log.Println("Error read multiple file")
	}

	var fileName string

	if file != nil {
		fileName = file.Filename
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/files/%s", fileName))
		if errSaveFile != nil {
			log.Println(errSaveFile.Error())
		}
	} else {
		log.Println("tidak ada file upload")
	}

	ctx.Locals("fileName", fileName)

	return ctx.Next()
}

func HandleRemoveFile(fileName string, pathFile ...string) error {
	var removeName = DefaultPathAssetImage + fileName

	if len(pathFile) > 0 {
		removeName = pathFile[0] + fileName
	}

	err := os.Remove(removeName)
	if err != nil {
		log.Println("Failed to remove file")
		return err
	}

	return nil
}
