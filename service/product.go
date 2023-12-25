package service

import (
	"fmt"
	"io"
	"os"

	"github.com/rizkiramadhan201987/learn-grpc/protogen/go/bfi/productservice/product"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type ProductService struct {
	product *product.Product
}

func NewProductService(productData *product.Product) *ProductService {
	return &ProductService{product: productData}
}

func (p *ProductService) GetProduct() *product.Product {
	return p.product
}

func (p *ProductService) ProductToJSON() ([]byte, error) {
	json, err := protojson.Marshal(p.product)
	return json, err
}
func (p *ProductService) JsonToProduct(jsonData *[]byte) (*product.Product, error) {
	if err := protojson.Unmarshal(*jsonData, p.product); err != nil {
		return &product.Product{}, err
	}
	return p.product, nil
}
func (p *ProductService) WriteFileProduct() error {
	data, err := proto.Marshal(p.product)
	if err != nil {
		return fmt.Errorf("error saat serialisasi : %v", err)
	}
	file, err := os.Create("product.proto")
	if err != nil {
		return fmt.Errorf("error saat membuat file : %v", err)
	}
	defer file.Close()
	err = os.WriteFile("product.proto", data, 0644)
	if err != nil {
		return fmt.Errorf("error saat menulis file : %v", err)
	}
	return nil
}
func (p *ProductService) ReadProductFile() (*product.Product, error) {
	productProto := &product.Product{}
	file, err := os.Open("product.proto")
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = proto.Unmarshal(data, productProto)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return productProto, nil
}
func (p *ProductService) WriteFileJson() error {
	data, err := protojson.Marshal(p.product)
	if err != nil {
		return fmt.Errorf("error saat marshal json : %v", err)
	}
	file, err := os.Create("product.json")
	if err != nil {
		return fmt.Errorf("error saat membuat file json : %v", err)
	}
	defer file.Close()
	err = os.WriteFile("product.json", data, 0644)
	if err != nil {
		return fmt.Errorf("error saat menulis file json : %v", err)
	}
	return nil
}
func (p *ProductService) ReadFileJson() (*product.Product, error) {
	productProto := &product.Product{}
	file, err := os.Open("product.json")
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	err = protojson.Unmarshal(data, productProto)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return productProto, nil
}
