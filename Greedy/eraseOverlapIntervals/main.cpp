#include <vector>

using namespace std;

class Solution {
 public:
  int eraseOverlapIntervals(vector<vector<int>>& intervals) {
    sort(
        intervals.begin(), intervals.end(),
        [](const vector<int>& a, const vector<int>& b) { return a[0] < b[0]; });

    int res = 0;
    int prevEnd = intervals[0][1];

    for (int i = 1; i < intervals.size(); i++) {
      int start = intervals[i][0];
      int end = intervals[i][1];
      if (start >= prevEnd) {
        prevEnd = end;
        continue;
      }

      prevEnd = min(prevEnd, end);
      res++;
    }

    return res;
  }
};