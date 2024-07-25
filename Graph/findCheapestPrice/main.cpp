#include <queue>
#include <unordered_map>
#include <vector>

using namespace std;

struct flightInfo {
  int stop;
  int price;
  int flight;
};

class Solution {
 public:
  int findCheapestPrice(int n, vector<vector<int>>& flights, int src, int dst,
                        int k) {
    // create adjacency list
    unordered_map<int, vector<vector<int>>> adj;
    for (const vector<int>& f : flights) {
      int from = f[0], to = f[1], price = f[2];
      adj[from].push_back({to, price});
    }

    queue<flightInfo> q({{0, 0, src}});
    int minPrice = INT_MAX;

    while (!q.empty()) {
      flightInfo f = q.front();
      q.pop();

      // if stop is greater than k, then skip
      if (f.stop - 1 > k) continue;

      // if price is greater than minPrice, then skip
      if (f.price > minPrice) continue;

      // if destination is reached, then update minPrice
      if (f.flight == dst) {
        minPrice = min(minPrice, f.price);
        continue;
      }

      // loop through all the flights(adjacent flights) from the current flight
      // a[0] = to, a[1] = price
      for (const vector<int>& adjFlight : adj[f.flight]) {
        int newStops = f.stop + 1;
        int newPrice = f.price + adjFlight[1];
        int to = adjFlight[0];

        q.push({newStops, newPrice, to});
      }
    }

    return minPrice == INT_MAX ? -1 : minPrice;
  }
};

class Solution {
 public:
  int findCheapestPrice(int n, vector<vector<int>>& flights, int src, int dst,
                        int k) {
    // create adjacency list
    unordered_map<int, vector<vector<int>>> adj;
    for (const vector<int>& f : flights) {
      int from = f[0], to = f[1], price = f[2];
      adj[from].push_back({to, price});
    }

    // create prices array
    // prices[i] represents the minimum price to reach location i
    vector<int> prices(n, INT_MAX);

    queue<flightInfo> q({{0, 0, src}});
    int minPrice = INT_MAX;

    while (!q.empty()) {
      flightInfo f = q.front();
      q.pop();

      // if stop is greater than k, then skip
      if (f.stop - 1 > k) continue;

      // if destination is reached, then update minPrice
      if (f.flight == dst) {
        minPrice = min(minPrice, f.price);
        continue;
      }

      // loop through all the flights(adjacent flights) from the current flight
      // a[0] = to, a[1] = price
      for (const vector<int>& adjFlight : adj[f.flight]) {
        int newStops = f.stop + 1;
        int newPrice = f.price + adjFlight[1];
        int to = adjFlight[0];

        // note that prices[i] is the minimum price to reach location i
        // if the new price is greater than the existing price
        // then there's no need to explore this path
        if (newPrice > prices[to]) continue;

        q.push({newStops, newPrice, to});

        // update the price to reach location
        prices[to] = newPrice;
      }
    }

    return minPrice == INT_MAX ? -1 : minPrice;
  }
};

struct flightInfo1 {
  int price;
  int flight;
};

class Solution {
 public:
  int findCheapestPrice(int n, vector<vector<int>>& flights, int src, int dst,
                        int k) {
    // create adjacency list
    unordered_map<int, vector<vector<int>>> adj;
    for (const vector<int>& f : flights) {
      int from = f[0], to = f[1], price = f[2];
      adj[from].push_back({to, price});
    }

    // create prices array
    // prices[i] represents the minimum price to reach location i
    vector<int> prices(n, INT_MAX);

    queue<flightInfo1> q({{0, src}});
    int minPrice = INT_MAX;
    int stops = -1;

    // note that stops <= k is used to ensure that we don't exceed the number of
    // stops
    while (!q.empty() && stops <= k) {
      int size = q.size();
      for (int i = 0; i < size; i++) {
        flightInfo1 f = q.front();
        q.pop();

        // if destination is reached, then update minPrice
        if (f.flight == dst) {
          minPrice = min(minPrice, f.price);
          continue;
        }

        // loop through all the flights(adjacent flights) from the current
        // flight a[0] = to, a[1] = price
        for (const vector<int>& adjFlight : adj[f.flight]) {
          int newPrice = f.price + adjFlight[1];
          int to = adjFlight[0];

          // note that prices[i] is the minimum price to reach location i
          // if the new price is greater than the existing price
          // then there's no need to explore this path
          if (newPrice > prices[to]) continue;

          q.push({newPrice, to});

          // update the price to reach location a[0]
          prices[to] = newPrice;
        }
      }

      stops++;
    }

    return minPrice == INT_MAX ? -1 : minPrice;
  }
};

// Using Bellman Ford Algorithm
class Solution {
 public:
  int findCheapestPrice(int n, vector<vector<int>>& flights, int src, int dst,
                        int k) {
    // Initialize the prices array
    vector<int> prices(n, INT_MAX);

    // Mark the source price as 0
    prices[src] = 0;

    // Loop through the number of stops
    // Note that we loop through k + 1 times
    for (int i = 0; i <= k; i++) {
      // Create a temporary prices array(copy of the original prices array)
      vector<int> tempPrices = prices;

      // Loop through all the adjacent flights
      for (const vector<int>& flight : flights) {
        int from = flight[0], to = flight[1], price = flight[2];

        // If the price of the source is INT_MAX, then we skip this flight
        // It means the `from` node is not reachable yet
        if (prices[from] == INT_MAX) continue;

        // Update the price of the destination node
        tempPrices[to] = min(tempPrices[to], prices[from] + price);
      }

      // Update the prices array
      prices = tempPrices;
    }

    return prices[dst] == INT_MAX ? -1 : prices[dst];
  }
};