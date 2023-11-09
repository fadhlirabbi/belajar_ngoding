#include <iostream>

using namespace std;

int a = 10;

int tambah (){
    return a + 10;
}

void kata (){
    cout << "hasil dari a + 10 adalah : ";
}

int main (){

    int angka;

    kata();
    cout << tambah() << endl;

    cout << "\nmasukkan angka : ";
    cin >> angka;

    cout << "angka yang anda masukkan adalah : " << angka << endl;

    cin.get();

    return 0;
}