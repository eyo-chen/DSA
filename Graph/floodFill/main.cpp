#include <iostream>
#include <queue>
#include <set>
#include <stack>
#include <vector>

using namespace std;

class Solution {
 public:
  vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc,
                                int color) {
    int rowLen = image.size();
    int colLen = image[0].size();
    int startColor = image[sr][sc];

    queue<pair<int, int>> q;
    set<pair<int, int>> explored;
    q.push({sr, sc});

    while (!q.empty()) {
      pair<int, int> location = q.front();
      q.pop();

      int r = location.first;
      int c = location.second;

      // check if the location is valid
      if (r < 0 || c < 0 || r >= rowLen || c >= colLen) continue;
      if (explored.find(location) != explored.end()) continue;
      if (image[r][c] != startColor) continue;

      // push the adjacent locations into the data structure
      q.push({r + 1, c});
      q.push({r - 1, c});
      q.push({r, c - 1});
      q.push({r, c + 1});

      // change the color of the location
      image[r][c] = color;

      // mark the location as explored
      explored.insert(location);
    }

    return image;
  }
};

class Solution {
 public:
  vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc,
                                int color) {
    int rowLen = image.size();
    int colLen = image[0].size();
    int startColor = image[sr][sc];

    stack<pair<int, int>> s;
    set<pair<int, int>> explored;
    s.push({sr, sc});

    while (!s.empty()) {
      pair<int, int> location = s.top();
      s.pop();

      int r = location.first;
      int c = location.second;

      // check if the location is valid
      if (r < 0 || c < 0 || r >= rowLen || c >= colLen) continue;
      if (explored.find(location) != explored.end()) continue;
      if (image[r][c] != startColor) continue;

      // push the adjacent locations into the data structure
      s.push({r + 1, c});
      s.push({r - 1, c});
      s.push({r, c - 1});
      s.push({r, c + 1});

      // change the color of the location
      image[r][c] = color;

      // mark the location as explored
      explored.insert(location);
    }

    return image;
  }
};

class Solution {
 public:
  vector<vector<int>> floodFill(vector<vector<int>>& image, int sr, int sc,
                                int color) {
    set<pair<int, int>> explored;

    helper(image, sr, sc, color, image[sr][sc], explored);

    return image;
  }

 private:
  void helper(vector<vector<int>>& image, int r, int c, int color,
              int startColor, set<pair<int, int>>& explored) {
    pair<int, int> location = {r, c};

    // check if the location is valid
    if (r < 0 || c < 0 || r >= image.size() || c >= image[0].size()) return;
    if (explored.find(location) != explored.end()) return;
    if (image[r][c] != startColor) return;

    // change the color of the location
    image[r][c] = color;

    // mark the location as explored
    explored.insert(location);

    // explore the adjacent locations
    helper(image, r - 1, c, color, startColor, explored);
    helper(image, r + 1, c, color, startColor, explored);
    helper(image, r, c - 1, color, startColor, explored);
    helper(image, r, c + 1, color, startColor, explored);
  }
};