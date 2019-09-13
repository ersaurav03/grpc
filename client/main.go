package main

import (
	"JSMPJ_grpc/proto"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("This is client")
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := proto.NewAddServiceClient(conn)
	g := gin.Default()
	g.GET("/Add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater B"})
			return
		}
		req := &proto.Request{Num1: int64(a), Num2: int64(b)}
		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	})
	// g.GET("/Sub/:a/:b", func(ctx *gin.Context) {
	// 	a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater A"})
	// 		return
	// 	}

	// 	b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
	// 	if err != nil {
	// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater B"})
	// 		return
	// 	}
	// 	req := &proto.Request{Num1: int64(a), Num2: int64(b)}
	// 	if response, err := client.Sub(ctx, req); err == nil {
	// 		ctx.JSON(http.StatusOK, gin.H{
	// 			"result": fmt.Sprint(response.Result),
	// 		})
	// 	} else {
	// 		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// })
	g.GET("/Mul/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater B"})
			return
		}
		req := &proto.Request{Num1: int64(a), Num2: int64(b)}
		if response, err := client.Mul(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	})

	g.GET("/Div/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater A"})
			return
		}

		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Paramater B"})
			return
		}
		req := &proto.Request{Num1: int64(a), Num2: int64(b)}
		if response, err := client.Div(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
