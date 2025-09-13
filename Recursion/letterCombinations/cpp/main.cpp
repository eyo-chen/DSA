#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<string> letterCombinations(string digits) {
    vector<string> result;
    if (digits.length() == 0) return result;

    string mapping[10] = {"",    "",    "abc",  "def", "ghi",
                          "jkl", "mno", "pqrs", "tuv", "wxyz"};
    string ans;
    helper(digits, 0, mapping, ans, result);
    return result;
  }

 private:
  void helper(string digits, int index, string mapping[], string& ans,
              vector<string>& result) {
    if (index == digits.length()) {
      result.push_back(ans);
      return;
    }

    char digit = digits[index];
    string strs = mapping[digit - '0'];

    for (char s : strs) {
      ans.push_back(s);
      helper(digits, index + 1, mapping, ans, result);
      ans.pop_back();
    }

    return;
  }
};

int main() {
  Solution solution;
  vector<string> result = solution.letterCombinations("23");
  cout << "Result: ";
  for (const string& combination : result) {
    cout << combination << " ";
  }
  cout << endl;

  return 0;
}
