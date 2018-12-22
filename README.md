# Logger
Easy and convenient logging method when developing projects in the Golang language

--------------------------
Install: go get https://github.com/Cartmanis/logger.git

--------------------------
Usage

Для использования необходимо инициализировать логер. Существуют два способа инициализации:
    1.) Инициализация основного логера, для использования его в разных частях проекта
    Пример: 
```go
    func main() {
    	if err := logger.NewMainLogger("mainLogger.log", false); err!= nil {
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
    	someLogger, err := logger.NewLogger("someLogger.log", false)
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

Если вторым параметром передать true, то кроме записи лога в файл будет вывод лога в консоль.
Также существует возможно вывода только в консоль, без записи в файл
Для этого необходимо изменить значение OutToFile в false 
Пример: 
```go 
    func main() {
        logger.NewMainLogger("mainLogger.log", true)
        logger.OutToFile = false
    } 
```

