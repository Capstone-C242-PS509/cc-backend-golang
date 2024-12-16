package pkg

import (
	"backend-bangkit/pkg/errs"
	"time"

	"github.com/asaskevich/govalidator"
)

func ValidateStruct(payload interface{}) errs.MessageErr {
	_, err := govalidator.ValidateStruct(payload)

	if err != nil {
		return errs.NewBadRequest(err.Error())
	}
	return nil
}

func ValidateDate(i interface{}, context interface{}) bool {

	Date, ok := i.(string)
	if !ok {
		return false
	}
	// Define the expected date format
	dateFormat := "2006-01-02"

	// Try parsing the date string using the specified format
	_, err := time.Parse(dateFormat, Date)

	print(err.Error())

	// Check if parsing was successful
	// Check if parsing was successful
	if err == nil {
		return true
	} else {
		return false
	}

}

func ValidateJenisAkun(i interface{}, context interface{}) bool {
	jenisAkun, ok := i.(string)
	if !ok {
		return false
	}
	return jenisAkun == "admin" || jenisAkun == "siswa" || jenisAkun == "pakar"
}

func ValidateStatusAlarm(i interface{}, context interface{}) bool {
	statusAlarm, ok := i.(string)
	if !ok {
		return false
	}
	return statusAlarm == "1" || statusAlarm == "0"
}

func init() {
	// Register the custom validation function
	govalidator.CustomTypeTagMap.Set("jenisAkunValidator", govalidator.CustomTypeValidator(ValidateJenisAkun))
	govalidator.CustomTypeTagMap.Set("date", govalidator.CustomTypeValidator(ValidateDate))
	govalidator.CustomTypeTagMap.Set("statusAlarm", govalidator.CustomTypeValidator(ValidateStatusAlarm))
}

// func DeleteImage(publicID string) errs.MessageErr {

// 	cld, contextImg := Credentials()
// 	// Menghapus gambar dari Cloudinary
// 	_, err := cld.Admin.DeleteAssetsByPrefix(contextImg, admin.DeleteAssetsByPrefixParams{
// 		Prefix: []string{publicID}})

// 	if err != nil {
// 		return errs.NewInternalServerError(err.Error())
// 	}

// 	return nil
// }

// func UploadImage(formFile string, key string, ctx *gin.Context) (*string, errs.MessageErr) {
// 	// Handle file upload

// 	file, err := ctx.FormFile(formFile)

// 	if err != nil && err != http.ErrMissingFile {

// 		return nil, errs.NewInternalServerError(err.Error())
// 	}

// 	if file == nil {

// 		kosong := ""
// 		return &kosong, nil
// 	}

// 	// Generate a filename using the user's email and keep the original file extension
// 	extension := filepath.Ext(file.Filename)

// 	if !(extension == ".png" || extension == ".jpg" || extension == ".jpeg" || extension == ".webp" || extension == ".JPG") {

// 		return nil, errs.NewBadRequest("Format file not suported, please upload only image type")

// 	}
// 	srcFile, _ := file.Open()

// 	defer srcFile.Close()

// 	cld, contextImg := Credentials()
// 	//uploadImage(cld, ctx)

// 	UploadImageCloud(cld, contextImg, srcFile, key)

// 	imageURL, err := GetImageInfo(cld, contextImg, key)

// 	if err != nil {
// 		fmt.Println("error getting image info:", err)
// 		return nil, errs.NewInternalServerError(err.Error())
// 	}

// 	return &imageURL, nil
// }

// // RenameImage renames the temporary uploaded image to the desired filename.
// func RenameImage(tempFilename, newFilename string) errs.MessageErr {

// 	tempFilename = "." + tempFilename
// 	newFilename = "." + newFilename
// 	if err := os.Rename(tempFilename, newFilename); err != nil {
// 		return errs.NewInternalServerError("Cannot rename file")
// 	}
// 	return nil
// }

// func CreateIndex(indexPath string, indexMapping mapping.IndexMapping) (bleve.Index, error) {
// 	index, err := bleve.New(indexPath, indexMapping)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return index, nil
// }

// // deleteAllDocuments menghapus semua dokumen dari indeks
// func DeleteAllDocuments(index bleve.Index) error {
// 	// Retrieve all documents in the index
// 	batch := index.NewBatch()
// 	searchRequest := bleve.NewSearchRequest(bleve.NewMatchAllQuery())
// 	searchResult, err := index.Search(searchRequest)
// 	if err != nil {
// 		return err
// 	}

// 	// Add all documents to the batch for deletion
// 	for _, hit := range searchResult.Hits {
// 		batch.Delete(hit.ID)
// 	}

// 	// Execute the batch delete operation
// 	err = index.Batch(batch)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func UploadImagePemeriksaan(formFile string, editFileName string, ctx *gin.Context) (*string, errs.MessageErr) {
// 	// Handle file upload

// 	file, err := ctx.FormFile(formFile)

// 	if err != nil && err != http.ErrMissingFile {

// 		return nil, errs.NewInternalServerError(err.Error())
// 	}

// 	if file == nil {

// 		kosong := ""
// 		return &kosong, nil
// 	}

// 	// Generate a filename using the user's email and keep the original file extension
// 	extension := filepath.Ext(file.Filename)

// 	if !(extension == ".png" || extension == ".jpg" || extension == ".jpeg" || extension == ".webp" || extension == ".JPG") {

// 		return nil, errs.NewBadRequest("Format file not suported, please upload only image type")

// 	}

// 	newFilename := fmt.Sprintf("%s%s", editFileName, extension)
// 	newFilename = strings.ReplaceAll(newFilename, " ", "")

// 	// Save the file to the server with the new filename
// 	if err := ctx.SaveUploadedFile(file, "./uploads/pemeriksaan/"+newFilename+"-temp"); err != nil {

// 		return nil, errs.NewInternalServerError(err.Error())
// 	}

// 	finalRoutes := "/uploads/pemeriksaan/" + newFilename + "-temp"

// 	return &finalRoutes, nil
// }

// func Credentials() (*cloudinary.Cloudinary, context.Context) {
// 	// Add your Cloudinary credentials, set configuration parameter
// 	// Secure=true to return "https" URLs, and create a context
// 	//===================
// 	cld, _ := cloudinary.NewFromParams("dciv82xna", "843286689852972", "rBKo_sWQWaVO61GZfBfzutYl-zY")
// 	cld.Config.URL.Secure = true
// 	ctx := context.Background()
// 	return cld, ctx
// }

// func UploadImageCloud(cld *cloudinary.Cloudinary, ctx context.Context, file interface{}, publicID string) {

// 	// Upload the image.
// 	// Set the asset's public ID and allow overwriting the asset with new versions
// 	_, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
// 		PublicID:       publicID,
// 		UniqueFilename: api.Bool(false),
// 		Overwrite:      api.Bool(true)})
// 	if err != nil {
// 		fmt.Println("error")
// 	}

// }

// func GetImageInfo(cld *cloudinary.Cloudinary, ctx context.Context, publicID string) (string, error) {
// 	// Mendapatkan informasi tentang asset yang sudah di-upload
// 	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: publicID})
// 	if err != nil {
// 		return "", err
// 	}
// 	return resp.SecureURL, nil
// }

// func DownloadImage(url string, filePath string) error {
// 	// Membuka HTTP request untuk mengunduh gambar
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Membuat file lokal untuk menyimpan gambar
// 	file, err := os.Create(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	// Menulis konten dari respon HTTP ke file lokal
// 	_, err = io.Copy(file, resp.Body)
// 	return err
// }
