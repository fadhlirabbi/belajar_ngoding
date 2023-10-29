#include <iostream>
#include <cstdlib>
#include <conio.h>

using namespace std;

int foo1(){
    cout<<"ini adalah sebuah fucntion kembalikan nilai :";
    return 7;
}

void foo2(){
    cout<<"ini adalah procedure!!"<<endl;
}

//overloading function
int foo1(int a){
    cout<<"ini adalah overloading function kembalikan nilai :";
    return a;
}

//overloading procedure
void foo2(int a){
    cout<<"ini adalah overloading procedure nilai parameter adalah : "<<a<<endl;
}

/*
overloading adalah sebuah cara untuk membuat sebuah fungsi atau procedur memiliki berbagai bentuk/ algorithma
tergantung dari parameter yang dideklarasi dari main.
*/

//recursion 
unsigned long long faktorial(int n) {
    if (n <= 1) {   // basis rekursi
        return 1;
    }
    return n * faktorial(n - 1);  // langkah rekursi
}

/*recursion adalah sebuah teknik agar sebuah fungsi dapat memanggil dirinya sendiri yang mengakibatkan
program akan berjalan "looping" sampai waktu yang ditentukan*/

main(){
    int x = 10;
    int y;

    y = foo1();
    cout<<y<<endl;

    foo2();

    y = foo1(x);
    cout<< y<<endl;

    foo2(x);

    unsigned long long f = faktorial(x);
    cout<<"faktorial "<<x<<" adalah "<<f;
}