#include <iostream>
#include <string>

class Buku {
private:
    std::string judul;
    std::string penulis;
    
public:
    Buku(std::string j, std::string p) : judul(j), penulis(p) {}

    void tampilkanInfo() {
        std::cout << "Judul Buku: " << judul << "\nPenulis: " << penulis << std::endl;
    }
};



class Ebook : public Buku {
private:
    double ukuranFile; // dalam MB

public:
    Ebook(std::string j, std::string p, double ukuran) : Buku(j, p), ukuranFile(ukuran) {}

    void tampilkanInfo() {
        Buku::tampilkanInfo();
        std::cout << "Ukuran File: " << ukuranFile << " MB" << std::endl;
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

    return 0;
}
