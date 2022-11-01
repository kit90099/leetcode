class GFG {

    // Structure of the node
    static class Node {
        int value;
        Node L, R;

    }

    // Structure to get the newly formed
    // node
    static Node getnode() {
        Node temp = new Node();
        temp.value = 0;
        temp.L = null;
        temp.R = null;
        return temp;
    }

    // Creating the Root node
    static Node root = new Node();

    // Function to perform the point update
    // on the dynamic segment tree
    static void UpdateHelper(Node curr, int index, int L, int R, int val) {

        // If the index is not overlapping
        // with the index
        if (L > index || R < index)
            return;

        // If the index is completely overlapping
        // with the index
        if (L == R && L == index) {

            // Update the value of the node
            // to the given value
            curr.value = val;
            return;
        }

        // Computing the middle index if none
        // of the above base cases are satisfied
        int mid = L - (L - R) / 2;
        int sum1 = 0, sum2 = 0;

        // If the index is in the left subtree
        if (index <= mid) {

            // Create a new node if the left
            // subtree is is null
            if (curr.L == null)
                curr.L = getnode();

            // Recursively call the function
            // for the left subtree
            UpdateHelper(curr.L, index, L, mid, val);
        }

        // If the index is in the right subtree
        else {

            // Create a new node if the right
            // subtree is is null
            if (curr.R == null)
                curr.R = getnode();

            // Recursively call the function
            // for the right subtree
            UpdateHelper(curr.R, index, mid + 1, R, val);
        }

        // Storing the sum of the left subtree
        if (curr.L != null)
            sum1 = curr.L.value;

        // Storing the sum of the right subtree
        if (curr.R != null)
            sum2 = curr.R.value;

        // Storing the sum of the children into
        // the node's value
        curr.value = sum1 + sum2;
        return;
    }

    // Function to find the sum of the
    // values given by the range
    static int queryHelper(Node curr, int a, int b, int L, int R) {

        // Return 0 if the root is null
        if (curr == null)
            return 0;

        // If the index is not overlapping
        // with the index, then the node
        // is not created. So sum is 0
        if (L > b || R < a)
            return 0;

        // If the index is completely overlapping
        // with the index, return the node's value

        if (L >= a && R <= b)
            return curr.value;

        int mid = L - (L - R) / 2;

        // Return the sum of values stored
        // at the node's children
        return queryHelper(curr.L, a, b, L, mid) + queryHelper(curr.R, a, b, mid + 1, R);
    }

    // Function to call the queryHelper
    // function to find the sum for
    // the query
    static int query(int L, int R) {
        return queryHelper(root, L, R, 1, 10);
    }

    // Function to call the UpdateHelper
    // function for the point update
    static void update(int index, int value) {
        UpdateHelper(root, index, 1, 10, value);
    }

    // Function to perform the operations
    // on the tree
    static void operations() {
        // Creating an empty tree
        root = getnode();

        // Update the value at position 1 to 10
        update(1, 10);

        // Update the value at position 3 to 5
        update(3, 5);

        // Finding sum for the range [2, 8]
        System.out.println(query(2, 8));

        // Finding sum for the range [1, 10]
        System.out.println(query(1, 10));

    }

    // Driver code
    public static void main(String[] args) {
        operations();
    }
}