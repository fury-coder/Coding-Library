# Intro
Golang is a procedural and statically typed programming language having the syntax similar to C programming language. Sometimes it is termed as Go Programming Language. It provides a rich standard library, garbage collection, and dynamic-typing capability. It was developed in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson at Google but launched in 2009 as an open-source programming language and mainly used in Google’s production systems. Golang is one of the most trending programming languages among developers.
* Advantages:
    - Flexible- It is concise, simple and easy to read.
    - Concurrency- It allows multiple process running simultaneously and effectively.
    - Quick Outcome- Its compilation time is very fast.
    - Library- It provides a rich standard library.
    - Garbage collection- It is a key feature of go. Go excels in giving a lot of control over memory allocation and has dramatically reduced latency in the most recent versions of the garbage collector.
    - It validates for the interface and type embedding.

* Disadvantages: 
    - It has no support for generics, even if there are many discussions about it.
    - The packages distributed with this programming language is quite useful but Go is not so object-oriented in the conventional sense.
    - There is absence of some libraries especially a UI tool kit.

# How to Install
## Windows
Head to https://golang.org/doc/install and download the file prompted
1. Open the MSI file you downloaded and follow the prompts to install Go.
    By default, the installer will install Go to Program Files or Program Files (x86). You can change the location as needed. After installing, you will need to close and reopen any open command prompts so that changes to the environment made by the installer are reflected at the command prompt.
2. Verify that you've installed Go.
    1. In Windows, click the Start menu.
    2. In the menu's search box, type cmd, then press the Enter key.
    3. In the Command Prompt window that appears, type the following command:
        `$ go version`
    4. Confirm that the command prints the installed version of Go.

## Mac
Head to https://golang.org/doc/install and download the file prompted
1. Open the package file you downloaded and follow the prompts to install Go.
    The package installs the Go distribution to /usr/local/go. The package should put the /usr/local/go/bin directory in your PATH environment variable. You may need to restart any open Terminal sessions for the change to take effect.
2. Verify that you've installed Go by opening a command prompt and typing the following command:
    `$ go version`
3. Confirm that the command prints the installed version of Go.

## Linux
1. Use your distro's Package manager to download the `go` package
    - For example if you use Arch you would do
    `sudo pacman -S go`
