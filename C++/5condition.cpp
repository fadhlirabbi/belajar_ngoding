#include <iostream>
#include <cstdlib>
#include <conio.h>

using namespace std;

main(){
    int x;
    cin >>x;

    //if else statement
    
    if(x>0 && x<5){
        cout<<"x lebih besar dari 0 dan kurang dari lima"<<endl;
    }
    else if (x>=5 && x<10)
    {
        cout<<"x lebih besar sama dengan dari 5 dan kurang dari 10"<<endl;
    }
    else{
        cout<<"x diluar 1-9"<<endl;
    }

    //?: operator
    (x%2==0)?cout<<"genap"<<endl:cout<<"ganjil"<<endl;
    

    //switch case
    switch (x){
        case 1 :
            cout<<"satu"<<endl;
            break;
        case 2 :
            cout<<"dua"<<endl;
            break;
        case 3 :
            cout<<"tiga"<<endl;
            break;
        case 4 :
            cout<<"empat"<<endl;
            break;
        case 5 :
            cout<<"lima"<<endl;
            break;
        case 6 :
            cout<<"enam"<<endl;
            break;
        case 7 :
            cout<<"tujuh"<<endl;
            break;
        case 8 :
            cout<<"delapan"<<endl;
            break;
        case 9 :
            cout<<"sembilan"<<endl;
            break;
        default:
            cout<<"angka bukan 1-9";
            break;

    getch();
    }
}