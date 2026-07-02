-- 1. 테이블 생성
CREATE TABLE IF NOT EXISTS posts (
     id SERIAL PRIMARY KEY,
     title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
                             );

-- 2. 테스트 데이터 삽입
INSERT INTO posts (title, content, created_at) VALUES
('첫 번째 게시글', 'Go 언어와 PostgreSQL을 연동한 게시판입니다.', NOW()),
('두 번째 게시글', 'Docker Compose로 DB를 띄우니 정말 편하네요.', NOW()),
('세 번째 게시글', '스프링 개발자의 Go 적응기!', NOW());