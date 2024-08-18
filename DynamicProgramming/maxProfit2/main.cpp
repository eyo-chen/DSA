#include <vector>

using namespace std;

class Solution {
 public:
  int maxProfit(vector<int>& prices) {
    vector<int> maxProfitTable = vector<int>(prices.size(), 0);

    for (int i = 1; i < prices.size(); i++) {
      int sellToday = prices[i] - prices[i - 1] + maxProfitTable[i - 1];
      int doNotSellToday = maxProfitTable[i - 1];
      maxProfitTable[i] = max(sellToday, doNotSellToday);
    }

    return maxProfitTable[prices.size() - 1];
  }
};

class Solution {
 public:
  int maxProfit(vector<int>& prices) {
    int maxProfit = 0;

    for (int i = 1; i < prices.size(); i++) {
      int sellToday = prices[i] - prices[i - 1] + maxProfit;
      int doNotSellToday = maxProfit;
      maxProfit = max(sellToday, doNotSellToday);
    }

    return maxProfit;
  }
};