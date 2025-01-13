package main

import "fmt"
import t "time"

func main() {
  gopher := `              .-::::::-. 
          .:-::::::::::::::-:.
          _:::    ::    :::_ 
          .:( ^   :: ^   ):.
           :::   (..)   :::. 
           :::::::UU:::::::
          .::::::::::::::::.
          O::::::::::::::::O
          -::::::::::::::::-
           ::::::::::::::::
           .::::::::::::::.
             oO:::::::Oo`
              
  time := t.Now()
            
  fmt.Printf(gopher)
  fmt.Println()
  fmt.Println(time.Format("01-02-2006"))

}
