class Solution {
    public String addStrings(String num1, String num2) {
        StringBuilder res = new StringBuilder();
        int carry = 0;
        int partSum = 0;
        int i = num1.length() - 1, j = num2.length() - 1;

        while (i >= 0 || j >= 0) {
            if (i >= 0 && j >= 0) {
                // char to int
                partSum = (num1.charAt(i) - '0') + (num2.charAt(j) - '0') + carry;
            }

            else if (i >= 0 && j < 0) {
                partSum = (num1.charAt(i) - '0') + carry;
            }

            else if (i < 0 && j >= 0) {
                partSum = (num2.charAt(j) - '0') + carry;
            }

            if (partSum >= 10) {
                carry = 1;
                partSum = partSum - 10;
            } else {
                carry = 0;
            }

            // int to char
            res.insert(0, (char) (partSum + '0'));

            i--;
            j--;
        }

        if (carry != 0) {
            res.insert(0, (char) (carry + '0'));
        }

        return res.toString();
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.addStrings("11", "123"));
        System.out.println(solution.addStrings("99", "1"));
        System.out.println(solution.addStrings("456", "77"));
        System.out.println(solution.addStrings("0", "0"));
    }

}