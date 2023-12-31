#include <iostream>
#include <stack>
#include <string>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<string> generateParenthesis(int n) {
    vector<string> ans;

    helper(ans, "", 0, 0, n);

    return ans;
  }

  void helper(vector<string>& ans, string temp, int leftIndex, int rightIndex,
              int n) {
    if (temp.length() == n * 2) {
      ans.push_back(temp);
      return;
    }

    if (leftIndex < n) {
      helper(ans, temp + "(", leftIndex + 1, rightIndex, n);
    }

    if (leftIndex > rightIndex) {
      helper(ans, temp + ")", leftIndex, rightIndex + 1, n);
    }
  }
};

/*
This is the solution I came up with at first
Because my initial thought is to use a stack to validate the parenthesis

The genral idea is to use a stack as the index to check when to add ")" to the
But I forgot the unchoose part

Note that we have to restore the prior state when we backtrack

The space complexity is not efficient because we have to store the stack
*/
class Solution {
 public:
  vector<string> generateParenthesis(int n) {
    vector<string> ans;
    stack<string> sta;

    helper(ans, "", sta, 0, n);

    return ans;
  }

  void helper(vector<string>& ans, string temp, stack<string>& sta,
              int leftIndex, int n) {
    if (temp.length() == n * 2) {
      ans.push_back(temp);
      return;
    }

    if (leftIndex < n) {
      sta.push("(");
      helper(ans, temp + "(", sta, leftIndex + 1, n);
      sta.pop();  // Do Not Forget to unshoose
    }

    if (!sta.empty()) {
      sta.pop();
      helper(ans, temp + ")", sta, leftIndex, n);
      sta.push("(");  // Do Not Forget to unshoose
    }
  }
};