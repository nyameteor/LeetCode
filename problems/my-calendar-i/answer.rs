use std::collections::BTreeMap;

struct MyCalendar {
    end_by_start: BTreeMap<i32, i32>,
}

impl MyCalendar {
    fn new() -> Self {
        MyCalendar {
            end_by_start: BTreeMap::new(),
        }
    }

    fn book(&mut self, start_time: i32, end_time: i32) -> bool {
        let prev = self.end_by_start.range(..start_time).next_back();
        if prev.is_some_and(|(_, prev_end)| *prev_end > start_time) {
            return false;
        }

        let next = self.end_by_start.range(start_time..).next();
        if next.is_some_and(|(next_start, _)| *next_start < end_time) {
            return false;
        }

        self.end_by_start.insert(start_time, end_time);
        true
    }
}

fn test_example() {
    let mut calendar = MyCalendar::new();

    assert_eq!(calendar.book(10, 20), true);
    assert_eq!(calendar.book(15, 25), false);
    assert_eq!(calendar.book(20, 30), true);
}

fn test_touching_endpoints() {
    let mut calendar = MyCalendar::new();

    assert_eq!(calendar.book(20, 30), true);
    assert_eq!(calendar.book(10, 20), true);
    assert_eq!(calendar.book(30, 40), true);
}

fn test_overlap() {
    let mut calendar = MyCalendar::new();

    assert_eq!(calendar.book(10, 20), true);
    assert_eq!(calendar.book(12, 18), false);
    assert_eq!(calendar.book(5, 25), false);
    assert_eq!(calendar.book(8, 12), false);
    assert_eq!(calendar.book(10, 15), false);
    assert_eq!(calendar.book(5, 10), true);
}

fn test_rejected_booking_is_not_stored() {
    let mut calendar = MyCalendar::new();

    assert_eq!(calendar.book(10, 20), true);
    assert_eq!(calendar.book(15, 25), false);
    assert_eq!(calendar.book(20, 25), true);
}

fn main() {
    test_example();
    test_touching_endpoints();
    test_overlap();
    test_rejected_booking_is_not_stored();
}
