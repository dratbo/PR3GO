<h1 align="center"> Привет! Я <a target="_blank"> Кармеев Артур из группы ЭФМО-01-25 </a> 
<img src="https://github.com/blackcater/blackcater/raw/main/images/Hi.gif" height="32"/></h1>
<h3 align="center"> Данная практика была выполнена со всеми доп заданиями и нервами!</h3>

## Структура проекта

    └── pz3-http/
        ├── go.mod
        ├── .idea/
        │   ├── .gitignore
        │   ├── modules.xml
        │   ├── pz3-http.iml
        │   └── workspace.xml
        ├── internal/
        │   ├── storage/
        │   │   └── memory.go
        │   └── api/
        │       ├── handlers.go
        │       ├── handlers_test.go
        │       ├── middleware.go
        │       └── responses.go
        └── cmd/
            └── server/
                └── main.go

## Поднимаем HTTP-сервер

GET /health

<img width="974" height="778" alt="image" src="https://github.com/user-attachments/assets/91eab57f-86cb-4609-b3da-af5606f7b24e" />

POST /tasks (создание)

<img width="974" height="134" alt="image" src="https://github.com/user-attachments/assets/317ee0aa-cb28-41d3-ba84-dbf7d41c9fab" />

GET /tasks (список)

<img width="974" height="819" alt="image" src="https://github.com/user-attachments/assets/f7e382ad-67f3-46df-87f5-a95cc65d9054" />

<img width="974" height="806" alt="image" src="https://github.com/user-attachments/assets/2352bfe5-c040-4677-914b-4eedd6046dcd" />

GET /tasks/{id}

<img width="974" height="810" alt="image" src="https://github.com/user-attachments/assets/b12e07c3-8732-4f3a-af41-942cd7dae90e" />

