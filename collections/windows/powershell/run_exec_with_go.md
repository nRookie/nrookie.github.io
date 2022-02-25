``` golang
package main

import "os/exec"
import "fmt"
import "bytes"  
func main() {   
ps, _ := exec.LookPath("powershell.exe")
args := []string{"Robocopy.exe", "D:\\", "D:\\ufssmb-zz2zni45" ,"/e"}   
args = append([]string{"-NoProfile", "-NonInteractive"}, args...)   
cmd := exec.Command(ps, args...)

var stdout bytes.Buffer 
var stderr bytes.Buffer 
cmd.Stdout = &stdout
cmd.Stderr = &stderr

err := cmd.Run()
if err != nil { 
fmt.Println(err.Error())
}   

stdOut, stdErr := stdout.String(), stderr.String()  
fmt.Println(stdOut) 
fmt.Println(stdErr) 

}

```