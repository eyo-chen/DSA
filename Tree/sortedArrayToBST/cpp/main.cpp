#include <vector>

using namespace std;

//  Definition for a binary tree node.
struct TreeNode {
  int val;
  TreeNode* left;
  TreeNode* right;
  TreeNode() : val(0), left(nullptr), right(nullptr) {}
  TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
  TreeNode(int x, TreeNode* left, TreeNode* right)
      : val(x), left(left), right(right) {}
};

class Solution {
 public:
  TreeNode* sortedArrayToBST(vector<int>& nums) {
    return build(nums, 0, nums.size() - 1);
  }

  TreeNode* build(vector<int>& nums, int left, int right) {
    if (right < left) return nullptr;

    int middle = left + ((right - left) / 2);
    int val = nums[middle];
    TreeNode* node = new TreeNode(val);

    node->left = build(nums, left, middle - 1);
    node->right = build(nums, middle + 1, right);

    return node;
  }
};