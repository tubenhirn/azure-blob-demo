package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func main() {

	ac := flag.String("a", "", "the action (create, delete)")
	sa := flag.String("s", "", "name of the azure storageaccount")
	co := flag.Int("n", 5, "# of iterations")
	flag.Parse()

	action := *ac
	storageAccountName := *sa
	blobCounter := *co
	if storageAccountName == "" {
		panic("Flat storageAccountName missing.")
	}

	blobNameTemplate := "what-a-time-to-be-a-blob"
	containerNameTemplate := "a-container"

	ctx := context.Background()

	url := fmt.Sprintf("https://%s.blob.core.windows.net/", storageAccountName)

	credential, _ := azidentity.NewDefaultAzureCredential(nil)
	client, _ := azblob.NewClient(url, credential, nil)

	if action == "create" {
		for i := 0; i < blobCounter; i++ {
			err := createData(ctx, i, client, containerNameTemplate, blobNameTemplate)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else if action == "delete" {
		for i := 0; i < blobCounter; i++ {
			err := cleanData(ctx, i, client, containerNameTemplate, blobNameTemplate)
			if err != nil {
				fmt.Println(err)
			}
		}
	} else {
		fmt.Println("No action provided. Doin nothin.")
	}
}

func createData(ctx context.Context, counter int, client *azblob.Client, containerNameTemplate string, blobNameTemplate string) error {
	_, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*3000))
	defer cancel()

	containerName := fmt.Sprintf("%s-%d", containerNameTemplate, counter)
	blobName := blobNameTemplate

	fmt.Printf("Creating a container named %s\n", containerName)
	_, err := client.CreateContainer(ctx, containerName, nil)

	data := []byte("\nIm a blob.\n")
	_, err = client.UploadBuffer(ctx, containerName, blobName, data, &azblob.UploadBufferOptions{})

	return err
}

func cleanData(ctx context.Context, counter int, client *azblob.Client, containerNameTemplate string, blobNameTemplate string) error {
	_, cancel := context.WithTimeout(ctx, time.Duration(time.Millisecond*3000))
	defer cancel()

	containerName := fmt.Sprintf("%s-%d", containerNameTemplate, counter)
	blobName := blobNameTemplate

	fmt.Printf("Deleting the blob " + blobName + "\n")

	_, err := client.DeleteBlob(ctx, containerName, blobName, nil)

	fmt.Printf("Deleting the container " + containerName + "\n")
	_, err = client.DeleteContainer(ctx, containerName, nil)

	return err
}
