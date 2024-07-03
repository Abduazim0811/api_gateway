package orderapi

import (
	order "Api/protos/orderproto"
	"context"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	ClientOrder order.OrderServiceClient
}

func (o *OrderHandler) CreateOrder(c *gin.Context) {
	var req order.Request
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := o.ClientOrder.CreateOrder(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (o *OrderHandler) CreateOrders(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	stream, err := o.ClientOrder.CreateOrders(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	sendOrder := func(req *order.Request) error {
		if err := stream.Send(req); err != nil {
			return err
		}
		return nil
	}

	receiveResponse := func() (*order.Response, error) {
		res, err := stream.Recv()
		if err != nil {
			return nil, err
		}
		return res, nil
	}

	var reqs []*order.Request
	if err := c.BindJSON(&reqs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, req := range reqs {
		if err := sendOrder(req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	stream.CloseSend()
	var responses []*order.Response
	for {
		res, err := receiveResponse()
		if err == io.EOF {
			break
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		responses = append(responses, res)
	}

	c.JSON(http.StatusOK, responses)
}

func (o *OrderHandler) GetAllOrders(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	req := &order.OrderRequest{}

	res, err := o.ClientOrder.GetAllOrders(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res.Order)
}
