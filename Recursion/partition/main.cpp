#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<string>> partition(string s) {
    vector<vector<string>> ans;
    vector<string> tmp;

    helper(ans, tmp, s, 0);

    return ans;
  }

  void helper(vector<vector<string>> &ans, vector<string> &tmp, string s,
              int index) {
    if (index == s.length()) {
      ans.push_back(tmp);
      return;
    }

    for (int i = index; i < s.length(); i++) {
      string subs = s.substr(index, i - index + 1);

      if (!isPalindrome(subs)) continue;

      tmp.push_back(subs);
      helper(ans, tmp, s, i + 1);
      tmp.pop_back();
    }
  }

  bool isPalindrome(string &s) {
    int start = 0;
    int end = s.length() - 1;

    while (start < end) {
      if (s[start] != s[end]) {
        return false;
      }

      start++;
      end--;
    }

    return true;
  }
};