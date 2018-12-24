# Logger
Easy and convenient logging method when developing projects in the Golang language

--------------------------
Install 

go get https://github.com/Cartmanis/logger.git

--------------------------
Usage

Для использования необходимо инициализировать логер. Существуют два способа инициализации:
    1.) Инициализация основного логера, для использования его в разных частях проекта
    Пример: 
```go
    func main() {
    	if err := logger.NewMainLogger("mainLogger.log", false, true); err!= nil {
    		fmt.Printf("ERROR: при инициализации logger. Текст ошибки: ", err)
    	}
    	defer logger.Close()
    	logger.LogInfo("Лог уровня Info")
    	logger.LogWarn("Лог уровня Warn")
    	logger.LogError("Лог уровня Error")
    }
```
   2.) Инициализация вспомогательного логера для решения определенной задачи.
    Пример:
```go
    func test() {
    	someLogger, err := logger.NewLogger("someLogger.log", false, true)
    	if err != nil {
    		fmt.Printf("ERROR: при инициализации logger. Текст ошибки: ", err)
    		return
    	}
    	defer someLogger.Close()
    	someLogger.LogInfo("Дополнительный лог уровня Info")
    	someLogger.LogWarn("Дополнительный лог уровня Warn")
    	someLogger.LogError("Дополнительный лог уровня Error")
    }
```
------------------------
Дополнительно

Первый параметр отвечает за путь до файла. Второй параметр вывод лога в консоль. Третий за вывод лога в файл.
Если передать false, false, то будет вывод в консоль только ошибочных записей.
При передачи третьего параметра, отвечающего за вывода в консоль, обязательно необъодимо передавать путь до файла.


