package main

import (
	"fmt"
	"log"
	"time"

	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/productservice/product"
	"github.com/rizkiramadhan201987/learn-grpc/service"
)

type logWriter struct {
}

func (writer logWriter) Write(bytes []byte) (int, error) {
	return fmt.Print(time.Now().Format("15:04:05" + " " + string(bytes)))
}
func main() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))

	// service.UserProtoToJson()
	// service.BasicUnmarshall()
	// service.BasicOneOf()
	// service.BasicMap()
	// service.WriteProto()
	// service.WriteProtoToJsonBytes()
	// service.ReadJsonBytesToProto()
	// service.ReadProto()
	// service.BasicHello()
	// service.SayHelloUser()
	// service.UserProtoToJson()
	productProto := product.Product{
		ProductName: "Tepung Terigu",
		Price:       10000,
	}
	productService := service.NewProductService(&productProto)

	log.Println("Hasil inisialisasi : ", *productService)
	log.Println("Hasil pembuatan product : ", productService.GetProduct())
	productJSON, err := productService.ProductToJSON()
	if err != nil {
		panic(err)
	}
	log.Println("Hasil Product proto -> JSON : ", string(productJSON))
	var result *product.Product
	result, err = productService.JsonToProduct(&productJSON)
	if err != nil {
		panic(err)
	}
	log.Println("Hasil JSON -> Product proto : ", result)
	if err = productService.WriteFileProduct(); err != nil {
		panic(err)
	}
	if err = productService.WriteFileJson(); err != nil {
		panic(err)
	}
	result, err = productService.ReadProductFile()
	if err != nil {
		panic(err)
	}
	log.Println("Hasil Read Proto File : ", result)
	result, err = productService.ReadFileJson()
	if err != nil {
		panic(err)
	}
	log.Println("Hasil Read Proto JSON : ", result)
	// productServic.
	// service.UserJsonToProto()
	// service.AddUserToUserGroup()
	// service.JobSearch()
	// service.JobSearchSoftware()
}
