#include <iostream>
#include <conio.h>
using namespace std;

main(){
    int x  = 214125135131415;// << ilegal kerena value melebihi batas range value yang dapat disimpan

    cout<<"besar memory x = "<<sizeof(x)<<" byte";

    char y = 'B' ;

    cout<<"\nbesar memory x = "<<sizeof(y)<<" byte";

    float z = 2.5;

    cout<<"\nbesar memory x = "<<sizeof(z)<<" byte";

    getch();
}