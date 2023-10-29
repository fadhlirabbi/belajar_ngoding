#include <iostream> //header file 
#include <cstdio>

// global scope


using namespace std; //namespace

int globalvar = 7; 

//prosedure
void foo1(){
    //local scope 1 

    cout<<"ini adalah prosedure\n";
    cout<<"gobal variable :"<< globalvar;
}

//fuction
int foo2(){
    //local scope 2

    cout<<"\nini adalah fungsi nilai kembali adalah : ";
    return 5;
}

main(){

    //local scope 3

    int localvar;
    foo1();
    localvar = foo2();
    cout<<localvar;

    std::getchar();
}