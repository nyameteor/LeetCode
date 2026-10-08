use std::collections::{HashMap, HashSet};

struct Twitter {
    follows: HashMap<i32, HashSet<i32>>,
    tweets: Vec<(i32, i32)>,
}

impl Twitter {
    fn new() -> Self {
        Twitter {
            follows: HashMap::new(),
            tweets: Vec::new(),
        }
    }

    fn post_tweet(&mut self, user_id: i32, tweet_id: i32) {
        self.tweets.push((user_id, tweet_id));
    }

    fn get_news_feed(&self, user_id: i32) -> Vec<i32> {
        let followees = self.follows.get(&user_id);

        self.tweets
            .iter()
            .rev()
            .filter(|(author, _)| {
                *author == user_id || followees.is_some_and(|set| set.contains(author))
            })
            .map(|(_, tweet_id)| *tweet_id)
            .take(10)
            .collect()
    }

    fn follow(&mut self, follower_id: i32, followee_id: i32) {
        self.follows
            .entry(follower_id)
            .or_default()
            .insert(followee_id);
    }

    fn unfollow(&mut self, follower_id: i32, followee_id: i32) {
        self.follows
            .entry(follower_id)
            .or_default()
            .remove(&followee_id);
    }
}

fn main() {
    let mut twitter = Twitter::new();

    twitter.post_tweet(1, 5);
    assert_eq!(twitter.get_news_feed(1), vec![5]);

    twitter.follow(1, 2);
    twitter.post_tweet(2, 6);
    assert_eq!(twitter.get_news_feed(1), vec![6, 5]);

    twitter.unfollow(1, 2);
    assert_eq!(twitter.get_news_feed(1), vec![5]);
}
