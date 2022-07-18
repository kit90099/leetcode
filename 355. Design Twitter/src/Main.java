import java.util.List;

public class Main {
    public static void main(String... args){
        Twitter twitter = new Twitter();
        twitter.postTweet(1,5);
        twitter.postTweet(1,3);
        twitter.postTweet(1,101);
        twitter.postTweet(1,13);
        twitter.postTweet(1,10);
        twitter.postTweet(1,2);
        twitter.postTweet(1,94);
        twitter.postTweet(1,505);
        twitter.postTweet(1,333);
        twitter.postTweet(1,22);
        twitter.postTweet(1,11);
        List<Integer> feeds = twitter.getNewsFeed(1);
        System.out.println(twitter.getNewsFeed(1));
    }
}
