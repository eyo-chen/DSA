#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;

int main() {
  string s = "123456789";

  string a = s.substr(3, 3);

  cout << a << endl;

  return 0;
}