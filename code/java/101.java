class Solution {
  public boolean isSymmetric(TreeNode root) {
    // judge the subtree
    return symmetric(root.left, root.right);
  }

  private boolean symmetric(TreeNode p, TreeNode q) {
    if (p == null && q == null) {
      return true;
    } else if ((p != null && q == null) || (p == null && q != null)) {
      return false;
    } else {
      // if val is equal and subtree is symmetric, return true
      return p.val == q.val && symmetric(p.left, q.right) && symmetric(p.right, q.left);
    }
  }
}

class TreeNode {
  int val;
  TreeNode left;
  TreeNode right;

  TreeNode() {
  }

  TreeNode(int val) {
    this.val = val;
  }

  TreeNode(int val, TreeNode left, TreeNode right) {
    this.val = val;
    this.left = left;
    this.right = right;
  }
}