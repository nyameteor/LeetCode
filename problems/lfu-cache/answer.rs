use std::{
    cell::RefCell,
    collections::HashMap,
    rc::{Rc, Weak},
};

// =============================================================================
// Doubly-Linked List
// =============================================================================

type NodeRef = Rc<RefCell<Node>>;
type WeakNodeRef = Weak<RefCell<Node>>;
type Key = i32;

struct Node {
    key: Key,
    val: i32,
    freq: u32,
    prev: Option<WeakNodeRef>,
    next: Option<NodeRef>,
}

fn new_node(key: Key, val: i32, freq: u32) -> NodeRef {
    Rc::new(RefCell::new(Node {
        key,
        val,
        freq: freq,
        prev: None,
        next: None,
    }))
}

struct List {
    head: NodeRef,
    tail: NodeRef,
}

impl List {
    fn new() -> List {
        let list = List {
            head: new_node(-1, -1, 0),
            tail: new_node(-1, -1, 0),
        };
        list.head.borrow_mut().next = Some(Rc::clone(&list.tail));
        list.tail.borrow_mut().prev = Some(Rc::downgrade(&list.head));
        list
    }

    fn first(&self) -> NodeRef {
        let head = self.head.borrow();
        let first = head.next.as_ref().unwrap();
        return first.clone();
    }

    fn is_empty(&self) -> bool {
        Rc::ptr_eq(&self.first(), &self.tail)
    }

    fn pop_front(&mut self) -> Option<NodeRef> {
        let first = self.first();
        if Rc::ptr_eq(&first, &self.tail) {
            None
        } else {
            self.remove(&first);
            Some(first)
        }
    }

    fn push_back(&mut self, node: NodeRef) {
        let last = self.tail.borrow().prev.as_ref().unwrap().upgrade().unwrap();
        self.insert_after(&last, node);
    }

    fn remove(&mut self, node: &NodeRef) {
        let (prev, next) = {
            let node = node.borrow();
            (
                node.prev.as_ref().unwrap().upgrade().unwrap(),
                node.next.as_ref().unwrap().clone(),
            )
        };

        prev.borrow_mut().next = Some(next.clone());
        next.borrow_mut().prev = Some(Rc::downgrade(&prev));

        let mut node = node.borrow_mut();
        node.prev = None;
        node.next = None;
    }

    fn insert_after(&mut self, prev: &NodeRef, node: NodeRef) {
        let prev_next = prev.borrow().next.as_ref().unwrap().clone();

        node.borrow_mut().prev = Some(Rc::downgrade(prev));
        node.borrow_mut().next = Some(prev_next.clone());

        prev_next.borrow_mut().prev = Some(Rc::downgrade(&node));
        prev.borrow_mut().next = Some(node);
    }
}

// =============================================================================
// LFU Cache
// =============================================================================

struct LFUCache {
    capacity: usize,
    min_freq: u32,
    nodes_by_key: HashMap<Key, NodeRef>,
    lists_by_freq: HashMap<u32, List>,
}

impl LFUCache {
    fn new(capacity: i32) -> Self {
        LFUCache {
            capacity: capacity.try_into().unwrap(),
            min_freq: u32::MAX,
            nodes_by_key: HashMap::new(),
            lists_by_freq: HashMap::new(),
        }
    }

    fn len(&self) -> usize {
        self.nodes_by_key.len()
    }

    fn get(&mut self, key: i32) -> i32 {
        let node = self.nodes_by_key.get(&key).cloned();
        match node {
            Some(node) => {
                self.touch(&node);
                node.borrow().val
            }
            None => -1,
        }
    }

    fn put(&mut self, key: i32, val: i32) {
        let node = self.nodes_by_key.get(&key).cloned();
        match node {
            Some(node) => {
                node.borrow_mut().val = val;
                self.touch(&node);
            }
            None => {
                if self.len() == self.capacity {
                    self.evict();
                }
                let freq = 1;
                let node = new_node(key, val, freq);
                self.nodes_by_key.insert(key, node.clone());

                let list = self.lists_by_freq.entry(freq).or_insert(List::new());
                list.push_back(node.clone());

                self.min_freq = freq;
            }
        }
    }

    fn touch(&mut self, node: &NodeRef) {
        let old_freq = node.borrow().freq;
        let new_freq = old_freq + 1;
        node.borrow_mut().freq = new_freq;

        let old_list = self.lists_by_freq.get_mut(&old_freq).unwrap();
        old_list.remove(node);

        let old_list_empty = old_list.is_empty();
        if old_list_empty {
            self.lists_by_freq.remove(&old_freq);
        }

        if self.min_freq == old_freq && old_list_empty {
            self.min_freq = new_freq;
        }

        let list = self.lists_by_freq.entry(new_freq).or_insert(List::new());
        list.push_back(node.clone());
    }

    fn evict(&mut self) {
        let list = self.lists_by_freq.get_mut(&self.min_freq).unwrap();
        let node = list.pop_front().unwrap();
        let freq = node.borrow().freq;

        self.nodes_by_key.remove(&node.borrow().key);
        if list.is_empty() {
            self.lists_by_freq.remove(&freq);
        }
    }
}

fn test_example() {
    let mut lfu = LFUCache::new(2);

    lfu.put(1, 1);
    lfu.put(2, 2);
    assert_eq!(lfu.get(1), 1);

    lfu.put(3, 3);
    assert_eq!(lfu.get(2), -1);
    assert_eq!(lfu.get(3), 3);

    lfu.put(4, 4);
    assert_eq!(lfu.get(1), -1);
    assert_eq!(lfu.get(3), 3);
    assert_eq!(lfu.get(4), 4);
}

fn test_capacity_one() {
    let mut lfu = LFUCache::new(1);

    lfu.put(1, 1);
    assert_eq!(lfu.get(1), 1);

    lfu.put(2, 2);
    assert_eq!(lfu.get(1), -1);
    assert_eq!(lfu.get(2), 2);

    lfu.put(3, 3);
    assert_eq!(lfu.get(2), -1);
    assert_eq!(lfu.get(3), 3);
}

fn test_same_freq_evicts_least_recent() {
    let mut lfu = LFUCache::new(2);

    lfu.put(1, 1);
    lfu.put(2, 2);
    lfu.put(3, 3);

    assert_eq!(lfu.get(1), -1);
    assert_eq!(lfu.get(2), 2);
    assert_eq!(lfu.get(3), 3);
}

fn test_update_existing() {
    let mut lfu = LFUCache::new(2);

    lfu.put(1, 1);
    lfu.put(2, 2);
    lfu.put(1, 10);

    lfu.put(3, 3);
    assert_eq!(lfu.get(1), 10);
    assert_eq!(lfu.get(2), -1);
    assert_eq!(lfu.get(3), 3);
}

fn test_promoted_key_evicts_older_tie() {
    let mut lfu = LFUCache::new(2);

    lfu.put(1, 1);
    lfu.put(2, 2);
    assert_eq!(lfu.get(1), 1);
    assert_eq!(lfu.get(2), 2);

    lfu.put(3, 3);
    assert_eq!(lfu.get(1), -1);
    assert_eq!(lfu.get(2), 2);
    assert_eq!(lfu.get(3), 3);
}

fn test_missing_key() {
    let mut lfu = LFUCache::new(2);

    assert_eq!(lfu.get(1), -1);
    lfu.put(1, 1);
    lfu.put(2, 2);
    assert_eq!(lfu.get(9), -1);

    lfu.put(3, 3);
    assert_eq!(lfu.get(1), -1);
    assert_eq!(lfu.get(2), 2);
    assert_eq!(lfu.get(3), 3);
}

fn main() {
    test_example();
    test_capacity_one();
    test_same_freq_evicts_least_recent();
    test_update_existing();
    test_promoted_key_evicts_older_tie();
    test_missing_key();
}
