package post

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) GetPosts(c *gin.Context) {
	posts, err := h.repo.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "데이터 조회 실패"})
		return
	}
	c.JSON(http.StatusOK, posts)
}

func (h *Handler) GetPost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID 형식"})
		return
	}

	post, err := h.repo.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "게시글을 찾을 수 없음"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (h *Handler) CreatePost(c *gin.Context) {
	var p Post
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.repo.Save(&p); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "게시글 저장 실패"})
		return
	}
	c.JSON(http.StatusCreated, p)
}

func (h *Handler) DeletePost(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "잘못된 ID 형식"})
		return
	}

	if err := h.repo.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "게시글 삭제 실패"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "삭제 완료"})
}

// HTML 화면을 보여주는 메인 페이지 (GET /web/posts)
func (h *Handler) IndexView(c *gin.Context) {
	posts, err := h.repo.FindAll()
	if err != nil {
		c.String(http.StatusInternalServerError, "데이터 조회 실패")
		return
	}
	// 스프링의 ModelAndView처럼 index.html에 posts 데이터를 담아 렌더링합니다.
	c.HTML(http.StatusOK, "index.html", gin.H{
		"posts": posts,
	})
}

// 폼 서브밋을 처리하는 메서드 (POST /web/posts)
func (h *Handler) CreatePostForm(c *gin.Context) {
	// JSON 바인딩 대신 Form 데이터를 가져옵니다.
	title := c.PostForm("title")
	content := c.PostForm("content")

	p := Post{Title: title, Content: content}
	if err := h.repo.Save(&p); err != nil {
		c.String(http.StatusInternalServerError, "저장 실패")
		return
	}
	// 성공 시 메인 화면으로 리다이렉트
	c.Redirect(http.StatusMovedPermanently, "/web/posts")
}

// 웹 브라우저용 삭제 요청 처리 (GET /web/posts/delete?id=X)
func (h *Handler) DeletePostForm(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "잘못된 ID")
		return
	}

	if err := h.repo.Delete(id); err != nil {
		c.String(http.StatusInternalServerError, "삭제 실패")
		return
	}
	c.Redirect(http.StatusMovedPermanently, "/web/posts")
}
