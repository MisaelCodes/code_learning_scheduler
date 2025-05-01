package components

import "os/exec"


type Action interface {
    Execute() error
    Abort() error
    Describe() string
}

type FileOpener struct {
    path string
}

func NewFileOpener(path string) *FileOpener{
    return &FileOpener{path}
}

func (fo *FileOpener) Execute() error{
    cmd := exec.Command(fo.path)
    if err := cmd.Run();err != nil{
        return err
    }
    return nil
}

type CodeEditingSession struct{
    name string
    launcher string
    path string
    venvPath string
}

func NewCodeEditingSession(name, launcher, path, venvPath string) *CodeEditingSession{
    return &CodeEditingSession{name, launcher, path, venvPath}
}

type Browser struct{
    name string
    laucher string
    url string
}
