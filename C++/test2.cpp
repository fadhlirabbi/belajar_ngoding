#include <iostream>

using namespace std;

void printAnyting (){ // ini adalah prosedure
    cout << "nilai dari ";
}

int sizeOfByte (int a){
    return sizeof(a);
}

int hitungAngka (int a, int b, int c){
    return a + b + c;
}

unsigned long long faktorial (int a){
    if (a == 1){
        return 1;
    }
    return a * faktorial(a - 1);
}

int main() {
    int angka_a, angka_b, angka_c;
    int bilanganFaktorial;

    cout << "masukkan 3 angka yang ingin ditambahkan : ";
    cin >> angka_a >> angka_b >> angka_c;

    int hitungTotal = hitungAngka(angka_a, angka_b, angka_c);

    printf("hasil dari %d + %d + %d adalah : %d\n", angka_a, angka_b, angka_c, hitungTotal);

    printf("byte dari %d adalah : %d\n", angka_a, sizeOfByte(angka_a));
    
    cout << "masukkan bilangan yang ingin difaktorkan : ";
    cin >> bilanganFaktorial;

    cout << "hasil faktor adalah : " << faktorial(bilanganFaktorial) << endl;

    for (int i = 1;  i <= 5; i++){
        cout << "ini adalah looping for ke : "<< i << endl;
        }

        int b = 1;
        while (b <= 5){
            cout << "ini adalah while loop ke : " << b << endl;
            b++;
        }

        int c = 1;
        do {
            cout << "ini adalah do while loop ke : " << c << endl;
            c++;
        } while (c <= 5);

        if (angka_a < 5){
            cout << "angka a lebih kecil dari 5" << endl;
        }
        else if (angka_a >= 5 && angka_b <= 10){
            cout << "angka a sama dengan diatas 5 atau sama dengan dibawah 10" << endl;
        }
        else{
        cout << "angka a lebih dari 10" << endl;
    }

    (angka_a % 2 == 0) ? cout << "angka_a genap\n" << endl : cout << "angka_a ganjil\n" << endl;

    int pilihanSwitch;
    cout << "pilih 1-3 : ";
    cin >> pilihanSwitch;

    switch (pilihanSwitch)
    {
    case 1:
        cout << "anda memilih case 1\n";
        break;

    case 2:
        cout << "anda memilih case 2\n";
        break;
    
    case 3:
        cout << "anda memilih case 3\n";
        break;

    default:
        cout << "pilihan anda salah\n";
        break;
    }

    system("pause");
    return 0;
}


    // Menggunakan scanf untuk membaca input dari pengguna
    // printf("Masukkan sebuah angka: ");
    // scanf("%d", &angka);

    // Menampilkan nilai yang diinput oleh pengguna
    // printf("Anda memasukkan: %d\n", angka);