package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	company_gin_adapter_in "github.com/limiu82214/CleanArchitectureHexagonal/internal/company/adapter/in/gin"
	"github.com/limiu82214/CleanArchitectureHexagonal/pkg/sig"
)

func main() {
	fmt.Println("Hello")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r := gin.Default()
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	{ // 初始化
		v1 := r.Group("/")
		{
			NCGA := company_gin_adapter_in.NewCompanyGinAdapter()
			v1.GET("get_siteinfo", NCGA.GetSiteInfo)
			v1.GET("ping", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
		}
		go (func() {
			if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
				log.Printf("listen: %s\n", err)
			}
			r.Run(":8080")
		})()
	}

	sig.ServerNotify()
	log.Println("伺服器開始關閉...")

	{ // 關閉服務
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalln("伺服器錯誤退出:", err)
		}
	}
	log.Println("伺服器正常運行結束")
}
