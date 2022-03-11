/*
   @Author:huolun
   @Date:2022/3/11
   @Description
*/
package auth

import (
	"log"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	token := Sign(12, "ddddd", "nika", "www.who.com", time.Minute*60, 0)
	log.Println(token)
	time.Sleep(time.Second * 2)
	i, err := Parse(token, "ddddd")
	log.Println(i)
	log.Println(err)
}
