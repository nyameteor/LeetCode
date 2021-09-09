import java.util.ArrayDeque;
import java.util.LinkedHashMap;
import java.util.Map;
import java.util.Queue;

// Hash Table + Queue
class LRUCache {

    // store key to specific priority
    // queue.remove(o) -> O(n)
    // queue.offer -> O(1)
    Queue<Integer> queue = new ArrayDeque<>();

    // store key-value pairs
    // LinkedHashMap: Hash table and linked list implementation of the Map
    // interface, with predictable iteration order.
    // if use HashMap, will get `Time Limit Exceeded`
    Map<Integer, Integer> map = new LinkedHashMap<>();

    int capacity;

    public LRUCache(int capacity) {
        this.capacity = capacity;
    }

    public int get(int key) {
        // according to LRU principal, if key is present in the map,
        // key hash to be searched and deleted,
        // and then enqueue to update the priority
        if (map.containsKey(key)) {
            queue.remove(key);
            queue.offer(key);
            return map.get(key);
        } else {
            return -1;
        }
    }

    public void put(int key, int value) {
        if (map.containsKey(key)) {
            queue.remove(key);
            queue.offer(key);
            map.put(key, value);
        } else {
            if (queue.size() < capacity) {
                queue.offer(key);
            } else {
                // remove the least recently used key
                map.remove(queue.poll());
                queue.offer(key);
            }
            map.put(key, value);
        }
    }
}
