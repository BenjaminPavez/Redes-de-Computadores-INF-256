package main

import (
	"os"
	"net"
)

func main(){
   // Pruebas libreria net
   strUDP := ":1800"   

   str, err := net.ResolveUDPAddr("udp4", strUDP)
   if err != nil {
      os.Exit(1)
   }

   udp, err := net.ListenUDP("udp", str)
   defer udp.Close()
   if err != nil {
      os.Exit(2)
   }

   for{
      var message [256]byte
      n, addr, err := udp.ReadFromUDP(message[0:])
      if err != nil {
         continue
      }
      udp.WriteToUDP(message[0:n], addr)
   }
}
