package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var (
	// test용 시크릿 키
	secretKey = "JWT_SECRET_KEY"

	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients = make(map[*websocket.Conn]*ClientInfo)
	mutex   sync.Mutex
)

// Message 구조체 정의
type Message struct {
	User    string `json:"user"`
	Content string `json:"content"`
	Token   string `json:"token"`
}

// ClientInfo 구조체 정의
type ClientInfo struct {
	Conn        *websocket.Conn
	Index       int
	ConnectedAt time.Time
}

func HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		conn.Close()
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
	}()

	mutex.Lock()
	client := &ClientInfo{
		Conn:        conn,
		Index:       len(clients) + 1, // 현재 클라이언트 수 + 1을 인덱스로 사용
		ConnectedAt: time.Now(),
	}
	clients[conn] = client
	mutex.Unlock()

	for {
		// 웹 소켓에서 메시지 읽기
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		// JSON 형식의 메시지를 Message 구조체로 파싱
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			fmt.Println(err)
			return
		}

		tokenString := message.Token

		token, err := jwt.Parse(strings.Split(tokenString, " ")[1], func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		if err != nil || !token.Valid {
			fmt.Println("유효하지 않은 토큰")
			if err := conn.WriteJSON(gin.H{"error": err.Error()}); err != nil {
				fmt.Println(err)
			}
			return
		}

		// 받은 메시지를 다른 클라이언트들에게 전송
		mutex.Lock()
		for _, info := range clients {
			if info.Conn != conn {
				if err := info.Conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					fmt.Println(err)
				}
			}
		}

		// 자기 자신에게도 메시지 전송
		if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
			fmt.Println(err)
		}
		mutex.Unlock()
	}
}
