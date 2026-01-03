package main

import "sort"

// post represents a single tweet with metadata for timeline ordering
type post struct {
	userID      int // ID of the user who created this post
	tweetID     int // Unique identifier for the tweet
	timeCounter int // Monotonic counter for chronological ordering
}

// Twitter implements a simplified social media platform with posting and following features
type Twitter struct {
	userPosts   map[int][]post       // Maps userID to their list of posts
	userFollows map[int]map[int]bool // Maps followerID to a set of followeeIDs
	timeCounter int                  // Global counter to track post creation order
}

// Constructor initializes a new Twitter instance
// Time: O(1), Space: O(1)
func Constructor() Twitter {
	return Twitter{
		userPosts:   map[int][]post{},
		userFollows: map[int]map[int]bool{},
		timeCounter: 0,
	}
}

// PostTweet creates a new tweet for the given user
// Time: O(1) - appending to slice is amortized O(1)
// Space: O(1) - stores one post object
func (t *Twitter) PostTweet(userId int, tweetId int) {
	// Create a new post with current timestamp
	newPost := post{
		userID:      userId,
		tweetID:     tweetId,
		timeCounter: t.timeCounter,
	}

	// Add the post to the user's post list
	t.userPosts[userId] = append(t.userPosts[userId], newPost)

	// Increment the global time counter for next post
	t.timeCounter++
}

// GetNewsFeed retrieves the 10 most recent tweets from the user and their followees
// Time: O(N log N) where N is total posts from user + all followees
// Space: O(N) for collecting and sorting all relevant posts
func (t *Twitter) GetNewsFeed(userId int) []int {
	// Get the set of users this user follows
	followeeSet := t.userFollows[userId]

	// Collect all relevant posts (user's own posts + followees' posts)
	allRelevantPosts := []post{}

	// Add the user's own posts
	allRelevantPosts = append(allRelevantPosts, t.userPosts[userId]...)

	// Add posts from all users that this user follows
	for followeeID, isFollowing := range followeeSet {
		if isFollowing {
			allRelevantPosts = append(allRelevantPosts, t.userPosts[followeeID]...)
		}
	}

	// Sort posts by time counter in descending order (most recent first)
	sort.Slice(allRelevantPosts, func(i, j int) bool {
		return allRelevantPosts[i].timeCounter > allRelevantPosts[j].timeCounter
	})

	// Extract the top 10 (or fewer) tweet IDs
	feedSize := min(10, len(allRelevantPosts))
	newsFeed := make([]int, feedSize)
	for i := range feedSize {
		newsFeed[i] = allRelevantPosts[i].tweetID
	}

	return newsFeed
}

// Follow creates a follow relationship from follower to followee
// Time: O(1) - map lookup and insertion
// Space: O(1) - stores one boolean in nested map
func (t *Twitter) Follow(followerId int, followeeId int) {
	// Initialize the follower's follow map if it doesn't exist
	if _, exists := t.userFollows[followerId]; !exists {
		t.userFollows[followerId] = map[int]bool{}
	}

	// Mark the followee as followed
	t.userFollows[followerId][followeeId] = true
}

// Unfollow removes a follow relationship from follower to followee
// Time: O(1) - map deletion
// Space: O(1) - no additional space used
func (t *Twitter) Unfollow(followerId int, followeeId int) {
	// Remove the followee from the follower's follow set
	// Safe to call even if the relationship doesn't exist
	delete(t.userFollows[followerId], followeeId)
}
