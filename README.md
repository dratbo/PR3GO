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

## Доп задания конечно же без GPT (нет) :sweat_smile:

CORS (минимально): добавить заголовки Access-Control-Allow-Origin: * для GET/POST (в отдельной middleware).

<img width="1862" height="676" alt="image" src="https://github.com/user-attachments/assets/0a1b9c37-78fd-4675-8199-fdbade632e26" />

<img width="1431" height="931" alt="image" src="https://github.com/user-attachments/assets/d743164a-8d0e-4b23-8fbc-613abc50a702" />

<img width="1432" height="945" alt="image" src="https://github.com/user-attachments/assets/02403fe6-11b8-4049-a22e-c02638bf6f14" />

Валидация длины title (например, 1…140 символов).

<img width="1813" height="276" alt="image" src="https://github.com/user-attachments/assets/6cdd84ba-5cc7-4cba-9bfb-984457dab61f" />

Метод PATCH /tasks/{id} для отметки Done=true.

<img width="1434" height="941" alt="image" src="https://github.com/user-attachments/assets/a738c658-6c62-480a-9d9d-39010ad1f0ac" />

Метод DELETE /tasks/{id}`.

<img width="1432" height="946" alt="image" src="https://github.com/user-attachments/assets/01586c7b-5b5d-4b2c-bdff-b14c22e9dfe5" />

Юнит-тесты обработчиков с httptest.

<img width="1871" height="745" alt="image" src="https://github.com/user-attachments/assets/1af022f4-07b0-4284-b81d-a98ef30902cf" />
