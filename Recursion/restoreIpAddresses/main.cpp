#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<string> restoreIpAddresses(string s) {
    vector<string> res;
    vector<string> ans;

    helper(s, res, ans, 0);
    return res;
  }

 private:
  void helper(string& s, vector<string>& res, vector<string>& ans, int index) {
    if (ans.size() == 4 && index == s.length()) {
      string conS = concatenateStrings(ans);
      res.push_back(conS);
      return;
    }

    for (int i = 1; i <= 3 && index + i <= s.length(); i++) {
      string subS = s.substr(index, i);
      if (!isValid(subS)) break;

      ans.push_back(subS);
      helper(s, res, ans, index + i);
      ans.pop_back();
    }
  }

  bool isValid(string subS) {
    // check leading zero
    if (subS.length() > 1 && subS[0] == '0') return false;
    int subI = stoi(subS);

    return subI >= 0 && subI <= 255;
  }

  string concatenateStrings(vector<string>& ans) {
    string output;

    for (const string& str : ans) {
      if (output.length() != 0) {
        output += '.';
      }
      output += str;
    }

    return output;
  }
};

int main() {
  Solution solution;
  string s = "25525511135";
  vector<string> result = solution.restoreIpAddresses(s);
  cout << "Result: ";
  for (const string& combination : result) {
    cout << combination << " ";
  }
  cout << endl;

  return 0;
}