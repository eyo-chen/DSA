#include <vector>

using namespace std;

class Solution {
 public:
  int findMinArrowShots(vector<vector<int>>& points) {
    sort(
        points.begin(), points.end(),
        [](const vector<int>& a, const vector<int>& b) { return a[0] < b[0]; });

    int count = points.size();
    vector<int> prev = points[0];

    // start at 2nd index because we use the first index as the initial prev
    for (int i = 1; i < points.size(); i++) {
      vector<int> cur = points[i];

      // not overlapping
      // prev[1] -> end of the previous interval
      // cur[0]  -> start of the current interval
      if (prev[1] < cur[0]) {
        prev = cur;
        continue;
      }

      // overlapping
      count--;
      // merge the intervals
      prev = {max(prev[0], cur[0]), min(prev[1], cur[1])};
    }

    return count;
  }
};