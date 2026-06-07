public class Solution {
    public TreeNode createBinaryTree(int[][] descriptions) {
        Map<Integer, TreeNode> map = new HashMap<>();
        Set<Integer> children = new HashSet<>();

        // Create nodes and build the tree connections
        for (int[] desc : descriptions) {
            int pVal = desc[0];
            int cVal = desc[1];
            int isLeft = desc[2];

            map.putIfAbsent(pVal, new TreeNode(pVal));
            map.putIfAbsent(cVal, new TreeNode(cVal));

            TreeNode parent = map.get(pVal);
            TreeNode child = map.get(cVal);

            if (isLeft == 1) {
                parent.left = child;
            } else {
                parent.right = child;
            }

            children.add(cVal);
        }

        // Find the root (exists in map, but not in children set)
        TreeNode root = null;
        for (int val : map.keySet()) {
            if (!children.contains(val)) {
                root = map.get(val);
                break;
            }
        }

        return root;
    }
}
