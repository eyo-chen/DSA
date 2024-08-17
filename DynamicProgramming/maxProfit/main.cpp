#include <vector>

using namespace std;

class Solution {
 public:
  int maxProfit(vector<int>& prices) {
    int maxProfit = 0;

    // Try each day as a buying day
    for (int i = 0; i < prices.size(); i++) {
      int buyPrice = prices[i];

      // Try each day after the buying day as a selling day
      for (int j = i + 1; j < prices.size(); j++) {
        int sellPrice = prices[j];
        maxProfit = max(maxProfit, sellPrice - buyPrice);
      }
    }

    return maxProfit;
  }
};

class Solution {
 public:
  int maxProfit(vector<int>& prices) {
    int minPrice = prices[0];
    int maxProfit = 0;

    for (int p : prices) {
      // What is the minimum price so far?
      minPrice = min(minPrice, p);

      // What is the maximum profit so far?
      maxProfit = max(maxProfit, p - minPrice);
    }

    return maxProfit;
  }
};