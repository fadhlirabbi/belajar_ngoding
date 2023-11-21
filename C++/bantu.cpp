#include <iostream>

using namespace std;

int main() {
    int baris = 3, kolom = 3;

    // Deklarasi array 2D dengan ukuran dinamis
    int** array = new int*[baris];
    for (int i = 0; i < baris; i++) {
        array[i] = new int[kolom];
    }

    // Meminta input dari user untuk nilai array
    cout << "Masukkan nilai array:\n";
    for (int i = 0; i < baris; i++) {
        for (int j = 0; j < kolom; j++) {
            cin >> array[i][j];
        }
    }

    // Menampilkan nilai array
    cout << "Nilai array:\n";
    for (int i = 0; i < baris; i++) {
        for (int j = 0; j < kolom; j++) {
            cout << array[i][j] << " ";
        }
        cout << "\n";
    }

    // Dealokasi memori
    for (int i = 0; i < baris; i++) {
        delete[] array[i];
    }
    delete[] array;

    return 0;
}
