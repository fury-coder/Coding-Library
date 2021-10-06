introduction to c

About C Language

C is a very powerful and widely used language. It is used in many scientific programming situations. It forms (or is the basis for) the core of the modern languages Java and C++. It allows you access to the bare bones of your computer. 

Our 1st Hello World program 

#include <stdio.h>
int main()
{
    printf("Hello World");
}

How accually its work??

The #include is a preprocessor command that tells the compiler to include the contents of stdio.h (standard input and output) file in the program.
The stdio.h file contains functions such as scanf() and printf() to take input and display output respectively.
If you use the printf() function without writing #include <stdio.h>, the program will not compile.
The execution of a C program starts from the main() function.
printf() is a library function to send formatted output to the screen. In this program, printf() displays Hello, World! text on the screen.
The return 0; statement is the "Exit status" of the program. In simple terms, the program ends with this statement.
