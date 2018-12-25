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
    	logger.Info("Лог уровня Info")
    	logger.Warn("Лог уровня Warn")
    	logger.Error("Лог уровня Error")
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
    	someLogger.Info("Дополнительный лог уровня Info")
    	someLogger.Warn("Дополнительный лог уровня Warn")
    	someLogger.Error("Дополнительный лог уровня Error")
    }
```
------------------------
Дополнительно

Первый параметр отвечает за путь до файла. Второй параметр вывод лога в консоль. Третий за вывод лога в файл.
Если передать false, false, то будет вывод в консоль только лог уровня Error - Error()
При передачи третьего параметра, отвечающего за вывода в файл, обязательно необходимо передать путь до файла.


