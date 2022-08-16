class Solution {
    int[][] dp;

    public static void main(String[] args){
        Solution s = new Solution();
        s.KMP("ABCAB");
    }

    public void KMP(String pat) {
        int m = pat.length();

        dp = new int[m][256];
        dp[0][pat.charAt(0)] = 1;
        int x = 0;
        for (int i = 1; i < m; i++) {
            for (int c = 0; c < 256; c++) {
                dp[i][c] = dp[x][c];
            }
            dp[i][pat.charAt(i)] = i + 1;
            x = dp[x][pat.charAt(i)]; // match from head
            /**
             * eg ABCAB
             * if i = 3
             * pat[0] = pat[3] = 'A'
             * so x = 1
             */
        }
    }

    public int strStr(String haystack, String needle) {
        KMP(needle);

        int m = haystack.length();
        int n = needle.length();
        int j = 0;
        for (int i = 0; i < m; i++) {
            j = dp[j][needle.charAt(i)];
            if (j == n){
                return i - n +1;
            }
        }
        return -1;
    }
}