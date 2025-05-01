# Coding tasks CLI

Creates todo list of task, each task needs to have a name,
a time in the day to start it and an optional file for configuration for automation.

The file for configuration can be something like a bash file 
that opens a terminal ready to code with your tmux session 
an your vim editor in the file that you want.

When time parameter execute the given configuration in that moment
if a file was given.

When all tasks are finished it asks you if you want different task for tomorrow.
you can add breaks to it to be set at a certain time.

[*] I need to interact with the command line for control, color and text insertion
[ ] I need to create the tasks and tasks list logic
[ ] I need to figure out how to reproduce audio
[ ] I need to figure out how to connect a sqllite
[ ] I need to create test
[ ] Avoid chatgpt and AI in general.


First Stage

Getting to know ansi and familiarization with the golang's term package,
we are gonna create an application that is both interactive and commandable.

First I need sort of ansi handlers, some for cursor control, text enhancement and coloring.
