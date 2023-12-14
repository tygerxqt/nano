package api

import (
	"io"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/xid"
	"github.com/spf13/viper"
	"go.etcd.io/bbolt"
)

func CreateApiRoutes(app *fiber.App, db *bbolt.DB) {
	v1 := app.Group("/api")

	v1.Get("/files", func(c *fiber.Ctx) error {
		return GetAllFiles(c, db)
	})

	v1.Get("/files/:id", func(c *fiber.Ctx) error {
		return GetFile(c, db)
	})

	v1.Post("/upload", func(c *fiber.Ctx) error {
		return UploadFile(c, db)
	})
}

func GetAllFiles(c *fiber.Ctx, db *bbolt.DB) error {
	var files []string

	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("uploads"))
		return err
	})

	if err != nil {
		return err
	}

	err = db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("uploads"))

		b.ForEach(func(k, v []byte) error {
			files = append(files, string(k))
			return nil
		})

		return nil
	})

	if err != nil {
		return err
	}

	if len(files) == 0 {
		return c.Status(200).JSON([]string{})
	}

	return c.Status(200).JSON(files)
}

func GetFile(c *fiber.Ctx, db *bbolt.DB) error {
	id := c.Params("id")

	var fileName string

	err := db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte("uploads"))

		fileName = string(b.Get([]byte(id)))

		return nil
	})

	if err != nil {
		return err
	}

	return c.SendFile("./" + viper.GetString("data_dir") + "/uploads/" + fileName)
}

func UploadFile(c *fiber.Ctx, db *bbolt.DB) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	defer src.Close()

	// check if the nano-data directory exists
	if _, err := os.Stat("./" + viper.GetString("data_dir")); os.IsNotExist(err) {
		os.Mkdir("./"+viper.GetString("data_dir"), 0755)
	}

	// check if the uploads directory exists
	if _, err := os.Stat("./" + viper.GetString("data_dir") + "/uploads"); os.IsNotExist(err) {
		os.Mkdir("./"+viper.GetString("data_dir")+"/uploads", 0755)
	}

	key := xid.New()
	imageName := strings.ReplaceAll(file.Filename, " ", "-")

	dst, err := os.Create("./" + viper.GetString("data_dir") + "/uploads/" + imageName)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	// add the id to the uploads database
	err = db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("uploads"))
		if err != nil {
			return err
		}

		err = b.Put([]byte(key.String()), []byte(imageName))

		return err
	})

	if err != nil {
		return err
	}

	return c.Redirect("/")
}
