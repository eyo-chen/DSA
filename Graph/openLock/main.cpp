#include <queue>
#include <string>
#include <unordered_set>
#include <vector>

using namespace std;

class Solution {
 public:
  int openLock(vector<string>& deadends, string target) {
    unordered_set<string> deadendSet(deadends.begin(), deadends.end());
    if (deadendSet.find("0000") != deadendSet.end()) return -1;

    int steps = 0;
    queue<string> q;
    unordered_set<string> seenSet;
    q.push("0000");
    seenSet.insert("0000");

    while (!q.empty()) {
      int size = q.size();

      for (int i = 0; i < size; i++) {
        string digit = q.front();
        q.pop();

        if (digit == target) return steps;

        vector<string> nextDigits = genNextDigits(digit);
        for (const string d : nextDigits) {
          if (deadendSet.find(d) != deadendSet.end()) continue;
          if (seenSet.find(d) != seenSet.end()) continue;

          q.push(d);
          seenSet.insert(d);
        }
      }

      steps++;
    }

    return -1;
  }

  vector<string> genNextDigits(string digit) {
    vector<string> result;

    for (int i = 0; i < 4; i++) {
      vector<string> inAndDe = genInAndDe(digit, i);
      result.push_back(inAndDe[0]);
      result.push_back(inAndDe[1]);
    }

    return result;
  }

  // using modulo to handle increment and decrement
  // it's more clever and concise way to handle increment and decrement
  vector<string> genInAndDe(string digit, int index) {
    string in = digit;
    string de = digit;

    in[index] = (in[index] - '0' + 1) % 10 + '0';
    de[index] = (de[index] - '0' + 9) % 10 + '0';

    return {in, de};
  }

  // using if-statement to handle increment and decrement
  vector<string> genInAndDe(string digit, int index) {
    string in = digit;
    string de = digit;

    // handle increment
    if (in[index] == '9') {
      in[index] = '0';
    } else {
      in[index] = in[index] + 1;
    }

    // handle decrement
    if (de[index] == '0') {
      de[index] = '9';
    } else {
      de[index] = de[index] - 1;
    }

    return {in, de};
  }
};

// This is more efficient than the previous one
// Because there's less function call, and less memory allocation
// (e.g. nextDigits)
class Solution {
 public:
  int openLock(vector<string>& deadends, string target) {
    unordered_set<string> deadendSet(deadends.begin(), deadends.end());
    if (deadendSet.find("0000") != deadendSet.end()) return -1;

    int steps = 0;
    queue<string> q;
    q.push("0000");
    deadendSet.insert("0000");

    while (!q.empty()) {
      int size = q.size();

      for (int i = 0; i < size; i++) {
        string digit = q.front();
        q.pop();

        if (digit == target) return steps;

        for (int i = 0; i < 4; i++) {
          string in = digit;
          string de = digit;

          in[i] = (in[i] - '0' + 1) % 10 + '0';
          de[i] = (de[i] - '0' + 9) % 10 + '0';

          if (deadendSet.find(in) == deadendSet.end()) {
            q.push(in);
            deadendSet.insert(in);
          }

          if (deadendSet.find(de) == deadendSet.end()) {
            q.push(de);
            deadendSet.insert(de);
          }
        }
      }

      steps++;
    }

    return -1;
  }
};