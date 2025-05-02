package components

type ResourceManager interface{
    Create() *ResourceManager
    Retrieve() *ResourceManager
    Update() *ResourceManager
    Delete() error
}


