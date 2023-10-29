#include <iostream>
#include <cstdlib>
#include <conio.h>

using namespace std;

main(){
    int x = 10;
    int *ptr = &x; // ptr adalah pointer yang menunjuk ke alamat memori dari x
    
    cout << "Alamat dari x: " << &x << endl;
    cout << "Nilai yang ditunjuk oleh ptr: " << *ptr <<endl;

    int y = 10;
    int &ref = y; // ref adalah reference untuk x

    cout << "Nilai dari x: " << y << std::endl;
    cout << "Nilai dari ref: " << ref << std::endl;

    ref = 20; // mengubah x melalui ref
    cout << "Nilai dari x setelah diubah: " << x << std::endl;
}