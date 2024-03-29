#include <map>
#include <vector>

using namespace std;

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
  TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
    map<int, int> hashMap;
    for (int i = 0; i < inorder.size(); i++) {
      hashMap[inorder[i]] = i;
    }

    int index = 0;

    return helper(preorder, 0, inorder.size() - 1, index, hashMap);
  }

  TreeNode* helper(vector<int>& preorder, int left, int right, int& index,
                   map<int, int>& hashMap) {
    if (left > right) return nullptr;

    int curVal = preorder[index];
    index++;

    TreeNode* curRoot = new TreeNode(curVal);

    curRoot->left = helper(preorder, left, hashMap[curVal] - 1, index, hashMap);
    curRoot->right =
        helper(preorder, hashMap[curVal] + 1, right, index, hashMap);

    return curRoot;
  }
};
