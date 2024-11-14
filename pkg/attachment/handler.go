package attachment

import (
	"os"
	"path/filepath"
	"strconv"
	"time"
	"top-gun-app-services/internal/handlers"
	"top-gun-app-services/pkg/models"

	"github.com/gofiber/fiber/v2"
	helpers "github.com/zercle/gofiber-helpers"
)

type attachmentHandler struct {
	service AttachmentService
	auth    *handlers.RouterResources
}

func NewWorkshopHandler(route fiber.Router, service AttachmentService, router *handlers.RouterResources) {
	handler := &attachmentHandler{service: service, auth: router}
	//CRUD
	route.Post("/file", handler.auth.ReqAuthHandler(0), handler.CreateAttachment())
	route.Get("/file/:attach_id", handler.auth.ReqAuthHandler(0), handler.GetAttachment())
	route.Get("/", handler.auth.ReqAuthHandler(0), handler.GetDatas())
	route.Get("/:id", handler.auth.ReqAuthHandler(0), handler.GetData())
}

// @Summary Get File data
// @Tags Integration
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /api/v1/attachment/{id} [get]
func (h *attachmentHandler) GetData() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusBadRequest,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		file, err := h.service.GetData(idInt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		return c.Status(fiber.StatusOK).JSON(helpers.ResponseForm{
			Success: true,
			Data:    file,
		})
	}
}

// @Summary Get File data
// @Tags Integration
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Router /api/v1/attachment/ [get]
func (h *attachmentHandler) GetDatas() fiber.Handler {
	return func(c *fiber.Ctx) error {
		paginate := models.Paginate{}
		err := c.QueryParser(&paginate)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusBadRequest,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		files, err := h.service.GetDatas(paginate)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		return c.Status(fiber.StatusOK).JSON(helpers.ResponseForm{
			Success: true,
			Data:    files,
		})
	}
}

// @Summary Send Attachment
// @Description file_type has two values model or sound
// @Tags Integration
// @Accept json
// @Produce json
// @Param file formData file true "File"
// @Param file_type formData string true "File Type"
// @Security ApiKeyAuth
// @Router /api/v1/attachment/file [post]
func (h *attachmentHandler) CreateAttachment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		file, err := c.FormFile("file")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusBadRequest,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		//get file type value
		fileType := c.FormValue("file_type")
		if fileType == "" {
			return c.Status(fiber.StatusBadRequest).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusBadRequest,
						Message: "file_type is required (model or sound)",
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		//get file name and ext
		fileName := file.Filename
		//save file
		// Define the upload directory
		var uploadDir string
		if fileType == "model" {
			uploadDir = "./upload/model"
		} else if fileType == "sound" {
			uploadDir = "./upload/sound"
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusBadRequest,
						Message: "file_type is invalid use model or sound",
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: "Unable to create upload directory",
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}

		// Define the path to save the file
		savePath := filepath.Join(uploadDir, fileName)

		// Save the file to the upload directory
		if err := c.SaveFile(file, savePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: "Unable to save file",
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		//save to db
		attachFile := AttachFile{
			FileName:  fileName,
			FilePath:  savePath,
			FileType:  fileType,
			CreatedAt: time.Now(),
		}
		err = h.service.CreateFile(attachFile)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		return c.Status(fiber.StatusCreated).JSON(helpers.ResponseForm{
			Success:  true,
			Messages: []string{"File uploaded successfully"},
		})
	}
}

// @Summary Get Attachment
// @Tags Integration
// @Accept json
// @Produce json
// @Param attach_id path string true "Attach ID"
// @Security ApiKeyAuth
// @Router /api/v1/attachment/file/{attach_id} [get]
func (h *attachmentHandler) GetAttachment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		fileID := c.Params("attach_id")
		fileIDInt, err := strconv.Atoi(fileID)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusBadRequest,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		files, err := h.service.GetFile(fileIDInt)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(helpers.ResponseForm{
				Errors: []helpers.ResponseError{
					{
						Code:    fiber.StatusInternalServerError,
						Message: err.Error(),
						Source:  helpers.WhereAmI(),
					},
				},
			})
		}
		return c.Status(fiber.StatusOK).SendFile(files.FilePath)
	}
}
