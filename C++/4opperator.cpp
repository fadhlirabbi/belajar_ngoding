#include <iostream>
using namespace std;

int main() {
    int a, b;
    a = 7;
    b = 2;

    //arithmatic operator
    cout<<"arithmatic operator : \n";
    cout << "a + b = " << (a + b) << endl;
    cout << "a - b = " << (a - b) << endl;
    cout << "a * b = " << (a * b) << endl;
    cout << "a / b = " << (a / b) << endl;
    cout << "a % b = " << (a % b) << endl;

    //assignment operator
    cout << endl;
    cout<<"assignment operator : \n";
    a = 7;
    b = 2;    
    cout << "a = b " << (a = b) << endl;
    a = 7;
    b = 2;
    cout << "a += b " << (a += b) << endl;
    a = 7;
    b = 2;
    cout << "a -= b " << (a -= b) << endl;
    a = 7;
    b = 2;
    cout << "a *= b " << (a *= b) << endl;
    a = 7;
    b = 2;
    cout << "a /= b " << (a /= b) << endl;
    a = 7;
    b = 2;
    cout << "a %= b " << (a %= b) << endl;
    a = 7;
    b = 2;
    cout << "a++ " << (a++) << endl;
    a = 7;
    b = 2;
    cout << "b-- " << (b--) << endl;

    //relational operator
    cout << endl;
    cout<<"relational operator : \n";
    a = 3;
    b = 5;
    bool result;

    result = (a == b);   // false
    cout << "3 == 5 is " << result << endl;
    result = (a != b);  // true
    cout << "3 != 5 is " << result << endl;
    result = a > b;   // false
    cout << "3 > 5 is " << result << endl;
    result = a < b;   // true
    cout << "3 < 5 is " << result << endl;
    result = a >= b;  // false
    cout << "3 >= 5 is " << result << endl;
    result = a <= b;  // true
    cout << "3 <= 5 is " << result << endl;

    //logical operator
    cout << endl;
    cout<<"logical operator : \n";
    result = (3 != 5) && (3 < 5);     // true
    cout << "(3 != 5) && (3 < 5) is " << result << endl;

    result = (3 == 5) && (3 < 5);    // false
    cout << "(3 == 5) && (3 < 5) is " << result << endl;

    result = (3 == 5) && (3 > 5);    // false
    cout << "(3 == 5) && (3 > 5) is " << result << endl;

    result = (3 != 5) || (3 < 5);    // true
    cout << "(3 != 5) || (3 < 5) is " << result << endl;

    result = (3 != 5) || (3 > 5);    // true
    cout << "(3 != 5) || (3 > 5) is " << result << endl;

    result = (3 == 5) || (3 > 5);    // false
    cout << "(3 == 5) || (3 > 5) is " << result << endl;

    result = !(5 == 2);    // true
    cout << "!(5 == 2) is " << result << endl;

    result = !(5 == 5);    // false
    cout << "!(5 == 5) is " << result << endl;


    //bitwise operator
    cout << endl;
    cout<<"bitwise operator : \n";
    a = 3;
    b = 7;
    //AND
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;
    cout << "a & b = " << (a & b) << endl;
    //OR
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;
    cout << "a | b = " << (a | b) << endl;
    //XOR
    cout << "a = " << a << endl;
    cout << "b = " << b << endl;
    cout << "a ^ b = " << (a ^ b) << endl;

    return 0;
}