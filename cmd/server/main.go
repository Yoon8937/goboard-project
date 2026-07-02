package main

import (
	"awesomeProject/internal/database"
	"awesomeProject/internal/post"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewPostgresDB()
	if err != nil {
		log.Fatalf("DB 연결 에러: %v", err)
	}

	postRepo := post.NewRepository(db)
	postHandler := post.NewHandler(postRepo)

	r := gin.Default()

	// ⭐ 중요: HTML 템플릿 폴더 위치 설정 (Spring의 Thymeleaf 설정과 유사)
	r.LoadHTMLGlob("templates/*")

	// 기존 REST API 주소들
	r.GET("/posts", postHandler.GetPosts)
	r.GET("/posts/:id", postHandler.GetPost)
	r.POST("/posts", postHandler.CreatePost)
	r.DELETE("/posts/:id", postHandler.DeletePost)

	// 🌐 [추가] 웹 브라우저용 주소들
	r.GET("/web/posts", postHandler.IndexView)
	r.POST("/web/posts", postHandler.CreatePostForm)
	r.GET("/web/posts/delete", postHandler.DeletePostForm)

	log.Println("서버 시작 (:8080)... 브라우저에서 http://localhost:8080/web/posts 로 접속하세요!")
	r.Run(":8080")
}
