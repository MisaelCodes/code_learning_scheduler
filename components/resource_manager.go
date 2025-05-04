package components

import "fmt"

// the resource manager have functions that help with the crud
// operations rendering, this can be thought like the forms

// first we need an interface that takes a resource
// we should be able to have the fields and the type of each field

type ResourceRenderer interface{
    Fields() ([][2]string, error)
}


func CreateForm(r ResourceRenderer) error{
    // clear the current menu
    // show a message that tells what resource is
    // going to be created
    fields, err := r.Fields()
    if err != nil{
        return err
    }
    for _, field := range fields{

        fmt.Println(field)
    }
    return nil
}
