#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;

void tmp(vector<int> &aaa, vector<vector<int>> &bbb) {
  // print aaa[0], aaa[1], aaa[2]
  cout << "aaa[0]: " << aaa[0] << endl;
  cout << "aaa[1]: " << aaa[1] << endl;
  cout << "aaa[2]: " << aaa[2] << endl;

  bbb.push_back(aaa);
  aaa.pop_back();
}

void tmp2(vector<int> &aaa) {
  aaa.pop_back();
  aaa.push_back(10);
}

int main() {
  // vector<int> aaa = {7, 8, 9};
  // vector<vector<int>> bbb = {{1, 2, 3}, {4, 5, 6}};

  // tmp(aaa, bbb);

  // cout << "bbb: " << endl;
  // for (auto &i : bbb) {
  //   for (auto &j : i) {
  //     cout << j << " ";
  //   }
  //   cout << endl;
  // }

  // cout << "aaa: ";
  // for (auto &i : aaa) {
  //   cout << i << " ";
  // }
  // cout << endl;

  vector<int> aaa = {7, 8, 9};
  // print aaa
  cout << "aaa: ";
  for (auto &i : aaa) {
    cout << i << " ";
  }
  cout << endl;

  tmp2(aaa);

  // print aaa
  cout << "aaa: ";
  for (auto &i : aaa) {
    cout << i << " ";
  }
  cout << endl;

  return 0;
}