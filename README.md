# go_pointer_vs_value_demo

## introduction 

    This is a demo show about & operator and * operator  how to use and how to work

## import Concept to use goroutine while with pointer reference

type email struct
```golang===
  type email struct {
	from string
	to string
}
```
while use value setter, need return orignal value
```golang===
func(e email) From(s string) email{
	e.from = s
	return e
}
func(e email) To(s string) email{
	e.to = s
	return e
}
```
use the return struct Object to reference
```golang===
  
    e := &email{}
	for i:=1; i<=10; i++ {
		go func(i int) {
			// e := &email{}
			e.From(fmt.Sprintf("example%02d@example.com", i)).To(fmt.Sprintf("example%02d@example.com", i+1)).Send()
			// e.To(fmt.Sprintf("example%02d@example.com", i+1))
			// e.Send()
		}(i)
	}
```