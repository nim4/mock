# mock
net.Conn mock for Testing


ex:
```
func TestX(t *testing.T) {
	var s mock.ConnMock
	err := X(s)
	if err != nil {
		t.Fatal(err)
	}

	mock.ReturnError(true)
	defer mock.ReturnError(false)
  err = X(s)
	if err == nil {
		t.Fatal("Expected error")
	}
}
```
