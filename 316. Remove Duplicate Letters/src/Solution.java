import java.nio.charset.StandardCharsets;
import java.util.*;
import java.util.stream.Collectors;

class Solution {
    public String removeDuplicateLetters(String s) {
        boolean[] contains = new boolean[256];

        ArrayDeque<Byte> queue = new ArrayDeque<>();
        int[] counter = new int[256];
        for (byte c : s.getBytes()) {
            counter[c]++;
        }

        for (byte c : s.getBytes()) {
            counter[c]--;
            if (contains[c]) {
                continue;
            }

            while (!queue.isEmpty()) {
                Byte last = queue.getLast();
                if (last > c && counter[last] > 0) {
                    queue.pollLast();
                    contains[last] = false;
                } else {
                    break;
                }
            }

            queue.addLast(c);
            contains[c] = true;
        }

        return queue.stream().map((b) -> String.valueOf((char) b.byteValue())).collect(Collectors.joining());
    }
}