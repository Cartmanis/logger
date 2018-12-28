# Logger
Easy and convenient logging method when developing projects in the Golang language

--------------------------
Install 

go get github.com/cartmanis/logger.git

--------------------------
Usage

 To use, you must initialize the logger. There are two ways to initialize:    
 * Initialization of the main logger for use in different parts of the project     
```go
    func main() {
    	if err := logger.NewMainLogger("mainLogger.log", false, true); err != nil {
    		fmt.Prinln(err)
    	}
    	defer logger.Close()
    	logger.Info("Info level log")
    	logger.Warn("Warn level log")
    	logger.Error("Error level log")
    }
```
 * Initializing auxiliary logger to solve a specific task.
```go
    func test() {
    	someLogger, err := logger.NewLogger("someLogger.log", false, true)
    	if err != nil {
    		fmt.Printf(err)
    		return
    	}
    	defer someLogger.Close()
    	someLogger.Info("Additional level log Info")
    	someLogger.Warn("Additional level log Warn")
    	someLogger.Error("Additional level log Error")
    }
```
------------------------
Additionally

The first parameter is responsible for the path to the file. The second parameter is the switch output log to the console.
The third parameter is a switch to output the log to a file.
If the second and third parameters are false, false, then there will be output only to the Error level log console
When passing the third parameter, it is imperative that you pass the path to the file.


