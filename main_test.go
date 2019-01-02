package main

import "testing"

func TestReadFile(t *testing.T) {
  result := readFile("test1.xml")

  if(result.Data[0] != first_test{"Point 1", "3", "5"}){
    t.Errorf("Got incorrect value back")
  }
  if(result.Data[1] != first_test{"Point 2", "7", "9"}){
    t.Errorf("Got incorrect value back")
  }
  if(result.Data[2] != first_test{"Point 3", "11", "13"}){
    t.Errorf("Got incorrect value back")
  }
  if(len(result.Data)!=3){
    t.Error("Got too many results back")
  }
}
