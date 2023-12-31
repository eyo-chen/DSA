#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;

int main() {
  string sta;

  sta += "(";
  cout << sta << endl;

  sta += ")";
  cout << sta << endl;

  sta.pop_back();
  cout << sta << endl;
}