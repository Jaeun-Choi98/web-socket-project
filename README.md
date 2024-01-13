# Web Socket Project - [링크](http://175.45.205.116/chat)

실시간 소통을 위한 웹 채팅 기능 구현(웹 프로젝트[dohyeong](https://github.com/Jaeun-Choi98/dohyeong?tab=readme-ov-file)의 side project)

<br>

## 계획

- 추가적인 기능 개발과 배포 과정의 편리성을 위해 주 프로젝트는 프론트 엔드와 연동하여 사용자에게 완성된 서비스를 제공하고, 사이드 프로젝트는 도커 컨테이너로 구현된 독립적인 백엔드로 서비스

<br>

## 설계

- 프론트 엔드 부분은 [dohyeong](https://github.com/Jaeun-Choi98/dohyeong?tab=readme-ov-file)에서 개발하고, 백엔드 부분은 다른 작업 트리에서 개발
- 도커를 활용하여 배포 및 서비스

<br>

## 개발

### Frontend

- 리액트를 이용한 채팅방 페이지
- 클라이언트에서 데이터를 보낼 시, jwt도 함께 송신

<br>

### Backend

- gorilla/websocket 패키지를 이용한 웹 소켓 코드 작성
- gin 프레임워크를 이용한 서버(라우터) 구축
- 서버에서는 데이터 수신을 하고, jwt 검증 후 연결된 클라이언트에게 모두 데이터 송신

<br>

## 배포

- 서버는 web-project의 application-server 사용
- 도커파일과 GitAction을 이용한 서버에 배포
