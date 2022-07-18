import java.util.*;
import java.util.stream.Collectors;

class Twitter {

    User[] allUsers;
    int uuid = 0;

    public Twitter() {
        allUsers = new User[500];
    }

    public void postTweet(int userId, int tweetId) {
        User user = getUser(userId);
        Tweet tweet = new Tweet(tweetId, uuid++);
        user.feeds.addFirst(tweet);
    }

    public List<Integer> getNewsFeed(int userId) {
        User user = getUser(userId);
        return user.follows.stream().flatMap(u -> u.feeds.stream()).sorted((a,b)-> b.uuid - a.uuid).map(t -> t.id).limit(10).collect(Collectors.toList());
    }

    public void follow(int followerId, int followeeId) {
        User follower = getUser(followerId);
        User followee = getUser(followeeId);
        follower.follows.add(followee);
    }

    public void unfollow(int followerId, int followeeId) {
        User follower = getUser(followerId);
        User followee = getUser(followeeId);
        follower.follows.remove(followee);
    }

    User getUser(int userId) {
        User user = allUsers[userId];
        if (user == null) {
            user = new User(userId);
            allUsers[userId] = user;
        }
        return user;
    }
}

class User {
    int id;
    LinkedList<Tweet> feeds = new LinkedList();
    HashSet<User> follows = new HashSet<>();

    User(int id) {
        this.id = id;
        this.follows.add(this);
    }
}

class Tweet {
    int id;
    int uuid;

    Tweet(int id, int uuid) {
        this.id = id;
        this.uuid = uuid;
    }
}

/**
 * Your Twitter object will be instantiated and called as such:
 * Twitter obj = new Twitter();
 * obj.postTweet(userId,tweetId);
 * List<Integer> param_2 = obj.getNewsFeed(userId);
 * obj.follow(followerId,followeeId);
 * obj.unfollow(followerId,followeeId);
 */