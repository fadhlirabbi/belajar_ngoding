#include <iostream>
#include <string>

using namespace std;
class Buku {
private:
    string judul;
    string penulis;
    
public:
    Buku(string j, string p) : judul(j), penulis(p) {}

    void tampilkanInfo() {
        cout << "Judul Buku: " << judul << "\nPenulis: " << penulis << endl;
    }
};

class Ebook : public Buku {
private:
    double ukuranFile; // dalam MB

public:
    Ebook(string j, string p, double ukuran) : Buku(j, p), ukuranFile(ukuran) {}

    void tampilkanInfo() {
        Buku::tampilkanInfo();
        cout << "Ukuran File: " << ukuranFile << " MB" << endl;
    }
};

void cetakInfoBuku(Buku &buku) {
    buku.tampilkanInfo();
}

int main() {

    Buku buku1("Belajar C++", "John Doe");
    Ebook ebook1("Belajar C++ Edisi Digital", "John Doe", 5.2);

    cetakInfoBuku(buku1);
    cetakInfoBuku(ebook1);  // Polimorfisme di sini!

    getchar();
    return 0;
}
