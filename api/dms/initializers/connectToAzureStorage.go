package initializers

import (
	"context"
	"fmt"
	"github.com/Azure/azure-storage-blob-go/azblob"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
)

var AzContainer azblob.ContainerURL

func ConnectToAzureStorage() {
	accountName := os.Getenv("AZURE_STORAGE_ACCOUNT_NAME")
	accountKey := os.Getenv("AZURE_STORAGE_ACCOUNT_KEY")
	credential, err := azblob.NewSharedKeyCredential(accountName, accountKey)
	if err != nil {
		panic(err)
	}
	p := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	URL, _ := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net", accountName))
	serviceURL := azblob.NewServiceURL(*URL, p)
	ctx := context.Background()

	// Create new container only if it does not exist
	AzContainer = serviceURL.NewContainerURL("firmware")
	_, err = AzContainer.GetProperties(ctx, azblob.LeaseAccessConditions{})
	if err != nil {
		_, err = AzContainer.Create(ctx, azblob.Metadata{}, azblob.PublicAccessNone)
		if err != nil {
			panic(err)
		}
	}
	log.Info("Connected to Azure Storage")
}
