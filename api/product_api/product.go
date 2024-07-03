package productapi

import (
	"Api/protos/productproto"
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	ClientProduct productproto.ProductsClient
}


func (p *ProductHandler) CreateProduct(c *gin.Context) {
	var req productproto.ProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.CreateProduct(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (p *ProductHandler) GetByIdProduct(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.GetbyIdProduct(ctx, &productproto.CreateProductResponse{Id: int32(Id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (p *ProductHandler) GetAllProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := p.ClientProduct.GetAllProducts(ctx, &productproto.ProductEmpty{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var products []*productproto.Product
	for {
		product, err := stream.Recv()
		if err != nil {
			break
		}
		products = append(products, product)
	}
	c.JSON(http.StatusOK, products)
}

func (p *ProductHandler) UpdateProduct(c *gin.Context) {
	var req productproto.UpdateProductRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.UpdateProduct(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (p *ProductHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	Id, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := p.ClientProduct.DeleteProduct(ctx, &productproto.CreateProductResponse{Id: int32(Id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
