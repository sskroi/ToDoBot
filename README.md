# Description
Бот предоставляет реализацию простого ToDo списка в телеграм боте. С помощью бота можно добавить задачу в свой список задач, смотреть невыполненные задачи и время до их дедлайна, отмечать задачи как выполненные, просматривать архив выполненных задач с датой их выполнения, удалять задачи.

# Launch
* `docker build -t todobot {path/to/project/directory}`
* `docker run --rm -e TG_TOKEN={your telegtam bot token} todobot:latest`

or
* `docker run --rm -e TG_TOKEN={your telegtam bot token} sskroi/todobot:0.1`

# Mount database directory
Используйте `-v {path/to/your/database/dir}:/app/database`
когда запускаете контейнер с помощью `docker run` для сохранения файла с базой данных при перезапуске контейнера.
#### Example
`docker run --rm -e TG_TOKEN=00000:XXXXXX -v ./database:/app/database`

# Demo
![Demo start](https://github.com/sskroi/ToDoBot1/blob/master/demo/images/start.png?raw=true)
![Demo adding](https://github.com/sskroi/ToDoBot1/blob/master/demo/images/adding.png?raw=true)
![Demo completing](https://github.com/sskroi/ToDoBot1/blob/master/demo/images/completing.png?raw=true)
![Demo menu](https://github.com/sskroi/ToDoBot1/blob/master/demo/images/menu.png?raw=true)