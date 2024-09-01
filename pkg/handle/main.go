package handle

import (
	utilsstorage "object_storage/utils/stoarage"

	"github.com/gofiber/fiber/v2"
)


func Store(ctx *fiber.Ctx) error {
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	objectName, err := utilsstorage.UploadFile(file)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": objectName})
}

func Show(ctx *fiber.Ctx) error {

	objectName := ctx.Params("filename")
	if objectName == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "No filename provided"})
	}

	objectReader,err := utilsstorage.ShowFile(objectName)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	
	return ctx.Status(fiber.StatusOK).SendStream(objectReader)
}

