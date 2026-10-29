use std::collections::{BTreeSet, HashMap};

#[derive(Debug, PartialEq, Eq, PartialOrd, Ord, Clone)]
struct Task {
    priority: i32,
    task_id: i32,
    user_id: i32,
}

struct TaskManager {
    task_by_id: HashMap<i32, Task>,
    task_set: BTreeSet<Task>,
}

impl TaskManager {
    fn new(tasks: Vec<Vec<i32>>) -> Self {
        let mut task_by_id = HashMap::new();
        let mut task_set = BTreeSet::new();
        for t in tasks.into_iter() {
            let task = Task {
                user_id: t[0],
                task_id: t[1],
                priority: t[2],
            };
            task_by_id.insert(task.task_id, task.clone());
            task_set.insert(task);
        }
        TaskManager {
            task_by_id,
            task_set,
        }
    }

    fn add(&mut self, user_id: i32, task_id: i32, priority: i32) {
        let task = Task {
            user_id,
            task_id,
            priority,
        };
        self.task_by_id.insert(task.task_id, task.clone());
        self.task_set.insert(task);
    }

    fn edit(&mut self, task_id: i32, new_priority: i32) {
        let task = self.task_by_id.get(&task_id).unwrap();
        let user_id = task.user_id;
        self.task_set.remove(task);
        self.task_by_id.remove(&task_id);

        let new_task = Task {
            priority: new_priority,
            task_id: task_id,
            user_id: user_id,
        };
        self.task_by_id.insert(task_id, new_task.clone());
        self.task_set.insert(new_task);
    }

    fn rmv(&mut self, task_id: i32) {
        let task = self.task_by_id.get(&task_id).unwrap();
        self.task_set.remove(task);
        self.task_by_id.remove(&task_id);
    }

    fn exec_top(&mut self) -> i32 {
        let Some(task) = self.task_set.pop_last() else {
            return -1;
        };
        self.task_by_id.remove(&task.task_id);
        task.user_id
    }
}

fn test_example() {
    let mut manager = TaskManager::new(vec![vec![1, 101, 10], vec![2, 102, 20], vec![3, 103, 15]]);

    manager.add(4, 104, 5);
    manager.edit(102, 8);
    assert_eq!(manager.exec_top(), 3);

    manager.rmv(101);
    manager.add(5, 105, 15);
    assert_eq!(manager.exec_top(), 5);
}

fn test_tie_breaks_by_task_id() {
    let mut manager = TaskManager::new(vec![vec![1, 1, 5], vec![2, 9, 5]]);

    assert_eq!(manager.exec_top(), 2);
    assert_eq!(manager.exec_top(), 1);
    assert_eq!(manager.exec_top(), -1);
}

fn test_edit_changes_top() {
    let mut manager = TaskManager::new(vec![vec![1, 1, 10], vec![2, 2, 5]]);

    manager.edit(2, 20);
    assert_eq!(manager.exec_top(), 2);
    assert_eq!(manager.exec_top(), 1);
}

fn test_rmv_removes_top() {
    let mut manager = TaskManager::new(vec![vec![1, 10, 5], vec![2, 20, 5]]);

    manager.rmv(20);
    assert_eq!(manager.exec_top(), 1);
    assert_eq!(manager.exec_top(), -1);
}

fn test_one_user_many_tasks() {
    let mut manager = TaskManager::new(vec![vec![1, 1, 3], vec![1, 2, 3]]);

    assert_eq!(manager.exec_top(), 1);
    assert_eq!(manager.exec_top(), 1);
    assert_eq!(manager.exec_top(), -1);
}

fn main() {
    test_example();
    test_tie_breaks_by_task_id();
    test_edit_changes_top();
    test_rmv_removes_top();
    test_one_user_many_tasks();
}
