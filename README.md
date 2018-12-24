# Logger
Easy and convenient logging method when developing projects in the Golang language

--------------------------
Install 

go get github.com/cartmanis/logger.git

--------------------------
Usage

 Для использования необходимо инициализировать логер. Существует два способа инициализации:    
 * Инициализация основного логера, для использования его в разных частях проекта     
```go
    func main() {
    	if err := logger.NewMainLogger("mainLogger.log", false, true); err != nil {
    		fmt.Prinln(err)
    	}
    	defer logger.Close()
    	logger.LogInfo("Лог уровня Info")
    	logger.LogWarn("Лог уровня Warn")
    	logger.LogError("Лог уровня Error")
    }
```
 *  Инициализация вспомогательного логера для решения определенной задачи.
```go
    func test() {
    	someLogger, err := logger.NewLogger("someLogger.log", false, true)
    	if err != nil {
    		fmt.Printf(err)
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
Если передать false, false, то будет вывод в консоль только лога уровня Error - LogError()
При передачи третьего параметра, отвечающего за вывода в файл, обязательно необходимо передать путь до файла.


